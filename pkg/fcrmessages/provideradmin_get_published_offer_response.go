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
)

// providerAdminGetPublishedOfferResponse is the response to providerAdminGetPublishedOfferRequest
type providerAdminGetPublishedOfferResponse struct {
	Exists bool                `json:"exists"`
	Offers []cidoffer.CIDOffer `json:"cid_offers"`
}

// EncodeProviderAdminGetPublishedOfferResponse is used to get the FCRMessage of providerAdminGetPublishedOfferResponse
func EncodeProviderAdminGetPublishedOfferResponse(
	exists bool,
	offers []cidoffer.CIDOffer,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerAdminGetPublishedOfferResponse{
		Exists: exists,
		Offers: offers,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderAdminGetPublishedOfferResponseType, body), nil
}

// DecodeProviderAdminGetPublishedOfferResponse is used to get the fields from FCRMessage of providerAdminGetPublishedOfferResponse
func DecodeProviderAdminGetPublishedOfferResponse(fcrMsg *FCRMessage) (
	bool, // exists
	[]cidoffer.CIDOffer, // cid offers
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderAdminGetPublishedOfferResponseType {
		return false, nil, errors.New("message type mismatch")
	}
	msg := providerAdminGetPublishedOfferResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return false, nil, err
	}
	return msg.Exists, msg.Offers, nil
}
