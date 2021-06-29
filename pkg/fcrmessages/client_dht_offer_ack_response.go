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

// clientDHTOfferAckResponse is the response to clientDHTOfferAckRequest
type clientDHTOfferAckResponse struct {
	PieceCID                string     `json:"piece_cid"`
	GatewayID               string     `json:"gateway_id"`
	Found                   bool       `json:"found"`
	PublishDHTOfferRequest  FCRMessage `json:"publish_dht_offer_request"`
	PublishDHTOfferResponse FCRMessage `json:"publish_dht_offer_response"`
}

// EncodeClientDHTOfferAckResponse is used to get the FCRMessage of clientDHTOfferAckResponse
func EncodeClientDHTOfferAckResponse(
	pieceCID *cid.ContentID,
	gatewayID *nodeid.NodeID,
	found bool,
	publishDHTOfferRequest *FCRMessage,
	publishDHTOfferResponse *FCRMessage,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientDHTOfferAckResponse{
		PieceCID:                pieceCID.ToString(),
		GatewayID:               gatewayID.ToString(),
		Found:                   found,
		PublishDHTOfferRequest:  *publishDHTOfferRequest,
		PublishDHTOfferResponse: *publishDHTOfferResponse,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientDHTOfferAckResponseType, body), nil
}

// DecodeClientDHTOfferAckResponse is used to get the fields from FCRMessage of clientDHTOfferAckResponse
func DecodeClientDHTOfferAckResponse(fcrMsg *FCRMessage) (
	*cid.ContentID, // piece cid
	*nodeid.NodeID, // gateway id
	bool, // found
	*FCRMessage, // publish dht offer request
	*FCRMessage, // publish dht offer resposne
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientDHTOfferAckResponseType {
		return nil, nil, false, nil, nil, errors.New("message type mismatch")
	}
	msg := clientDHTOfferAckResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, false, nil, nil, err
	}
	contentID, _ := cid.NewContentIDFromHexString(msg.PieceCID)
	nodeID, _ := nodeid.NewNodeIDFromHexString(msg.GatewayID)
	return contentID, nodeID, msg.Found, &msg.PublishDHTOfferRequest, &msg.PublishDHTOfferResponse, nil
}
