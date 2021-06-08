package fcrp2pserver

/*
 * Copyright 2020 ConsenSys Software Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import (
	"errors"
	"net"
	"sync"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrregistermgr"
	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// FCRP2PServer represents a server handling p2p connection using tcp.
type FCRP2PServer struct {
	start       bool
	listenAddrs []string
	timeout     time.Duration

	// Connection pool
	pool *communicationPool

	// handlers for different message type
	handlers   map[string]map[int32]func(reader *FCRServerReader, writer *FCRServerWriter, request *fcrmessages.FCRMessage) error
	requesters map[int32]func(reader *FCRServerReader, writer *FCRServerWriter, args ...interface{}) (*fcrmessages.FCRMessage, error)
}

// NewFCRP2PServer creates an empty FCRP2PServer.
func NewFCRP2PServer(
	listenAddrs []string,
	registerMgr *fcrregistermgr.FCRRegisterMgr,
	defaultTimeout time.Duration) *FCRP2PServer {
	s := &FCRP2PServer{
		start:       false,
		listenAddrs: listenAddrs,
		timeout:     defaultTimeout,
		pool: &communicationPool{
			registerMgr:         registerMgr,
			activeGateways:      make(map[string]*communicationChannel),
			activeGatewaysLock:  sync.RWMutex{},
			activeProviders:     make(map[string]*communicationChannel),
			activeProvidersLock: sync.RWMutex{},
		},
		handlers:   make(map[string]map[int32]func(reader *FCRServerReader, writer *FCRServerWriter, request *fcrmessages.FCRMessage) error),
		requesters: make(map[int32]func(reader *FCRServerReader, writer *FCRServerWriter, args ...interface{}) (*fcrmessages.FCRMessage, error)),
	}
	for _, listenAddr := range listenAddrs {
		s.handlers[listenAddr] = make(map[int32]func(reader *FCRServerReader, writer *FCRServerWriter, request *fcrmessages.FCRMessage) error)
	}
	return s
}

// AddHandler is used to add a handler to the server for a given type.
func (s *FCRP2PServer) AddHandler(listenAddr string, msgType int32, handler func(reader *FCRServerReader, writer *FCRServerWriter, request *fcrmessages.FCRMessage) error) *FCRP2PServer {
	if s.start {
		return s
	}
	addrHandler, ok := s.handlers[listenAddr]
	if !ok {
		return s
	}
	addrHandler[msgType] = handler
	return s
}

// AddRequester is used to add a requester to the server for a given type.
func (s *FCRP2PServer) AddRequester(msgType int32, requester func(reader *FCRServerReader, writer *FCRServerWriter, args ...interface{}) (*fcrmessages.FCRMessage, error)) *FCRP2PServer {
	if s.start {
		return s
	}
	s.requesters[msgType] = requester
	return s
}

// Start is used to start the server.
func (s *FCRP2PServer) Start() error {
	// Start server
	if s.start {
		return errors.New("server already started")
	}
	for _, listenAddr := range s.listenAddrs {
		ln, err := net.Listen("tcp", ":"+listenAddr)
		if err != nil {
			return err
		}
		go func(ln net.Listener, listenAddr string) {
			for {
				conn, err := ln.Accept()
				if err != nil {
					logging.Error("P2P server has error accepting connection: %s", err.Error())
					continue
				}
				logging.Info("P2P server has incoming connection from :%s", conn.RemoteAddr())
				go s.handleIncomingConnection(conn, s.handlers[listenAddr])
			}
		}(ln, listenAddr)
		logging.Info("P2P server starts listening on %s for connections.", listenAddr)
	}
	s.start = true
	return nil
}

// handleIncomingConnection handles incomming connection using given handlers.
func (s *FCRP2PServer) handleIncomingConnection(conn net.Conn, handlers map[int32]func(reader *FCRServerReader, writer *FCRServerWriter, request *fcrmessages.FCRMessage) error) {
	// Close connection on exit.
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()

	// Loop until error occurs and connection is dropped.
	for {
		message, err := readTCPMessage(conn, s.timeout)
		if err != nil && !isTimeoutError(err) {
			// Error in tcp communication, drop the connection.
			logging.Error("P2P Server has error reading message from %s: %s", conn.RemoteAddr(), err.Error())
			return
		}
		// TODO: discard a connection if it doesn’t give a valid response for a really long time
		if err != nil && isTimeoutError(err) {
			continue
		}
		handler := handlers[message.GetMessageType()]
		if handler != nil {
			// Call handler to handle the request
			writer := &FCRServerWriter{conn: conn}
			reader := &FCRServerReader{conn: conn}
			err = handler(reader, writer, message)
			if err != nil {
				// Error that couldn't ignore, drop the connection.
				logging.Error("P2P Server has error handling message from %s: %s", conn.RemoteAddr(), err.Error())
				return
			}
		} else {
			// Message is invalid.
			err = sendInvalidMessage(conn, s.timeout)
			if err != nil {
				// Error in tcp communication, drop the connection.
				logging.Error("P2P Server has error responding to %s: %s", conn.RemoteAddr(), err.Error())
				return
			}
		}
	}
}

// RequestGatewayFromGateway uses a given requester to send a request to a given gateway from gateway.
func (s *FCRP2PServer) RequestGatewayFromGateway(id *nodeid.NodeID, msgType int32, args ...interface{}) (*fcrmessages.FCRMessage, error) {
	if !s.start {
		return nil, errors.New("server not started")
	}
	requester := s.requesters[msgType]
	if requester == nil {
		return nil, errors.New("no available requester found for given type")
	}
	comm, err := s.pool.getGatewayConn(id, accessFromGateway)
	if err != nil {
		logging.Error("P2P Server has error get gateway connection to %s: %s", id.ToString(), err.Error())
		return nil, err
	}
	comm.lock.Lock()
	// Call requester to request
	writer := &FCRServerWriter{conn: comm.conn}
	reader := &FCRServerReader{conn: comm.conn}
	response, err := requester(reader, writer, args...)
	comm.lock.Unlock()
	if err != nil {
		// Error that couldn't ignore, remove the connection.
		s.pool.removeActiveGateway(id)
		return nil, err
	}
	return response, err
}

// RequestGatewayFromProvider uses a given requester to send a request to a given gateway from provider.
func (s *FCRP2PServer) RequestGatewayFromProvider(id *nodeid.NodeID, msgType int32, args ...interface{}) (*fcrmessages.FCRMessage, error) {
	if !s.start {
		return nil, errors.New("server not started")
	}
	requester := s.requesters[msgType]
	if requester == nil {
		return nil, errors.New("no available requester found for given type")
	}
	comm, err := s.pool.getGatewayConn(id, accessFromProvider)
	if err != nil {
		logging.Error("P2P Server has error get gateway connection to %s: %s", id.ToString(), err.Error())
		return nil, err
	}
	comm.lock.Lock()
	// Call requester to request
	writer := &FCRServerWriter{conn: comm.conn}
	reader := &FCRServerReader{conn: comm.conn}
	response, err := requester(reader, writer, args...)
	comm.lock.Unlock()
	if err != nil {
		// Error that couldn't ignore, remove the connection.
		s.pool.removeActiveGateway(id)
		return nil, err
	}
	return response, err
}

// RequestProvider uses a given requester to send a request to a given provider. (Only possible from gateway)
func (s *FCRP2PServer) RequestProvider(id *nodeid.NodeID, msgType int32, args ...interface{}) (*fcrmessages.FCRMessage, error) {
	if !s.start {
		return nil, errors.New("server not started")
	}
	requester := s.requesters[msgType]
	if requester == nil {
		return nil, errors.New("no available requester found for given type")
	}
	comm, err := s.pool.getProviderConn(id)
	if err != nil {
		logging.Error("P2P Server has error get provider connection to %s: %s", id.ToString(), err.Error())
		return nil, err
	}
	comm.lock.Lock()
	// Call requester to request
	writer := &FCRServerWriter{conn: comm.conn}
	reader := &FCRServerReader{conn: comm.conn}
	response, err := requester(reader, writer, args...)
	comm.lock.Unlock()
	if err != nil {
		// Error that couldn't ignore, remove the connection.
		s.pool.removeActiveProvider(id)
		return nil, err
	}
	return response, err
}
