package fcrmessages

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestEncodeProtocolChangeResponse success test
func TestEncodeProtocolChangeResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       901,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"success":true}`),
		signature:         "",
	}
	msg, err := EncodeProtocolChangeResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProtocolChangeResponse success test
func TestDecodeProtocolChangeResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       901,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{}`),
		signature:         "",
	}
	msg, err := DecodeProtocolChangeResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, msg, false)
}

// TestDecodeProtocolChangeResponse error type test
func TestDecodeProtocolChangeResponseErrorType(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       -1,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{}`),
		signature:         "",
	}
	msg, err := DecodeProtocolChangeResponse(validMsg)
	assert.Empty(t, msg)
	assert.Equal(t, err, errors.New("message type mismatch"))
}
