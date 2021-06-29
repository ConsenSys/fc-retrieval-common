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

	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// providerPublishGroupOfferRequest is the request from provider to gateway to publish group cid offer
type providerPublishGroupOfferRequest struct {
	ProviderID string            `json:"provider_id"`
	Nonce      int64             `json:"nonce"`
	Offer      cidoffer.CIDOffer `json:"offer"`
}

// EncodeProviderPublishGroupOfferRequest is used to get the FCRMessage of ProviderPublishGroupCIDRequest
func EncodeProviderPublishGroupOfferRequest(
	providerID *nodeid.NodeID,
	nonce int64,
	offer *cidoffer.CIDOffer,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerPublishGroupOfferRequest{
		ProviderID: providerID.ToString(),
		Nonce:      nonce,
		Offer:      *offer,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderPublishGroupOfferRequestType, body), nil
}

// DecodeProviderPublishGroupOfferRequest is used to get the fields from FCRMessage of providerPublishGroupOfferRequest
func DecodeProviderPublishGroupOfferRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // provider id
	int64, // nonce
	*cidoffer.CIDOffer, // offer
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderPublishGroupOfferRequestType {
		return nil, 0, nil, errors.New("message type mismatch")
	}
	msg := providerPublishGroupOfferRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, 0, nil, err
	}
	nodeID, _ := nodeid.NewNodeIDFromHexString(msg.ProviderID)
	return nodeID, msg.Nonce, &msg.Offer, nil
}
