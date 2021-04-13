package fcrmessages

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
)

// // TestEncodeProviderPublishGroupOfferResponse success test
func TestEncodeProviderPublishGroupOfferResponse(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)
	mockMsgDigest := mockOffer.GetMessageDigest()
	validMsg := &FCRMessage{
		messageType:301,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","digest":[179,39,63,55,115,199,133,123,164,143,223,117,88,234,218,129,141,20,110,11,165,108,175,4,83,15,0,120,195,74,216,203]}`), 
		signature:"",
	}

	msg, err := EncodeProviderPublishGroupOfferResponse(*mockProviderID, mockMsgDigest)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}


// TestDecodeProviderPublishGroupOfferResponse success test
func TestDecodeProviderPublishGroupOfferResponse(t *testing.T) {
	mockProviderID, _ := nodeid.NewNodeIDFromHexString("42")
	contentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockCids := []cid.ContentID{*contentID}
	var mockPrice uint64 = 41
	var mockExpiry int64 = 42
	var mockQos uint64 = 43
	mockOffer, err := cidoffer.NewCIDOffer(mockProviderID, mockCids, mockPrice, mockExpiry, mockQos)
	mockMsgDigest := mockOffer.GetMessageDigest()

	mockOfferResponse, _ := json.Marshal(providerPublishGroupOfferResponse{
		GatewaydID: *mockProviderID,
		Digest:     mockMsgDigest,
	})

	validMsg := &FCRMessage{
		messageType:301,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:mockOfferResponse, 
		signature:"",
	}
	nodeID, digest, err := DecodeProviderPublishGroupOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockProviderID)
	assert.Equal(t, digest, mockMsgDigest)
}