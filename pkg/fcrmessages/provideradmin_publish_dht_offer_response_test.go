package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


// TestEncodeProviderAdminPublishDHTOfferResponse success test
func TestEncodeProviderAdminPublishDHTOfferResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:505,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"received":true}`), 
		signature:"",
	}
	version, err := EncodeProviderAdminPublishDHTOfferResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, version, validMsg)
}

// TestDecodeProviderAdminPublishDHTOfferResponse success test
func TestDecodeProviderAdminPublishDHTOfferResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:505,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{}`), 
		signature:"",
	}
	msg, err := DecodeProviderAdminPublishDHTOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, false)
}


