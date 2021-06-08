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
)

// clientDHTDiscoverRequest is the request from client to gateway to ask for cid offer using DHT
type clientDHTDiscoverRequest struct {
	PieceCID           cid.ContentID `json:"piece_cid"`
	Nonce              int64         `json:"nonce"`
	TTL                int64         `json:"ttl"`
	NumDHT             int64         `json:"num_dht"`
	IncrementalResults bool          `json:"incremental_results"`
	PaychAddr          string        `json:"payment_channel_address"`
	Voucher            string        `json:"voucher"`
}

// EncodeClientDHTDiscoverRequest is used to get the FCRMessage of clientDHTDiscoverRequest
func EncodeClientDHTDiscoverRequest(
	pieceCID *cid.ContentID,
	nonce int64,
	ttl int64,
	numDHT int64,
	incrementalResults bool,
	paychAddr string,
	voucher string,
) (*FCRMessage, error) {
	body, err := json.Marshal(clientDHTDiscoverRequest{
		PieceCID:           *pieceCID,
		Nonce:              nonce,
		TTL:                ttl,
		NumDHT:             numDHT,
		IncrementalResults: incrementalResults,
		PaychAddr:          paychAddr,
		Voucher:            voucher,
	})
	if err != nil {
		return nil, err
	}
	return CreateFCRMessage(ClientDHTDiscoverRequestType, body), nil
}

// DecodeClientDHTDiscoverRequest is used to get the fields from FCRMessage of clientDHTDiscoverRequest
func DecodeClientDHTDiscoverRequest(fcrMsg *FCRMessage) (
	*cid.ContentID, // piece cid
	int64, // nonce
	int64, // ttl
	int64, // num dht
	bool, // incremental results
	string, // payment channel address
	string, // voucher
	error, // error
) {
	if fcrMsg.GetMessageType() != ClientDHTDiscoverRequestType {
		return nil, 0, 0, 0, false, "", "", errors.New("message type mismatch")
	}
	msg := clientDHTDiscoverRequest{}
	err := json.Unmarshal(fcrMsg.GetMessageBody(), &msg)
	if err != nil {
		return nil, 0, 0, 0, false, "", "", err
	}
	return &msg.PieceCID, msg.Nonce, msg.TTL, msg.NumDHT, msg.IncrementalResults, msg.PaychAddr, msg.Voucher, nil
}
