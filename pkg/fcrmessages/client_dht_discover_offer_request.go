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
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// clientDHTDiscoverOfferRequest is the requset from client to gateway to ask for cid offer via DHT
type clientDHTDiscoverOfferRequest struct {
	PieceCID        cid.ContentID                         `json:"piece_cid"`
	Nonce           int64                                 `json:"nonce"`
	GatewaysDigests [][][cidoffer.CIDOfferDigestSize]byte `json:"gateways_digests"`
	GatewayIDs      []nodeid.NodeID                       `json:"gateway_ids"`
	PaychAddr       string                                `json:"payment_channel_address"`
	Voucher         string                                `json:"voucher"`
}

// EncodeClientDHTDiscoverOfferRequest is used to get the FCRMessage of clientDHTDiscoverOfferRequest
func EncodeClientDHTDiscoverOfferRequest(
	pieceCID *cid.ContentID,
	nonce int64,
	gatewaysDigests [][][cidoffer.CIDOfferDigestSize]byte,
	gatewayIDs []nodeid.NodeID,
	paychAddr string,
	voucher string,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientDHTDiscoverOfferRequest{
		PieceCID:        *pieceCID,
		Nonce:           nonce,
		GatewaysDigests: gatewaysDigests,
		GatewayIDs:      gatewayIDs,
		PaychAddr:       paychAddr,
		Voucher:         voucher,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientDHTDiscoverOfferRequestType, body), nil
}

// DecodeClientDHTDiscoverOfferRequest is used to get the fields from FCRMessage of clientDHTDiscoverOfferRequest
func DecodeClientDHTDiscoverOfferRequest(fcrMsg *FCRMessage) (
	*cid.ContentID, // piece cid
	int64, // nonce
	[][][cidoffer.CIDOfferDigestSize]byte, // gateway offer digests
	[]nodeid.NodeID, // gateway id
	string, // payment channel address
	string, // voucher
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientDHTDiscoverOfferRequestType {
		return nil, 0, nil, nil, "", "", errors.New("message type mismatch")
	}
	msg := clientDHTDiscoverOfferRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, 0, nil, nil, "", "", err
	}
	return &msg.PieceCID, msg.Nonce, msg.GatewaysDigests, msg.GatewayIDs, msg.PaychAddr, msg.Voucher, nil
}
