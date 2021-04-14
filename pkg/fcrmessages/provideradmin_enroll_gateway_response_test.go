package fcrmessages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderAdminEnrollGatewayResponse success test
func TestEncodeProviderAdminEnrollGatewayResponse(t *testing.T) {

	validMsg := &FCRMessage{
		messageType:       509,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	msg, err := EncodeProviderAdminEnrollGatewayResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, validMsg, msg)
}

// TestDecodeProviderAdminEnrollGatewayResponse success test
func TestDecodeProviderAdminEnrollGatewayResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       509,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	enrolled, err := DecodeProviderAdminEnrollGatewayResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, true, enrolled)
}
