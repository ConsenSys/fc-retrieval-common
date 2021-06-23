package fcrp2pserver

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
	"net"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
)

// FCRServerWriter stores the connection to write to.
type FCRServerWriter struct {
	conn net.Conn
}

// Write writes a given message.
func (w *FCRServerWriter) Write(msg *fcrmessages.FCRMessage, timeout time.Duration) error {
	return sendTCPMessage(w.conn, msg, timeout)
}

// WriteProtocolChanged writes a protocol changed message.
func (w *FCRServerWriter) WriteProtocolChanged(timeout time.Duration) error {
	fcrMsg, _ := fcrmessages.EncodeProtocolChangeResponse(true)
	return sendTCPMessage(w.conn, fcrMsg, timeout)
}

// WriteProtocolMismatch sends a protocol mistmatch message.
func (w *FCRServerWriter) WriteProtocolMismatch(timeout time.Duration) error {
	fcrMsg, _ := fcrmessages.EncodeProtocolChangeResponse(false)
	return sendTCPMessage(w.conn, fcrMsg, timeout)
}

// WriteInvalidMessage sends an invalid message.
func (w *FCRServerWriter) WriteInvalidMessage(timeout time.Duration) error {
	fcrMsg, _ := fcrmessages.EncodeInvalidMessageResponse()
	return sendTCPMessage(w.conn, fcrMsg, timeout)
}
