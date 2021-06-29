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

// providerPublishDHTOfferRequest is the request from provider to gateway to publish dht offer
type providerPublishDHTOfferRequest struct {
	ProviderID string              `json:"provider_id"`
	Nonce      int64               `json:"nonce"`
	NumOffers  int64               `json:"num_of_offers"`
	Offers     []cidoffer.CIDOffer `json:"single_offers"`
}

// EncodeProviderPublishDHTOfferRequest is used to get the FCRMessage of providerPublishDHTOfferRequest
func EncodeProviderPublishDHTOfferRequest(
	providerID *nodeid.NodeID,
	nonce int64,
	offers []cidoffer.CIDOffer,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerPublishDHTOfferRequest{
		ProviderID: providerID.ToString(),
		Nonce:      nonce,
		NumOffers:  int64(len(offers)),
		Offers:     offers,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderPublishDHTOfferRequestType, body), nil
}

// DecodeProviderPublishDHTOfferRequest is used to get the fields from FCRMessage of providerPublishDHTOfferRequest
func DecodeProviderPublishDHTOfferRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // provider id
	int64, // nonce
	[]cidoffer.CIDOffer, // offers
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderPublishDHTOfferRequestType {
		return nil, 0, nil, errors.New("message type mismatch")
	}
	msg := providerPublishDHTOfferRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, 0, nil, err
	}
	// Check every offer is single offer
	for _, offer := range msg.Offers {
		if len(offer.GetCIDs()) != 1 {
			return nil, 0, nil, errors.New("offers contain group offer")
		}
	}
	nodeID, _ := nodeid.NewNodeIDFromHexString(msg.ProviderID)
	return nodeID, msg.Nonce, msg.Offers, nil
}
