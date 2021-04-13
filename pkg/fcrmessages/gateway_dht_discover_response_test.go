package fcrmessages

import (
	"testing"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayDHTDiscoverResponse success test
func TestEncodeGatewayDHTDiscoverResponse(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)

	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockFound := true
	mockSubOffer, _ := mockOffer.GenerateSubCIDOffer(mockContentID)
	mockSubOffers := []cidoffer.SubCIDOffer{*mockSubOffer}
	mockFPCs := []bool{true}
	
	validMsg := &FCRMessage{
		messageType:204,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"found":true,"sub_cid_offers":[{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","sub_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","merkle_root":"c3c3a46684c07d12a9c238787df3049a6f258e7af203e5ddb66a8bd66637e108","merkle_proof":"AAAAMFsiQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBRT0iXQAAAANbMV0=","price":41,"expiry":42,"qos":43,"signature":""}],"funded_payment_channel":[true]}`), 
		signature:"",
	}

	msg, err := EncodeGatewayDHTDiscoverResponse(mockContentID, mockNonce, mockFound, mockSubOffers, mockFPCs)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayDHTDiscoverResponse success test
func TestDecodeGatewayDHTDiscoverResponse(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	mockNonce := int64(42)
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)

	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockFound := true
	mockSubOffer, _ := mockOffer.GenerateSubCIDOffer(mockContentID)
	mockSubOffers := []cidoffer.SubCIDOffer{*mockSubOffer}
	mockFPCs := []bool{true}
	
	validMsg := &FCRMessage{
		messageType:204,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"found":true,"sub_cid_offers":[{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","sub_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","merkle_root":"c3c3a46684c07d12a9c238787df3049a6f258e7af203e5ddb66a8bd66637e108","merkle_proof":"AAAAMFsiQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBRT0iXQAAAANbMV0=","price":41,"expiry":42,"qos":43,"signature":""}],"funded_payment_channel":[true]}`), 
		signature:"",
	}

	contentID, nonce, found, subOffers, FPCs, err := DecodeGatewayDHTDiscoverResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, contentID, mockContentID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, found, mockFound)
	assert.Equal(t, subOffers, mockSubOffers)
	assert.Equal(t, FPCs, mockFPCs)
}