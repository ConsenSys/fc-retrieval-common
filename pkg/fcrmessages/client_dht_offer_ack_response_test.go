package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeClientDHTOfferAckResponse success test
func TestEncodeClientDHTOfferAckResponse(t *testing.T) {
	mockPieceID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockFound := true
	mockPublishDHTOfferRequest := &FCRMessage{
		messageType:       106,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="}`),
		signature:         "",
	}
	mockPublishDHTOfferResponse := &FCRMessage{
		messageType:       106,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="}`),
		signature:         "",
	}
	validMsg := &FCRMessage{
		messageType:       107,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","found":true,"publish_dht_offer_request":{"message_type":106,"protocol_version":1,"protocol_supported":[1,1],"message_body":"eyJwaWVjZV9jaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFFPSIsImdhdGV3YXlfaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUVJPSJ9","message_signature":""},"publish_dht_offer_response":{"message_type":106,"protocol_version":1,"protocol_supported":[1,1],"message_body":"eyJwaWVjZV9jaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFFPSIsImdhdGV3YXlfaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUVJPSJ9","message_signature":""}}`),
		signature:         "",
	}

	msg, err := EncodeClientDHTOfferAckResponse(
		mockPieceID,
		mockNodeID,
		mockFound,
		mockPublishDHTOfferRequest,
		mockPublishDHTOfferResponse,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientDHTOfferAckResponse success test
func TestDecodeClientDHTOfferAckResponse(t *testing.T) {
	mockPieceID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockFound := true
	mockPublishDHTOfferRequest := &FCRMessage{
		messageType:       106,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="}`),
		signature:         "",
	}
	mockPublishDHTOfferResponse := &FCRMessage{
		messageType:       106,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI="}`),
		signature:         "",
	}
	validMsg := &FCRMessage{
		messageType:       107,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","found":true,"publish_dht_offer_request":{"message_type":106,"protocol_version":1,"protocol_supported":[1,1],"message_body":"eyJwaWVjZV9jaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFFPSIsImdhdGV3YXlfaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUVJPSJ9","message_signature":""},"publish_dht_offer_response":{"message_type":106,"protocol_version":1,"protocol_supported":[1,1],"message_body":"eyJwaWVjZV9jaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFFPSIsImdhdGV3YXlfaWQiOiJBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUVJPSJ9","message_signature":""}}`),
		signature:         "",
	}

	pieceCID, gatewayID, found, publishDHTOfferRequest, publishDHTOfferResponse, err := DecodeClientDHTOfferAckResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, pieceCID, mockPieceID)
	assert.Equal(t, gatewayID, mockNodeID)
	assert.Equal(t, found, mockFound)
	assert.Equal(t, publishDHTOfferRequest, mockPublishDHTOfferRequest)
	assert.Equal(t, publishDHTOfferResponse, mockPublishDHTOfferResponse)
}
