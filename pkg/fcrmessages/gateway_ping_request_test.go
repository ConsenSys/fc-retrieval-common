package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayDHTDiscoverRequest success test
func TestEncodeGatewayPingRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	mockTTL := int64(43)

	validMsg := &FCRMessage{
		messageType:       205,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","nonce":42,"ttl":43}`),
		signature:         "",
	}

	msg, err := EncodeGatewayPingRequest(mockNodeID, mockNonce, mockTTL)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayPingRequest success test
func TestDecodeGatewayPingRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	mockTTL := int64(43)

	validMsg := &FCRMessage{
		messageType:       205,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","nonce":42,"ttl":43}`),
		signature:         "",
	}

	nodeID, nonce, TTL, err := DecodeGatewayPingRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, TTL, mockTTL)
}

// TestDecodeGatewayPingRequest failure test
func TestDecodeGatewayPingRequestWrongMessageType(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       205000,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"dummy":43}`),
		signature:         "",
	}

	nodeID, nonce, TTL, err := DecodeGatewayPingRequest(validMsg)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "message type mismatch")
	assert.Empty(t, nodeID)
	assert.Empty(t, nonce)
	assert.Empty(t, TTL)
}

// TestDecodeGatewayPingRequest failure test
func TestDecodeGatewayPingRequestUnmarshalError(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       205,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`invalid_message`),
		signature:         "",
	}

	nodeID, nonce, TTL, err := DecodeGatewayPingRequest(validMsg)
	assert.NotNil(t, err)
	assert.NotEmpty(t, err.Error())
	assert.Empty(t, nodeID)
	assert.Empty(t, nonce)
	assert.Empty(t, TTL)
}
