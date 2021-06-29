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

// providerAdminGetPublishedOfferRequest is the requset from provideradmin to provider to ask for published  offers for given gateway ids
type providerAdminGetPublishedOfferRequest struct {
	GatewayIDs []string `json:"gateway_id"`
}

// EncodeProviderAdminGetPublishedOfferRequest is used to get the FCRMessage of providerAdminGetPublishedOfferRequest
func EncodeProviderAdminGetPublishedOfferRequest(
	gatewayIDs []nodeid.NodeID,
) (*FCRMessage, error) {
	body, err := json.Marshal(providerAdminGetPublishedOfferRequest{
		GatewayIDs: nodeid.MapNodeIDToString(gatewayIDs),
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderAdminGetPublishedOfferRequestType, body), nil
}

// DecodeProviderAdminGetPublishedOfferRequest is used to get the fields from FCRMessage of providerAdminGetPublishedOfferRequest
func DecodeProviderAdminGetPublishedOfferRequest(fcrMsg *FCRMessage) (
	[]nodeid.NodeID, // piece cids
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderAdminGetPublishedOfferRequestType {
		return nil, errors.New("message type mismatch")
	}
	msg := providerAdminGetPublishedOfferRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, err
	}
	return nodeid.MapStringToNodeID(msg.GatewayIDs), nil
}
