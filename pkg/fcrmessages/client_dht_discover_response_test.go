package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeClientDHTDiscoverResponse success test
func TestEncodeClientDHTDiscoverResponse(t *testing.T) {
	mockContactedMsg := &FCRMessage{
		messageType:105,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(``), 
		signature:"",
	}
	mockContactedMsgs := make([]FCRMessage, 0)
	mockContactedMsgs = append(mockContactedMsgs, *mockContactedMsg)

	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeIDs := make([]nodeid.NodeID, 0)
	mockNodeIDs = append(mockNodeIDs, *mockNodeID)

	mockNonce := int64(42)

	validMsg := &FCRMessage{
		messageType:105,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"contacted_gateways":[{"message_type":105,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}],"uncontactable_gateways":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="],"nonce":42}`), 
		signature:"",
	}

	msg, err := EncodeClientDHTDiscoverResponse(
		mockContactedMsgs,
		mockNodeIDs,
		mockNonce,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientDHTDiscoverResponse success test
func TestDecodeClientDHTDiscoverResponse(t *testing.T) {
	mockContactedMsg := &FCRMessage{
		messageType:105,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(``), 
		signature:"",
	}
	mockContactedMsgs := make([]FCRMessage, 0)
	mockContactedMsgs = append(mockContactedMsgs, *mockContactedMsg)

	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeIDs := make([]nodeid.NodeID, 0)
	mockNodeIDs = append(mockNodeIDs, *mockNodeID)

	mockNonce := int64(42)

	validMsg := &FCRMessage{
		messageType:105,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"contacted_gateways":[{"message_type":105,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}],"uncontactable_gateways":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="],"nonce":42}`), 
		signature:"",
	}

	contactedMsg, nodeIDs, nonce, err := DecodeClientDHTDiscoverResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, contactedMsg, mockContactedMsgs)
	assert.Equal(t, nodeIDs, mockNodeIDs)
	assert.Equal(t, nonce, mockNonce)
}
