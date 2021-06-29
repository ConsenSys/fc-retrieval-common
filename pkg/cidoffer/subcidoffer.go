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
	"encoding/json"
	"errors"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmerkletree"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// SubCIDOffer represents a sub CID Offer. That is, part of a CID offer.
// It contains one sub cid and a merkle proof showing that this sub cid
// is part of the cid array in the original cid offer.
type SubCIDOffer struct {
	providerID  *nodeid.NodeID
	subCID      *cid.ContentID
	merkleRoot  string
	merkleProof *fcrmerkletree.FCRMerkleProof
	price       uint64
	expiry      int64
	qos         uint64
	signature   string
}

// subCIDOfferJson is used to parse to and from json.
type subCIDOfferJson struct {
	ProviderID  string                       `json:"provider_id"`
	SubCID      string                       `json:"sub_cid"`
	MerkleRoot  string                       `json:"merkle_root"`
	MerkleProof fcrmerkletree.FCRMerkleProof `json:"merkle_proof"`
	Price       uint64                       `json:"price"`
	Expiry      int64                        `json:"expiry"`
	QoS         uint64                       `json:"qos"`
	Signature   string                       `json:"signature"`
}

// subCIDOfferSigning is used to generate and verify signature.
type subCIDOfferSigning struct {
	ProviderID nodeid.NodeID `json:"provider_id"`
	MerkleRoot string        `json:"merkle_root"`
	Price      uint64        `json:"price"`
	Expiry     int64         `json:"expiry"`
	QoS        uint64        `json:"qos"`
}

// NewSubCIDOffer creates a sub CID Offer.
func NewSubCIDOffer(providerID *nodeid.NodeID, subCID *cid.ContentID, merkleRoot string, merkleProof *fcrmerkletree.FCRMerkleProof, price uint64, expiry int64, qos uint64, signature string) *SubCIDOffer {
	return &SubCIDOffer{
		providerID:  providerID,
		subCID:      subCID,
		merkleRoot:  merkleRoot,
		merkleProof: merkleProof,
		price:       price,
		expiry:      expiry,
		qos:         qos,
		signature:   signature,
	}
}

// GetProviderID returns the provider ID of this offer.
func (c *SubCIDOffer) GetProviderID() *nodeid.NodeID {
	return c.providerID
}

// GetSubCID returns the sub cid of this offer.
func (c *SubCIDOffer) GetSubCID() *cid.ContentID {
	return c.subCID
}

// GetMerkleRoot returns the merkle root of this offer.
func (c *SubCIDOffer) GetMerkleRoot() string {
	return c.merkleRoot
}

// GetMerkleProof returns the merkle proof of this offer.
func (c *SubCIDOffer) GetMerkleProof() *fcrmerkletree.FCRMerkleProof {
	return c.merkleProof
}

// GetPrice returns the price of this offer.
func (c *SubCIDOffer) GetPrice() uint64 {
	return c.price
}

// GetExpiry returns the expiry of this offer.
func (c *SubCIDOffer) GetExpiry() int64 {
	return c.expiry
}

// GetQoS returns the quality of service of this offer.
func (c *SubCIDOffer) GetQoS() uint64 {
	return c.qos
}

// GetSignature returns the signature of this offer.
func (c *SubCIDOffer) GetSignature() string {
	return c.signature
}

// HasExpired returns true if the offer expiry date is in the past.
func (c *SubCIDOffer) HasExpired() bool {
	expiryTime := time.Unix(c.expiry, 0)
	now := time.Now()
	return expiryTime.Before(now)
}

// Verify is used to verify the offer with a given public key.
func (c *SubCIDOffer) Verify(pubKey *fcrcrypto.KeyPair) error {
	raw, err := c.MarshalToSign()
	if err != nil {
		return err
	}
	res, err := fcrcrypto.VerifyMessage(pubKey, c.signature, raw)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("Offer does not pass signature verification")
	}
	return nil
}

// VerifyMerkleProof is used to verify the sub cid is part of the merkle trie
func (c *SubCIDOffer) VerifyMerkleProof() error {
	if c.merkleProof.VerifyContent(c.subCID, c.merkleRoot) {
		return nil
	}
	return errors.New("Offer does not pass merkle proof verification")
}

// MarshalJSON is used to marshal offer into bytes.
func (c SubCIDOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(subCIDOfferJson{
		ProviderID:  c.providerID.ToString(),
		SubCID:      c.subCID.ToString(),
		MerkleRoot:  c.merkleRoot,
		MerkleProof: *c.merkleProof,
		Price:       c.price,
		Expiry:      c.expiry,
		QoS:         c.qos,
		Signature:   c.signature,
	})
}

// UnmarshalJSON is used to unmarshal bytes into offer.
func (c *SubCIDOffer) UnmarshalJSON(p []byte) error {
	cJson := subCIDOfferJson{}
	err := json.Unmarshal(p, &cJson)
	if err != nil {
		return err
	}
	providerID, _ := nodeid.NewNodeIDFromHexString(cJson.ProviderID)
	c.providerID = providerID
	subCID, _ := cid.NewContentIDFromHexString(cJson.SubCID)
	c.subCID = subCID
	c.merkleRoot = cJson.MerkleRoot
	c.merkleProof = &cJson.MerkleProof
	c.price = cJson.Price
	c.expiry = cJson.Expiry
	c.qos = cJson.QoS
	c.signature = cJson.Signature
	return nil
}

func (c *SubCIDOffer) MarshalToSign() ([]byte, error) {
	return json.Marshal(subCIDOfferSigning{
		ProviderID: *c.providerID,
		MerkleRoot: c.merkleRoot,
		Price:      c.price,
		Expiry:     c.expiry,
		QoS:        c.qos,
	})
}
