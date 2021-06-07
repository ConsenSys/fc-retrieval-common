/*
Package fcrmessages - contains all RPC messages exist in FileCoin Secondary Retrieval Market network and used to
communicate between the following types of nodes: Retrieval Providers, Retrieval Providers Admins, Retrieval Gateways,
Retrieval Gateways Admins, Retrieval Registers and Retrieval Clients.

All the messages could be split in two group: request and response.
*/
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

// insufficientFundsResponse message is sent to indicate the there are not sufficient funds to finish the request
type insufficientFundsResponse struct {
	PaymentChannelID int64 `json:"payment_channel_id"`
}

// EncodeInsufficientFundsResponse is used to get the FCRMessage of insufficientFundsResponse
func EncodeInsufficientFundsResponse(paymentChannelID int64) (*FCRMessage, error) {
	body, err := json.Marshal(insufficientFundsResponse{
		PaymentChannelID: paymentChannelID,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(InsufficientFundsResponseType, body), nil
}

// DecodeInsufficientFundsResponse is used to get the fields from FCRMessage of insufficientFundsResponse
func DecodeInsufficientFundsResponse(fcrMsg *FCRMessage) (
	int64, // payment channel id
	error, // error
) {
	if fcrMsg.GetMessageType() != InsufficientFundsResponseType {
		return 0, errors.New("Message type mismatch")
	}
	msg := insufficientFundsResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return 0, err
	}
	return msg.PaymentChannelID, nil
}
