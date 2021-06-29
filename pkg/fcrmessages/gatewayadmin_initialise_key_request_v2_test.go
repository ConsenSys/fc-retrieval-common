package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayAdminInitialiseKeyRequestV2 success test
func TestEncodeGatewayAdminInitialiseKeyRequestV2(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockPrivateKey, _ := fcrcrypto.GenerateRetrievalV1KeyPair()
	mockKeyVersion := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:       412,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody: []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","private_key":"` + mockPrivateKey.EncodePrivateKey() +
			`","private_key_version":1,"wallet_private_key":"wallet_private_key` +
			`","lotus_ap":"lotus_ap","lotus_auth_token":"lotus_auth_token"}`),
		signature: "",
	}

	msg, err := EncodeGatewayAdminInitialiseKeyRequestV2(mockNodeID, mockPrivateKey, mockKeyVersion, "wallet_private_key", "lotus_ap", "lotus_auth_token")
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayAdminInitialiseKeyRequestV2 success test
func TestDecodeGatewayAdminInitialiseKeyRequestV2(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockPrivateKey, _ := fcrcrypto.GenerateRetrievalV1KeyPair()
	mockKeyVersion := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:       412,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody: []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","private_key":"` + mockPrivateKey.EncodePrivateKey() +
			`","private_key_version":1,"wallet_private_key":"wallet_private_key02` +
			`","lotus_ap":"lotus_ap02","lotus_auth_token":"lotus_auth_token02"}`),
		signature: "",
	}

	nodeID, keyPair, keyVersion, walletPrivateKey, lotusAP, lotusAuthToken, err := DecodeGatewayAdminInitialiseKeyRequestV2(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, keyPair, mockPrivateKey)
	assert.Equal(t, keyVersion, mockKeyVersion)
	assert.Equal(t, walletPrivateKey, "wallet_private_key02")
	assert.Equal(t, lotusAP, "lotus_ap02")
	assert.Equal(t, lotusAuthToken, "lotus_auth_token02")
}
