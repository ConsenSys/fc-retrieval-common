package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeClientDHTDiscoverResponse success test
func TestEncodeClientDHTDiscoverResponseV2(t *testing.T) {
	mockContactedMsg := &FCRMessage{
		messageType:       113,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(``),
		signature:         "",
	}
	mockContactedMsgs := make([]FCRMessage, 0)
	mockContactedMsgs = append(mockContactedMsgs, *mockContactedMsg)

	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeIDs := make([]nodeid.NodeID, 0)
	mockNodeIDs = append(mockNodeIDs, *mockNodeID)

	mockNonce := int64(42)
	fakePaymentRequired := true
	fakePaymentChannel := int64(43)

	validMsg := &FCRMessage{
		messageType:       113,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"contacted_gateways":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="],"response":[{"message_type":113,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}],"uncontactable_gateways":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="],"nonce":42,"payment_required":true,"payment_channel":43}`),
		signature:         "",
	}

	msg, err := EncodeClientDHTDiscoverResponseV2(
		mockNodeIDs,
		mockContactedMsgs,
		mockNodeIDs,
		mockNonce,
		fakePaymentRequired,
		fakePaymentChannel,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientDHTDiscoverResponse success test
func TestDecodeClientDHTDiscoverResponseV2(t *testing.T) {

	mockContactedMsg := &FCRMessage{
		messageType:       113,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(``),
		signature:         "",
	}
	mockContactedMsgs := make([]FCRMessage, 0)
	mockContactedMsgs = append(mockContactedMsgs, *mockContactedMsg)

	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeIDs := make([]nodeid.NodeID, 0)
	mockNodeIDs = append(mockNodeIDs, *mockNodeID)

	mockNonce := int64(42)
	fakePaymentRequired := true
	fakePaymentChannel := int64(43)

	validMsg := &FCRMessage{
		messageType:       113,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"contacted_gateways":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="],"response":[{"message_type":113,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}],"uncontactable_gateways":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="],"nonce":42,"payment_required":true,"payment_channel":43}`),
		signature:         "",
	}

	contacted, contactedMsg, nodeIDs, nonce, paymentRequired, paymentChannel, err := DecodeClientDHTDiscoverResponseV2(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, contacted, mockNodeIDs)
	assert.Equal(t, contactedMsg, mockContactedMsgs)
	assert.Equal(t, nodeIDs, mockNodeIDs)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, fakePaymentRequired, paymentRequired)
	assert.Equal(t, fakePaymentChannel, paymentChannel)
}
