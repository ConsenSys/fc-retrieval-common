package fcrmessages

import (
	"encoding/json"
	"errors"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrcrypto"
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

const (
	defaultProtocolVersion            = 1
	defaultAlternativeProtocolVersion = 1
)

var protocolVersion int32 = defaultProtocolVersion
var protocolSupported []int32 = []int32{defaultProtocolVersion, defaultAlternativeProtocolVersion}

// FCRMessage is the message used in communication between filecoin retrieval entities.
type FCRMessage struct {
	messageType       int32
	protocolVersion   int32
	protocolSupported []int32
	messageBody       []byte
	signature         string
}

// fcrMessageJson is used to parse to and from json.
type fcrMessageJson struct {
	MessageType       int32   `json:"message_type"`
	ProtocolVersion   int32   `json:"protocol_version"`
	ProtocolSupported []int32 `json:"protocol_supported"`
	MessageBody       []byte  `json:"message_body"`
	Signature         string  `json:"message_signature"`
}

// CreateFCRMessage is used to create an unsigned message
func CreateFCRMessage(msgType int32, msgBody []byte) *FCRMessage {
	return &FCRMessage{
		messageType:       msgType,
		messageBody:       msgBody,
		protocolVersion:   protocolVersion,
		protocolSupported: protocolSupported,
		signature:         "",
	}
}

// GetMessageType is used to get the message type of the message.
func (fcrMsg *FCRMessage) GetMessageType() int32 {
	return fcrMsg.messageType
}

// GetProtocolVersion is used to get the protocol version of peer.
func (fcrMsg *FCRMessage) GetProtocolVersion() int32 {
	return fcrMsg.protocolVersion
}

// GetProtocolSupported is used to get the protocol supported of peer.
func (fcrMsg *FCRMessage) GetProtocolSupported() []int32 {
	return fcrMsg.protocolSupported
}

// GetMessageBody is used to get the message body.
func (fcrMsg *FCRMessage) GetMessageBody() []byte {
	return fcrMsg.messageBody
}

// GetSignature is used to get the signature.
func (fcrMsg *FCRMessage) GetSignature() string {
	return fcrMsg.signature
}

// Sign is used to sign the message with a given private key and a key version.
func (fcrMsg *FCRMessage) Sign(privKey *fcrcrypto.KeyPair, keyVer *fcrcrypto.KeyVersion) error {
	// Clear signature
	fcrMsg.signature = ""
	sig, err := fcrcrypto.SignMessage(privKey, keyVer, fcrMsg)
	if err != nil {
		return err
	}
	fcrMsg.signature = sig
	return nil
}

// Verify is used to verify the offer with a given public key.
func (fcrMsg *FCRMessage) Verify(pubKey *fcrcrypto.KeyPair) error {
	// Clear signature
	sig := fcrMsg.signature
	fcrMsg.signature = ""
	res, err := fcrcrypto.VerifyMessage(pubKey, sig, fcrMsg)
	if err != nil {
		return err
	}
	if !res {
		return errors.New("message does not pass signature verification")
	}
	// Recover signature
	fcrMsg.signature = sig
	return nil
}

// FCRMsgToBytes converts a FCRMessage to bytes
func (fcrMsg *FCRMessage) FCRMsgToBytes() ([]byte, error) {
	return json.Marshal(fcrMsg)
}

// FCRMsgFromBytes converts a bytes to FCRMessage
func FCRMsgFromBytes(data []byte) (*FCRMessage, error) {
	res := FCRMessage{}
	err := json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// MarshalJSON is used to marshal offer into bytes.
func (fcrMsg FCRMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(fcrMessageJson{
		MessageType:       fcrMsg.messageType,
		ProtocolVersion:   fcrMsg.protocolVersion,
		ProtocolSupported: fcrMsg.protocolSupported,
		MessageBody:       fcrMsg.messageBody,
		Signature:         fcrMsg.signature,
	})
}

// UnmarshalJSON is used to unmarshal bytes into offer.
func (fcrMsg *FCRMessage) UnmarshalJSON(p []byte) error {
	msgJson := fcrMessageJson{}
	err := json.Unmarshal(p, &msgJson)
	if err != nil {
		return err
	}
	fcrMsg.messageType = msgJson.MessageType
	fcrMsg.protocolVersion = msgJson.ProtocolVersion
	fcrMsg.protocolSupported = msgJson.ProtocolSupported
	fcrMsg.messageBody = msgJson.MessageBody
	fcrMsg.signature = msgJson.Signature
	return nil
}

// GetCurrentProtocolVersion gets the current protocol version of all messages.
func GetCurrentProtocolVersion() (int32, []int32) {
	return protocolVersion, protocolSupported
}

// SetCurrentProtocolVersion sets the current protocol version of all messages.
func SetCurrentProtocolVersion(newProtocolVersion int32, newProtocolSupported []int32) {
	protocolVersion = newProtocolVersion
	protocolSupported = newProtocolSupported
}

// DumpMessage return a message with ASCII and hex.
func (fcrMsg *FCRMessage) DumpMessage() string {
	msgBytes, err := fcrMsg.MarshalJSON()
	if err != nil {
		return "Error processing message"
	}
	return dumpMessage(msgBytes)
}
