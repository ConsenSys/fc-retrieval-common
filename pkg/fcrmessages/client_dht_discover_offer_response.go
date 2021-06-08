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

// clientDHTDiscoverOfferResponse is the response to clientDHTDiscoverOfferRequest
type clientDHTDiscoverOfferResponse struct {
	PieceCID   cid.ContentID   `json:"piece_cid"`
	Nonce      int64           `json:"nonce"`
	GatewayIDs []nodeid.NodeID `json:"gateway_ids"`
	Response   []FCRMessage    `json:"response"`
}

// EncodeClientDHTDiscoverOfferResponse is used to get the FCRMessage of clientDHTDiscoverOfferResponse
func EncodeClientDHTDiscoverOfferResponse(
	pieceCID *cid.ContentID,
	nonce int64,
	gatewayIDs []nodeid.NodeID,
	response []FCRMessage,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientDHTDiscoverOfferResponse{
		PieceCID:   *pieceCID,
		Nonce:      nonce,
		GatewayIDs: gatewayIDs,
		Response:   response,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientDHTDiscoverOfferResponseType, body), nil
}

// DecodeClientDHTDiscoverOfferResponse is used to get the fields from FCRMessage of clientDHTDiscoverOfferResponse
func DecodeClientDHTDiscoverOfferResponse(fcrMsg *FCRMessage) (
	*cid.ContentID, // piece cid
	int64, // nonce
	[]nodeid.NodeID, // gateway ids
	[]FCRMessage, // respones slice
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientDHTDiscoverOfferResponseType {
		return nil, 0, nil, nil, errors.New("message type mismatch")
	}
	msg := clientDHTDiscoverOfferResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, 0, nil, nil, err
	}
	return &msg.PieceCID, msg.Nonce, msg.GatewayIDs, msg.Response, nil
}
