package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/stretchr/testify/assert"
)

// TestEncodeClientDHTDiscoverRequestV2 success test
func TestEncodeClientDHTDiscoverRequestV2(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockTTL := int64(100)
	mockNumDHT := int64(42)
	mockIncrementalResults := true
	mockPaychAddr := "0x42"
	mockVoucher := "1"
	validMsg := &FCRMessage{
		messageType:       112,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"ttl":100,"num_dht":42,"incremental_results":true,"payment_channel_address":"0x42","voucher":"1"}`),
		signature:         "",
	}
	msg, err := EncodeClientDHTDiscoverRequestV2(
		mockContentID,
		mockNonce,
		mockTTL,
		mockNumDHT,
		mockIncrementalResults,
		mockPaychAddr,
		mockVoucher,
	)
	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientDHTDiscoverRequestV2 success test
func TestDecodeClientDHTDiscoverRequestV2(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockTTL := int64(100)
	mockNumDHT := int64(42)
	mockIncrementalResults := true
	mockPaychAddr := "0x42"
	mockVoucher := "1"

	validMsg := &FCRMessage{
		messageType:       112,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"ttl":100,"num_dht":42,"incremental_results":true,"payment_channel_address":"0x42","voucher":"1"}`),
		signature:         "",
	}
	PieceCID, Nonce, TTL, NumDHT, IncrementalResults, PaychAddr, Voucher, err := DecodeClientDHTDiscoverRequestV2(validMsg)

	assert.Empty(t, err)
	assert.Equal(t, PieceCID, mockContentID)
	assert.Equal(t, Nonce, mockNonce)
	assert.Equal(t, TTL, mockTTL)
	assert.Equal(t, NumDHT, mockNumDHT)
	assert.Equal(t, IncrementalResults, mockIncrementalResults)
	assert.Equal(t, PaychAddr, mockPaychAddr)
	assert.Equal(t, Voucher, mockVoucher)
}
