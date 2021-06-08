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
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Hexdump provides utility functions to display binary slices as
// hex and printable ASCII.
// From https://github.com/augustoroman/hexdump/blob/master/hexdump.go

// Dump the byte slice to a human-readable hex dump using the default
// configuration.
func dumpMessage(data []byte) string {
	var out bytes.Buffer
	state := dumpState{32, 0, 0}
	state.dump(&out, data)
	return out.String()
}

type dumpState struct {
	width       int
	rowIndex    int
	maxRowWidth int
}

func (s *dumpState) dump(out io.Writer, buf []byte) {
	for i := 0; i*s.width < len(buf); i++ {
		a, b := i*s.width, (i+1)*s.width
		if b > len(buf) {
			b = len(buf)
		}
		row := buf[a:b]
		hex, ascii := printable(row)

		if len(row) < s.maxRowWidth {
			padding := s.maxRowWidth*2 + s.maxRowWidth/4 - len(row)*2 - len(row)/4
			hex += strings.Repeat(" ", padding)
		}
		s.maxRowWidth = len(row)

		if _, err := fmt.Fprintf(out, "%5d: %s | %s\n", s.rowIndex*s.width, hex, ascii); err != nil {
			panic(err)
		}
		s.rowIndex++
	}
}

func printable(data []byte) (hex, ascii string) {
	s := string(data)
	for i := 0; i < len(s); i++ {
		if s[i] < 32 || s[i] >= 127 {
			ascii += "░"
		} else {
			ascii += string(s[i])
		}
		hex += fmt.Sprintf("%02x ", s[i])

		if i%4 == 3 {
			hex += " "
		}
	}
	return hex, ascii
}
