package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

// TestEncodeProviderPublishDHTOfferResponse success test
func TestEncodeProviderPublishDHTOfferResponse(t *testing.T) {
	mockNonce := int64(42)
	mockSignature := "00000001e79fdaa275888de3b3171ddf219d61cd19df42f1fa942d88150722438efcdcaf55195e10fd8267536eb9e807061460bec41d8b04a88cf53a68a0c17f383c625801"
	validMsg := &FCRMessage{
		messageType:303,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"nonce":42,"signature":"00000001e79fdaa275888de3b3171ddf219d61cd19df42f1fa942d88150722438efcdcaf55195e10fd8267536eb9e807061460bec41d8b04a88cf53a68a0c17f383c625801"}`), 
		signature:"",
	}

	msg, err := EncodeProviderPublishDHTOfferResponse(mockNonce, mockSignature)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderPublishDHTOfferResponse success test
func TestDecodeProviderPublishDHTOfferResponse(t *testing.T) {
	mockNonce := int64(42)
	mockSignature := "MySignature"
	validMsg := &FCRMessage{
		messageType:303,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"nonce":42,"signature":"MySignature"}`), 
		signature:"",
	}

	nonce, signature, err := DecodeProviderPublishDHTOfferResponse(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, signature, mockSignature)
	
}