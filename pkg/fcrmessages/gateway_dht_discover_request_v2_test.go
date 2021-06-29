package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeGatewayDHTDiscoverRequestV2 success test
func TestEncodeGatewayDHTDiscoverRequestV2(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockTTL := int64(43)
	mockPaychAddr := "t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq"
	mockVoucher := "i1UCnYNY6cC8M4VO8IJjXfwn-"

	validMsg := &FCRMessage{
		messageType:       207,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`),
		signature:         "",
	}

	msg, err := EncodeGatewayDHTDiscoverRequestV2(mockNodeID, mockContentID, mockNonce, mockTTL, mockPaychAddr, mockVoucher)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeGatewayDHTDiscoverRequestV2 success test
func TestDecodeGatewayDHTDiscoverRequestV2(t *testing.T) {
	mockNodeID, _ := nodeid.NewNodeIDFromHexString("42")
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockTTL := int64(43)
	mockPaychAddr := "t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq"
	mockVoucher := "i1UCnYNY6cC8M4VO8IJjXfwn-"

	validMsg := &FCRMessage{
		messageType:       207,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"gateway_id":"0000000000000000000000000000000000000000000000000000000000000042","piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"ttl":43,"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-"}`),
		signature:         "",
	}

	nodeID, contentID, nonce, TTL, paychAddr, voucher, err := DecodeGatewayDHTDiscoverRequestV2(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, nodeID, mockNodeID)
	assert.Equal(t, contentID, mockContentID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, TTL, mockTTL)
	assert.Equal(t, paychAddr, mockPaychAddr)
	assert.Equal(t, voucher, mockVoucher)
}
