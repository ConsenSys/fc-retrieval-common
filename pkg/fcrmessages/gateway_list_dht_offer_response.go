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
)

// gatewayListDHTOfferResponse is the repsonse to gatewayListDHTOfferRequest
type gatewayListDHTOfferResponse struct {
	PublishedDHTOffers []FCRMessage `json:"published_dht_offers"`
}

// EncodeGatewayListDHTOfferResponse is used to get the FCRMessage of GatewayListDHTOfferResponse
func EncodeGatewayListDHTOfferResponse(
	publishedDHTOffers []FCRMessage,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayListDHTOfferResponse{
		PublishedDHTOffers: publishedDHTOffers,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayListDHTOfferResponseType, body), nil
}

// DecodeGatewayListDHTOfferResponse is used to get the fields from FCRMessage of GatewayListDHTOfferResponse
func DecodeGatewayListDHTOfferResponse(fcrMsg *FCRMessage) (
	[]FCRMessage, // published dht cid offers
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayListDHTOfferResponseType {
		return nil, errors.New("message type mismatch")
	}
	msg := gatewayListDHTOfferResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, err
	}
	return msg.PublishedDHTOffers, nil
}
