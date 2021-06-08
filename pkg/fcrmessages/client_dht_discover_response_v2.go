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

// EncodeClientDHTDiscoverResponseV2 is used to get the FCRMessage of ClientDHTDiscoverResponse
func EncodeClientDHTDiscoverResponseV2(
	contacted []nodeid.NodeID,
	response []FCRMessage,
	unContactable []nodeid.NodeID,
	nonce int64,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientDHTDiscoverResponse{
		Contacted:     contacted,
		Response:      response,
		UnContactable: unContactable,
		Nonce:         nonce,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientDHTDiscoverResponseV2Type, body), nil
}

// DecodeClientDHTDiscoverResponseV2 is used to get the fields from FCRMessage of ClientDHTDiscoverResponse
func DecodeClientDHTDiscoverResponseV2(fcrMsg *FCRMessage) (
	[]nodeid.NodeID, // contacted
	[]FCRMessage, // response
	[]nodeid.NodeID, // uncontactable
	int64, // nonce
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientDHTDiscoverResponseV2Type {
		return nil, nil, nil, 0, errors.New("message type mismatch")
	}
	msg := clientDHTDiscoverResponse{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, nil, 0, err
	}
	return msg.Contacted, msg.Response, msg.UnContactable, msg.Nonce, nil
}
