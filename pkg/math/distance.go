/*
Package math - contains common mathematical operations like distance calculation.
*/
package math

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
	"math/big"
)

// GetDistance gets the distance between two bytes given a cardinality in big.Int
func GetDistance(a []byte, b []byte, cardinality *big.Int) (*big.Int, error) {
	x := new(big.Int).SetBytes(a)
	y := new(big.Int).SetBytes(b)
	diff1 := big.NewInt(0).Sub(y, x)
	dist1 := big.NewInt(0).Abs(diff1)

	// We need to compute the opposite distance because values can be traversed in both directions
	diff2 := big.NewInt(0).Sub(cardinality, dist1)
	dist2 := big.NewInt(0).Abs(diff2)

	c := dist1.Cmp(dist2)
	if c < 0 {
		return dist1, nil
	} else {
		return dist2, nil
	}
}
