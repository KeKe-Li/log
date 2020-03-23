package v2

import (
	"crypto/sha1"
)

func New(ns [16]byte, name []byte) (uuid [16]byte) {
	h := sha1.New()
	h.Write(ns[:])
	h.Write(name)
	copy(uuid[:], h.Sum(nil))
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // set variant rfc4122
	uuid[6] = (uuid[6] & 0x0f) | 5<<4 // set version 5
	return
}
