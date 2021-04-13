package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderAdminInitialiseKeyResponse success test
func TestEncodeProviderAdminInitialiseKeyResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:501,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"success":true}`), 
		signature:"",
	}
	version, err := EncodeProviderAdminInitialiseKeyResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, version, validMsg)
}

// TestDecodeProviderAdminInitialiseKeyResponse success test
func TestDecodeProviderAdminInitialiseKeyResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:501,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{}`), 
		signature:"",
	}
	msg, err := DecodeProviderAdminInitialiseKeyResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, false)
}