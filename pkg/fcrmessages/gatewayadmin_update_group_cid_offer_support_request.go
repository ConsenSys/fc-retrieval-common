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
	GatewayID              nodeid.NodeID `json:"gateway_id"`
	GroupCIDOfferSupported bool          `json:"group_cid_offer_supported"`
}

func EncodeUpdateGatewayGroupCIDOfferSupportRequest(
	nodeID *nodeid.NodeID,
	groupCIDOfferSupported bool,
) (*FCRMessage, error) {
	body, err := json.Marshal(updateGatewayGroupCIDOfferSupportRequest{
		*nodeID,
		groupCIDOfferSupported,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminUpdateGatewayGroupCIDOfferSupportRequestType, body), nil
}

func DecodeUpdateGatewayGroupCIDOfferSupportRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway id
	bool, // group CID offer is supported by the gateway
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminUpdateGatewayGroupCIDOfferSupportRequestType {
		return nil, false, errors.New("message type mismatch")
	}
	msg := updateGatewayGroupCIDOfferSupportRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, false, err
	}
	return &msg.GatewayID, msg.GroupCIDOfferSupported, nil
}
