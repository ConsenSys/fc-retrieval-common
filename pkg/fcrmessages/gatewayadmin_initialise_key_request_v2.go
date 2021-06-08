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

// gatewayAdminInitialiseKeyRequestV2 is the request from a gateway admin to a gateway to initialise with a key pair,
// and lotus access point and lotus auth token
type gatewayAdminInitialiseKeyRequestV2 struct {
	GatewayID         nodeid.NodeID `json:"gateway_id"`
	PrivateKey        string        `json:"private_key"`
	PrivateKeyVersion uint32        `json:"private_key_version"`
	WalletPrivateKey  string        `json:"wallet_private_key"`
	LotusAP           string        `json:"lotus_ap"`
	LotusAuthToken    string        `json:"lotus_auth_token"`
}

// EncodeGatewayAdminInitialiseKeyRequestV2 is used to get the FCRMessage of gatewayAdminInitialiseKeyRequestV2
func EncodeGatewayAdminInitialiseKeyRequestV2(
	nodeID *nodeid.NodeID,
	privateKey *fcrcrypto.KeyPair,
	keyVersion *fcrcrypto.KeyVersion,
	walletPrivateKey string,
	lotusAP string,
	lotusAuthToken string,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayAdminInitialiseKeyRequestV2{
		*nodeID,
		privateKey.EncodePrivateKey(),
		keyVersion.EncodeKeyVersion(),
		walletPrivateKey,
		lotusAP,
		lotusAuthToken,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminInitialiseKeyRequestV2Type, body), nil
}

// DecodeGatewayAdminInitialiseKeyRequestV2 is used to get the fields from FCRMessage of gatewayAdminInitialiseKeyRequestV2
func DecodeGatewayAdminInitialiseKeyRequestV2(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway id
	*fcrcrypto.KeyPair, // private key
	*fcrcrypto.KeyVersion, // private key version
	string, // wallet private key
	string, // lotus ap
	string, // lotus auth token
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminInitialiseKeyRequestV2Type {
		return nil, nil, nil, "", "", "", errors.New("message type mismatch")
	}
	msg := gatewayAdminInitialiseKeyRequestV2{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, nil, "", "", "", err
	}
	privKey, err := fcrcrypto.DecodePrivateKey(msg.PrivateKey)
	if err != nil {
		return nil, nil, nil, "", "", "", errors.New("fail to decode private key")
	}
	privKeyVer := fcrcrypto.DecodeKeyVersion(msg.PrivateKeyVersion)
	return &msg.GatewayID, privKey, privKeyVer, msg.WalletPrivateKey, msg.LotusAP, msg.LotusAuthToken, nil
}
