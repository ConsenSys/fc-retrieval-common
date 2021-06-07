/*
Package fcrmerkletree - provides methods to work with Merkle Tree data structure
*/
package fcrmerkletree

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
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/cbergoon/merkletree"
)

// FCRMerkleProof is the proof of a single cid in a merkle tree.
type FCRMerkleProof struct {
	path  [][]byte
	index []int64
}

// VerifyContent is used to verify a given content and a given root matches the proof.
func (mp *FCRMerkleProof) VerifyContent(content merkletree.Content, root string) bool {
	currentHash, _ := content.CalculateHash()
	for i, path := range mp.path {
		hashFunc := sha256.New()
		if mp.index[i] == 1 {
			hashFunc.Write(append(currentHash, path...))
		} else {
			hashFunc.Write(append(path, currentHash...))
		}
		currentHash = hashFunc.Sum(nil)
	}
	return hex.EncodeToString(currentHash) == root
}

// MarshalJSON is used to marshal FCRMerkleProof into bytes.
func (mp FCRMerkleProof) MarshalJSON() ([]byte, error) {
	// Encode path
	pathBytes, err := json.Marshal(mp.path)
	if err != nil {
		return nil, err
	}
	// Put length
	length1 := make([]byte, 4)
	binary.BigEndian.PutUint32(length1, uint32(len(pathBytes)))
	// Encode index
	indexBytes, err := json.Marshal(mp.index)
	if err != nil {
		return nil, err
	}
	// Put length
	length2 := make([]byte, 4)
	binary.BigEndian.PutUint32(length2, uint32(len(indexBytes)))

	// Append result
	res := append(length1, pathBytes...)
	res = append(res, length2...)
	res = append(res, indexBytes...)
	return json.Marshal(res)
}

// UnmarshalJSON is used to unmarshal bytes into FCRMerkleProof.
func (mp *FCRMerkleProof) UnmarshalJSON(p []byte) error {
	var current []byte
	err := json.Unmarshal(p, &current)
	if err != nil {
		return err
	}
	// Decode path
	if len(current) <= 4 {
		return fmt.Errorf("FCRMerkleProof: Incorrect size")
	}
	data := current[:4]
	current = current[4:]
	length1 := int(binary.BigEndian.Uint32(data))
	if len(current) <= length1 {
		return fmt.Errorf("FCRMerkleProof: Incorrect size")
	}
	data = current[:length1]
	current = current[length1:]
	var path [][]byte
	err = json.Unmarshal(data, &path)
	if err != nil {
		return err
	}
	// Decode index
	if len(current) <= 4 {
		return fmt.Errorf("FCRMerkleProof: Incorrect size")
	}
	data = current[:4]
	current = current[4:]
	length2 := int(binary.BigEndian.Uint32(data))
	if len(current) != length2 {
		return fmt.Errorf("FCRMerkleProof: Incorrect size")
	}
	var index []int64
	err = json.Unmarshal(current, &index)
	if err != nil {
		return err
	}
	mp.path = path
	mp.index = index
	return nil
}
