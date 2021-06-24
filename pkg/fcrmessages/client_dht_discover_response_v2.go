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

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

type encodeClientDHTDiscoverResponseV2 struct {
	Contacted            []nodeid.NodeID `json:"contacted"`
	Response             []FCRMessage    `json:"response"`
	UnContactable        []nodeid.NodeID `json:"un_contactable"`
	Nonce                int64           `json:"nonce"`
	FundedPaymentChannel []bool          `json:"funded_payment_channel"`
	PaymentRequired      bool            `json:"payment_required"` // when true means caller have to pay first, using the PaymentChannel field
	PaymentChannel       int64           `json:"payment_channel"`  // payment channel address used in conjunction with PaymentRequired field
}

// EncodeClientDHTDiscoverResponseV2 is used to get the FCRMessage of ClientDHTDiscoverResponse
func EncodeClientDHTDiscoverResponseV2(
	contacted []nodeid.NodeID,
	response []FCRMessage,
	unContactable []nodeid.NodeID,
	nonce int64,
	paymentRequired bool,
	paymentChannel int64,
) (*FCRMessage, error) {
	body, err := json.Marshal(encodeClientDHTDiscoverResponseV2{
		Contacted:       contacted,
		Response:        response,
		UnContactable:   unContactable,
		Nonce:           nonce,
		PaymentRequired: paymentRequired,
		PaymentChannel:  paymentChannel,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientDHTDiscoverResponseV2Type, body), nil
}

// DecodeClientDHTDiscoverResponseV2 is used to get the fields from FCRMessage of ClientDHTDiscoverResponse
func DecodeClientDHTDiscoverResponseV2(fcrMsg *FCRMessage) (
	[]nodeid.NodeID, // contacted
	[]FCRMessage, // response
	[]nodeid.NodeID, // uncontactable
	int64, // nonce
	bool, // paymentRequired
	int64, // paymentChannel
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientDHTDiscoverResponseV2Type {
		return nil, nil, nil, 0, false, 0, errors.New("message type mismatch")
	}
	msg := encodeClientDHTDiscoverResponseV2{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, nil, 0, false, 0, err
	}
	return msg.Contacted, msg.Response, msg.UnContactable, msg.Nonce, msg.PaymentRequired, msg.PaymentChannel, nil
}
