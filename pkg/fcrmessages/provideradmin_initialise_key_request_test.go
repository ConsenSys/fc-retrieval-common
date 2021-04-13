package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
)

// TestEncodeProviderAdminInitialiseKeyRequest success test
func TestEncodeProviderAdminInitialiseKeyRequest(t *testing.T) {
	mockNodeId, _ := nodeid.NewNodeIDFromHexString("42")
	mockPrivateKey, _ := fcrcrypto.GenerateRetrievalV1KeyPair()
	mockKeyVersion := fcrcrypto.InitialKeyVersion()

	validMsg := &FCRMessage{
		messageType:500,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","private_key":"` + mockPrivateKey.EncodePrivateKey() + `","private_key_version":1}`), 
		signature:"",
	}
	msg, err := EncodeProviderAdminInitialiseKeyRequest(mockNodeId, mockPrivateKey, mockKeyVersion)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeProviderAdminInitialiseKeyRequest success test
func TestDecodeProviderAdminInitialiseKeyRequest(t *testing.T) {
	mockNodeId, _ := nodeid.NewNodeIDFromHexString("42")
	mockPrivateKey, _ := fcrcrypto.GenerateRetrievalV1KeyPair()
	mockKeyVersion := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:500,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"provider_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","private_key":"` + mockPrivateKey.EncodePrivateKey() + `","private_key_version":1}`), 
		signature:"",
	}
	nodeID, keyPair, keyVersion, err := DecodeProviderAdminInitialiseKeyRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeId)
	assert.Equal(t, keyPair, mockPrivateKey)
	assert.Equal(t, keyVersion, mockKeyVersion)
}