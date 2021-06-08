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

// providerAdminPublishGroupOfferResponse is the response to providerAdminPublishGroupOfferRequest
type providerAdminPublishGroupOfferResponse struct {
	Received bool `json:"received"`
}

// EncodeProviderAdminPublishGroupOfferResponse is used to get the FCRMessage of providerAdminPublishGroupOfferResponse
func EncodeProviderAdminPublishGroupOfferResponse(
	received bool,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerAdminPublishGroupOfferResponse{
		Received: received,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderAdminPublishGroupOfferResponseType, body), nil
}

// DecodeProviderAdminPublishGroupOfferResponse is used to get the fields from FCRMessage of providerAdminPublishGroupOfferResponse
func DecodeProviderAdminPublishGroupOfferResponse(fcrMsg *FCRMessage) (
	bool, // received
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderAdminPublishGroupOfferResponseType {
		return false, errors.New("message type mismatch")
	}
	msg := providerAdminPublishGroupOfferResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return false, err
	}
	return msg.Received, nil
}
