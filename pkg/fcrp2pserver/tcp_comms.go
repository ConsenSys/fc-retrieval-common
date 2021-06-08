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
	"bufio"
	"encoding/binary"
	"io"
	"net"
	"time"

	"github.com/ConsenSys/fc-retrieval-common/pkg/fcrmessages"
)

// isTimeoutError checks if the given error is a timeout error
func isTimeoutError(err error) bool {
	neterr, ok := err.(net.Error)
	return ok && neterr.Timeout()
}

// readTCPMessage read the tcp message from a given connection.
func readTCPMessage(conn net.Conn, timeout time.Duration) (*fcrmessages.FCRMessage, error) {
	// Initialise a reader
	reader := bufio.NewReader(conn)
	// Read the length
	length := make([]byte, 4)
	// Set timeout
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		panic(err)
	}
	_, err := io.ReadFull(reader, length)
	if err != nil {
		return nil, err
	}
	// Read the data
	data := make([]byte, int(binary.BigEndian.Uint32(length)))
	// Set timeout
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		panic(err)
	}
	_, err = io.ReadFull(reader, data)
	if err != nil {
		return nil, err
	}
	return fcrmessages.FCRMsgFromBytes(data)
}

// sendTCPMessage sends a tcp message to a given connection.
func sendTCPMessage(conn net.Conn, fcrMsg *fcrmessages.FCRMessage, timeout time.Duration) error {
	// Get data
	data, err := fcrMsg.FCRMsgToBytes()
	if err != nil {
		return err
	}
	// Initialise a writer
	writer := bufio.NewWriter(conn)
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(data)))
	// Set timeout
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		panic(err)
	}
	_, err = writer.Write(append(length, data...))
	if err != nil {
		return err
	}
	// Set timeout
	if err := conn.SetDeadline(time.Now().Add(timeout)); err != nil {
		panic(err)
	}
	return writer.Flush()
}

// sendProtocolChanged sends a protocol changed message to a given connection.
func sendProtocolChanged(conn net.Conn, timeout time.Duration) error {
	fcrMsg, _ := fcrmessages.EncodeProtocolChangeResponse(true)
	return sendTCPMessage(conn, fcrMsg, timeout)
}

// sendProtocolMismatch sends a protocol mistmatch message to a given connection.
func sendProtocolMismatch(conn net.Conn, timeout time.Duration) error {
	fcrMsg, _ := fcrmessages.EncodeProtocolChangeResponse(false)
	return sendTCPMessage(conn, fcrMsg, timeout)
}

// sendInvalidMessage sends an invalid message to a given connection.
func sendInvalidMessage(conn net.Conn, timeout time.Duration) error {
	fcrMsg, _ := fcrmessages.EncodeInvalidMessageResponse()
	return sendTCPMessage(conn, fcrMsg, timeout)
}

// sendInsufficientFunds sends a insufficient payment message to a given connection.
func sendInsufficientFunds(conn net.Conn, timeout time.Duration, paymentChannelID int64) error {
	fcrMsg, _ := fcrmessages.EncodeInsufficientFundsResponse(paymentChannelID)
	return sendTCPMessage(conn, fcrMsg, timeout)
}
