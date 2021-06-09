/*
Package dhtring - provides operations like find a closest node, add new and remove for a Distributed Hash Table Ring data structure
*/
package dhtring

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
	"fmt"
	"math/big"

	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
)

// ringNode is a node inside the ring
type ringNode struct {
	prv     *ringNode
	distPrv *big.Int

	key *big.Int
	val string

	distNext *big.Int
	next     *ringNode
}

// Ring is a struct to store the DHT Ring to store 32-bytes hex
type Ring struct {
	head *ringNode
	size int
}

// CreateRing creates a new ring data structure
func CreateRing() *Ring {
	return &Ring{
		head: nil,
		size: 0,
	}
}

// Insert inserts a hex string into the ring
func (r *Ring) Insert(hex string) {
	if !validateInput(hex) {
		logging.Error("Ring invalid hex: %v", hex)
		return
	}
	hexKey, _ := new(big.Int).SetString(hex, 16)
	newNode := &ringNode{
		prv:      nil,
		distPrv:  nil,
		key:      hexKey,
		val:      hex,
		distNext: nil,
		next:     nil,
	}
	if r.size == 0 {
		r.head = newNode
		r.size++
	}
	// If head is bigger than newNode
	cmp := r.head.key.Cmp(newNode.key)
	if cmp == 0 {
		// Already existed
		return
	}
	if cmp > 0 {
		// Head is bigger than newNode, put between tail and head
		tail := r.head.prv
		if tail == nil {
			tail = r.head
		}
		// Put newNode in front of head
		r.head.prv = newNode
		r.head.distPrv = getDist(newNode.key, r.head.key)
		newNode.next = r.head
		newNode.distNext = getDist(newNode.key, r.head.key)
		// Put newNode behind tail
		tail.next = newNode
		tail.distNext = getDist(tail.key, newNode.key)
		newNode.prv = tail
		newNode.distPrv = getDist(tail.key, newNode.key)
		// Update ring's head to the new node
		r.head = newNode
		r.size++
		return
	}
	// Head is smaller than newNode, start search
	prv := r.head
	current := r.head.next
	for current != nil && current.val != r.head.val {
		cmp = current.key.Cmp(newNode.key)
		if cmp == 0 {
			// Already existed
			return
		}
		if cmp > 0 {
			// Current is bigger than newNode, put between previous and current
			current.prv = newNode
			current.distPrv = getDist(newNode.key, current.key)
			newNode.next = current
			newNode.distNext = getDist(newNode.key, current.key)
			prv.next = newNode
			prv.distNext = getDist(prv.key, newNode.key)
			newNode.prv = prv
			newNode.distPrv = getDist(prv.key, newNode.key)
			r.size++
			return
		}
		// current is smaller than newNode, next
		prv = current
		current = current.next
	}
	if current == nil {
		// If current is nil, it means there is only one node now
		// We update the tail to be the head too.
		current = prv
	}
	// Update tail
	current.prv = newNode
	current.distPrv = getDist(newNode.key, current.key)
	newNode.next = current
	newNode.distNext = getDist(newNode.key, r.head.key)
	prv.next = newNode
	prv.distNext = getDist(prv.key, newNode.key)
	newNode.prv = prv
	newNode.distPrv = getDist(prv.key, newNode.key)
	r.size++
	return
}

// Remove inserts a given hex string out of the ring
func (r *Ring) Remove(hex string) {
	if !validateInput(hex) {
		logging.Error("Ring invalid hex: %v", hex)
		return
	}
	if r.size == 0 {
		return
	}
	if r.size == 1 {
		r.head = nil
		r.size = 0
		return
	}
	if r.size == 2 {
		if r.head.val == hex {
			r.head = r.head.next
			r.head.next = nil
			r.head.prv = nil
			r.head.distPrv = big.NewInt(0)
			r.head.distNext = big.NewInt(0)
			r.size = 1
		} else if r.head.next.val == hex {
			r.head.next = nil
			r.head.prv = nil
			r.head.distPrv = big.NewInt(0)
			r.head.distNext = big.NewInt(0)
			r.size = 1
		}
		return
	}
	if r.head.val == hex {
		oldHead := r.head
		// first becomes the head
		r.head = r.head.next
		// tail links to the new head
		oldHead.prv = r.head
		// update distances
		tailToNewHeadDistance := getDist(r.head.prv.key, r.head.key)
		r.head.distPrv = tailToNewHeadDistance
		r.head.prv.distNext = tailToNewHeadDistance
		// help GC
		oldHead = nil
		r.size--
		return
	}
	hexKey, _ := new(big.Int).SetString(hex, 16)
	current := r.head.next
	for current != nil && current.val != r.head.val {
		// Loop until we reach nil or we go back to head
		if current.val == hex {
			// We need to remove current
			prv := current.prv
			next := current.next
			if prv != nil && next != nil {
				if prv.val == next.val {
					next.prv = nil
					next.distPrv = nil
					next.next = nil
					next.distNext = nil
				} else {
					prv.next = next
					prv.distNext = getDist(prv.key, next.key)
					next.prv = prv
					next.distPrv = getDist(prv.key, next.key)
				}
			}
			if r.head.val == current.val {
				r.head = next
			}
			r.size--
			return
		}
		// Not equal
		if current.key.Cmp(hexKey) > 0 {
			// Return now
			return
		}
		current = current.next
	}
}

// GetClosest gets the closest hexes close to the given hex
func (r *Ring) GetClosest(hex string, num int, exclude string) ([]string, error) {
	if !validateInput(hex) || (exclude != "" && !validateInput(exclude)) {
		logging.Error("Ring invalid hex: %v %v", hex, exclude)
		return nil, errors.New("invalid input")
	}
	res := make([]string, 0)
	if num == 0 || r.size == 0 || (exclude != "" && r.size == 1) {
		return res, nil
	}
	// Consider exclusion
	if exclude != "" {
		before := r.size
		r.Remove(exclude)
		if r.size != before {
			defer r.Insert(exclude)
		}
	}
	if num > r.size {
		if r.head.val == hex {
			res = append(res, r.head.val)
		}
		current := r.head.next
		for current != nil && current.val != r.head.val {
			res = append(res, current.val)
			current = current.next
		}
		return res, nil
	}
	// Insert & Search & Remove
	before := r.size
	r.Insert(hex)
	if r.size == before {
		// This already exists
		res = append(res, hex)
		if num == 1 {
			// Return immediately if only requires one
			return res, nil
		}
	} else {
		defer r.Remove(hex)
	}
	// Now search
	current := r.head
	for ok := true; ok; ok = current.val != r.head.val {
		if current.val == hex {
			// We found it
			break
		}
		current = current.next
	}
	// Now current is the hex inserted
	prv := current.prv
	distToPrv := big.NewInt(0)
	distToPrv.Add(distToPrv, current.distPrv)
	next := current.next
	distToNext := big.NewInt(0)
	distToNext.Add(distToNext, current.distNext)
	for {
		// If prv and next is the same thing
		// Add it and return
		if prv.val == next.val {
			res = append(res, prv.val)
			break
		}
		// fmt.Printf("\n\nprv %v\ndist to prv %v\nnext %v\ndist to next %v\n\n", prv.val, distToPrv, next.val, distToNext)
		cmp := distToPrv.Cmp(distToNext)
		// If equal, we choose the previous one
		if cmp <= 0 {
			res = append([]string{prv.val}, res...)
			distToPrv.Add(distToPrv, prv.distPrv)
			prv = prv.prv
		} else {
			res = append(res, next.val)
			distToNext.Add(distToNext, next.distNext)
			next = next.next
		}
		if len(res) == num {
			// We have enough
			break
		}
	}
	return res, nil
}

// GetWithinRange gets all entries within a range
func (r *Ring) GetWithinRange(startHex string, endHex string) ([]string, error) {
	if !validateInput(startHex) || !validateInput(endHex) {
		logging.Error("Ring invalid hex: %v %v", startHex, endHex)
		return nil, errors.New("invalid input")
	}
	res := make([]string, 0)
	var startNode *ringNode
	var endNode *ringNode
	addStart := false
	addEnd := false
	if temp := r.get(startHex); temp != nil {
		startNode = temp
		addStart = true
	} else {
		r.Insert(startHex)
		startNode = r.get(startHex)
		if startNode == nil {
			return res, errors.New("internal error")
		}
		defer r.Remove(startHex)
	}

	if temp := r.get(endHex); temp != nil {
		endNode = temp
		addEnd = true
	} else {
		r.Insert(endHex)
		endNode = r.get(endHex)
		if endNode == nil {
			return res, errors.New("internal error")
		}
		defer r.Remove(endHex)
	}

	if addStart {
		res = append(res, startHex)
	}

	next := startNode.next
	for next.val != endNode.val {
		res = append(res, next.val)
		next = next.next
	}

	if addEnd {
		res = append(res, endHex)
	}

	return res, nil
}

// Size gets the size of the ring
func (r *Ring) Size() int {
	return r.size
}

// Dump is for debugging use ONLY
func (r *Ring) Dump() {
	fmt.Printf("\nSize: %v [\n", r.size)
	if r.head == nil {
		fmt.Println("]")
		return
	}
	fmt.Println(r.head.val)
	fmt.Printf("\t%v\n", r.head.distPrv)
	fmt.Printf("\t%v\n", r.head.distNext)
	current := r.head.next
	for current != nil && current.val != r.head.val {
		fmt.Println(current.val)
		fmt.Printf("\t%v\n", current.distPrv)
		fmt.Printf("\t%v\n", current.distNext)
		current = current.next
	}
	fmt.Printf("]\n\n")
}

// get gets the ringNode inside this ring, nil if not found
func (r *Ring) get(hex string) *ringNode {
	if !validateInput(hex) {
		logging.Error("Ring invalid hex: %v", hex)
		return nil
	}
	if r.size == 0 {
		return nil
	}
	if r.head.val == hex {
		return r.head
	}
	current := r.head.next
	for current != nil && current.val != r.head.val {
		// Loop until we reach nil or we go back to head
		if current.val == hex {
			return current
		}
		current = current.next
	}
	return nil
}

// getDist gets the distance from one to another, clockwise
func getDist(from *big.Int, to *big.Int) *big.Int {
	// So from is always smaller than to
	if from.Cmp(to) < 0 {
		return big.NewInt(0).Sub(to, from)
	} else {
		// It has across the max/min boundary
		max, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF", 16)
		min, _ := new(big.Int).SetString("0000000000000000000000000000000000000000000000000000000000000000", 16)
		dist1 := big.NewInt(0).Sub(max, from)
		dist2 := big.NewInt(0).Sub(to, min)
		sum := big.NewInt(0).Add(dist1, dist2)
		return sum.Add(sum, big.NewInt(1))
	}
}

// validateInput makes sure the given hex string is 32 bytes hex string
func validateInput(hex string) bool {
	if len(hex) != 64 {
		return false
	}
	for _, char := range hex {
		if (char < '0' || char > '9') && (char < 'A' || char > 'F') && (char < 'a' || char > 'f') {
			return false
		}
	}
	return true
}
