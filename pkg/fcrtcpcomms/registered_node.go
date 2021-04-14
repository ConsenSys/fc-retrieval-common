package fcrtcpcomms

import "github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"

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
