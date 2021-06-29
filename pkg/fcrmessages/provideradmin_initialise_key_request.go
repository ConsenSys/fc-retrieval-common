package fcrmessages

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

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// providerAdminInitialiseKeyRequest is the request from a provider admin to a provider to initialise with a key pair.
type providerAdminInitialiseKeyRequest struct {
	ProviderID        string `json:"provider_id"`
	PrivateKey        string `json:"private_key"`
	PrivateKeyVersion uint32 `json:"private_key_version"`
}

// EncodeProviderAdminInitialiseKeyRequest is used to get the FCRMessage of providerAdminInitialiseKeyRequest
func EncodeProviderAdminInitialiseKeyRequest(
	nodeID *nodeid.NodeID,
	privateKey *fcrcrypto.KeyPair,
	keyVersion *fcrcrypto.KeyVersion,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerAdminInitialiseKeyRequest{
		nodeID.ToString(),
		privateKey.EncodePrivateKey(),
		keyVersion.EncodeKeyVersion(),
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderAdminInitialiseKeyRequestType, body), nil
}

// DecodeProviderAdminInitialiseKeyRequest is used to get the fields from FCRMessage of providerAdminInitialiseKeyRequest
func DecodeProviderAdminInitialiseKeyRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // provider id
	*fcrcrypto.KeyPair, // private key
	*fcrcrypto.KeyVersion, // private key version
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderAdminInitialiseKeyRequestType {
		return nil, nil, nil, errors.New("message type mismatch")
	}
	msg := providerAdminInitialiseKeyRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, nil, err
	}
	privKey, err := fcrcrypto.DecodePrivateKey(msg.PrivateKey)
	if err != nil {
		return nil, nil, nil, errors.New("fail to decode private key")
	}
	privKeyVer := fcrcrypto.DecodeKeyVersion(msg.PrivateKeyVersion)
	nodeID, _ := nodeid.NewNodeIDFromHexString(msg.ProviderID)
	return nodeID, privKey, privKeyVer, nil
}
