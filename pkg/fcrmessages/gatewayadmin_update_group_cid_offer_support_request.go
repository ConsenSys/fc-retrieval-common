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

type updateGatewayGroupCIDOfferSupportRequest struct {
	GatewayID   nodeid.NodeID   `json:"gateway_id"`
	ProviderIDs []nodeid.NodeID `json:"provider_ids"` // the Gateway supports Group CID Offers only from these Providers
}

func EncodeUpdateGatewayGroupCIDOfferSupportRequest(
	nodeID *nodeid.NodeID,
	providerIDs []nodeid.NodeID,
) (*FCRMessage, error) {
	body, err := json.Marshal(updateGatewayGroupCIDOfferSupportRequest{
		*nodeID,
		providerIDs,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminUpdateGatewayGroupCIDOfferSupportRequestType, body), nil
}

func DecodeUpdateGatewayGroupCIDOfferSupportRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway ID
	[]nodeid.NodeID, // providers IDs
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminUpdateGatewayGroupCIDOfferSupportRequestType {
		return nil, nil, errors.New("message type mismatch")
	}
	msg := updateGatewayGroupCIDOfferSupportRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, err
	}
	return &msg.GatewayID, msg.ProviderIDs, nil
}
