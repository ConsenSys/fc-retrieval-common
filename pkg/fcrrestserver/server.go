/*
Package fcrrestserver - common methods to create HTTP endpoints and handle HTTP requests across the application.
*/
package fcrrestserver

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
	"io/ioutil"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
)

// FCRRESTServer represents a REST server handling http requests.
type FCRRESTServer struct {
	start       bool
	listenAddrs []string

	// handlers for different message type
	handlers map[string]map[int32]func(rw rest.ResponseWriter, request *fcrmessages.FCRMessage)
}

// NewFCRRESTServer creates an empty FCRRESTServer
func NewFCRRESTServer(
	listenAddrs []string,
) *FCRRESTServer {
	s := &FCRRESTServer{
		start:       false,
		listenAddrs: listenAddrs,
		handlers:    make(map[string]map[int32]func(rw rest.ResponseWriter, request *fcrmessages.FCRMessage)),
	}
	for _, listenAddr := range listenAddrs {
		s.handlers[listenAddr] = make(map[int32]func(rw rest.ResponseWriter, request *fcrmessages.FCRMessage))
	}
	return s
}

// AddHandler is used to add a handler to the server for a given type.
func (s *FCRRESTServer) AddHandler(listenAddr string, msgType int32, handler func(rw rest.ResponseWriter, request *fcrmessages.FCRMessage)) *FCRRESTServer {
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

// Start is used to start the server.
func (s *FCRRESTServer) Start() error {
	// Start server
	if s.start {
		return errors.New("server already started")
	}
	for _, listenAddr := range s.listenAddrs {
		errChan := make(chan bool)
		go func(addr string, errChan chan bool) {
			api := rest.NewApi()
			api.Use(rest.DefaultDevStack...)
			router, err := rest.MakeRouter(
				rest.Post("/v1", func(w rest.ResponseWriter, r *rest.Request) {
					s.msgRouter(w, r, addr)
				}),
			)
			if err != nil {
				logging.Error(err.Error())
				errChan <- true
				return
			}
			api.SetApp(router)
			errChan <- false
			logging.Error(http.ListenAndServe(":"+addr, api.MakeHandler()).Error())
		}(listenAddr, errChan)
		if <-errChan {
			return errors.New("fail to start REST Server")
		}
		logging.Info("REST server starts listening on %s for connections.", listenAddr)
	}
	s.start = true
	return nil
}

// msgRouter routes message
func (s *FCRRESTServer) msgRouter(w rest.ResponseWriter, r *rest.Request, listenAddr string) {
	logging.Trace("Received request via /v1 API")
	content, err := ioutil.ReadAll(r.Body)

	if closeErr := r.Body.Close(); closeErr != nil {
		logging.Error("msgRouter can't close request body")
	}

	if err != nil {
		logging.Error("Error reading request: %s.", err.Error())
		rest.Error(w, "Error reading request", http.StatusBadRequest)
		return
	}
	if len(content) == 0 {
		logging.Error("Error empty request")
		rest.Error(w, "Error empty request", http.StatusBadRequest)
		return
	}
	message, err := fcrmessages.FCRMsgFromBytes(content)
	if err != nil {
		logging.Error("Failed to decode payload: %s.", err.Error())
		rest.Error(w, "Failed to decode payload: "+err.Error(), http.StatusBadRequest)
		return
	}
	handler := s.handlers[listenAddr][message.GetMessageType()]
	if handler != nil {
		handler(w, message)
	} else {
		logging.Warn("Client Request: Unknown message type: %d", message.GetMessageType())
		rest.Error(w, "Unknown message type", http.StatusBadRequest)
	}
}
