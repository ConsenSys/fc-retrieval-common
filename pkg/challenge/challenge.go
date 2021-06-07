/*
Package challenge - is used in tests to generate fake challenges.

The challenge is a 32 byte long string sent by a caller to a receiver.
The caller expects to receive the string back.
*/
package challenge

import (
	"crypto/rand"
	"encoding/base64"
)

// NewRandomChallenge generates and returns 32 byte long, base64 encoded random string
func NewRandomChallenge() string {
	random := make([]byte, 32)
	rand.Read(random)
	challenge := make([]byte, base64.StdEncoding.EncodedLen(len(random)))
	base64.StdEncoding.Encode(challenge, random[:])
	return string(challenge)
}
