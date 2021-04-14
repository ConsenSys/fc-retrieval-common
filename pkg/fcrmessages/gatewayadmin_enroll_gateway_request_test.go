package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeGatewayAdminEnrollGatewayRequest success test
func TestEncodeGatewayAdminEnrollGatewayRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockAddress := "address"
	mockRootSigningKey := "root signing key"
	mockSigningKey := "signing key"
	mockRegionCode := "region code"
	mockNetworkInfoGateway := "network info gateway"
  mockNetworkInfoProvider := "network info provider"
	mockNetworkInfoClient := "network info client"
	mockNetworkInfoAdmin := "network info admin"

	validMsg := &FCRMessage{
		messageType:       406,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"node_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","address":"address","root_signing_key":"root signing key","signing_key":"signing key","region_code":"region code","network_info_gateway":"network info gateway","network_info_provider":"network info provider","network_info_client":"network info client","network_info_admin":"network info admin"}`),
	}

	msg, err := EncodeGatewayAdminEnrollGatewayRequest(mockNodeID, mockAddress, mockRootSigningKey, mockSigningKey, mockRegionCode, mockNetworkInfoGateway, mockNetworkInfoProvider, mockNetworkInfoClient, mockNetworkInfoAdmin)
	assert.Empty(t, err)
	assert.Equal(t, validMsg, msg)
}

// TestDecodeGatewayAdminEnrollGatewayRequest success test
func TestDecodeGatewayAdminEnrollGatewayRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockAddress := "address"
	mockRootSigningKey := "root signing key"
	mockSigningKey := "signing key"
	mockRegionCode := "region code"
	mockNetworkInfoGateway := "network info gateway"
  mockNetworkInfoProvider := "network info provider"
	mockNetworkInfoClient := "network info client"
	mockNetworkInfoAdmin := "network info admin"

	validMsg := &FCRMessage{
		messageType:       406,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"node_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","address":"address","root_signing_key":"root signing key","signing_key":"signing key","region_code":"region code","network_info_gateway":"network info gateway","network_info_provider":"network info provider","network_info_client":"network info client","network_info_admin":"network info admin"}`),
	}

	nodeID, address, rootSigningKey, signingKey, regionCode, networkInfoGateway, networkInfoProvider, networkInfoClient,
		networkInfoAdmin, err := DecodeGatewayAdminEnrollGatewayRequest(validMsg)

	assert.Empty(t, err)
	assert.Equal(t, mockNodeID, nodeID)
	assert.Equal(t, mockAddress, address)
	assert.Equal(t, mockRootSigningKey, rootSigningKey)
	assert.Equal(t, mockSigningKey, signingKey)
	assert.Equal(t, mockRegionCode, regionCode)
	assert.Equal(t, mockNetworkInfoGateway, networkInfoGateway)
	assert.Equal(t, mockNetworkInfoProvider, networkInfoProvider)
	assert.Equal(t, mockNetworkInfoClient, networkInfoClient)
	assert.Equal(t, mockNetworkInfoAdmin, networkInfoAdmin)
}
