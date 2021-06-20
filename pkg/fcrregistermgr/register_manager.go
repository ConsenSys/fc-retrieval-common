/*
Package fcrregistermgr - provides network API methods to work with FileCoin Secondary Retrieval Manager.
This package uses core Register functionality from `register` package internally.

Retrieval Register is a central node, holding information about Retrieval Gateways and Retrieval Providers.
*/
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
  "encoding/json"
  "errors"
  "sync"
  "time"

  "github.com/ConsenSys/fc-retrieval-common/pkg/cid"
  "github.com/ConsenSys/fc-retrieval-common/pkg/dhtring"
  "github.com/ConsenSys/fc-retrieval-common/pkg/logging"
  "github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
  "github.com/ConsenSys/fc-retrieval-common/pkg/register"
  "github.com/ConsenSys/fc-retrieval-common/pkg/request"
)

// FCRRegisterMgr Register Manager manages the internal storage of registered nodes.
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

	// Channels to control the threads
	gatewayShutdownCh  chan bool
	providerShutdownCh chan bool
	gatewayRefreshCh   chan bool
	providerRefreshCh  chan bool

	// registeredGatewaysMap stores mapping from gateway id (big int in string repr) to its registration info
	registeredGatewaysMap     map[string]register.GatewayRegistrar
	registeredGatewaysMapLock sync.RWMutex

	// closestGateways stores the mapping from gateway closest for DHT network sorted clockwise
	closestGatewaysIDs     *dhtring.Ring
	closestGatewaysIDsLock sync.RWMutex

	// registeredProvidersMap stores mapping from provider id (big int in string repr) to its registration info
	registeredProvidersMap     map[string]register.ProviderRegistrar
	registeredProvidersMapLock sync.RWMutex

  httpCommunicator    request.HttpCommunications
}

// NewFCRRegisterMgr creates a new register manager.
func NewFCRRegisterMgr(registerAPI string, providerDiscv bool, gatewayDiscv bool, refreshDuration time.Duration) *FCRRegisterMgr {
	res := &FCRRegisterMgr{
		start:           false,
		registerAPI:     registerAPI,
		refreshDuration: refreshDuration,
		gatewayDiscv:    gatewayDiscv,
		providerDiscv:   providerDiscv,
		httpCommunicator: request.NewHttpCommunicator(),
	}
	if gatewayDiscv {
		res.registeredGatewaysMap = make(map[string]register.GatewayRegistrar)
		res.registeredGatewaysMapLock = sync.RWMutex{}
		res.gatewayShutdownCh = make(chan bool)
		res.gatewayRefreshCh = make(chan bool)
		res.closestGatewaysIDs = dhtring.CreateRing()
		res.closestGatewaysIDsLock = sync.RWMutex{}
	}
	if providerDiscv {
		res.registeredProvidersMap = make(map[string]register.ProviderRegistrar)
		res.registeredProvidersMapLock = sync.RWMutex{}
		res.providerShutdownCh = make(chan bool)
		res.providerRefreshCh = make(chan bool)
	}
	return res
}

// Start starts a thread to auto update the internal map every given duration.
func (mgr *FCRRegisterMgr) Start() error {
	if mgr.start {
		return errors.New("manager has already started")
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

// Refresh refreshes the internal map immediately.
func (mgr *FCRRegisterMgr) Refresh() {
	if !mgr.start {
		return
	}
	if mgr.gatewayDiscv {
		mgr.gatewayRefreshCh <- true
		<-mgr.gatewayRefreshCh
	}
	if mgr.providerDiscv {
		mgr.providerRefreshCh <- true
		<-mgr.providerRefreshCh
	}
}

// RegisterGateway to register a gateway
func (mgr *FCRRegisterMgr) RegisterGateway(gatewayRegistrar register.GatewayRegistrar) error {
  url := mgr.registerAPI + "/registers/gateway"
  return mgr.httpCommunicator.SendJSON(url, gatewayRegistrar.Serialize())
}

// RegisterProvider to register a provider
func (mgr *FCRRegisterMgr) RegisterProvider(providerRegistrar register.ProviderRegistrar) error {
  url := mgr.registerAPI + "/registers/provider"
  return mgr.httpCommunicator.SendJSON(url, providerRegistrar.Serialize())
}

// GetGateway returns a gateway register if found.
func (mgr *FCRRegisterMgr) GetGateway(id *nodeid.NodeID) register.GatewayRegistrar {
	if !mgr.start || !mgr.gatewayDiscv {
    logging.Error("method GetGateway called while Register Manager is not started or gateway discovery is not enabled")
		return nil
	}
	mgr.registeredGatewaysMapLock.RLock()
	gateway, ok := mgr.registeredGatewaysMap[id.ToString()]
	mgr.registeredGatewaysMapLock.RUnlock()
	if !ok {
		mgr.Refresh()
		mgr.registeredGatewaysMapLock.RLock()
		defer mgr.registeredGatewaysMapLock.RUnlock()
		gateway, ok = mgr.registeredGatewaysMap[id.ToString()]
		if !ok {
			return nil
		}
	}
	return gateway
}

// GetProvider returns a provider register if found.
func (mgr *FCRRegisterMgr) GetProvider(id *nodeid.NodeID) register.ProviderRegistrar {
	if !mgr.start || !mgr.providerDiscv {
		return nil
	}
	mgr.registeredProvidersMapLock.RLock()
	provider, ok := mgr.registeredProvidersMap[id.ToString()]
	if !ok {
		mgr.registeredProvidersMapLock.RUnlock()
		// TODO: Do we call refresh here, if can't find a provider?
		// mgr.Refresh()
		mgr.registeredProvidersMapLock.RLock()
		provider, ok = mgr.registeredProvidersMap[id.ToString()]
		if !ok {
			return nil
		}
	}
	defer mgr.registeredProvidersMapLock.RUnlock()
	return provider
}

// GetAllGateways returns  all discovered gateways.
func (mgr *FCRRegisterMgr) GetAllGateways() []register.GatewayRegistrar {
  if !mgr.start || !mgr.gatewayDiscv {
    return nil
  }
  res := make([]register.GatewayRegistrar, 0)
  mgr.registeredGatewaysMapLock.RLock()
  defer mgr.registeredGatewaysMapLock.RUnlock()
  for _, gateway := range mgr.registeredGatewaysMap {
    res = append(res, gateway)
  }
  return res
}

// GetAllProviders returns  all discovered providers.
func (mgr *FCRRegisterMgr) GetAllProviders() []register.ProviderRegistrar {
  if !mgr.start || !mgr.providerDiscv {
    return nil
  }
  res := make([]register.ProviderRegistrar, 0)
  mgr.registeredProvidersMapLock.RLock()
  defer mgr.registeredProvidersMapLock.RUnlock()
  for _, provider := range mgr.registeredProvidersMap {
    res = append(res, provider)
  }
  return res
}

// pullGatewaysFromRegisterSrv calls remote service to synchronize discovered Gateway nodes
func (mgr *FCRRegisterMgr) pullGatewaysFromRegisterSrv() []register.GatewayRegistrar {
	if !mgr.start || !mgr.gatewayDiscv {
		return nil
	}
  url := mgr.registerAPI + "/registers/gateway/"
  rspBytes, err := mgr.httpCommunicator.GetJSON(url)
  if err != nil {
    logging.Error("error updating gateways: %s", err.Error())
    return nil
  }
  var gateways []*register.GatewayRegister
  if err := json.Unmarshal(rspBytes, &gateways); err != nil {
    logging.Error("error updating gateways, invalid response")
    return nil
  }

  var result []register.GatewayRegistrar
  for _, g := range gateways {
    result = append(result, register.NewGatewayRegister(g.NodeID, g.Address, g.RootSigningKey, g.SigningKey, g.RegionCode, g.NetworkInfoGateway, g.NetworkInfoProvider, g.NetworkInfoClient, g.NetworkInfoAdmin))
  }
  return result
}

// pullProvidersFromRegisterSrv calls remote service to synchronize discovered Provider nodes
func (mgr *FCRRegisterMgr) pullProvidersFromRegisterSrv() []register.ProviderRegistrar {
	if !mgr.start || !mgr.providerDiscv {
		return nil
	}
  url := mgr.registerAPI + "/registers/provider/"
  var providers []*register.ProviderRegister
  rspBytes, err := mgr.httpCommunicator.GetJSON(url)
  if err != nil {
    logging.Error("error updating providers: %s", err.Error())
    return nil
  }
  if err := json.Unmarshal(rspBytes, &providers); err != nil {
    logging.Error("error updating providers, invalid response")
    return nil
  }
  var result []register.ProviderRegistrar
  for _, g := range providers {
    result = append(result, register.NewProviderRegister(g.NodeID, g.Address, g.RootSigningKey, g.SigningKey, g.RegionCode, g.NetworkInfoGateway, g.NetworkInfoClient, g.NetworkInfoAdmin))
  }
  return result
}

// GetGatewayCIDRange gets the cid max and cid min of this gateway at start up
func (mgr *FCRRegisterMgr) GetGatewayCIDRange(gatewayID *nodeid.NodeID) (*cid.ContentID, *cid.ContentID, error) {
	mgr.closestGatewaysIDsLock.RLock()
	defer mgr.closestGatewaysIDsLock.RUnlock()

	cID, err := cid.NewContentIDFromHexString(gatewayID.ToString())
	if err != nil {
		return nil, nil, err
	}
	temp, err := mgr.GetGatewaysNearCID(cID, 16, gatewayID)
	if err != nil {
		return nil, nil, err
	}
	if len(temp) < 2 {
		// TODO: What is the cid max and cid min if there is only 1 gw or 2.
		return nil, nil, errors.New("not enough gateways")
	}
	cidMin, err := cid.NewContentIDFromHexString(temp[0].GetNodeID())
	if err != nil {
		return nil, nil, err
	}
	cidMax, err := cid.NewContentIDFromHexString(temp[len(temp)-1].GetNodeID())
	if err != nil {
		return nil, nil, err
	}
	return cidMin, cidMax, nil
}

// GetGatewaysNearCID returns a list of gatewayRegisters whose id is close to the given cid.
func (mgr *FCRRegisterMgr) GetGatewaysNearCID(cID *cid.ContentID, numDHT int, notAllowed *nodeid.NodeID) ([]register.GatewayRegistrar, error) {
	if numDHT > 16 {
		numDHT = 16
	}
	mgr.closestGatewaysIDsLock.RLock()
	defer mgr.closestGatewaysIDsLock.RUnlock()
	mgr.registeredGatewaysMapLock.RLock()
	defer mgr.registeredGatewaysMapLock.RUnlock()

	var ids []string
	var err error
	if notAllowed == nil {
		ids, err = mgr.closestGatewaysIDs.GetClosest(cID.ToString(), numDHT, "")
	} else {
		ids, err = mgr.closestGatewaysIDs.GetClosest(cID.ToString(), numDHT, notAllowed.ToString())
	}
	if err != nil {
		return nil, err
	}
	res := make([]register.GatewayRegistrar, 0)
	for _, id := range ids {
		gateway, exists := mgr.registeredGatewaysMap[id]
		if !exists {
			return nil, errors.New("internal error, register map not match dht ring")
		}
		res = append(res, gateway)
	}
	return res, nil
}

func (mgr *FCRRegisterMgr) mapNodeIDs(nodsIDs []*nodeid.NodeID) map[string]string {
	mapNodeIDs := make(map[string]string)
	for _, v := range nodsIDs {
		mapNodeIDs[v.ToString()] = v.ToString()
	}
	return mapNodeIDs
}

// updateGateways updates gateways.
// TODO: Support removal of gateways
func (mgr *FCRRegisterMgr) updateGateways() {
	refreshForce := false
	for {
		gateways := mgr.pullGatewaysFromRegisterSrv()
    // Check for update
    for _, gateway := range gateways {
      mgr.registeredGatewaysMapLock.RLock()
      storedInfo, ok := mgr.registeredGatewaysMap[gateway.GetNodeID()]
      mgr.registeredGatewaysMapLock.RUnlock()
      if !ok {
        // Not exist, we need to add a new entry
        mgr.registeredGatewaysMapLock.Lock()
        mgr.registeredGatewaysMap[gateway.GetNodeID()] = gateway
        mgr.registeredGatewaysMapLock.Unlock()
        mgr.closestGatewaysIDsLock.Lock()
        mgr.closestGatewaysIDs.Insert(gateway.GetNodeID())
        mgr.closestGatewaysIDsLock.Unlock()
      } else {
        // Exist, check if need update
        if gateway != storedInfo {
          // Need update
          mgr.registeredGatewaysMapLock.Lock()
          mgr.registeredGatewaysMap[gateway.GetNodeID()] = gateway
          mgr.registeredGatewaysMapLock.Unlock()
        }
      }
    }

		if refreshForce {
			mgr.gatewayRefreshCh <- true
			refreshForce = false
		}
		afterChan := time.After(mgr.refreshDuration)
		select {
		case <-mgr.gatewayRefreshCh:
			// Need to refresh
			logging.Error("Register manager force update internal gateway map.")
			refreshForce = true
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
	refreshForce := false
	for {
		providers := mgr.pullProvidersFromRegisterSrv()
    // Check for update
    for _, provider := range providers {
      mgr.registeredProvidersMapLock.RLock()
      storedInfo, ok := mgr.registeredProvidersMap[provider.GetNodeID()]
      mgr.registeredProvidersMapLock.RUnlock()
      if !ok {
        // Not exist, we need to add a new entry
        mgr.registeredProvidersMapLock.Lock()
        mgr.registeredProvidersMap[provider.GetNodeID()] = provider
        mgr.registeredProvidersMapLock.Unlock()
      } else {
        // Exist, check if need update
        if provider != storedInfo {
          // Need update
          mgr.registeredProvidersMapLock.Lock()
          mgr.registeredProvidersMap[provider.GetNodeID()] = provider
          mgr.registeredProvidersMapLock.Unlock()
        }
      }
    }

		if refreshForce {
			mgr.providerRefreshCh <- true
			refreshForce = false
		}
		afterChan := time.After(mgr.refreshDuration)
		select {
		case <-mgr.providerRefreshCh:
			// Need to refresh
			logging.Info("Register manager force update internal provider map.")
			refreshForce = true
		case <-afterChan:
			// Need to refresh
		case <-mgr.providerShutdownCh:
			// Need to shutdown
			logging.Info("Register manager shutdown provider routine.")
			return
		}
	}
}
