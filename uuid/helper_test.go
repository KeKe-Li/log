package uuid

import (
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	u1 := NewV1()
	encoded := u1.Encode()
	u2, err := Decode(encoded)
	if err != nil {
		t.Error("decode failed")
		return
	}
	if u1 != u2 {
		t.Error("encode and decode mismatch")
		return
	}
}

//var (
//	NamespaceDNS, _  = Decode([]byte("6ba7b810-9dad-11d1-80b4-00c04fd430c8"))
//	NamespaceURL, _  = Decode([]byte("6ba7b811-9dad-11d1-80b4-00c04fd430c8"))
//	NamespaceOID, _  = Decode([]byte("6ba7b812-9dad-11d1-80b4-00c04fd430c8"))
//	NamespaceX500, _ = Decode([]byte("6ba7b814-9dad-11d1-80b4-00c04fd430c8"))
//)

func TestNamespace(t *testing.T) {
	if NamespaceDNS.String() != "6ba7b810-9dad-11d1-80b4-00c04fd430c8" {
		t.Error("decode and encode mismatch")
		return
	}
	if NamespaceURL.String() != "6ba7b811-9dad-11d1-80b4-00c04fd430c8" {
		t.Error("decode and encode mismatch")
		return
	}
	if NamespaceOID.String() != "6ba7b812-9dad-11d1-80b4-00c04fd430c8" {
		t.Error("decode and encode mismatch")
		return
	}
	if NamespaceX500.String() != "6ba7b814-9dad-11d1-80b4-00c04fd430c8" {
		t.Error("decode and encode mismatch")
		return
	}
}
