package fcrmessages

import (
	"testing"
	// "errors"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminInitialiseKeyResponse success test
func TestEncodeGatewayAdminInitialiseKeyResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:401,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"success":true}`), 
		signature:"",
	}
	version, err := EncodeGatewayAdminInitialiseKeyResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, version, validMsg)
}

// TestDecodeGatewayAdminInitialiseKeyResponse success test
func TestDecodeGatewayAdminInitialiseKeyResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:401,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{}`), 
		signature:"",
	}
	msg, err := DecodeGatewayAdminInitialiseKeyResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, false)
}