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
)

// gatewayDHTDiscoverResponse is the response to gatewayDHTDiscoverRequest
type gatewayDHTDiscoverResponse struct {
	PieceCID             cid.ContentID          `json:"piece_cid"`
	Nonce                int64                  `json:"nonce"`
	Found                bool                   `json:"found"`
	SubCIDOffers         []cidoffer.SubCIDOffer `json:"sub_cid_offers"`
	FundedPaymentChannel []bool                 `json:"funded_payment_channel"`
}

// EncodeGatewayDHTDiscoverResponse is used to get the FCRMessage of gatewayDHTDiscoverResponse
func EncodeGatewayDHTDiscoverResponse(
	pieceCID *cid.ContentID,
	nonce int64,
	found bool,
	offers []cidoffer.SubCIDOffer,
	fundedPaymentChannel []bool,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayDHTDiscoverResponse{
		PieceCID:             *pieceCID,
		Nonce:                nonce,
		Found:                found,
		SubCIDOffers:         offers,
		FundedPaymentChannel: fundedPaymentChannel,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayDHTDiscoverResponseType, body), nil
}

// DecodeGatewayDHTDiscoverResponse is used to get the fields from FCRMessage of GatewayDHTDiscoverResponse
func DecodeGatewayDHTDiscoverResponse(fcrMsg *FCRMessage) (
	*cid.ContentID, // piece cid
	int64, // nonce
	bool, // found
	[]cidoffer.SubCIDOffer, // sub cid offers
	[]bool, // fundedPaymentChannel
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayDHTDiscoverResponseType {
		return nil, 0, false, nil, nil, errors.New("message type mismatch")
	}
	msg := gatewayDHTDiscoverResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, 0, false, nil, nil, err
	}
	return &msg.PieceCID, msg.Nonce, msg.Found, msg.SubCIDOffers, msg.FundedPaymentChannel, nil
}
