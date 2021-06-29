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
	"fmt"

	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// gatewayPingRequest is the request from gateway to gateway to check if is alive
type gatewayPingRequest struct {
	GatewayID string `json:"gateway_id"`
	Nonce     int64  `json:"nonce"`
	TTL       int64  `json:"ttl"`
}

// EncodeGatewayPingRequest is used to get the FCRMessage of gatewayPingRequest
func EncodeGatewayPingRequest(gatewayID *nodeid.NodeID, nonce, ttl int64) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayPingRequest{
		GatewayID: gatewayID.ToString(),
		Nonce:     nonce,
		TTL:       ttl,
	})
	if err != nil {
		return nil, err
	}

	return CreateFCRMessage(GatewayPingRequestType, body), nil
}

// DecodeGatewayPingRequest is used to get the fields from FCRMessage of gatewayPingRequest
func DecodeGatewayPingRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gateway id
	int64, // nonce
	int64, // ttl
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayPingRequestType {
		return nil, 0, 0, errors.New("message type mismatch")
	}

	msg := gatewayPingRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)

	if err != nil {
		return nil, 0, 0, fmt.Errorf("invalid message: %s", err)
	}

	nodeID, _ := nodeid.NewNodeIDFromHexString(msg.GatewayID)
	return nodeID, msg.Nonce, msg.TTL, nil
}
