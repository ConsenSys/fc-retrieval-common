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
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

// KeyPair holds information related to a key pair. Either of the private key or
// public key may be nil.
type KeyPair struct {
	pKey   []byte
	pubKey []byte
	alg    KeySigAlg
}

// GenerateRetrievalV1KeyPair generates a key pair for use with v1 of the Filecoin Retrieval protocol.
func GenerateRetrievalV1KeyPair() (*KeyPair, error) {
	pkey, err := secp256k1GenerateKey()
	key := KeyPair{}
	key.pKey = pkey
	key.alg = DecodeSigAlg(SigAlgEcdsaSecP256K1Blake2b)
	return &key, err
}

// EncodePrivateKey encodes the algorithm and private key as a hex string.
func (k *KeyPair) EncodePrivateKey() string {
	data := k.alg.EncodeSigAlgAsBytes()
	all := append(data, k.pKey...)
	return hex.EncodeToString(all)
}

// EncodeRawPrivateKey encodes the private key as a hex string. Note that the
// algorithm is not stored with the key. As such, the code will return an error
// if the key algorithm is anything other than SigAlgEcdsaSecP256K1Blake2b.
func (k *KeyPair) EncodeRawPrivateKey() (string, error) {
	if k.alg.IsNot(SigAlgEcdsaSecP256K1Blake2b) {
		return "", fmt.Errorf("can not raw encode private key with algorithm: %d", k.alg.EncodeSigAlg())
	}
	return hex.EncodeToString(k.pKey), nil
}

// DecodePrivateKey decodes the algorithm and private key from a hex string.
func DecodePrivateKey(encoded string) (*KeyPair, error) {
	algKeyBytes, err := hex.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	alg := algKeyBytes[0]
	switch alg {
	case SigAlgEcdsaSecP256K1Blake2b:
		return decodeSecP256K1PrivateKey(algKeyBytes[1:])
	default:
		return nil, fmt.Errorf("unknown private key algorithm: %d", alg)
	}
}

// DecodeRawPrivateKey decodes the private key from a hex string, and assumes the key algorithm
// is the default one for Filecoin: SigAlgEcdsaSecP256K1Blake2b
func DecodeRawPrivateKey(encoded string) (*KeyPair, error) {
	keyBytes, err := hex.DecodeString(encoded)
	if err != nil {
		return nil, err
	}
	return decodeSecP256K1PrivateKey(keyBytes)
}

func decodeSecP256K1PrivateKey(keyBytes []byte) (*KeyPair, error) {
	key := KeyPair{}
	key.alg = DecodeSigAlg(SigAlgEcdsaSecP256K1Blake2b)
	key.pKey = keyBytes
	if len(key.pKey) != secp256k1PrivateKeyBytes {
		return nil, fmt.Errorf("incorrect secp256k1 private key length: %d", len(key.pKey))
	}
	return &key, nil
}

// EncodePublicKey encodes the algorithm and public key as a hex string.
func (k *KeyPair) EncodePublicKey() (string, error) {
	if k.alg.IsNot(SigAlgEcdsaSecP256K1Blake2b) {
		return "", fmt.Errorf("unsupported key algorithm: %d", k.alg.EncodeSigAlg())
	}
	if k.pubKey == nil {
		k.pubKey = secp256k1PublicKey(k.pKey)
	}
	data := k.alg.EncodeSigAlgAsBytes()
	all := append(data, k.pubKey...)
	return hex.EncodeToString(all), nil
}

// DecodePublicKey decodes the algorithm and public key from a hex string.
func DecodePublicKey(encoded string) (*KeyPair, error) {
	algKeyBytes, err := hex.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	alg := algKeyBytes[0]
	switch alg {
	case SigAlgEcdsaSecP256K1Blake2b:
		return decodeSecP256K1PublicKey(algKeyBytes[1:])
	default:
		return nil, fmt.Errorf("unknown private key algorithm: %d", alg)
	}
}

func decodeSecP256K1PublicKey(keyBytes []byte) (*KeyPair, error) {
	key := KeyPair{}
	key.alg = DecodeSigAlg(SigAlgEcdsaSecP256K1Blake2b)
	key.pubKey = keyBytes
	if len(key.pubKey) != secp256k1PublicKeyBytes {
		return nil, fmt.Errorf("incorrect secp256k1 public key length: %d", len(key.pubKey))
	}
	return &key, nil
}

// Sign some data.
func (k *KeyPair) Sign(toBeSigned []byte) ([]byte, error) {
	if k.alg.IsNot(SigAlgEcdsaSecP256K1Blake2b) {
		return nil, fmt.Errorf("unsupported key algorithm: %d", k.alg.EncodeSigAlg())
	}
	digest := RetrievalV1Hash(toBeSigned)
	return secp256k1Sign(k.pKey, digest)
}

// Verify a signature across some data.
func (k *KeyPair) Verify(signature, toBeSigned []byte) (bool, error) {
	if k.alg.IsNot(SigAlgEcdsaSecP256K1Blake2b) {
		return false, fmt.Errorf("unsupported key algorithm: %d", k.alg.EncodeSigAlg())
	}
	if k.pubKey == nil {
		k.pubKey = secp256k1PublicKey(k.pKey)
	}

	digest := RetrievalV1Hash(toBeSigned)
	return secp256k1Verify(k.pubKey, digest, signature), nil
}

// RetrievalV1Verify verifies a signature across some data assuming algorithms
// used for Retrieval V1.
func RetrievalV1Verify(signature, toBeSigned, hashedPublicKey []byte) (bool, error) {
	digest := RetrievalV1Hash(toBeSigned)
	pubKey, err := secp256k1EcRecover(digest, signature)
	if err != nil {
		return false, err
	}

	alg := DecodeSigAlg(SigAlgEcdsaSecP256K1Blake2b)
	all := append(alg.EncodeSigAlgAsBytes(), pubKey...)
	calculatedHash := RetrievalV1Hash(all)
	return bytes.Compare(calculatedHash, hashedPublicKey) == 0, nil
}

// HashPublicKey generates a message digest that matches the public key.
func (k *KeyPair) HashPublicKey() ([]byte, error) {
	if k.alg.IsNot(SigAlgEcdsaSecP256K1Blake2b) {
		return nil, fmt.Errorf("unsupported key algorithm: %d", k.alg.EncodeSigAlg())
	}

	if k.pubKey == nil {
		k.pubKey = secp256k1PublicKey(k.pKey)
	}

	data := k.alg.EncodeSigAlgAsBytes()
	all := append(data, k.pubKey...)

	return RetrievalV1Hash(all), nil
}

// MarshalJSON is used to marshal public Keypair into bytes
func (k KeyPair) MarshalJSON() ([]byte, error) {
	pKeyStr, err := k.EncodePublicKey()
	if err != nil {
		return nil, err
	}
	return json.Marshal(pKeyStr)
}

// UnmarshalJSON is used to unmarshal bytes into Keypair
func (k *KeyPair) UnmarshalJSON(p []byte) error {
	var pKeyStr string
	err := json.Unmarshal(p, &pKeyStr)
	if err != nil {
		return err
	}
	res, err := DecodePublicKey(pKeyStr)
	if err != nil {
		return err
	}
	k.alg = res.alg
	k.pubKey = res.pubKey
	return nil
}
