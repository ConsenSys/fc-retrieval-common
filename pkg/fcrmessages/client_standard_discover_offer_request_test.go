package fcrmessages

import (
	"testing"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/stretchr/testify/assert"
)

// TestEncodeClientStandardDiscoverOfferRequest success test
func TestEncodeClientStandardDiscoverOfferRequest(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockTTL := int64(100)

	mockOfferDigests := [][cidoffer.CIDOfferDigestSize]byte{{1, 2, 4}, {5, 6, 7}}
	mockPayChAddr := "t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq"
	mockVoucher := "i1UCnYNY6cC8M4VO8IJjXfwn-bJT0g4AAED2AAFJAEVjkYJE9AAAAIBYYQK3pJLhIR8XTVSmQzsEiE7NIId2-2DPbWF396mBPBJCdoSQ_ctibPesW-YMnzKhGAEScF09H_sldF1nTfizTbsjWea9MN6R3T0Ew0Lb4znHtJnucGAkcbdlIyDAHCScOXE"
	validMsg := &FCRMessage{
		messageType:       110,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"ttl":100,"offer_digests":[[1,2,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[5,6,7,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-bJT0g4AAED2AAFJAEVjkYJE9AAAAIBYYQK3pJLhIR8XTVSmQzsEiE7NIId2-2DPbWF396mBPBJCdoSQ_ctibPesW-YMnzKhGAEScF09H_sldF1nTfizTbsjWea9MN6R3T0Ew0Lb4znHtJnucGAkcbdlIyDAHCScOXE"}`),
		signature:         "",
	}

	msg, err := EncodeClientStandardDiscoverOfferRequest(
		mockContentID,
		mockNonce,
		mockTTL,
		mockOfferDigests,
		mockPayChAddr,
		mockVoucher,
	)

	assert.Empty(t, err)
	assert.Equal(t, msg, validMsg)
}

// TestDecodeClientStandardDiscoverOfferRequest success test
func TestDecodeClientStandardDiscoverOfferRequest(t *testing.T) {
	mockContentID, _ := cid.NewContentIDFromBytes([]byte{1})
	mockNonce := int64(42)
	mockTTL := int64(100)
	mockOfferDigests := [][cidoffer.CIDOfferDigestSize]byte{{1, 2, 4}, {5, 6, 7}}
	mockPayChAddr := "t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq"
	mockVoucher := "i1UCnYNY6cC8M4VO8IJjXfwn-bJT0g4AAED2AAFJAEVjkYJE9AAAAIBYYQK3pJLhIR8XTVSmQzsEiE7NIId2-2DPbWF396mBPBJCdoSQ_ctibPesW-YMnzKhGAEScF09H_sldF1nTfizTbsjWea9MN6R3T0Ew0Lb4znHtJnucGAkcbdlIyDAHCScOXE"
	validMsg := &FCRMessage{
		messageType:       110,
		protocolVersion:   1,
		protocolSupported: []int32{1, 1},
		messageBody:       []byte(`{"piece_cid":"0000000000000000000000000000000000000000000000000000000000000001","nonce":42,"ttl":100,"offer_digests":[[1,2,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0],[5,6,7,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"payment_channel_address":"t2twbvr2oaxqzyktxqqjrv37bh7gzfhuqonfioayq","voucher":"i1UCnYNY6cC8M4VO8IJjXfwn-bJT0g4AAED2AAFJAEVjkYJE9AAAAIBYYQK3pJLhIR8XTVSmQzsEiE7NIId2-2DPbWF396mBPBJCdoSQ_ctibPesW-YMnzKhGAEScF09H_sldF1nTfizTbsjWea9MN6R3T0Ew0Lb4znHtJnucGAkcbdlIyDAHCScOXE"}`),
		signature:         "",
	}

	pieceCID, nonce, ttl, offersDigest, paychAddr, voucher, err := DecodeClientStandardDiscoverOfferRequest(validMsg)
	assert.Empty(t, err)
	assert.Equal(t, pieceCID, mockContentID)
	assert.Equal(t, nonce, mockNonce)
	assert.Equal(t, ttl, mockTTL)
	assert.Equal(t, offersDigest, mockOfferDigests)
	assert.Equal(t, paychAddr, mockPayChAddr)
	assert.Equal(t, voucher, mockVoucher)
}
