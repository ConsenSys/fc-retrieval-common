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

// providerPublishDHTOfferResponse is the acknowledgement to providerPublishDHTOfferRequest
type providerPublishDHTOfferResponse struct {
	Nonce     int64  `json:"nonce"`
	Signature string `json:"signature"`
}

// EncodeProviderPublishDHTOfferResponse is used to get the FCRMessage of ProviderPublishDHTOfferResponse
func EncodeProviderPublishDHTOfferResponse(
	nonce int64,
	signature string,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerPublishDHTOfferResponse{
		Nonce:     nonce,
		Signature: signature,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderPublishDHTOfferResponseType, body), nil
}

// DecodeProviderPublishDHTOfferResponse is used to get the fields from FCRMessage of ProviderPublishDHTOfferResponse
func DecodeProviderPublishDHTOfferResponse(fcrMsg *FCRMessage) (
	int64, // nonce
	string, // signature
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderPublishDHTOfferResponseType {
		return 0, "", errors.New("message type mismatch")
	}
	msg := providerPublishDHTOfferResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return 0, "", err
	}
	return msg.Nonce, msg.Signature, nil
}
