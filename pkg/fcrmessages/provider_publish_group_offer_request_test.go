package fcrmessages

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
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
		messageType:300,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","nonce":42,"offer":{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","cids":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE="],"price":41,"expiry":42,"qos":43,"signature":""}}`), 
		signature:"",
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
		ProviderID: *mockProviderID,
		Nonce:     mockNonce,
		Offer:     *mockOffer,
	})

	validMsg := &FCRMessage{
		messageType:300,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:mockOfferRequest, 
		signature:"",
	}
	
	nodeID, nonce, CIDOffer, err := DecodeProviderPublishGroupOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockProviderID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, CIDOffer.GetMessageDigest(), mockOffer.GetMessageDigest())
}