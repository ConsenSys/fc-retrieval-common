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

// gatewayAdminInitialiseKeyRequest is the request from a gateway admin to a gateway to initialise with a key pair.
type gatewayAdminInitialiseKeyRequest struct {
	GatewayID         nodeid.NodeID `json:"gateway_id"`
	PrivateKey        string        `json:"private_key"`
	PrivateKeyVersion uint32        `json:"private_key_version"`
}

// EncodeGatewayAdminInitialiseKeyRequest is used to get the FCRMessage of gatewayAdminInitialiseKeyRequest
func EncodeGatewayAdminInitialiseKeyRequest(
	nodeID *nodeid.NodeID,
	privateKey *fcrcrypto.KeyPair,
	keyVersion *fcrcrypto.KeyVersion,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayAdminInitialiseKeyRequest{
		*nodeID,
		privateKey.EncodePrivateKey(),
		keyVersion.EncodeKeyVersion(),
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminInitialiseKeyRequestType, body), nil
}

// DecodeGatewayAdminInitialiseKeyRequest is used to get the fields from FCRMessage of gatewayAdminInitialiseKeyRequest
func DecodeGatewayAdminInitialiseKeyRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway id
	*fcrcrypto.KeyPair, // private key
	*fcrcrypto.KeyVersion, // private key version
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminInitialiseKeyRequestType {
		return nil, nil, nil, errors.New("message type mismatch")
	}
	msg := gatewayAdminInitialiseKeyRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, nil, err
	}
	privKey, err := fcrcrypto.DecodePrivateKey(msg.PrivateKey)
	if err != nil {
		return nil, nil, nil, errors.New("fail to decode private key")
	}
	privKeyVer := fcrcrypto.DecodeKeyVersion(msg.PrivateKeyVersion)
	return &msg.GatewayID, privKey, privKeyVer, nil
}
