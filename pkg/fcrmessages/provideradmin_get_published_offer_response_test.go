package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
)

// TestEncodeProviderAdminGetPublishedOfferResponse success test
func TestEncodeProviderAdminGetPublishedOfferResponse(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockCidOffers := make([]cidoffer.CIDOffer, 0)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43

	cidOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)
	mockCidOffers = append(mockCidOffers, *cidOffer)
	validMsg := &FCRMessage{
		messageType:507,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"exists":true,"cid_offers":[{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","cids":["AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE="],"price":41,"expiry":42,"qos":43,"signature":""}]}`), 
		signature:"",
	}
	msg, err := EncodeProviderAdminGetPublishedOfferResponse(true, mockCidOffers)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderAdminGetPublishedOfferResponse success test
func TestDecodeProviderAdminGetPublishedOfferResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:507,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{}`), 
		signature:"",
	}
	exists, cidoffers, err := DecodeProviderAdminGetPublishedOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, exists, false)
	assert.Equal(t, cidoffers, []cidoffer.CIDOffer([]cidoffer.CIDOffer(nil)))
}