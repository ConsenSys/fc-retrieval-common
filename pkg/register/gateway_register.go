package register

import (
  "github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
  "github.com/ConsenSys/fc-retrieval-common/pkg/request"
)

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
  httpCommunicator    request.HttpCommunications
}


// GetNodeID gets the node id
func (r *GatewayRegister) GetNodeID() string {
  return r.NodeID
}

// GetAddress gets the address
func (r *GatewayRegister) GetAddress() string {
  return r.Address
}

// GetRegionCode gets the region code
func (r *GatewayRegister) GetRegionCode() string {
  return r.RegionCode
}

// GetNetworkInfoGateway gets the network gateway ap
func (r *GatewayRegister) GetNetworkInfoGateway() string {
  return r.NetworkInfoGateway
}

// GetNetworkInfoClient gets the network client ap
func (r *GatewayRegister) GetNetworkInfoClient() string {
  return r.NetworkInfoClient
}

// GetNetworkInfoAdmin gets the network admin ap
func (r *GatewayRegister) GetNetworkInfoAdmin() string {
  return r.NetworkInfoAdmin
}

// GetRootSigningKey gets the root signing key
func (r *GatewayRegister) GetRootSigningKey() (*fcrcrypto.KeyPair, error) {
  return fcrcrypto.DecodePublicKey(r.RootSigningKey)
}

// GetSigningKey gets the signing key
func (r *GatewayRegister) GetSigningKey() (*fcrcrypto.KeyPair, error) {
  return fcrcrypto.DecodePublicKey(r.SigningKey)
}

// RegisterGateway to register a gateway
func (r *GatewayRegister) RegisterGateway(registerURL string) error {
  url := registerURL + "/registers/gateway"
  return r.httpCommunicator.SendJSON(url, r)
}
