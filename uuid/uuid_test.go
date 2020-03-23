package uuid

import (
	"testing"
)

func TestVariantVersion(t *testing.T) {
	u1 := NewV1()
	u5 := NewV5(u1, []byte("test"))
	if u1.Variant() != VariantRFC4122 || u5.Variant() != VariantRFC4122 {
		t.Error("Variant is incorrect")
		return
	}
	if u1.Version() != 1 || u5.Version() != 5 {
		t.Error("Version is incorrect")
		return
	}
}
