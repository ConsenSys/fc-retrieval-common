/*
Package fcroffermgr - is dedicated for operations on FileCoin Retrieval Offers, including DHT (Distributed Hash Table structure)
Offers and Group Offers.

Offer is an agreement from a Storage Provider to deliver a file from their storage to a client.
*/
package fcroffermgr

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
	"errors"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
	"github.com/ConsenSys/fc-retrieval-common/pkg/dhtring"
	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
)

// FCROfferMgr manages offer storage
type FCROfferMgr struct {
	dhtOffers    *offerStorage
	dhtOfferRing *dhtring.Ring
	groupOffers  *offerStorage
}

// NewFCROfferMgr returns
func NewFCROfferMgr() *FCROfferMgr {
	return &FCROfferMgr{
		dhtOffers:    newOfferStorage(),
		dhtOfferRing: dhtring.CreateRing(),
		groupOffers:  newOfferStorage(),
	}
}

// AddGroupOffer stores a group offer
func (mgr *FCROfferMgr) AddGroupOffer(offer *cidoffer.CIDOffer) error {
	if len(offer.GetCIDs()) <= 1 {
		return errors.New("Not a group offer")
	}
	return mgr.groupOffers.add(offer)
}

// AddDHTOffer stores a dht offer
func (mgr *FCROfferMgr) AddDHTOffer(offer *cidoffer.CIDOffer) error {
	if len(offer.GetCIDs()) != 1 {
		return errors.New("Not a DHT offer")
	}
	mgr.dhtOfferRing.Insert(offer.GetCIDs()[0].ToString())
	return mgr.dhtOffers.add(offer)
}

// GetGroupOffers returns a list of group offers that contain the given cid
func (mgr *FCROfferMgr) GetGroupOffers(cid *cid.ContentID) ([]cidoffer.CIDOffer, bool) {
	res := mgr.groupOffers.get(cid)
	return res, len(res) > 0
}

// GetDHTOffers returns a list of dht offers that contain the given cid
func (mgr *FCROfferMgr) GetDHTOffers(cid *cid.ContentID) ([]cidoffer.CIDOffer, bool) {
	res := mgr.dhtOffers.get(cid)
	return res, len(res) > 0
}

// GetDHTOffersWithinRange returns a list of dht offers contains a cid within the given range
func (mgr *FCROfferMgr) GetDHTOffersWithinRange(cidMin, cidMax *cid.ContentID, maxOffers int) ([]cidoffer.CIDOffer, bool) {
	offers := make([]cidoffer.CIDOffer, 0)

	entries, err := mgr.dhtOfferRing.GetWithinRange(cidMin.ToString(), cidMax.ToString())
	if err != nil {
		return offers, false
	}

	for _, entry := range entries {
		cid, err := cid.NewContentIDFromHexString(entry)
		if err != nil {
			logging.Error("Internal error")
			return offers, false
		}
		offersTemp := mgr.dhtOffers.get(cid)
		for _, offer := range offersTemp {
			offers = append(offers, offer)
			if len(offers) >= maxOffers {
				return offers, len(offers) > 0
			}
		}
	}

	return offers, len(offers) > 0
}

// GetOffers returns a list of all offers (group or dht) that contain the given cid
func (mgr *FCROfferMgr) GetOffers(cid *cid.ContentID) ([]cidoffer.CIDOffer, bool) {
	res := append(mgr.groupOffers.get(cid), mgr.dhtOffers.get(cid)...)
	return res, len(res) > 0
}

// GetOfferByDigest allows a gateway to be able to respond to a query to search for an offer by the offer digest
func (mgr *FCROfferMgr) GetOfferByDigest(digest [cidoffer.CIDOfferDigestSize]byte) (result *cidoffer.CIDOffer, exist bool) {
	// first, look in DHT offers
	mgr.dhtOffers.lock.RLock()
	defer mgr.dhtOffers.lock.RUnlock()
	for _, digestAndOffer := range mgr.dhtOffers.cidMap {
		digestAndOffer.lock.RLock()
		result, exist = digestAndOffer.dMap[digest]
		digestAndOffer.lock.RUnlock()
		if exist {
			return result, true
		}
	}
	// look in Group offers
	for _, digestAndOffer := range mgr.groupOffers.cidMap {
		digestAndOffer.lock.RLock()
		result, exist = digestAndOffer.dMap[digest]
		digestAndOffer.lock.RUnlock()
		if exist {
			return result, true
		}
	}
	return nil, false
}
