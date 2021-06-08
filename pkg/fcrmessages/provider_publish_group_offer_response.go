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

// providerPublishGroupOfferResponse is the response to providerPublishGroupOfferRequest
type providerPublishGroupOfferResponse struct {
	GatewaydID nodeid.NodeID                     `json:"gateway_id"`
	Digest     [cidoffer.CIDOfferDigestSize]byte `json:"digest"`
}

// EncodeProviderPublishGroupOfferResponse is used to get the FCRMessage of ProviderPublishGroupOfferResponse
func EncodeProviderPublishGroupOfferResponse(
	gatewayID nodeid.NodeID,
	digest [cidoffer.CIDOfferDigestSize]byte,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerPublishGroupOfferResponse{
		GatewaydID: gatewayID,
		Digest:     digest,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderPublishGroupOfferResponseType, body), nil
}

// DecodeProviderPublishGroupOfferResponse is used to get the fields from FCRMessage of ProviderPublishGroupOfferResponse
func DecodeProviderPublishGroupOfferResponse(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gatewayID
	[cidoffer.CIDOfferDigestSize]byte, // digest
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderPublishGroupOfferResponseType {
		return nil, [cidoffer.CIDOfferDigestSize]byte{}, errors.New("message type mismatch")
	}
	msg := providerPublishGroupOfferResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, [cidoffer.CIDOfferDigestSize]byte{}, err
	}
	return &msg.GatewaydID, msg.Digest, nil
}
