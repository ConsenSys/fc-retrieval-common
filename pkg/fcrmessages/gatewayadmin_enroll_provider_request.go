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

// gatewayAdminEnrolProviderRequest is the request from an admin client to a gateway to enroll a retrieval provider
type gatewayAdminEnrolProviderRequest struct {
	NodeID             nodeid.NodeID `json:"node_id"`
	Address            string        `json:"address"`
	RootSigningKey     string        `json:"root_signing_key"`
	SigningKey         string        `json:"signing_key"`
	RegionCode         string        `json:"region_code"`
	NetworkInfoGateway string        `json:"network_info_gateway"`
	NetworkInfoClient  string        `json:"network_info_client"`
	NetworkInfoAdmin   string        `json:"network_info_admin"`
}

// EncodeGatewayAdminEnrollProviderRequest is used to get the FCRMessage of gatewayAdminEnrolProviderRequest
func EncodeGatewayAdminEnrollProviderRequest(
	nodeID *nodeid.NodeID,
	address string,
	rootSigningKey string,
	signingKey string,
	regionCode string,
	networkInfoGateway string,
	networkInfoClient string,
	networkInfoAdmin string,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayAdminEnrolProviderRequest{
		NodeID:             *nodeID,
		Address:            address,
		RootSigningKey:     rootSigningKey,
		SigningKey:         signingKey,
		RegionCode:         regionCode,
		NetworkInfoGateway: networkInfoGateway,
		NetworkInfoClient:  networkInfoClient,
		NetworkInfoAdmin:   networkInfoAdmin,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ProviderAdminEnrollProviderRequestType, body), nil
}

// DecodeGatewayAdminEnrollProviderRequest is used to get the fields from FCRMessage of gatewayAdminEnrolProviderRequest
func DecodeGatewayAdminEnrollProviderRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // node id
	string, // address
	string, // root signing key
	string, // signing key
	string, // region code
	string, // network info gateway
	string, // network info client
	string, // network info admin
	error, // error
) {
	if fcrMsg.GetMessageType() != ProviderAdminEnrollProviderRequestType {
		return nil, "", "", "", "", "", "", "", errors.New("message type mismatch")
	}
	msg := gatewayAdminEnrolProviderRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, "", "", "", "", "", "", "", err
	}
	return &msg.NodeID, msg.Address, msg.RootSigningKey, msg.SigningKey, msg.RegionCode,
		msg.NetworkInfoGateway, msg.NetworkInfoClient, msg.NetworkInfoAdmin, nil
}
