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

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// gatewayDHTDiscoverRequestV2 is the request from gateway to gateway to discover cid offer
type gatewayDHTDiscoverRequestV2 struct {
	GatewayID string `json:"gateway_id"`
	PieceCID  string `json:"piece_cid"`
	Nonce     int64  `json:"nonce"`
	TTL       int64  `json:"ttl"`
	PaychAddr string `json:"payment_channel_address"`
	Voucher   string `json:"voucher"`
}

// EncodeGatewayDHTDiscoverRequestV2 is used to get the FCRMessage of gatewayDHTDiscoverRequestV2
func EncodeGatewayDHTDiscoverRequestV2(
	gatewayID *nodeid.NodeID,
	pieceCID *cid.ContentID,
	nonce int64,
	ttl int64,
	paychAddr string,
	voucher string,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayDHTDiscoverRequestV2{
		GatewayID: gatewayID.ToString(),
		PieceCID:  pieceCID.ToString(),
		Nonce:     nonce,
		TTL:       ttl,
		PaychAddr: paychAddr,
		Voucher:   voucher,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayDHTDiscoverRequestV2Type, body), nil
}

// DecodeGatewayDHTDiscoverRequestV2 is used to get the fields from FCRMessage of GatewayDHTDiscoverRequestV2
func DecodeGatewayDHTDiscoverRequestV2(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway id
	*cid.ContentID, // piece cid
	int64, // nonce
	int64, // ttl
	string, // payment channel address
	string, // voucher
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayDHTDiscoverRequestV2Type {
		return nil, nil, 0, 0, "", "", errors.New("message type mismatch")
	}
	msg := gatewayDHTDiscoverRequestV2{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, 0, 0, "", "", err
	}
	nodeID, _ := nodeid.NewNodeIDFromHexString(msg.GatewayID)
	contentID, _ := cid.NewContentIDFromHexString(msg.PieceCID)
	return nodeID, contentID, msg.Nonce, msg.TTL, msg.PaychAddr, msg.Voucher, nil
}
