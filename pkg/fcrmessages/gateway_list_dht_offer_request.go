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

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmerkletree"
	"github.com/ConsenSys/fc-retrieval-common/pkg/nodeid"
)

// gatewayListDHTOfferRequest is the request from gateway to provider during start-up asking for dht offers
type gatewayListDHTOfferRequest struct {
	GatewayID          nodeid.NodeID                `json:"gateway_id"`
	CIDMin             cid.ContentID                `json:"cid_min"`
	CIDMax             cid.ContentID                `json:"cid_max"`
	BlockHash          string                       `json:"block_hash"`
	TransactionReceipt string                       `json:"transaction_receipt"`
	MerkleRoot         string                       `json:"merkle_root"`
	MerkleProof        fcrmerkletree.FCRMerkleProof `json:"merkle_proof"`
}

// EncodeGatewayListDHTOfferRequest is used to get the FCRMessage of gatewayListDHTOfferRequest
func EncodeGatewayListDHTOfferRequest(
	gatewayID *nodeid.NodeID,
	cidMin *cid.ContentID,
	cidMax *cid.ContentID,
	blockHash string,
	transactionReceipt string,
	merkleRoot string,
	merkleProof *fcrmerkletree.FCRMerkleProof,
) (*FCRMessage, error) {
	body, err := json.Marshal(gatewayListDHTOfferRequest{
		GatewayID:          *gatewayID,
		CIDMin:             *cidMin,
		CIDMax:             *cidMax,
		BlockHash:          blockHash,
		TransactionReceipt: transactionReceipt,
		MerkleRoot:         merkleRoot,
		MerkleProof:        *merkleProof,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(GatewayListDHTOfferRequestType, body), nil
}

// DecodeGatewayListDHTOfferRequest is used to get the fields from FCRMessage of gatewayListDHTOfferRequest
func DecodeGatewayListDHTOfferRequest(fcrMsg *FCRMessage) (
	*nodeid.NodeID, // gatewayID
	*cid.ContentID, // cid min
	*cid.ContentID, // cid max
	string, // block hash
	string, // transaction receipt
	string, // merkle root
	*fcrmerkletree.FCRMerkleProof, // merkle proof
	error, // error
) {
	if fcrMsg.GetMessageType() != GatewayListDHTOfferRequestType {
		return nil, nil, nil, "", "", "", nil, errors.New("Message type mismatch")
	}
	msg := gatewayListDHTOfferRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, nil, nil, "", "", "", nil, err
	}
	return &msg.GatewayID, &msg.CIDMin, &msg.CIDMax, msg.BlockHash, msg.TransactionReceipt, msg.MerkleRoot, &msg.MerkleProof, nil
}
