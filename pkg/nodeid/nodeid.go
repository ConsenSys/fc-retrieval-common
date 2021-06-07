/*
Package nodeid - combines common operation on FileCoin NodeID.

NodeID is a unique identifier for a node, participating in FileCoin network operations.
Example of the node might be a Retrieval Gateway or a Retrieval Provider.
*/
package nodeid

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
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"sort"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
)

const WordSize = 32 // the NodeID length is 32 bytes.

// NodeID represents a NodeID.
type NodeID struct {
	id []byte
}

// NewNodeID creates a NodeID object.
func NewNodeID(id *big.Int) (*NodeID, error) {
	var n = NodeID{}
	b := id.Bytes()
	l := len(b)
	if l > WordSize {
		return nil, fmt.Errorf("NodeID: Incorrect size: %d, should be fewer than %d", l, WordSize)
	}
	n.id = make([]byte, WordSize)
	copy(n.id[WordSize-l:], b)
	return &n, nil
}

// NewNodeIDFromBytes creates a NodeID object.
func NewNodeIDFromBytes(id []byte) (*NodeID, error) {
	var n = NodeID{}
	lenID := len(id)
	if lenID > WordSize {
		return nil, fmt.Errorf("NodeID: Incorrect size: %d, should be fewer than %d", lenID, WordSize)
	}
	n.id = make([]byte, WordSize)
	copy(n.id[WordSize-len(id):], id)
	return &n, nil
}

// NewNodeIDFromHexString creates a NodeID from a string.
func NewNodeIDFromHexString(id string) (*NodeID, error) {
	var n = NodeID{}
	bytes, err := hex.DecodeString(id)
	if err != nil {
		return nil, err
	}

	if len(bytes) > WordSize {
		return nil, fmt.Errorf("NodeID: Incorrect size: %d, should be fewer than %d", len(id), WordSize)
	}
	n.id = make([]byte, WordSize)
	copy(n.id[WordSize-len(bytes):], bytes)
	return &n, nil
}

// NewNodeIDFromPublicKey create a NodeID based on a public key.
func NewNodeIDFromPublicKey(pubKey *fcrcrypto.KeyPair) (*NodeID, error) {
	hashedPubKey, err := pubKey.HashPublicKey()
	if err != nil {
		return nil, err
	}
	return NewNodeIDFromBytes(hashedPubKey)
}

// NewRandomNodeID creates a random NodeID object.
func NewRandomNodeID() *NodeID {
	var n = NodeID{}
	n.id = make([]byte, WordSize)
	fcrcrypto.GeneratePublicRandomBytes(n.id)
	return &n
}

// ToString returns a string for the NodeID.
func (n *NodeID) ToString() string {
	str := hex.EncodeToString(n.id)
	if str == "" {
		str = "00"
	}
	return str
}

// ToBytes returns the byte array representation of the NodeID.
func (n *NodeID) ToBytes() []byte {
	return n.id
}

// AsBytes32 returns the NodeID as a [32]byte.
func (n *NodeID) AsBytes32() (result [WordSize]byte) {
	copy(result[:], n.id)
	return
}

// MarshalJSON is used to marshal NodeID into bytes.
func (n NodeID) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.id)
}

// UnmarshalJSON is used to unmarshal bytes into NodeID.
func (n *NodeID) UnmarshalJSON(p []byte) error {
	var id []byte
	err := json.Unmarshal(p, &id)
	if err != nil {
		return err
	}

	if len(id) != WordSize {
		return fmt.Errorf("NodeID: Incorrect size: %d, should be %d", len(id), WordSize)
	}
	n.id = make([]byte, WordSize)
	copy(n.id, id)
	return nil
}

func sortByteArrays(src []*NodeID) {
	sort.Slice(src,
		func(i, j int) bool {
			return bytes.Compare(src[i].ToBytes(), src[j].ToBytes()) < 0
		},
	)
}

// SortClockwise sort nodeIDs in clockwise mode starting on nodeID
func SortClockwise(nodeID *NodeID, nodeIDs []*NodeID) []*NodeID {
	if len(nodeIDs) < 2 {
		return nodeIDs
	}

	sortByteArrays(nodeIDs)

	startIndx := 0
	startBytes := nodeID.ToBytes()

	for i, v := range nodeIDs {
		if bytes.Compare(startBytes, v.ToBytes()) < 1 {
			startIndx = i
			break
		}
	}

	nodeIDs = append(nodeIDs[startIndx:], nodeIDs[0:startIndx]...)

	return nodeIDs
}
