package fcrmessages

import (
	"testing"
	// "github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayListDHTOfferAck success test
func TestEncodeGatewayListDHTOfferAck(t *testing.T) {
	mockDHTOffer := &FCRMessage{
		messageType:202,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(``), 
		signature:"",
	}
	mockDHTOffers := []FCRMessage{*mockDHTOffer}
	validMsg := &FCRMessage{
		messageType:202,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"published_dht_offers_ack":[{"message_type":202,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}]}`), 
		signature:"",
	}

	msg, err := EncodeGatewayListDHTOfferAck(mockDHTOffers)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayListDHTOfferAck success test
func TestDecodeGatewayListDHTOfferAck(t *testing.T) {
	mockDHTOffer := &FCRMessage{
		messageType:202,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(``), 
		signature:"",
	}
	mockDHTOffers := []FCRMessage{*mockDHTOffer}
	validMsg := &FCRMessage{
		messageType:202,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"published_dht_offers_ack":[{"message_type":202,"protocol_version":1,"protocol_supported":[1,1],"message_body":"","message_signature":""}]}`), 
		signature:"",
	}

	DHTOffers, err := DecodeGatewayListDHTOfferAck(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, DHTOffers, mockDHTOffers)
}