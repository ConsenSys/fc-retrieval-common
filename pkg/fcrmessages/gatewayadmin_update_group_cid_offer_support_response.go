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

type updateGatewayGroupCIDOfferSupportResponse struct {
	Success bool `json:"success"`
}

func EncodeUpdateGatewayGroupCIDOfferSupportResponse(
	success bool,
) (*FCRMessage, error) {
	body, err := json.Marshal(updateGatewayGroupCIDOfferSupportResponse{
		Success: success,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminUpdateGatewayGroupCIDOfferSupportResponseType, body), nil
}

func DecodeUpdateGatewayGroupCIDOfferSupportResponse(fcrMsg *FCRMessage) (
	bool, // success
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminUpdateGatewayGroupCIDOfferSupportResponseType {
		return false, errors.New("message type mismatch")
	}
	msg := updateGatewayGroupCIDOfferSupportResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return false, err
	}
	return msg.Success, nil
}
