package fcrmessages

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminEnrollGatewayResponse success test
func TestEncodeGatewayAdminEnrollGatewayResponse(t *testing.T) {

	validMsg := &FCRMessage{
		messageType:       407,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	msg, err := EncodeGatewayAdminEnrollGatewayResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, validMsg, msg)
}

// TestDecodeGatewayAdminEnrollGatewayResponse success test
func TestDecodeGatewayAdminEnrollGatewayResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       407,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	enrolled, err := DecodeGatewayAdminEnrollGatewayResponse(validMsg)

	assert.Empty(t, err)
	assert.Equal(t, true, enrolled)
}
