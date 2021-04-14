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

// Message types
const (
	// Message originating from client
	ClientEstablishmentRequestType     = 100
	ClientEstablishmentResponseType    = 101
	ClientStandardDiscoverRequestType  = 102
	ClientStandardDiscoverResponseType = 103
	ClientDHTDiscoverRequestType       = 104
	ClientDHTDiscoverResponseType      = 105
	ClientDHTOfferAckRequestType       = 106
	ClientDHTOfferAckResponseType      = 107

	// Message originating from gateway
	GatewayListDHTOfferRequestType  = 200
	GatewayListDHTOfferResponseType = 201
	GatewayListDHTOfferAckType      = 202
	GatewayDHTDiscoverRequestType   = 203
	GatewayDHTDiscoverResponseType  = 204

	// Message originating from provider
	ProviderPublishGroupOfferRequestType  = 300
	ProviderPublishGroupOfferResponseType = 301
	ProviderPublishDHTOfferRequestType    = 302
	ProviderPublishDHTOfferResponseType   = 303

	// Message originating from gateway admin
	GatewayAdminInitialiseKeyRequestType   = 400
	GatewayAdminInitialiseKeyResponseType  = 401
	GatewayAdminGetReputationRequestType   = 402
	GatewayAdminGetReputationResponseType  = 403
	GatewayAdminSetReputationRequestType   = 404
	GatewayAdminSetReputationResponseType  = 405
	GatewayAdminEnrollGatewayRequestType   = 406
	GatewayAdminEnrollGatewayResponseType  = 407
	GatewayAdminEnrollProviderRequestType  = 408
	GatewayAdminEnrollProviderResponseType = 409

	// Message originating from provider admin
	ProviderAdminInitialiseKeyRequestType      = 500
	ProviderAdminInitialiseKeyResponseType     = 501
	ProviderAdminPublishGroupOfferRequestType  = 502
	ProviderAdminPublishGroupOfferResponseType = 503
	ProviderAdminPublishDHTOfferRequestType    = 504
	ProviderAdminPublishDHTOfferResponseType   = 505
	ProviderAdminGetPublishedOfferRequestType  = 506
	ProviderAdminGetPublishedOfferResponseType = 507
	ProviderAdminEnrollGatewayRequestType      = 508
	ProviderAdminEnrollGatewayResponseType     = 509

	// Messages for basic protocol
	ProtocolChangeRequestType     = 900
	ProtocolChangeResponseType    = 901
	InvalidMessageResponseType    = 902
	InsufficientFundsResponseType = 903
)
