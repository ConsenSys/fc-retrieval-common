package fcrmessages

import (
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmerkletree"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

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

// Message types
// The enum should remain the same for client, admin clients,provider and gateway.
const (
	// Client
	ClientEstablishmentRequestType          = 100
	ClientEstablishmentResponseType         = 101
	ClientStandardDiscoverRequestType       = 102
	ClientStandardDiscoverResponseType      = 103
	ClientDHTDiscoverRequestType            = 104
	ClientDHTDiscoverResponseType           = 105
	ClientCIDGroupPublishDHTAckRequestType  = 106
	ClientCIDGroupPublishDHTAckResponseType = 107
	// Provider
	ProviderPublishGroupCIDRequestType    = 200
	ProviderPublishGroupCIDResponseType   = 201
	ProviderDHTPublishGroupCIDRequestType = 202
	ProviderDHTPublishGroupCIDAckType     = 203
	// Gateway
	GatewaySingleCIDOfferPublishRequestType     = 300
	GatewaySingleCIDOfferPublishResponseType    = 301
	GatewaySingleCIDOfferPublishResponseAckType = 302
	GatewayDHTDiscoverRequestType               = 303
	GatewayDHTDiscoverResponseType              = 304
	// GatewayAdmin
	GatewayAdminGetReputationChallengeType = 400
	GatewayAdminGetReputationResponseType  = 401
	GatewayAdminSetReputationChallengeType = 402
	GatewayAdminSetReputationResponseType  = 403
	GatewayAdminAcceptKeyChallengeType     = 404
	GatewayAdminAcceptKeyResponseType      = 405
	// ProviderAdmin
	ProviderAdminGetGroupCIDRequestType        = 500
	ProviderAdminGetGroupCIDResponseType       = 501
	ProviderAdminPublishGroupCIDRequestType    = 502
	ProviderAdminDHTPublishGroupCIDRequestType = 503
	ProviderAdminPublishOfferAckType           = 504
	ProviderAdminAcceptKeyChallengeType        = 505
	// Errors
	InvalidMessageResponseType    = 900
	InsufficientFundsResponseType = 901
	ProtocolChangeResponseType    = 902
	ProtocolMismatchResponseType  = 903
)

// CIDGroupInformation represents a cid group information
type CIDGroupInformation struct {
	ProviderID           nodeid.NodeID                `json:"provider_id"`
	Price                uint64                       `json:"price_per_byte"`
	Expiry               int64                        `json:"expiry_date"`
	QoS                  uint64                       `json:"qos"`
	Signature            string                       `json:"signature"`
	MerkleRoot           string                       `json:"merkle_root"`
	MerkleProof          fcrmerkletree.FCRMerkleProof `json:"merkle_proof"`
	FundedPaymentChannel bool                         `json:"funded_payment_channel"` // TODO: Is this boolean?
}
