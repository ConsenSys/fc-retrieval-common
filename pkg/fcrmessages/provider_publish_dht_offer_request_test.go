package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderPublishDHTOfferRequest success test
func TestEncodeProviderPublishDHTOfferRequest(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)
	mockCidOffers := make([]cidoffer.CIDOffer, 0)
	mockCidOffers = append(mockCidOffers, *mockOffer)

	validMsg := &FCRMessage{
		messageType:       302,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","nonce":42,"num_of_offers":1,"single_offers":[{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":41,"expiry":42,"qos":43,"signature":""}]}`),
		signature:         "",
	}

	msg, err := EncodeProviderPublishDHTOfferRequest(mockNodeID, mockNonce, mockCidOffers)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderPublishDHTOfferRequest success test
func TestDecodeProviderPublishDHTOfferRequest(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)
	mockCidOffers := make([]cidoffer.CIDOffer, 0)
	mockCidOffers = append(mockCidOffers, *mockOffer)

	validMsg := &FCRMessage{
		messageType:       302,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","nonce":42,"num_of_offers":1,"single_offers":[{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","cids":["0000000000000000000000000000000000000000000000000000000000000001"],"price":41,"expiry":42,"qos":43,"signature":""}]}`),
		signature:         "",
	}

	nodeID, nonce, CIDOffers, err := DecodeProviderPublishDHTOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, nonce, mockNonce)
	for i, CIDoffer := range CIDOffers {
		assert.Equal(t, CIDoffer.GetMessageDigest(), mockCidOffers[i].GetMessageDigest())
	}

}
