package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayListDHTOfferResponse success test
func TestEncodeGatewayListDHTOfferResponse(t *testing.T) {
	mockDHTOffer := &FCRMessage{
		messageType:402,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(``), 
		signature:"",
	}
	 mockDHTOffers := []FCRMessage{*mockDHTOffer}
	 validMsg := &FCRMessage{
		messageType:201,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"published_dht_offers":[{"message_type":402,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}]}`), 
		signature:"",
	}

	msg, err := EncodeGatewayListDHTOfferResponse(mockDHTOffers)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayListDHTOfferResponse success test
func TestDecodeGatewayListDHTOfferResponse(t *testing.T) {
	mockDHTOffer := &FCRMessage{
		messageType:402,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(``), 
		signature:"",
	}
	 mockDHTOffers := []FCRMessage{*mockDHTOffer}
	 validMsg := &FCRMessage{
		messageType:201,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"published_dht_offers":[{"message_type":402,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}]}`), 
		signature:"",
	}

	DHTOffers, err := DecodeGatewayListDHTOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, DHTOffers, mockDHTOffers)
}