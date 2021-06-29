package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderAdminPublishDHTOfferRequest success test
func TestEncodeProviderAdminPublishDHTOfferRequest(t *testing.T) {
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	mockPrices := []uint64{mockPrice}
	var mockExpiry int64 = 42
	mockExpiries := []int64{mockExpiry}
	var mockQos uint64 = 43
	mockQoss := []uint64{mockQos}

	validMsg := &FCRMessage{
		messageType:       504,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":[41],"expiry":[42],"qos":[43]}`),
		signature:         "",
	}
	msg, err := EncodeProviderAdminPublishDHTOfferRequest(mockCids, mockPrices, mockExpiries, mockQoss)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderAdminPublishDHTOfferRequest success test
func TestDecodeProviderAdminPublishDHTOfferRequest(t *testing.T) {
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}

	validMsg := &FCRMessage{
		messageType:       504,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":[41],"expiry":[42],"qos":[43]}`),
		signature:         "",
	}
	cids, price, expiry, qos, err := DecodeProviderAdminPublishDHTOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, cids, mockCids)
	assert.Equal(t, price, []uint64{41})
	assert.Equal(t, expiry, []int64{42})
	assert.Equal(t, qos, []uint64{43})
}
