package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeGatewayAdminEnrollProviderRequest success test
func TestEncodeGatewayAdminEnrollProviderRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockAddress := "address"
	mockRootSigningKey := "root signing key"
	mockSigningKey := "signing key"
	mockRegionCode := "region code"
	mockNetworkInfoGateway := "network info gateway"
	mockNetworkInfoClient := "network info client"
	mockNetworkInfoAdmin := "network info admin"

	validMsg := &FCRMessage{
		messageType:       508,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"node_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","address":"address","root_signing_key":"root signing key","signing_key":"signing key","region_code":"region code","network_info_gateway":"network info gateway","network_info_client":"network info client","network_info_admin":"network info admin"}`),
	}

	msg, err := EncodeGatewayAdminEnrollProviderRequest(mockNodeID, mockAddress, mockRootSigningKey, mockSigningKey, mockRegionCode, mockNetworkInfoGateway, mockNetworkInfoClient, mockNetworkInfoAdmin)
	assert.Empty(t, err)
	assert.Equal(t, validMsg, msg)
}

// TestDecodeGatewayAdminEnrollProviderRequest success test
func TestDecodeGatewayAdminEnrollProviderRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockAddress := "address"
	mockRootSigningKey := "root signing key"
	mockSigningKey := "signing key"
	mockRegionCode := "region code"
	mockNetworkInfoGateway := "network info gateway"
	mockNetworkInfoClient := "network info client"
	mockNetworkInfoAdmin := "network info admin"

	validMsg := &FCRMessage{
		messageType:       508,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"node_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","address":"address","root_signing_key":"root signing key","signing_key":"signing key","region_code":"region code","network_info_gateway":"network info gateway","network_info_client":"network info client","network_info_admin":"network info admin"}`),
	}

	nodeID, address, rootSigningKey, signingKey, regionCode, networkInfoGateway, networkInfoClient,
		networkInfoAdmin, err := DecodeGatewayAdminEnrollProviderRequest(validMsg)

	assert.Empty(t, err)
	assert.Equal(t, mockNodeID, nodeID)
	assert.Equal(t, mockAddress, address)
	assert.Equal(t, mockRootSigningKey, rootSigningKey)
	assert.Equal(t, mockSigningKey, signingKey)
	assert.Equal(t, mockRegionCode, regionCode)
	assert.Equal(t, mockNetworkInfoGateway, networkInfoGateway)
	assert.Equal(t, mockNetworkInfoClient, networkInfoClient)
	assert.Equal(t, mockNetworkInfoAdmin, networkInfoAdmin)
}
