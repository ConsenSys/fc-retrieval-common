package cidoffer

/*
 * Copyright 2020 ConsenSys Software Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with
 * the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 *
 * SPDX-License-Identifier: Apache-2.0
 */

import (
	"math/big"
	"testing"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
	"github.com/stretchr/testify/assert"
)

const (
	PrivKey     = "015ed053eab6fdf18c03954373ff7f89089992017d56beb8b05305b19800d6afe0"
	PubKey      = "01047799f37b014564e23578447d718e5c70a786b0e4e58ca25cb2a086b822434594d910b9b8c0fcbfe9f4c2db321e874819e0614be5b57fbb5080accd69adb2eaad"
	PubKeyWrong = "01047799f37b014564e23578447d718e5c70a786b0e4e58ca25cb2a086b822434594d910b9b8c0fcbfe9f4c2db321e874819e0614be5b57fbb5080accd69adb2eaaa"
)

func TestNewCIDOffer(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	aCid, err := cid.NewContentID(big.NewInt(7))
	assert.Empty(t, err)
	cids := []cid.ContentID{*aCid}
	price := uint64(5)
	expiry := int64(10)
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	assert.NotEmpty(t, offer)
	assert.Equal(t, aNodeID.ToString(), offer.GetProviderID().ToString())
	offerCIDs := offer.GetCIDs()
	assert.Equal(t, len(cids), len(offerCIDs))
	for i := 0; i < len(cids); i++ {
		assert.Equal(t, cids[i].ToString(), offerCIDs[i].ToString())
	}
	assert.Equal(t, price, offer.GetPrice())
	assert.Equal(t, expiry, offer.GetExpiry())
	assert.Equal(t, qos, offer.GetQoS())
}

func TestNewCIDOfferMultipleCIDs(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	aCid1, err := cid.NewContentID(big.NewInt(7))
	assert.Empty(t, err)
	aCid2, err := cid.NewContentID(big.NewInt(8))
	assert.Empty(t, err)
	aCid3, err := cid.NewContentID(big.NewInt(9))
	assert.Empty(t, err)
	cids := []cid.ContentID{*aCid1, *aCid2, *aCid3}
	price := uint64(5)
	expiry := int64(10)
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	assert.NotEmpty(t, offer)
	assert.Equal(t, aNodeID.ToString(), offer.GetProviderID().ToString())
	offerCIDs := offer.GetCIDs()
	assert.Equal(t, len(cids), len(offerCIDs))
	for i := 0; i < len(cids); i++ {
		assert.Equal(t, cids[i].ToString(), offerCIDs[i].ToString())
	}
	assert.Equal(t, price, offer.GetPrice())
	assert.Equal(t, expiry, offer.GetExpiry())
	assert.Equal(t, qos, offer.GetQoS())
}

func TestNewCIDOfferWithError(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	cids := []cid.ContentID{}
	price := uint64(5)
	expiry := int64(10)
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.NotEmpty(t, err)
	assert.Empty(t, offer)
}

func TestHasExpired(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	aCid, err := cid.NewContentID(big.NewInt(7))
	assert.Empty(t, err)
	cids := []cid.ContentID{*aCid}
	price := uint64(5)
	expiry := time.Now().Add(12 * time.Hour).Unix()
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	assert.NotEmpty(t, offer)
	assert.False(t, offer.HasExpired())
	expiry = time.Now().Add(-12 * time.Hour).Unix()
	offer, err = NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	assert.NotEmpty(t, offer)
	assert.True(t, offer.HasExpired())
}

func TestSigning(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	aCid, err := cid.NewContentID(big.NewInt(7))
	assert.Empty(t, err)
	cids := []cid.ContentID{*aCid}
	price := uint64(5)
	expiry := int64(9_223_372_030_000_000_000)
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	privKey, err := fcrcrypto.DecodePrivateKey(PrivKey)
	assert.Empty(t, err)
	err = offer.Sign(privKey, fcrcrypto.InitialKeyVersion())
	assert.Empty(t, err)
	assert.Equal(t, "000000011445d8467f41ecc6c452f5dbb0a3a0fa2dacc24f0b3dd3a110713541c4fa8b3a2582b7724eb13ffad92dd0c312150bc196cc3eb981aa9a332dc7871cb0da18eb01", offer.GetSignature())

	pubKey, err := fcrcrypto.DecodePublicKey(PubKey)
	assert.Empty(t, err)
	err = offer.Verify(pubKey)
	assert.Empty(t, err)

	pubKey, err = fcrcrypto.DecodePublicKey(PubKeyWrong)
	assert.Empty(t, err)
	err = offer.Verify(pubKey)
	assert.NotEmpty(t, err)
}

func TestDigest(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	aCid, err := cid.NewContentID(big.NewInt(7))
	assert.Empty(t, err)
	cids := []cid.ContentID{*aCid}
	price := uint64(5)
	expiry := int64(10)
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	privKey, err := fcrcrypto.DecodePrivateKey(PrivKey)
	assert.Empty(t, err)
	err = offer.Sign(privKey, fcrcrypto.InitialKeyVersion())
	assert.Empty(t, err)
	assert.Equal(t, [32]byte{0x9b, 0x85, 0xe9, 0x73, 0xf3,
		0x94, 0xd, 0x71, 0xcc, 0x43, 0x55, 0xe8, 0xdc,
		0x65, 0x2a, 0x53, 0x54, 0x4f, 0x40, 0x3d, 0xd5,
		0x1f, 0xb6, 0x2d, 0x77, 0x12, 0x52, 0xed, 0x6f,
		0x82, 0x27, 0x1e}, offer.GetMessageDigest())
}

func TestJSON(t *testing.T) {
	aNodeID, err := nodeid.NewNodeID(big.NewInt(7))
	assert.Empty(t, err)
	aCid, err := cid.NewContentID(big.NewInt(7))
	assert.Empty(t, err)
	cids := []cid.ContentID{*aCid}
	price := uint64(5)
	expiry := int64(1)
	qos := uint64(5)
	offer, err := NewCIDOffer(aNodeID, cids, price, expiry, qos)
	assert.Empty(t, err)
	p, err := offer.MarshalJSON()
	assert.Empty(t, err)
	assert.Equal(t, []byte{0x7b, 0x22, 0x70, 0x72, 0x6f, 0x76,
		0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x3a,
		0x22, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x37, 0x22, 0x2c, 0x22, 0x63, 0x69, 0x64, 0x73,
		0x22, 0x3a, 0x5b, 0x22, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30,
		0x30, 0x30, 0x30, 0x30, 0x37, 0x22, 0x5d, 0x2c, 0x22,
		0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x3a, 0x35, 0x2c,
		0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x22, 0x3a,
		0x31, 0x2c, 0x22, 0x71, 0x6f, 0x73, 0x22, 0x3a, 0x35,
		0x2c, 0x22, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75,
		0x72, 0x65, 0x22, 0x3a, 0x22, 0x22, 0x7d}, p)
	offer2 := CIDOffer{}
	err = offer2.UnmarshalJSON(p)
	assert.Empty(t, err)
	assert.Equal(t, offer.GetProviderID(), offer2.GetProviderID())
	assert.Equal(t, offer.GetCIDs(), offer2.GetCIDs())
	assert.Equal(t, offer.GetPrice(), offer2.GetPrice())
	assert.Equal(t, offer.GetExpiry(), offer2.GetExpiry())
	assert.Equal(t, offer.GetQoS(), offer2.GetQoS())
	assert.Equal(t, offer.GetSignature(), offer2.GetSignature())
	assert.Equal(t, offer.merkleRoot, offer2.merkleRoot)
	err = offer2.UnmarshalJSON([]byte{})
	assert.NotEmpty(t, err)
}
