package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminGetReputationRequest success test
func TestEncodeGatewayAdminGetReputationRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	validMsg := &FCRMessage{
		messageType:       402,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"client_id":"0000000000000000000000000000000000000000000000000000000000000042"}`),
		signature:         "",
	}

	msg, err := EncodeGatewayAdminGetReputationRequest(mockNodeID)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayAdminGetReputationRequest success test
func TestDecodeGatewayAdminGetReputationRequest(t *testing.T) {

	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	validMsg := &FCRMessage{
		messageType:       402,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"client_id":"0000000000000000000000000000000000000000000000000000000000000042"}`),
		signature:         "",
	}

	nodeID, err := DecodeGatewayAdminGetReputationRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
}
