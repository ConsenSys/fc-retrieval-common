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

// clientEstablishmentResponse is the response to clientEstablishmentRequest
type clientEstablishmentResponse struct {
	GatewayID nodeid.NodeID `json:"gateway_id"`
	Challenge string        `json:"challenge"`
}

// EncodeClientEstablishmentResponse is used to get the FCRMessage of clientEstablishmentResponse
func EncodeClientEstablishmentResponse(
	gatewayID *nodeid.NodeID,
	challenge string,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientEstablishmentResponse{
		GatewayID: *gatewayID,
		Challenge: challenge,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientEstablishmentResponseType, body), nil
}

// DecodeClientEstablishmentResponse is used to get the fields from FCRMessage of ClientEstablishmentResponse
func DecodeClientEstablishmentResponse(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway id
	string, // challenge
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientEstablishmentResponseType {
		return nil, "", errors.New("message type mismatch")
	}
	msg := clientEstablishmentResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, "", err
	}
	return &msg.GatewayID, msg.Challenge, nil
}
