package fcrmessages

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

// TestEncodeProviderAdminEnrollProviderResponse success test
func TestEncodeProviderAdminEnrollProviderResponse(t *testing.T) {

	validMsg := &FCRMessage{
		messageType:       509,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	msg, err := EncodeProviderAdminEnrollProviderResponse(true)
	assert.Empty(t, err)
	assert.Equal(t, validMsg, msg)
}

// TestDecodeProviderAdminEnrollProviderResponse success test
func TestDecodeProviderAdminEnrollProviderResponse(t *testing.T) {
	validMsg := &FCRMessage{
		messageType:       509,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"enrolled":true}`),
	}

	enrolled, err := DecodeProviderAdminEnrollProviderResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, true, enrolled)
}
