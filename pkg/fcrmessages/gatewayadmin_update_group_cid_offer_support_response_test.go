package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeUpdateGatewayGroupCIDOfferSupportResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       409,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"success":true}`),
		signature:         "",
	}
	msg, err := EncodeUpdateGatewayGroupCIDOfferSupportResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

func TestDecodeUpdateGatewayGroupCIDOfferSupportResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       409,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"success":true}`),
		signature:         "",
	}
	msg, err := DecodeUpdateGatewayGroupCIDOfferSupportResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, true)
}
