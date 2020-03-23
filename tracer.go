package log

import "github.com/KeKe-Li/log/trace"

var _ trace.Tracer = (*logger)(nil)

func (l *logger) TraceId() string {
	return l.getOptions().traceId
}
