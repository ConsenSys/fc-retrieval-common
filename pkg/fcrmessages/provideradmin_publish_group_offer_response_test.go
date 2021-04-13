package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderAdminPublishGroupOfferResponse success test
func TestEncodeProviderAdminPublishGroupOfferResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:503,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"received":true}`), 
		signature:"",
	}
	version, err := EncodeProviderAdminPublishGroupOfferResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, version, validMsg)
}

// TestDecodeProviderAdminPublishGroupOfferResponse success test
func TestDecodeProviderAdminPublishGroupOfferResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:503,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{}`), 
		signature:"",
	}
	msg, err := DecodeProviderAdminPublishGroupOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, false)
}
