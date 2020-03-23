package trace

import "github.com/KeKe-Li/log/uuid"

func NewTraceId() string {
	return string(uuid.NewV1().HexEncode())
}
