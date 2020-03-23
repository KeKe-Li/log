package uuid

import (
	"encoding/hex"
	"errors"
)

func HexEncode(uuid UUID) []byte {
	buf := make([]byte, 32)
	hex.Encode(buf, uuid[:])
	return buf
}

// Encode encodes UUID to "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" format.
func Encode(uuid UUID) []byte {
	const encodedUUIDLen = 36
	var buf [encodedUUIDLen]byte
	hex.Encode(buf[:8], uuid[:4])
	buf[8] = '-'
	hex.Encode(buf[9:13], uuid[4:6])
	buf[13] = '-'
	hex.Encode(buf[14:18], uuid[6:8])
	buf[18] = '-'
	hex.Encode(buf[19:23], uuid[8:10])
	buf[23] = '-'
	hex.Encode(buf[24:], uuid[10:])
	return buf[:]
}

var (
	NamespaceDNS, _  = Decode([]byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
	NamespaceURL, _  = Decode([]byte("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
	NamespaceOID, _  = Decode([]byte("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
	NamespaceX500, _ = Decode([]byte("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
)

// Decode decodes data with "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" format into UUID.
func Decode(data []byte) (uuid UUID, err error) {
	const encodedUUIDLen = 36
	if len(data) != encodedUUIDLen {
		err = errors.New("invalid UUID data")
		return
	}
	_ = data[encodedUUIDLen-1]
	if data[8] != '-' || data[13] != '-' || data[18] != '-' || data[23] != '-' {
		err = errors.New("invalid UUID data")
		return
	}
	if _, err = hex.Decode(uuid[:4], data[:8]); err != nil {
		return
	}
	if _, err = hex.Decode(uuid[4:6], data[9:13]); err != nil {
		return
	}
	if _, err = hex.Decode(uuid[6:8], data[14:18]); err != nil {
		return
	}
	if _, err = hex.Decode(uuid[8:10], data[19:23]); err != nil {
		return
	}
	_, err = hex.Decode(uuid[10:], data[24:])
	return
}
