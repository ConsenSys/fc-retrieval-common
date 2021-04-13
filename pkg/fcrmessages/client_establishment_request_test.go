package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/challenge"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// TestEncodeClientEstablishmentRequest success test
func TestEncodeClientEstablishmentRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockChallenge := challenge.NewRandomChallenge()
	mockTTL := int64(100)
	validMsg := &FCRMessage{
		messageType:100,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"client_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","challenge":"` + mockChallenge + `","ttl":100}`), 
		signature:"",
	}

	msg, err := EncodeClientEstablishmentRequest(
		mockNodeID,
		mockChallenge,
		mockTTL,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientEstablishmentRequest success test
func TestDecodeClientEstablishmentRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockChallenge := challenge.NewRandomChallenge()
	mockTTL := int64(100)
	validMsg := &FCRMessage{
		messageType:100,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"client_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","challenge":"` + mockChallenge + `","ttl":100}`), 
		signature:"",
	}

	clientID, challenge, ttl, err := DecodeClientEstablishmentRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, clientID, mockNodeID)
	assert.Equal(t, challenge, mockChallenge)
	assert.Equal(t, ttl, mockTTL)
}
