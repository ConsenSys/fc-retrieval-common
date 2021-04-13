package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
// TestEncodeInsufficientFundsResponseValidMessage success test
func TestEncodeInsufficientFundsResponseValidMessage(t *testing.T) {
	var mockPaymentChannelIDValid int64 = 42
	validMsg := &FCRMessage{
		messageType:903,
		protocolVersion:1, 
		protocolSupported:[]int32{1, 1}, 
		messageBody:[]byte(`{"payment_channel_id":42}`), 
		signature:"",
	}

	msg, err := EncodeInsufficientFundsResponse(mockPaymentChannelIDValid)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeInsufficientFundsResponse success test
func TestDecodeInsufficientFundsResponse(t *testing.T) {
	var mockPaymentChannelIDValid int64 = 42
	validMsg := &FCRMessage{
		messageType:903,
		protocolVersion:1, 
		protocolSupported:[]int32{1, 1}, 
		messageBody:[]byte(`{"payment_channel_id":42}`), 
		signature:"",
	}

	msg, err := DecodeInsufficientFundsResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, mockPaymentChannelIDValid)
}