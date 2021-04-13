package challenge

import (
	"crypto/rand"
	"encoding/base64"
)

func NewRandomChallenge() (string) {
	random := make([]byte, 32)
	rand.Read(random)
	challenge := make([]byte, base64.StdEncoding.EncodedLen(len(random)))
	base64.StdEncoding.Encode(challenge, random[:])
	return string(challenge)
}
