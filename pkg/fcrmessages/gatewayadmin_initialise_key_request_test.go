package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminInitialiseKeyRequest success test
func TestEncodeGatewayAdminInitialiseKeyRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockPrivateKey, _ := fcrcrypto.GenerateRetrievalV1KeyPair()
	mockKeyVersion := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:       400,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","private_key":"` + mockPrivateKey.EncodePrivateKey() + `","private_key_version":1}`),
		signature:         "",
	}

	msg, err := EncodeGatewayAdminInitialiseKeyRequest(mockNodeID, mockPrivateKey, mockKeyVersion)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayAdminInitialiseKeyRequest success test
func TestDecodeGatewayAdminInitialiseKeyRequest(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockPrivateKey, _ := fcrcrypto.GenerateRetrievalV1KeyPair()
	mockKeyVersion := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:       400,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","private_key":"` + mockPrivateKey.EncodePrivateKey() + `","private_key_version":1}`),
		signature:         "",
	}

	nodeID, keyPair, keyVersion, err := DecodeGatewayAdminInitialiseKeyRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, keyPair, mockPrivateKey)
	assert.Equal(t, keyVersion, mockKeyVersion)
}
