package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeGatewayAdminSetReputationRequest success test
func TestEncodeGatewayAdminSetReputationRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockReputation := int64(42)
	validMsg := &FCRMessage{
		messageType:404,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"client_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","reputation":42}`), 
		signature:"",
	}

	msg, err := EncodeGatewayAdminSetReputationRequest(mockNodeID, mockReputation)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayAdminSetReputationRequest success test
func TestDecodeGatewayAdminSetReputationRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockReputation := int64(42)
	validMsg := &FCRMessage{
		messageType:404,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"client_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","reputation":42}`), 
		signature:"",
	}

	nodeID, reputation, err := DecodeGatewayAdminSetReputationRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, reputation, mockReputation)
}