/*
Package fcrpaymentmgr - provides common methods facilitate payment process for parties.

Methods like pay, receive payment, top up balance and verify payment.
*/
package fcrpaymentmgr

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
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	"github.com/filecoin-project/go-jsonrpc"
	lotusbig "github.com/filecoin-project/go-state-types/big"
	crypto2 "github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/api/apistruct"
	"github.com/filecoin-project/lotus/chain/actors"
	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/lib/sigs"
	init3 "github.com/filecoin-project/specs-actors/v3/actors/builtin/init"
	"github.com/ipfs/go-cid"
	"github.com/minio/blake2b-simd"

	"github.com/ConsenSys/fc-retrieval-common/pkg/logging"
)

// Payment manager manages all payment related functions
type FCRPaymentMgr struct {
	privKey []byte
	address *address.Address

	authToken    string
	lotusAPIAddr string

	// Channel states.
	// map[recipient addr] -> channel state
	outboundChs     map[string]*channelState
	outboundChsLock sync.RWMutex
	// map[paych addr] -> channel state
	inboundChs     map[string]*channelState
	inboundChsLock sync.RWMutex
}

// channelState represents the state of a channel
type channelState struct {
	addr     address.Address
	balance  big.Int
	redeemed big.Int
	lock     sync.RWMutex

	// Lane States.
	// map[lane id] -> lane state
	laneStates map[uint64]*laneState
}

// laneState represents the state of a lane
type laneState struct {
	nonce    uint64
	redeemed big.Int
	vouchers []string
}

// NewFCRPaymentMgr creates a new payment manager.
func NewFCRPaymentMgr(privateKey, lotusAPIAddr, authToken string) (*FCRPaymentMgr, error) {
	// Register algorithm for signing and verification
	sigs.RegisterSignature(crypto2.SigTypeSecp256k1, SecpSigner{})
	// Get private key and address
	privKey, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}
	pubKey, err := sigs.ToPublic(crypto2.SigTypeSecp256k1, privKey)
	if err != nil {
		return nil, err
	}
	address, err := address.NewSecp256k1Address(pubKey)

	// TODO: Storage on payment channels needs to read from a DB. For now, everytime a new payment channel is called
	// new maps are created with no previous data.
	return &FCRPaymentMgr{
		privKey:         privKey,
		address:         &address,
		authToken:       authToken,
		lotusAPIAddr:    lotusAPIAddr,
		outboundChs:     make(map[string](*channelState)),
		outboundChsLock: sync.RWMutex{},
		inboundChs:      make(map[string](*channelState)),
		inboundChsLock:  sync.RWMutex{}}, nil
}

// Topup will topup a payment channel to recipient with given amount. (amount of value "1" equals 1 coin)
func (mgr *FCRPaymentMgr) Topup(recipient string, amount *big.Int) error {
	// Get recipient address
	recipientAddr, err := address.NewFromString(recipient)
	if err != nil {
		return err
	}
	// Get API
	api, closer, err := getLotusAPI(mgr.authToken, mgr.lotusAPIAddr)
	if err != nil {
		return err
	}
	defer closer()

	// Get channel state
	mgr.outboundChsLock.RLock()
	cs, ok := mgr.outboundChs[recipient]
	if !ok {
		// Need to create a channel
		mgr.outboundChsLock.RUnlock()
		mgr.outboundChsLock.Lock()
		defer mgr.outboundChsLock.Unlock()
		builder := paych.Message(actors.Version3, *mgr.address)
		msg, err := builder.Create(recipientAddr, lotusbig.NewFromGo(amount))
		if err != nil {
			return err
		}
		signedMsg, err := fillMsg(mgr.privKey, api, msg)
		if err != nil {
			return err
		}
		// Send request to lotus
		cid, err := api.MpoolPush(context.Background(), signedMsg)
		if err != nil {
			return err
		}
		receipt := waitReceipt(&cid, api)
		if receipt.ExitCode != 0 {
			logging.Error("Transaction fails to execute: %v", receipt.ExitCode.Error())
			return errors.New("Transaction fails to execute")
		}
		var decodedReturn init3.ExecReturn
		err = decodedReturn.UnmarshalCBOR(bytes.NewReader(receipt.Return))
		if err != nil {
			logging.Error("Payment manager has error unmarshal receipt: %v", receipt)
			return errors.New("Error unmarshal receipt")
		}
		// Create new channel
		mgr.outboundChs[recipient] = &channelState{
			addr:       decodedReturn.RobustAddress,
			balance:    *amount,
			redeemed:   *big.NewInt(0),
			lock:       sync.RWMutex{},
			laneStates: make(map[uint64]*laneState),
		}
	} else {
		// No need to create a channel
		defer mgr.outboundChsLock.RUnlock()
		cs.lock.Lock()
		defer cs.lock.Unlock()
		// Top up msg
		msg := &types.Message{
			To:     cs.addr,
			From:   *mgr.address,
			Value:  lotusbig.NewFromGo(amount),
			Method: 0,
		}
		signedMsg, err := fillMsg(mgr.privKey, api, msg)
		if err != nil {
			return err
		}
		// Send request to lotus
		cid, err := api.MpoolPush(context.Background(), signedMsg)
		if err != nil {
			return err
		}
		receipt := waitReceipt(&cid, api)
		if receipt.ExitCode != 0 {
			return errors.New("Transaction fail to execute")
		}
		// Need to update the balance of this payment channel
		cs.balance.Add(&cs.balance, amount)
	}
	return nil
}

// Pay will generate a voucher and pay the recipient a given amount.
// Return channel address, voucher, true if needs to top up, and error.
func (mgr *FCRPaymentMgr) Pay(recipient string, lane uint64, amount *big.Int) (string, string, bool, error) {
	zero, err := types.ParseFIL("0")
	if err != nil {
		return "", "", false, err
	}
	// Get channel state
	mgr.outboundChsLock.RLock()
	defer mgr.outboundChsLock.RUnlock()
	cs, ok := mgr.outboundChs[recipient]
	if !ok {
		// No existing channel, need to create one
		return "", "", true, nil
	}
	// Lock channel state to make sure only one thread can read/write channel state
	cs.lock.Lock()
	defer cs.lock.Unlock()
	// Check balance
	newRedeemed := big.NewInt(0).Add(&cs.redeemed, amount)
	if cs.balance.Cmp(newRedeemed) < 0 {
		// Balance not enough
		return "", "", true, nil
	}
	// Get lane state
	ls, ok := cs.laneStates[lane]
	if !ok {
		// Lane not existed, create a new lane
		ls = &laneState{
			nonce:    0,
			redeemed: *big.NewInt(0),
			vouchers: make([]string, 0),
		}
		cs.laneStates[lane] = ls
	}
	// Create a voucher
	zero.Add(&ls.redeemed, amount)
	sv := &paych.SignedVoucher{
		ChannelAddr: cs.addr,
		Lane:        lane,
		Nonce:       ls.nonce,
		Amount:      types.BigInt(zero),
	}
	vb, err := sv.SigningBytes()
	if err != nil {
		return "", "", false, err
	}
	sig, err := sigs.Sign(crypto2.SigTypeSecp256k1, mgr.privKey, vb)
	if err != nil {
		return "", "", false, err
	}
	sv.Signature = sig
	voucher, err := encodedVoucher(sv)
	if err != nil {
		return "", "", false, err
	}
	// Voucher created, update lane state
	ls.nonce++
	ls.redeemed.Add(&ls.redeemed, amount)
	ls.vouchers = append(ls.vouchers, voucher)
	// Update channel state
	cs.redeemed.Add(&cs.redeemed, amount)
	return cs.addr.String(), voucher, false, nil
}

// Receive will receive a given voucher at a given payment channel and return the amount received.
// Amount of 1000000000000000000 means 1 coin received.
func (mgr *FCRPaymentMgr) Receive(channel string, voucher string) (*big.Int, error) {
	// TODO: We can query the lane state from the chain via chain get object,
	// but don't know how to interpret the result.
	// It is entirely possible to do validation completely against local storage, just like below.
	// Get channel address
	channelAddr, err := address.NewFromString(channel)
	if err != nil {
		return nil, err
	}
	// Decode voucher
	sv, err := paych.DecodeSignedVoucher(voucher)
	if err != nil {
		return nil, err
	}
	// Get API
	api, closer, err := getLotusAPI(mgr.authToken, mgr.lotusAPIAddr)
	if err != nil {
		return nil, err
	}
	defer closer()

	// Get channel state from the chain
	state, err := api.StateReadState(context.Background(), channelAddr, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	// TODO, Need to make sure it is indeed paych actor state
	paychState, ok := state.State.(map[string]interface{})
	if !ok {
		return nil, errors.New("Not a paych state")
	}

	// Get channel state from local storage
	mgr.inboundChsLock.RLock()
	cs, ok := mgr.inboundChs[channel]
	if !ok {
		// Need to create a entry in local storage
		mgr.inboundChsLock.RUnlock()
		mgr.inboundChsLock.Lock()
		defer mgr.inboundChsLock.Unlock()
		cs = &channelState{
			addr:       channelAddr,
			balance:    *state.Balance.Int,
			redeemed:   *big.NewInt(0),
			lock:       sync.RWMutex{},
			laneStates: make(map[uint64]*laneState),
		}
		mgr.inboundChs[channel] = cs
	} else {
		// Need to update the channel state
		defer mgr.inboundChsLock.RUnlock()
		cs.lock.Lock()
		defer cs.lock.Unlock()
		if cs.balance.Cmp(state.Balance.Int) > 0 {
			// No possible to happen
			return nil, errors.New("On chain state has smaller balance than local chain state")
		} else {
			// Update local channel balance
			cs.balance = *state.Balance.Int
		}
	}

	// Verify recipient
	to, err := address.NewFromString(paychState["To"].(string))
	if err != nil {
		return nil, err
	}
	recipient, err := api.StateAccountKey(context.Background(), to, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	if recipient != *mgr.address {
		return nil, errors.New("Wrong recipient")
	}

	// Verify signature
	f, err := address.NewFromString(paychState["From"].(string))
	if err != nil {
		return nil, err
	}
	pubKey, err := api.StateAccountKey(context.Background(), f, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	vb, err := sv.SigningBytes()
	if err != nil {
		return nil, err
	}
	err = sigs.Verify(sv.Signature, pubKey, vb)
	if err != nil {
		return nil, err
	}

	// Verify lane state
	ls, ok := cs.laneStates[sv.Lane]
	if !ok {
		// Lane not existed, create a new lane
		ls = &laneState{
			nonce:    0,
			redeemed: *big.NewInt(0),
			vouchers: make([]string, 0),
		}
		cs.laneStates[sv.Lane] = ls
	}
	if ls.nonce > sv.Nonce {
		// Nonce not match.
		return nil, errors.New("Nonce is smaller than local stored value")
	}
	if ls.redeemed.Cmp(sv.Amount.Int) >= 0 {
		// Amount in voucher is smaller than redeemed in storage
		return nil, errors.New("Voucher has bad amount")
	}
	paymentValue := big.NewInt(0).Sub(sv.Amount.Int, &ls.redeemed)

	// Verify channel balance
	newRedeemed := big.NewInt(0).Add(&cs.redeemed, paymentValue)
	if cs.balance.Cmp(newRedeemed) < 0 {
		// Channel Balance not enough
		return nil, errors.New("Not enough channel balance")
	}
	// Voucher validated, update lane state
	ls.nonce = sv.Nonce + 1
	ls.redeemed = *sv.Amount.Int
	ls.vouchers = append(ls.vouchers, voucher)
	// Update channel state
	cs.redeemed.Add(&cs.redeemed, paymentValue)
	return paymentValue, nil
}

// Shutdown will safely shutdown the payment manager.
func (mgr *FCRPaymentMgr) Shutdown() {
	// TODO: Need to save the internal storage of the channel state
}

// Dump is used for debugging only, it returns the string repr of payment manager.
// This is not threadsafe and should only be called for debugging.
func (mgr *FCRPaymentMgr) Dump() string {
	var sb strings.Builder
	sb.WriteString("Payment Manager status:\n")
	sb.WriteString(fmt.Sprintf("Private key: %v\n", hex.EncodeToString(mgr.privKey)))
	sb.WriteString(fmt.Sprintf("Adress: %v\n", mgr.address.String()))
	sb.WriteString(fmt.Sprintf("Auth token: %v\n", mgr.authToken))
	sb.WriteString("Outbound channels:\n")
	for recipient, cs := range mgr.outboundChs {
		sb.WriteString(fmt.Sprintf("\t\tChannel address: %v (Recipient address %v)\n", cs.addr.String(), recipient))
		sb.WriteString(fmt.Sprintf("\t\tChannel balance: %v\n", cs.balance.String()))
		sb.WriteString(fmt.Sprintf("\t\tChannel redeemed: %v\n", cs.redeemed.String()))
		sb.WriteString("\t\tLane states:\n")
		for lane, ls := range cs.laneStates {
			sb.WriteString(fmt.Sprintf("\t\t\t\tLane: %v, Nonce: %v, Redeemed: %v\n", lane, ls.nonce, ls.redeemed))
			sb.WriteString(fmt.Sprintf("\t\t\t\tVouchers: "))
			for _, voucher := range ls.vouchers {
				sb.WriteString(voucher)
				sb.WriteString(" ")
			}
			sb.WriteString("\n")
		}
	}
	sb.WriteString("Inbound channels:\n")
	for _, cs := range mgr.inboundChs {
		sb.WriteString(fmt.Sprintf("\t\tChannel address: %v\n", cs.addr.String()))
		sb.WriteString(fmt.Sprintf("\t\tChannel balance: %v\n", cs.balance.String()))
		sb.WriteString(fmt.Sprintf("\t\tChannel redeemed: %v\n", cs.redeemed.String()))
		sb.WriteString("\t\tLane states:\n")
		for lane, ls := range cs.laneStates {
			sb.WriteString(fmt.Sprintf("\t\t\t\tLane: %v, Nonce: %v, Redeemed: %v\n", lane, ls.nonce, ls.redeemed))
			sb.WriteString(fmt.Sprintf("\t\t\t\tVouchers: "))
			for _, voucher := range ls.vouchers {
				sb.WriteString(voucher)
				sb.WriteString(" ")
			}
			sb.WriteString("\n")
		}
	}
	return sb.String()
}

// encodedVoucher returns the encoded string of a given signed voucher
func encodedVoucher(sv *paych.SignedVoucher) (string, error) {
	buf := new(bytes.Buffer)
	if err := sv.MarshalCBOR(buf); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(buf.Bytes()), nil
}

// fillMsg will fill the gas and sign a given message
func fillMsg(privKey []byte, api *apistruct.FullNodeStruct, msg *types.Message) (*types.SignedMessage, error) {
	// Get nonce
	nonce, err := api.MpoolGetNonce(context.Background(), msg.From)
	if err != nil {
		return nil, err
	}
	msg.Nonce = nonce

	// Calculate gas
	limit, err := api.GasEstimateGasLimit(context.Background(), msg, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	msg.GasLimit = int64(float64(limit) * 1.25)

	premium, err := api.GasEstimateGasPremium(context.Background(), 10, msg.From, msg.GasLimit, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	msg.GasPremium = premium

	feeCap, err := api.GasEstimateFeeCap(context.Background(), msg, 20, types.EmptyTSK)
	if err != nil {
		return nil, err
	}
	msg.GasFeeCap = feeCap

	// Sign message
	sig, err := sigs.Sign(crypto2.SigTypeSecp256k1, privKey, msg.Cid().Bytes())
	if err != nil {
		return nil, err
	}
	return &types.SignedMessage{
		Message:   *msg,
		Signature: *sig,
	}, nil
}

// wait receipt will wait until receipt is received for a given cid
func waitReceipt(cid *cid.Cid, api *apistruct.FullNodeStruct) *types.MessageReceipt {
	// Return until recipient is returned (transaction is processed)
	var receipt *types.MessageReceipt
	var err error
	for {
		receipt, err = api.StateGetReceipt(context.Background(), *cid, types.EmptyTSK)
		if err != nil {
			logging.Warn("Payment manager has error getting recipient of cid: %s", cid.String())
		}
		if receipt != nil {
			break
		}
		// TODO, Make the interval configurable
		time.Sleep(5 * time.Second)
	}
	return receipt
}

// get lotus api returns the api that interacts with lotus for a given lotus api addr and access token
func getLotusAPI(authToken, lotusAPIAddr string) (*apistruct.FullNodeStruct, jsonrpc.ClientCloser, error) {
	var api apistruct.FullNodeStruct
	headers := http.Header{"Authorization": []string{"Bearer " + authToken}}
	closer, err := jsonrpc.NewMergeClient(context.Background(), lotusAPIAddr, "Filecoin", []interface{}{&api.Internal, &api.CommonStruct.Internal}, headers)
	if err != nil {
		return nil, nil, err
	}
	return &api, closer, nil
}

// SecpSigner is used to sign, and the following code is taken from lotus source code.
type SecpSigner struct{}

// GenPrivate generates a private key
func (SecpSigner) GenPrivate() ([]byte, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return priv, nil
}

// ToPublic gets the public key of a given private key
func (SecpSigner) ToPublic(pk []byte) ([]byte, error) {
	// Check empty key to avoid segment fault
	if len(pk) == 0 {
		return nil, fmt.Errorf("Unable to get public key, private key is empty")
	}
	return crypto.PublicKey(pk), nil
}

// Sign signs the given msg using given private key
func (SecpSigner) Sign(pk []byte, msg []byte) ([]byte, error) {
	b2sum := blake2b.Sum256(msg)
	sig, err := crypto.Sign(pk, b2sum[:])
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// Verify verifies the given msg, using given signature and given public key
func (SecpSigner) Verify(sig []byte, a address.Address, msg []byte) error {
	b2sum := blake2b.Sum256(msg)
	pubk, err := crypto.EcRecover(b2sum[:], sig)
	if err != nil {
		return err
	}

	maybeaddr, err := address.NewSecp256k1Address(pubk)
	if err != nil {
		return err
	}

	if a != maybeaddr {
		return fmt.Errorf("signature did not match")
	}

	return nil
}
