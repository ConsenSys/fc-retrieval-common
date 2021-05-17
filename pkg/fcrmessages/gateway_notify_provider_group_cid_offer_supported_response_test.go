package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeGatewayNotifyProviderGroupCIDOfferSupportResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       206,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"acknowledged":true}`),
		signature:         "",
	}
	msg, err := EncodeGatewayNotifyProviderGroupCIDOfferSupportResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

func TestDecodeGatewayNotifyProviderGroupCIDOfferSupportResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       206,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"acknowledged":true}`),
		signature:         "",
	}
	msg, err := DecodeGatewayNotifyProviderGroupCIDOfferSupportResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, true)
}
