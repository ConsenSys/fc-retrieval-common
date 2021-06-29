package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeClientDHTOfferAckRequest success test
func TestEncodeClientDHTOfferAckRequest(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	validMsg := &FCRMessage{
		messageType:       106,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","gateway_id":"0000000000000000000000000000000000000000000000000000000000000042"}`),
		signature:         "",
	}

	msg, err := EncodeClientDHTOfferAckRequest(
		mockContentID,
		mockNodeID,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientDHTOfferAckRequest success test
func TestDecodeClientDHTOfferAckRequest(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	validMsg := &FCRMessage{
		messageType:       106,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","gateway_id":"0000000000000000000000000000000000000000000000000000000000000042"}`),
		signature:         "",
	}

	contentID, nodeID, err := DecodeClientDHTOfferAckRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, contentID, mockContentID)
	assert.Equal(t, nodeID, mockNodeID)
}
