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

// protocolChangeRequest message is sent to indicate that the entity is requesting the other entity change protocol version
type protocolChangeRequest struct {
	DesiredVersion int32 `json:"desired_version"`
}

// EncodeProtocolChangeRequest is used to get the FCRMessage of protocolChangeRequest
func EncodeProtocolChangeRequest(
	desiredVersion int32,
) (*FCRMessage, error) {
	body, err := json.Marshal(protocolChangeRequest{
		DesiredVersion: desiredVersion,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProtocolChangeRequestType, body), nil
}

// DecodeProtocolChangeRequest is used to get the fields from FCRMessage of protocolChangeRequest
func DecodeProtocolChangeRequest(fcrMsg *FCRMessage) (
	int32, // desired version
	error, // error
) {
	if fcrMsg.GetMessageType() != ProtocolChangeRequestType {
		return 0, errors.New("message type mismatch")
	}
	msg := protocolChangeRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return 0, err
	}
	return msg.DesiredVersion, nil
}
