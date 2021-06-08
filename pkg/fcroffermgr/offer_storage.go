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
	"sync"

	"github.com/ConsenSys/fc-retrieval-common/pkg/cid"
	"github.com/ConsenSys/fc-retrieval-common/pkg/cidoffer"
)

// digestOffer stores map from digest -> offer
type digestOffer struct {
	dMap map[[cidoffer.CIDOfferDigestSize]byte]*cidoffer.CIDOffer
	lock sync.RWMutex
}

// offerStorage stores all cid offers.
type offerStorage struct {
	// map from cid -> (digest, offer)
	cidMap map[string]*digestOffer
	lock   sync.RWMutex
}

// newOfferStorage create and return a new offers instance
func newOfferStorage() *offerStorage {
	return &offerStorage{
		cidMap: make(map[string]*digestOffer),
		lock:   sync.RWMutex{},
	}
}

// add adds a new cid offer to the offers and Ignore the offer if it already exists.
func (o *offerStorage) add(newOffer *cidoffer.CIDOffer) error {
	// If this offer exists
	testCIDs := newOffer.GetCIDs()
	if len(testCIDs) == 0 {
		return errors.New("this offer has no cid")
	}
	testCIDStr := testCIDs[0].ToString()
	digest := newOffer.GetMessageDigest()
	o.lock.RLock()
	digestMap, exists := o.cidMap[testCIDStr]
	o.lock.RUnlock()
	if exists {
		digestMap.lock.RLock()
		_, exists = digestMap.dMap[digest]
		digestMap.lock.RUnlock()
		if exists {
			// This offer is already in the system.
			return nil
		}
	}

	if newOffer.HasExpired() {
		return errors.New("offers: Attempt to add an expired offer")
	}

	for _, contentID := range newOffer.GetCIDs() {
		cidStr := contentID.ToString()
		o.lock.RLock()
		digestMap, exists = o.cidMap[cidStr]
		o.lock.RUnlock()
		if !exists {
			// Need a new entry in cid map
			o.lock.Lock()
			digestMap = &digestOffer{
				dMap: make(map[[cidoffer.CIDOfferDigestSize]byte]*cidoffer.CIDOffer),
				lock: sync.RWMutex{},
			}
			o.cidMap[cidStr] = digestMap
			o.lock.Unlock()
		}
		// Add offer to digest map
		digestMap.lock.Lock()
		digestMap.dMap[digest] = newOffer
		digestMap.lock.Unlock()
	}
	return nil
}

// get returns a list of offers that contains piece cid. It only returns offers that are yet expired.
func (o *offerStorage) get(cid *cid.ContentID) []cidoffer.CIDOffer {
	res := make([]cidoffer.CIDOffer, 0)
	cidStr := cid.ToString()

	o.lock.RLock()
	digestMap, exists := o.cidMap[cidStr]
	o.lock.RUnlock()
	if !exists {
		return res
	}
	toRemove := make([]cidoffer.CIDOffer, 0)
	digestMap.lock.RLock()
	for _, offer := range digestMap.dMap {
		if offer.HasExpired() {
			toRemove = append(toRemove, *offer)
		} else {
			res = append(res, *offer)
		}
	}
	digestMap.lock.RUnlock()

	// Before return the result
	// Remove all expired offer
	o.lock.RLock()
	for _, offer := range toRemove {
		digest := offer.GetMessageDigest()
		for _, contentID := range offer.GetCIDs() {
			digestMap, exists := o.cidMap[contentID.ToString()]
			if !exists {
				// Something is wrong
				panic("Internal error: Try to remove cid offer that are not existed in cid map")
			}
			digestMap.lock.Lock()
			delete(digestMap.dMap, digest)
			digestMap.lock.Unlock()
			// After this deletion, it is possible this cid entry has no offers.
			// This shouldn't be a problem, since this cid entry will be used in the future.
		}
	}
	o.lock.RUnlock()
	return res
}
