package fcrmessages

import (
	"testing"
	"errors"
	"github.com/stretchr/testify/assert"
)

// TestEncodeInvalidMessageResponse success test
func TestEncodeInvalidMessageResponse(t *testing.T) {
	emptyMsg := &FCRMessage{
		messageType:902,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody: []byte(`{}`),
		signature:"",
	}

	msg, err := EncodeInvalidMessageResponse()
	assert.Empty(t, err)
	assert.Equal(t, msg, emptyMsg)
}

// TestDecodeInvalidMessageResponseValid success test
func TestDecodeInvalidMessageResponseValid(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:902,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{}`), 
		signature:"",
	}

	err := DecodeInvalidMessageResponse(validMsg)
	assert.Equal(t, err, nil)
}

// TestDecodeInvalidMessageResponseInvalid error test
func TestDecodeInvalidMessageResponseInvalid(t *testing.T) {
	invalidTypeMsg := &FCRMessage{
		messageType:-1,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody: []byte(`{}`), 
		signature:"",
	}

	err := DecodeInvalidMessageResponse(invalidTypeMsg)
	assert.Equal(t, err, errors.New("Message type mismatch"))
}

