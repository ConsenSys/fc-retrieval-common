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

// clientEstablishmentRequest is the request from client to gateway to establish connection
type clientEstablishmentRequest struct {
	ClientID  nodeid.NodeID `json:"client_id"`
	Challenge string        `json:"challenge"`
	TTL       int64         `json:"ttl"`
}

// EncodeClientEstablishmentRequest is used to get the FCRMessage of clientEstablishmentRequest
func EncodeClientEstablishmentRequest(
	clientID *nodeid.NodeID,
	challenge string,
	ttl int64,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientEstablishmentRequest{
		ClientID:  *clientID,
		Challenge: challenge,
		TTL:       ttl,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientEstablishmentRequestType, body), nil
}

// DecodeClientEstablishmentRequest is used to get the fields from FCRMessage of clientEstablishmentRequest
func DecodeClientEstablishmentRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // client id
	string, // challenge
	int64, // ttl
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientEstablishmentRequestType {
		return nil, "", 0, errors.New("message type mismatch")
	}
	msg := clientEstablishmentRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, "", 0, err
	}
	return &msg.ClientID, msg.Challenge, msg.TTL, nil
}
