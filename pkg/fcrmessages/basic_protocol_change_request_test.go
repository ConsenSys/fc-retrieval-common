package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProtocolChangeRequest success test
func TestEncodeProtocolChangeRequest(t *testing.T) {
	var desiredVersion int32 = 1
	validMsg := &FCRMessage{
		messageType:900,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody: []byte(`{"desired_version":1}`), 
		signature:"",
	}
	version, err := EncodeProtocolChangeRequest(desiredVersion)
	assert.Empty(t, err)
	assert.Equal(t, version, validMsg)
}

// TestDecodeProtocolChangeRequest success test
func TestDecodeProtocolChangeRequest(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:900,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"payment_channel_id":42}`), 
		signature:"",
	}
	msg, err := DecodeProtocolChangeRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, int32(0))
}

