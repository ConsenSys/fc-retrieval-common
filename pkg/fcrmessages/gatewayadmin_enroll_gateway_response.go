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

// gatewayAdminEnrollGatewayResponse is the response to gatewayAdminEnrollGatewayRequest
type gatewayAdminEnrollGatewayResponse struct {
	Enrolled bool `json:"enrolled"`
}

// EncodeGatewayAdminEnrollGatewayResponse is used to get the FCRMessage of gatewayAdminEnrollGatewayResponse
func EncodeGatewayAdminEnrollGatewayResponse(
	enrolled bool,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayAdminEnrollGatewayResponse{
		Enrolled: enrolled,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminEnrollGatewayResponseType, body), nil
}

// DecodeGatewayAdminEnrollGatewayResponse is used to get the fields from FCRMessage of gatewayAdminEnrollGatewayResponse
func DecodeGatewayAdminEnrollGatewayResponse(fcrMsg *FCRMessage) (
	bool,  // enrolled
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminEnrollGatewayResponseType {
		return false, errors.New("message type mismatch")
	}
	msg := gatewayAdminEnrollGatewayResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return false, err
	}
	return msg.Enrolled, nil
}
