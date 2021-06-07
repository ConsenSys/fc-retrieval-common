/*
Package fcrp2pserver - contains common operations used by FileCoin Secondary Retrieval network nodes to establish
and perform TCP communications between the nodes.

Operations like get a Gateway or Provider by an identifier, send or read messages.
*/
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

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrregistermgr"
	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// Constants for identifying the correct access point
const (
	accessFromGateway  = 0
	accessFromProvider = 1
)

// communicationChannel holds the connection for sending outgoing TCP requests.
// lock is used to ensure only one thread can access the tcp connection at any time.
// conn is the net connection for sending outgoing TCP requests.
type communicationChannel struct {
	lock sync.RWMutex
	conn net.Conn
}

// communicationPool holds the node address map and active node connections.
type communicationPool struct {
	registerMgr *fcrregistermgr.FCRRegisterMgr

	activeGateways     map[string](*communicationChannel)
	activeGatewaysLock sync.RWMutex

	activeProviders     map[string](*communicationChannel)
	activeProvidersLock sync.RWMutex
}

// getGatewayConn gets a connection to a given gateway for sending request
func (c *communicationPool) getGatewayConn(id *nodeid.NodeID, accessFrom int) (*communicationChannel, error) {
	c.activeGatewaysLock.RLock()
	comm := c.activeGateways[id.ToString()]
	c.activeGatewaysLock.RUnlock()
	if comm != nil {
		return comm, nil
	}
	logging.Info("P2P server has no active connection to gateway %s, attempt connecting", id.ToString())
	gatewayInfo := c.registerMgr.GetGateway(id)
	if gatewayInfo == nil {
		return nil, errors.New("Gateway not found")
	}
	// Get address
	var address string
	switch accessFrom {
	case accessFromGateway:
		address = gatewayInfo.GetNetworkInfoGateway()
	case accessFromProvider:
		address = gatewayInfo.GetNetworkInfoProvider()
	}
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	// Get connection
	// Store the communication
	c.activeGatewaysLock.Lock()
	// It is possible that another thread creates a connection before this thread,
	// so do a final check here.
	if c.activeGateways[id.ToString()] == nil {
		comm = &communicationChannel{
			lock: sync.RWMutex{},
			conn: conn,
		}
		c.activeGateways[id.ToString()] = comm
	} else {
		comm = c.activeGateways[id.ToString()]
		conn.Close()
	}
	c.activeGatewaysLock.Unlock()
	return comm, nil
}

// getProviderConn gets a connection to a given provider for sending request
func (c *communicationPool) getProviderConn(id *nodeid.NodeID) (*communicationChannel, error) {
	c.activeProvidersLock.RLock()
	comm := c.activeProviders[id.ToString()]
	c.activeProvidersLock.RUnlock()
	if comm != nil {
		return comm, nil
	}
	logging.Info("P2P server has no active connection to provider %s, attempt connecting", id.ToString())
	providerInfo := c.registerMgr.GetProvider(id)
	if providerInfo == nil {
		return nil, errors.New("Provider not found")
	}
	// Get address
	address := providerInfo.GetNetworkInfoGateway()
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	// Get connection
	// Store the communication
	c.activeProvidersLock.Lock()
	// It is possible that another thread creates a connection before this thread,
	// so do a final check here.
	if c.activeProviders[id.ToString()] == nil {
		comm = &communicationChannel{
			lock: sync.RWMutex{},
			conn: conn,
		}
		c.activeProviders[id.ToString()] = comm
	} else {
		comm = c.activeProviders[id.ToString()]
		conn.Close()
	}
	c.activeProvidersLock.Unlock()
	return comm, nil
}

// removeActiveGateway remove a given gateway, and close the connection.
func (c *communicationPool) removeActiveGateway(id *nodeid.NodeID) {
	c.activeGatewaysLock.Lock()
	comm := c.activeGateways[id.ToString()]
	if comm != nil {
		comm.lock.Lock()
		comm.conn.Close()
		comm.lock.Unlock()
		delete(c.activeGateways, id.ToString())
	}
	c.activeGatewaysLock.Unlock()
}

// removeActiveProvider remove a given provider, and close the connection.
func (c *communicationPool) removeActiveProvider(id *nodeid.NodeID) {
	c.activeProvidersLock.Lock()
	comm := c.activeProviders[id.ToString()]
	if comm != nil {
		comm.lock.Lock()
		comm.conn.Close()
		comm.lock.Unlock()
		delete(c.activeProviders, id.ToString())
	}
	c.activeProvidersLock.Unlock()
}
