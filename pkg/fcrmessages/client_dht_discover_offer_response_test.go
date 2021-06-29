package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeClientDHTDiscoverOfferResponse success test
func TestEncodeClientDHTDiscoverOfferResponse(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")

	mockMessage := &FCRMessage{
		messageType:       115,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(``),
		signature:         "",
	}
	mockResponse := []FCRMessage{*mockMessage}

	validMsg := &FCRMessage{
		messageType:       115,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"gateway_ids":["0000000000000000000000000000000000000000000000000000000000000042"],"response":[{"message_type":115,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}],"payment_required":true,"payment_channel":42}`),
		signature:         "",
	}
	fakePaymentRequired := true
	fakePaymentChannel := int64(42)

	msg, err := EncodeClientDHTDiscoverOfferResponse(
		mockContentID,
		mockNonce,
		[]nodeid.NodeID{*mockNodeID},
		mockResponse,
		fakePaymentRequired,
		fakePaymentChannel,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientDHTDiscoverOfferResponse success test
func TestDecodeClientDHTDiscoverOfferResponse(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")

	mockMessage := &FCRMessage{
		messageType:       115,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(``),
		signature:         "",
	}
	mockResponse := []FCRMessage{*mockMessage}

	validMsg := &FCRMessage{
		messageType:       115,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"gateway_ids":["0000000000000000000000000000000000000000000000000000000000000042"],"response":[{"message_type":115,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}],"payment_required":true,"payment_channel":42}`),
		signature:         "",
	}
	fakePaymentRequired := true
	fakePaymentChannel := int64(42)

	contentID, nonce, gateway_ids, response, paymentRequired, paymentChannel, err := DecodeClientDHTDiscoverOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, contentID, mockContentID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, gateway_ids, []nodeid.NodeID{*mockNodeID})
	assert.Equal(t, response, mockResponse)
	assert.Equal(t, fakePaymentRequired, paymentRequired)
	assert.Equal(t, fakePaymentChannel, paymentChannel)
}
