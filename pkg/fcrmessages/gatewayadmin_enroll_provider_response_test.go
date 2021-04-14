package fcrmessages

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminEnrollProviderResponse success test
func TestEncodeGatewayAdminEnrollProviderResponse(t *testing.T) {

	validMsg := &FCRMessage{
		messageType:       409,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	msg, err := EncodeGatewayAdminEnrollProviderResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, validMsg, msg)
}

// TestDecodeGatewayAdminEnrollProviderResponse success test
func TestDecodeGatewayAdminEnrollProviderResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       409,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	enrolled, err := DecodeGatewayAdminEnrollProviderResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, true, enrolled)
}
