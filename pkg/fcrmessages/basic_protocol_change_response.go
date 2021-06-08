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

// protocolChangeResponse message is response to protocolChangeRequest
type protocolChangeResponse struct {
	Success bool `json:"success"`
}

// EncodeProtocolChangeResponse is used to get the FCRMessage of protocolChangeResponse
func EncodeProtocolChangeResponse(
	success bool,
) (*FCRMessage, error) {
	body, err := json.Marshal(protocolChangeResponse{
		Success: success,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProtocolChangeResponseType, body), nil
}

// DecodeProtocolChangeResponse is used to get the fields from FCRMessage of protocolChangeResponse
func DecodeProtocolChangeResponse(fcrMsg *FCRMessage) (
	bool, // success
	error, // error
) {
	if fcrMsg.GetMessageType() != ProtocolChangeResponseType {
		return false, errors.New("message type mismatch")
	}
	msg := protocolChangeResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return false, err
	}
	return msg.Success, nil
}
