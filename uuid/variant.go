package uuid

type Variant byte

const (
	VariantNCS Variant = iota
	VariantRFC4122
	VariantMicrosoft
	VariantFuture
)

func (v Variant) String() string {
	switch v {
	case VariantNCS:
		return "NCS-Variant"
	case VariantRFC4122:
		return "RFC4122-Variant"
	case VariantMicrosoft:
		return "Microsoft-Variant"
	case VariantFuture:
		return "Future-Variant"
	default:
		return "Unknown-Variant"
	}
}

func (uuid UUID) Variant() Variant {
	switch {
	case (uuid[8] & 0x80) == 0x00:
		return VariantNCS
	case (uuid[8] & 0xc0) == 0x80:
		return VariantRFC4122
	case (uuid[8] & 0xe0) == 0xc0:
		return VariantMicrosoft
	default:
		return VariantFuture
	}
}
