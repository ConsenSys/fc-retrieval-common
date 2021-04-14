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

// gatewayAdminEnrollGatewayRequest is the request from an admin client to a gateway to enroll a retrieval gateway
type gatewayAdminEnrollGatewayRequest struct {
	NodeID              nodeid.NodeID `json:"node_id"`
	Address             string        `json:"address"`
	RootSigningKey      string        `json:"root_signing_key"`
	SigningKey          string        `json:"signing_key"`
	RegionCode          string        `json:"region_code"`
	NetworkInfoGateway  string        `json:"network_info_gateway"`
	NetworkInfoProvider string        `json:"network_info_provider"`
	NetworkInfoClient   string        `json:"network_info_client"`
	NetworkInfoAdmin    string        `json:"network_info_admin"`
}

// EncodeGatewayAdminEnrollGatewayRequest is used to get the FCRMessage of gatewayAdminEnrollGatewayRequest
func EncodeGatewayAdminEnrollGatewayRequest(
	nodeID *nodeid.NodeID,
	address string,
	rootSigningKey string,
	signingKey string,
	regionCode string,
	networkInfoGateway string,
	networkInfoProvider string,
	networkInfoClient string,
	networkInfoAdmin string,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayAdminEnrollGatewayRequest{
		NodeID:              *nodeID,
		Address:             address,
		RootSigningKey:      rootSigningKey,
		SigningKey:          signingKey,
		RegionCode:          regionCode,
		NetworkInfoGateway:  networkInfoGateway,
		NetworkInfoProvider: networkInfoProvider,
		NetworkInfoClient:   networkInfoClient,
		NetworkInfoAdmin:    networkInfoAdmin,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayAdminEnrollGatewayRequestType, body), nil
}

// DecodeGatewayAdminEnrollGatewayRequest is used to get the fields from FCRMessage of gatewayAdminEnrollGatewayRequest
func DecodeGatewayAdminEnrollGatewayRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // node id
	string, // address
	string, // root signing key
	string, // signing key
	string, // region code
	string, // network info gateway
	string, // network info provider
	string, // network info client
	string, // network info admin
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayAdminEnrollGatewayRequestType {
		return nil, "", "", "", "", "", "", "", "", errors.New("message type mismatch")
	}
	msg := gatewayAdminEnrollGatewayRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, "", "", "", "", "", "", "", "", err
	}
	return &msg.NodeID, msg.Address, msg.RootSigningKey, msg.SigningKey, msg.RegionCode,
		msg.NetworkInfoGateway, msg.NetworkInfoProvider, msg.NetworkInfoClient, msg.NetworkInfoAdmin, nil
}
