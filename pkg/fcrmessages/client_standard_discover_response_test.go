package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeClientStandardDiscoverResponse success test
func TestEncodeClientStandardDiscoverResponse(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	mockFound := true
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*mockContentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	offer, _ := cidoffer.NewCIDOffer(mockNodeID, mockCids, mockPrice, mockExpiry, mockQos)
	subOffer, _ := offer.GenerateSubCIDOffer(mockContentID)
	mockSubCidOffers := []cidoffer.SubCIDOffer{*subOffer}
	mockFPCs := []bool{true}
	validMsg := &FCRMessage{
		messageType:       103,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"found":true,"sub_cid_offers":[{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","sub_cid":"0000000000000000000000000000000000000000000000000000000000000001","merkle_root":"c3c3a46684c07d12a9c238787df3049a6f258e7af203e5ddb66a8bd66637e108","merkle_proof":"AAAAMFsiQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBRT0iXQAAAANbMV0=","price":41,"expiry":42,"qos":43,"signature":""}],"funded_payment_channel":[true]}`),
		signature:         "",
	}

	msg, err := EncodeClientStandardDiscoverResponse(
		mockContentID,
		mockNonce,
		mockFound,
		mockSubCidOffers,
		mockFPCs,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientStandardDiscoverResponse success test
func TestDecodeClientStandardDiscoverResponse(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	mockFound := true
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*mockContentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	offer, _ := cidoffer.NewCIDOffer(mockNodeID, mockCids, mockPrice, mockExpiry, mockQos)
	subOffer, _ := offer.GenerateSubCIDOffer(mockContentID)
	mockSubCidOffers := []cidoffer.SubCIDOffer{*subOffer}
	mockFPCs := []bool{true}
	validMsg := &FCRMessage{
		messageType:       103,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"found":true,"sub_cid_offers":[{"provider_id":"0000000000000000000000000000000000000000000000000000000000000042","sub_cid":"0000000000000000000000000000000000000000000000000000000000000001","merkle_root":"c3c3a46684c07d12a9c238787df3049a6f258e7af203e5ddb66a8bd66637e108","merkle_proof":"AAAAMFsiQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBRT0iXQAAAANbMV0=","price":41,"expiry":42,"qos":43,"signature":""}],"funded_payment_channel":[true]}`),
		signature:         "",
	}

	contentID, nonce, found, subOffers, FPCs, err := DecodeClientStandardDiscoverResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, contentID, mockContentID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, found, mockFound)
	assert.Equal(t, subOffers, mockSubCidOffers)
	assert.Equal(t, FPCs, mockFPCs)
}
