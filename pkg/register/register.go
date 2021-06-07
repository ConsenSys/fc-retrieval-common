/*
Package register - provides methods for FileCoin Secondary Retrieval from Retrieval Register perspective.

Retrieval Register is a central node, holding information about Retrieval Gateways and Retrieval Providers.
*/
package register

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
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/request"
)

// RegisteredNode stored network information of a registered node
type RegisteredNode interface {
	GetNodeID() string
	GetAddress() string
	GetRegionCode() string
	GetRootSigningKey() (*fcrcrypto.KeyPair, error)
	GetSigningKey() (*fcrcrypto.KeyPair, error)
	GetNetworkInfoGateway() string
	GetNetworkInfoProvider() string
	GetNetworkInfoClient() string
	GetNetworkInfoAdmin() string
}

// GatewayRegister stores information of a registered gateway
type GatewayRegister struct {
	NodeID              string `json:"nodeId"`
	Address             string `json:"address"`
	RootSigningKey      string `json:"rootSigningKey"`
	SigningKey          string `json:"sigingKey"`
	RegionCode          string `json:"regionCode"`
	NetworkInfoGateway  string `json:"networkInfoGateway"`
	NetworkInfoProvider string `json:"networkInfoProvider"`
	NetworkInfoClient   string `json:"networkInfoClient"`
	NetworkInfoAdmin    string `json:"networkInfoAdmin"`
}

// ProviderRegister stores information of a registered provider
type ProviderRegister struct {
	NodeID             string `json:"nodeId"`
	Address            string `json:"address"`
	RootSigningKey     string `json:"rootSigningKey"`
	SigningKey         string `json:"sigingKey"`
	RegionCode         string `json:"regionCode"`
	NetworkInfoGateway string `json:"networkInfoGateway"`
	NetworkInfoClient  string `json:"networkInfoClient"`
	NetworkInfoAdmin   string `json:"networkInfoAdmin"`
}

// GetNodeID gets the node id
func (r *GatewayRegister) GetNodeID() string {
	return r.NodeID
}

// GetNodeID gets the node id
func (r *ProviderRegister) GetNodeID() string {
	return r.NodeID
}

// GetAddress gets the address
func (r *GatewayRegister) GetAddress() string {
	return r.Address
}

// GetAddress gets the node id
func (r *ProviderRegister) GetAddress() string {
	return r.Address
}

// GetRegionCode gets the region code
func (r *GatewayRegister) GetRegionCode() string {
	return r.RegionCode
}

// GetRegionCode gets the region code
func (r *ProviderRegister) GetRegionCode() string {
	return r.RegionCode
}

// GetNetworkInfoGateway gets the network gateway ap
func (r *GatewayRegister) GetNetworkInfoGateway() string {
	return r.NetworkInfoGateway
}

// GetNetworkInfoGateway gets the network gateway ap
func (r *ProviderRegister) GetNetworkInfoGateway() string {
	return r.NetworkInfoGateway
}

// GetNetworkInfoProvider gets the network provider ap
func (r *GatewayRegister) GetNetworkInfoProvider() string {
	return r.NetworkInfoProvider
}

// GetNetworkInfoProvider gets the network provider ap
func (r *ProviderRegister) GetNetworkInfoProvider() string {
	return ""
}

// GetNetworkInfoClient gets the network client ap
func (r *GatewayRegister) GetNetworkInfoClient() string {
	return r.NetworkInfoClient
}

// GetNetworkInfoClient gets the network client ap
func (r *ProviderRegister) GetNetworkInfoClient() string {
	return r.NetworkInfoClient
}

// GetNetworkInfoAdmin gets the network admin ap
func (r *GatewayRegister) GetNetworkInfoAdmin() string {
	return r.NetworkInfoAdmin
}

// GetNetworkInfoAdmin gets the network admin ap
func (r *ProviderRegister) GetNetworkInfoAdmin() string {
	return r.NetworkInfoAdmin
}

// GetRootSigningKey gets the root signing key
func (r *GatewayRegister) GetRootSigningKey() (*fcrcrypto.KeyPair, error) {
	return fcrcrypto.DecodePublicKey(r.RootSigningKey)
}

// GetRootSigningKey gets the root signing key
func (r *ProviderRegister) GetRootSigningKey() (*fcrcrypto.KeyPair, error) {
	return fcrcrypto.DecodePublicKey(r.RootSigningKey)
}

// GetSigningKey gets the signing key
func (r *GatewayRegister) GetSigningKey() (*fcrcrypto.KeyPair, error) {
	return fcrcrypto.DecodePublicKey(r.SigningKey)
}

// GetSigningKey gets the signing key
func (r *ProviderRegister) GetSigningKey() (*fcrcrypto.KeyPair, error) {
	return fcrcrypto.DecodePublicKey(r.SigningKey)
}

// RegisterGateway to register a gateway
func (r *GatewayRegister) RegisterGateway(registerURL string) error {
	url := registerURL + "/registers/gateway"
	return request.SendJSON(url, r)
}

// RegisterProvider to register a provider
func (r *ProviderRegister) RegisterProvider(registerURL string) error {
	url := registerURL + "/registers/provider"
	return request.SendJSON(url, r)
}

// GetRegisteredGateways returns registered gateways
func GetRegisteredGateways(registerURL string) ([]GatewayRegister, error) {
	url := registerURL + "/registers/gateway"
	gateways := []GatewayRegister{}
	err := request.GetJSON(url, &gateways)
	if err != nil {
		return gateways, err
	}
	return gateways, nil
}

// GetRegisteredProviders returns registered providers
func GetRegisteredProviders(registerURL string) ([]ProviderRegister, error) {
	url := registerURL + "/registers/provider"
	providers := []ProviderRegister{}
	err := request.GetJSON(url, &providers)
	if err != nil {
		return providers, err
	}
	return providers, nil
}

// GetGatewayByID gets the gateway register info by a given ID
func GetGatewayByID(registerURL string, nodeID *nodeid.NodeID) (GatewayRegister, error) {
	url := registerURL + "/registers/gateway/" + nodeID.ToString()
	gateway := GatewayRegister{}
	err := request.GetJSON(url, &gateway)
	if err != nil {
		return gateway, err
	}
	return gateway, nil
}

// GetProviderByID gets the provider register info by a given ID
func GetProviderByID(registerURL string, nodeID *nodeid.NodeID) (ProviderRegister, error) {
	url := registerURL + "/registers/provider/" + nodeID.ToString()
	provider := ProviderRegister{}
	err := request.GetJSON(url, &provider)
	if err != nil {
		return provider, err
	}
	return provider, nil
}
