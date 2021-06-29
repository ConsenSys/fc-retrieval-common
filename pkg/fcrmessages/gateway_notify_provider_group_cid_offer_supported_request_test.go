package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

func TestEncodeGatewayNotifyProviderGroupCIDOfferSupportRequest(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       205,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","group_cid_offer_supported":true}`),
		signature:         "",
	}
	fakeGatewayNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	msg, err := EncodeGatewayNotifyProviderGroupCIDOfferSupportRequest(fakeGatewayNodeID, true)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

func TestDecodeGatewayNotifyProviderGroupCIDOfferSupportRequest(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       205,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","group_cid_offer_supported":true}`),
		signature:         "",
	}
	fakeGatewayNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	gatewayNodeID, groupCIDOfferSupported, err := DecodeGatewayNotifyProviderGroupCIDOfferSupportRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, fakeGatewayNodeID, gatewayNodeID)
	assert.Equal(t, true, groupCIDOfferSupported)
}
