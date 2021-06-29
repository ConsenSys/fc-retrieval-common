package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderAdminPublishGroupOfferRequest success test
func TestEncodeProviderAdminPublishGroupOfferRequest(t *testing.T) {
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43

	validMsg := &FCRMessage{
		messageType:       502,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":41,"expiry":42,"qos":43}`),
		signature:         "",
	}

	msg, err := EncodeProviderAdminPublishGroupOfferRequest(mockCids, mockPrice, mockExpiry, mockQos)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderAdminPublishGroupOfferRequest success test
func TestDecodeProviderAdminPublishGroupOfferRequest(t *testing.T) {
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	validMsg := &FCRMessage{
		messageType:       502,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":41,"expiry":42,"qos":43}`),
		signature:         "",
	}
	cids, price, expiry, qos, err := DecodeProviderAdminPublishGroupOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, cids, mockCids)
	assert.Equal(t, price, uint64(41))
	assert.Equal(t, expiry, int64(42))
	assert.Equal(t, qos, uint64(43))
}
