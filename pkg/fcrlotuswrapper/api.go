/*
Package fcrlotuswrapper - is a wrapper API over FileCoin Lotus API
*/
package fcrlotuswrapper

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
	"context"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
)

type NodeApi interface {
	ChainHead(context.Context) (*types.TipSet, error)

	WalletNew(context.Context, types.KeyType) (address.Address, error)

	MpoolPushMessage(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec) (*types.SignedMessage, error)
}

type LotusNodeApi struct {
	closer jsonrpc.ClientCloser

	Methods struct {
		// MethodGroup: Chain
		// The Chain method group contains methods for interacting with the
		// blockchain, but that do not require any form of state computation.

		// ChainHead returns the current head of the chain.
		ChainHead func(context.Context) (*types.TipSet, error)

		// MethodGroup: Wallet

		// WalletNew creates a new address in the wallet with the given sigType.
		// Available key types: bls, secp256k1, secp256k1-ledger
		// Support for numerical types: 1 - secp256k1, 2 - BLS is deprecated
		WalletNew func(context.Context, types.KeyType) (address.Address, error)

		// MethodGroup: Mpool

		// MpoolPushMessage atomically assigns a nonce, signs, and pushes a message
		// to mempool.
		// maxFee is only used when GasFeeCap/GasPremium fields aren't specified
		//
		// When maxFee is set to 0, MpoolPushMessage will guess appropriate fee
		// based on current chain conditions
		MpoolPushMessage func(ctx context.Context, msg *types.Message, spec *api.MessageSendSpec) (*types.SignedMessage, error)
	}
}
