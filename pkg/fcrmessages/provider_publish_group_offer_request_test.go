package fcrmessages

import (
	"encoding/json"
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderPublishGroupOfferRequest success test
func TestEncodeProviderPublishGroupOfferRequest(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)

	validMsg := &FCRMessage{
		messageType:       300,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","nonce":42,"offer":{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":41,"expiry":42,"qos":43,"signature":""}}`),
		signature:         "",
	}

	msg, err := EncodeProviderPublishGroupOfferRequest(mockNodeID, mockNonce, mockOffer)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderPublishGroupOfferRequest success test
func TestDecodeProviderPublishGroupOfferRequest(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)

	mockOfferRequest, _ := json.Marshal(providerPublishGroupOfferRequest{
		ProviderID: mockProviderID.ToString(),
		Nonce:      mockNonce,
		Offer:      *mockOffer,
	})

	validMsg := &FCRMessage{
		messageType:       300,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       mockOfferRequest,
		signature:         "",
	}

	nodeID, nonce, CIDOffer, err := DecodeProviderPublishGroupOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockProviderID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, CIDOffer.GetMessageDigest(), mockOffer.GetMessageDigest())
}
