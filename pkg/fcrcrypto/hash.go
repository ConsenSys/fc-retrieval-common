/*
Package fcrcrypto - location for cryptographic tools to perform common operations on hashes, keys and signatures
*/
package fcrcrypto

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
	"hash"

	"golang.org/x/crypto/blake2b"
)

// GetBlockchainHasher returns a message digest implementation that hashes according to the
// algorithms used by the Filecoin blockchain.
func GetBlockchainHasher() hash.Hash {
	digestImpl, err := blake2b.New256(nil)
	if err != nil {
		// An error in getting a new message digest instance is catastrophic.
		panic(err)
	}
	return digestImpl
}

// BlockchainHash message digests some data using the algorithm used by the Filecoin blockchain.
func BlockchainHash(data []byte) []byte {
	hashSum := blake2b.Sum256(data)
	return hashSum[:]
}

// GetRetrievalV1Hasher returns a message digest implementation that hashes according to the
// algorithms used by version one of the Filecoin retrieval protocol.
func GetRetrievalV1Hasher() hash.Hash {
	digestImpl, err := blake2b.New256(nil)
	if err != nil {
		// An error in getting a new message digest instance is catastrophic.
		panic(err)
	}
	return digestImpl
}

// RetrievalV1Hash message digests some data using the algorithm used by version one of the
// Filecoin retrieval protocol.
func RetrievalV1Hash(data []byte) []byte {
	hashSum := blake2b.Sum256(data)
	return hashSum[:]
}

// The PRNG to be used as part of the PRF in the PRNG
func getPRNGHasher() hash.Hash {
	digestImpl, err := blake2b.New256(nil)
	if err != nil {
		// An error in getting a new message digest instance is catastrophic.
		panic(err)
	}
	return digestImpl
}

// The size of the message digest hash.
func getPRNGHasherDigestSize() int {
	return blake2b.Size256
}
