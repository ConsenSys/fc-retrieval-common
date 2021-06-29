package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminSetReputationResponse success test
func TestEncodeGatewayAdminSetReputationResponse(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockReputation := int64(42)
	mockExists := true

	validMsg := &FCRMessage{
		messageType:       405,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"client_id":"0000000000000000000000000000000000000000000000000000000000000042","reputation":42,"exists":true}`),
		signature:         "",
	}

	msg, err := EncodeGatewayAdminSetReputationResponse(mockNodeID, mockReputation, mockExists)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayAdminSetReputationResponse success test
func TestDecodeGatewayAdminSetReputationResponse(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockReputation := int64(42)
	mockExists := true

	validMsg := &FCRMessage{
		messageType:       405,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"client_id":"0000000000000000000000000000000000000000000000000000000000000042","reputation":42,"exists":true}`),
		signature:         "",
	}

	nodeID, reputation, exists, err := DecodeGatewayAdminSetReputationResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, reputation, mockReputation)
	assert.Equal(t, exists, mockExists)
}
