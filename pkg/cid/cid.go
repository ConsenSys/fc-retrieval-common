/*
Package cid - provides methods for ContentID struct.

ContentID is 32 bytes is a unique identifier of a file stored in a Filecoin blockchain network.
*/
package cid

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
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/cbergoon/merkletree"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
)

const WordSize = 32 // the ContentID length is 32 bytes.

// ContentID represents a CID.
type ContentID struct {
	id []byte
}

// NewContentID creates a ContentID object.
func NewContentID(id *big.Int) (*ContentID, error) {
	b := id.Bytes()
	l := len(b)
	if l > WordSize {
		return nil, fmt.Errorf("ContentID: Incorrect size: %d, should be fewer than %d", l, WordSize)
	}
	var n = ContentID{}
	n.id = make([]byte, WordSize)
	copy(n.id[WordSize-l:], b)
	return &n, nil
}

// NewContentIDFromBytes creates a ContentID object from bytes array.
func NewContentIDFromBytes(id []byte) (*ContentID, error) {
	l := len(id)
	if l > WordSize {
		return nil, fmt.Errorf("ContentID: Incorrect size: %d, should be fewer than %d", l, WordSize)
	}
	var n = ContentID{}
	n.id = make([]byte, WordSize)
	copy(n.id[WordSize-l:], id)
	return &n, nil
}

// NewContentIDFromHexString creates a ContentID object from hex string.
func NewContentIDFromHexString(id string) (*ContentID, error) {
	b, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}
	return NewContentIDFromBytes(b)
}

// NewRandomContentID creates a random ContentID object.
func NewRandomContentID() *ContentID {
	var n = ContentID{}
	n.id = make([]byte, WordSize)
	fcrcrypto.GeneratePublicRandomBytes(n.id)
	return &n
}

// ToString returns a string for the ContentID.
func (n *ContentID) ToString() string {
	str := hex.EncodeToString(n.id)
	if str == "" {
		str = "00"
	}
	return str
}

// ToBytes returns the byte array representation of the ContentID.
func (n *ContentID) ToBytes() []byte {
	return n.id
}

// MarshalJSON is used to marshal CID into bytes.
func (n ContentID) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.id)
}

// UnmarshalJSON is used to unmarshal bytes into ContentID.
func (n *ContentID) UnmarshalJSON(p []byte) error {
	var id []byte
	err := json.Unmarshal(p, &id)
	if err != nil {
		return err
	}

	if len(id) != WordSize {
		return fmt.Errorf("ContentID: Incorrect size: %d", len(id))
	}
	n.id = make([]byte, WordSize)
	copy(n.id, id)
	return nil
}

// CalculateHash hashes the values of a ContentID.
func (n ContentID) CalculateHash() ([]byte, error) {
	return n.id, nil
}

// Equals tests for equality of two ContentIDs.
func (n ContentID) Equals(other merkletree.Content) (bool, error) {
	return n.ToString() == other.(*ContentID).ToString(), nil
}
