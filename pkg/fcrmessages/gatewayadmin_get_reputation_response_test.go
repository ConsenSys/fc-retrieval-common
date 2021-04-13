package fcrmessages

import (
	"testing"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminGetReputationResponse success test
func TestEncodeGatewayAdminGetReputationResponse(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockReputation := int64(42)
	mockExists := true

	validMsg := &FCRMessage{
		messageType:403,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"client_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","reputation":42,"exists":true}`), 
		signature:"",
	}

	msg, err := EncodeGatewayAdminGetReputationResponse(mockNodeID, mockReputation, mockExists)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayAdminGetReputationResponse success test
func TestDecodeGatewayAdminGetReputationResponse(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockReputation := int64(42)
	mockExists := true
	validMsg := &FCRMessage{
		messageType:403,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"client_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","reputation":42,"exists":true}`), 
		signature:"",
	}

	nodeID, reputation, exists, err := DecodeGatewayAdminGetReputationResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, reputation, mockReputation)
	assert.Equal(t, exists, mockExists)
}