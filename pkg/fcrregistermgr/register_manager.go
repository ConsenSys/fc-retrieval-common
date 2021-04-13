package fcrregistermgr

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
	"sync"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-register/pkg/register"
)

// To implement register manager.
// 1. NewFCRRegisterMgr() initialises a register manager.
//		If given gateway param is true, initialise a map[nodeid](register.GatewayRegister) <- this can be a db
//		If given provider param is true, initialise a map[nodeid](register.ProviderRegister) <- this can be a db
// 2. Start() starts a thread to auto update the internal map every given duration.
//		Make sure Start() only works at first time, later attempt to call should fail.
// 3. Refresh() refreshs the internal map immediately.
// 4. GetGateway() returns a gateway register if found.
// 5. GetProvider() returns a provider register if found.
// Note to handle access from multiple threads (RWMutex is required to access internal storage for example).

// Register Manager manages the internal storage of registered nodes.
type FCRRegisterMgr struct {
	// Boolean indicates if the manager has started
	start bool

	// API for the register
	registerAPI string

	// Duration to wait between two updates
	refreshDuration time.Duration

	// Boolean indicates if to discover gateway/provider
	gatewayDiscv  bool
	providerDiscv bool

	// Gateway ID of self (only required when gateway is true)
	gatewayID string

	// Channels to control the threads
	gatewayShutdownCh  chan bool
	providerShutdownCh chan bool
	gatewayRefreshCh   chan bool
	providerRefreshCh  chan bool

	// registeredGatewaysMap stores mapping from gateway id (big int in string repr) to its registration info
	registeredGatewaysMap     map[string]register.GatewayRegister
	registeredGatewaysMapLock sync.RWMutex

	// registeredProvidersMap stores mapping from provider id (big int in string repr) to its registration info
	registeredProvidersMap     map[string]register.ProviderRegister
	registeredProvidersMapLock sync.RWMutex
}

// NewFCRRegisterMgr creates a new register manager.
func NewFCRRegisterMgr(registerAPI string, providerDiscv bool, gatewayDiscv bool, gatewayID nodeid.NodeID, refreshDuration time.Duration) *FCRRegisterMgr {
	res := &FCRRegisterMgr{
		start:           false,
		registerAPI:     registerAPI,
		refreshDuration: refreshDuration,
		gatewayDiscv:    gatewayDiscv,
		providerDiscv:   providerDiscv,
		gatewayID:       gatewayID.ToString(),
	}
	if gatewayDiscv {
		res.registeredGatewaysMap = make(map[string]register.GatewayRegister)
		res.registeredGatewaysMapLock = sync.RWMutex{}
		res.gatewayShutdownCh = make(chan bool)
		res.gatewayRefreshCh = make(chan bool)
	}
	if providerDiscv {
		res.registeredProvidersMap = make(map[string]register.ProviderRegister)
		res.registeredProvidersMapLock = sync.RWMutex{}
		res.providerShutdownCh = make(chan bool)
		res.providerRefreshCh = make(chan bool)
	}
	return res
}

// Start starts a thread to auto update the internal map every given duration.
func (mgr *FCRRegisterMgr) Start() error {
	if mgr.start {
		return errors.New("Manager has already started.")
	}
	mgr.start = true
	if mgr.gatewayDiscv {
		go mgr.updateGateways()
	}
	if mgr.providerDiscv {
		go mgr.updateProviders()
	}
	return nil
}

// Shutdown will shutdown the register manager.
func (mgr *FCRRegisterMgr) Shutdown() {
	if !mgr.start {
		return
	}
	if mgr.gatewayDiscv {
		mgr.gatewayShutdownCh <- true
	}
	if mgr.providerDiscv {
		mgr.providerShutdownCh <- true
	}
	mgr.start = false
}

// Refresh refreshs the internal map immediately.
func (mgr *FCRRegisterMgr) Refresh() {
	if !mgr.start {
		return
	}
	if mgr.gatewayDiscv {
		mgr.gatewayRefreshCh <- true
	}
	if mgr.providerDiscv {
		mgr.providerRefreshCh <- true
	}
}

// GetGateway returns a gateway register if found.
func (mgr *FCRRegisterMgr) GetGateway(id *nodeid.NodeID) *register.GatewayRegister {
	if !mgr.start || !mgr.gatewayDiscv {
		return nil
	}
	mgr.registeredGatewaysMapLock.RLock()
	gateway, ok := mgr.registeredGatewaysMap[id.ToString()]
	mgr.registeredGatewaysMapLock.RUnlock()
	if !ok {
		// TODO: Do we call refresh here, if can't find a gateway?
		// mgr.Refresh()
		mgr.registeredGatewaysMapLock.RLock()
		gateway = mgr.registeredGatewaysMap[id.ToString()]
		mgr.registeredGatewaysMapLock.RUnlock()
	}
	return &gateway
}

// GetProvider returns a provider register if found.
func (mgr *FCRRegisterMgr) GetProvider(id *nodeid.NodeID) *register.ProviderRegister {
	if !mgr.start || !mgr.providerDiscv {
		return nil
	}
	mgr.registeredProvidersMapLock.RLock()
	provider, ok := mgr.registeredProvidersMap[id.ToString()]
	mgr.registeredProvidersMapLock.RUnlock()
	if !ok {
		// TODO: Do we call refresh here, if can't find a provider?
		// mgr.Refresh()
		mgr.registeredProvidersMapLock.RLock()
		provider = mgr.registeredProvidersMap[id.ToString()]
		mgr.registeredProvidersMapLock.RUnlock()
	}
	return &provider
}

// updateGateways updates gateways.
func (mgr *FCRRegisterMgr) updateGateways() {
	for {
		gateways, err := register.GetRegisteredGateways(mgr.registerAPI)
		if err != nil {
			logging.Error("Register manager has error in getting registered gateways: %s", err.Error())
		} else {
			// Check for update
			for _, gateway := range gateways {
				// Skip itself
				if gateway.NodeID == mgr.gatewayID {
					continue
				}
				mgr.registeredGatewaysMapLock.RLock()
				storedInfo, ok := mgr.registeredGatewaysMap[gateway.NodeID]
				mgr.registeredGatewaysMapLock.RUnlock()
				if !ok {
					// Not exist, we need to add a new entry
					mgr.registeredGatewaysMapLock.Lock()
					mgr.registeredGatewaysMap[gateway.NodeID] = gateway
					mgr.registeredGatewaysMapLock.Unlock()
				} else {
					// Exist, check if need update
					if gateway.Address != storedInfo.Address ||
						gateway.RootSigningKey != storedInfo.RootSigningKey ||
						gateway.SigningKey != storedInfo.SigningKey ||
						gateway.RegionCode != storedInfo.RegionCode ||
						gateway.NetworkInfoGateway != storedInfo.NetworkInfoGateway ||
						gateway.NetworkInfoProvider != storedInfo.NetworkInfoProvider ||
						gateway.NetworkInfoClient != storedInfo.NetworkInfoClient ||
						gateway.NetworkInfoAdmin != storedInfo.NetworkInfoAdmin {
						// Need update
						mgr.registeredGatewaysMapLock.Lock()
						mgr.registeredGatewaysMap[gateway.NodeID] = gateway
						mgr.registeredGatewaysMapLock.Unlock()
					}
				}
			}
		}
		afterChan := time.After(mgr.refreshDuration)
		select {
		case <-mgr.gatewayRefreshCh:
			// Need to refresh
			logging.Error("Register manager force update internal gateway map.")
		case <-afterChan:
			// Need to refresh
		case <-mgr.gatewayShutdownCh:
			// Need to shutdown
			logging.Error("Register manager shutdown gateway routine.")
			return
		}
	}
}

// updateProviders updates providers.
func (mgr *FCRRegisterMgr) updateProviders() {
	for {
		providers, err := register.GetRegisteredProviders(mgr.registerAPI)
		if err != nil {
			logging.Error("Register manager has error in getting registered providers: %s", err.Error())
		} else {
			// Check for update
			for _, provider := range providers {
				mgr.registeredProvidersMapLock.RLock()
				storedInfo, ok := mgr.registeredProvidersMap[provider.NodeID]
				mgr.registeredProvidersMapLock.RUnlock()
				if !ok {
					// Not exist, we need to add a new entry
					mgr.registeredProvidersMapLock.Lock()
					mgr.registeredProvidersMap[provider.NodeID] = provider
					mgr.registeredProvidersMapLock.Unlock()
				} else {
					// Exist, check if need update
					if provider.Address != storedInfo.Address ||
						provider.RootSigningKey != storedInfo.RootSigningKey ||
						provider.SigningKey != storedInfo.SigningKey ||
						provider.RegionCode != storedInfo.RegionCode ||
						provider.NetworkInfoGateway != storedInfo.NetworkInfoGateway ||
						provider.NetworkInfoClient != storedInfo.NetworkInfoClient ||
						provider.NetworkInfoAdmin != storedInfo.NetworkInfoAdmin {
						// Need update
						mgr.registeredProvidersMapLock.Lock()
						mgr.registeredProvidersMap[provider.NodeID] = provider
						mgr.registeredProvidersMapLock.Unlock()
					}
				}
			}
		}
		afterChan := time.After(mgr.refreshDuration)
		select {
		case <-mgr.providerRefreshCh:
			// Need to refresh
			logging.Error("Register manager force update internal provider map.")
		case <-afterChan:
			// Need to refresh
		case <-mgr.providerShutdownCh:
			// Need to shutdown
			logging.Error("Register manager shutdown provider routine.")
			return
		}
	}
}
