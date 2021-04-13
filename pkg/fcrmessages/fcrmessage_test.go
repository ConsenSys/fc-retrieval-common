package fcrmessages

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
)

const (
	PrivKey     = "015ed053eab6fdf18c03954373ff7f89089992017d56beb8b05305b19800d6afe0"
	PubKey      = "01047799f37b014564e23578447d718e5c70a786b0e4e58ca25cb2a086b822434594d910b9b8c0fcbfe9f4c2db321e874819e0614be5b57fbb5080accd69adb2eaad"
	PubKeyWrong = "01047799f37b014564e23578447d718e5c70a786b0e4e58ca25cb2a086b822434594d910b9b8c0fcbfe9f4c2db321e874819e0614be5b57fbb5080accd69adb2eaaa"
)

// TestFCRMessage success test
func TestFCRMessage(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"",
	}

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)

	assert.Equal(t, msg, validMsg)
	assert.Equal(t, msg.GetMessageType(), validMsg.GetMessageType())
	assert.Equal(t, msg.GetProtocolVersion(), validMsg.GetProtocolVersion())
	assert.Equal(t, msg.GetProtocolSupported(), validMsg.GetProtocolSupported())
	assert.Equal(t, msg.GetMessageBody(), validMsg.GetMessageBody())
	assert.Equal(t, msg.GetSignature(), validMsg.GetSignature())
}

// TestSign success test
func TestSign(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	mockPrivKey, _ := fcrcrypto.DecodePrivateKey(PrivKey)
	mockKeyVer := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"00000001595e1c39d5334601e160c77a786ef088cb254a56de2669ab0534640d3d7d79bf2b2ef1fb5ac7b0ce4bc3e960cfdf3c40a0efbe95c3b0756fe931721be6098e1b01",
	}

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	msg.Sign(mockPrivKey, mockKeyVer)

	assert.Equal(t, msg, validMsg)
	assert.Equal(t, msg.GetSignature(), validMsg.GetSignature())
}

// TestVerify success test
func TestVerify(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	mockPrivKey, _ := fcrcrypto.DecodePrivateKey(PrivKey)
	mockPubKey, _ := fcrcrypto.DecodePublicKey(PubKey)
	mockKeyVer := fcrcrypto.InitialKeyVersion()
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"00000001595e1c39d5334601e160c77a786ef088cb254a56de2669ab0534640d3d7d79bf2b2ef1fb5ac7b0ce4bc3e960cfdf3c40a0efbe95c3b0756fe931721be6098e1b01",
	}

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	msg.Sign(mockPrivKey, mockKeyVer)
	err := msg.Verify(mockPubKey)

	assert.Equal(t, msg, validMsg)
	assert.Equal(t, msg.GetSignature(), validMsg.GetSignature())
	assert.Empty(t, err)
}

// TestFCRMsgDump success test
func TestFCRMsgDump(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"",
	}
	validDump := validMsg.DumpMessage()

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	dump := msg.DumpMessage()
	
	assert.Equal(t, msg, validMsg)
	assert.Equal(t, dump, validDump)
}

// TestFCRMsgToBytes success test
func TestFCRMsgToBytes(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"",
	}
	validFCRMsgToBytes, _ := validMsg.FCRMsgToBytes()

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	FCRMsgToBytes, _ := msg.FCRMsgToBytes()
	
	assert.Equal(t, msg, validMsg)
	assert.Equal(t, FCRMsgToBytes, validFCRMsgToBytes)
}

// TestFCRMsgFromBytes success test
func TestFCRMsgFromBytes(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"",
	}

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	FCRMsgToBytes, _ := msg.FCRMsgToBytes()
	FCRMsgFromBytes, _ := FCRMsgFromBytes(FCRMsgToBytes)
	
	assert.Equal(t, msg, validMsg)
	assert.Equal(t, FCRMsgFromBytes, validMsg)
}

// TestMarshalJSON success test
func TestMarshalJSON(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"",
	}
	validJson, _ := validMsg.MarshalJSON()

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	json, _ := msg.MarshalJSON()
	
	assert.Equal(t, msg, validMsg)
	assert.Equal(t, json, validJson)
}

// TestUnmarshalJSON success test
func TestUnmarshalJSON(t *testing.T) {
	mockMsgType := int32(203)
	mockMsgBody := []byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`)
	validMsg := &FCRMessage{
		messageType:203,
		protocolVersion:1,
		protocolSupported:[]int32{1, 1},
		messageBody:[]byte(`{"gateway_id":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAEI=","piece_cid":"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAE=","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`), 
		signature:"",
	}

	msg := CreateFCRMessage(mockMsgType, mockMsgBody)
	json, _ := msg.MarshalJSON()
	FCRMsg := &FCRMessage{}
	err := FCRMsg.UnmarshalJSON(json)

	assert.Equal(t, msg, validMsg)
	assert.Empty(t, err)
	assert.Equal(t, FCRMsg, validMsg)
	assert.Equal(t, FCRMsg.GetMessageType(), validMsg.GetMessageType())
	assert.Equal(t, FCRMsg.GetProtocolVersion(), validMsg.GetProtocolVersion())
	assert.Equal(t, FCRMsg.GetProtocolSupported(), validMsg.GetProtocolSupported())
	assert.Equal(t, FCRMsg.GetMessageBody(), validMsg.GetMessageBody())
	assert.Equal(t, FCRMsg.GetSignature(), validMsg.GetSignature())
}