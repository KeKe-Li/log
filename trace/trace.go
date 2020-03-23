package trace

import (
	"context"
	"net/http"
)

type Tracer interface {
	TraceId() string
}

type traceIdContextKey struct{}

var _traceIdContextKey traceIdContextKey

func NewContext(ctx context.Context, traceId string) context.Context {
	if traceId == "" {
		return ctx
	}
	if ctx == nil {
		return context.WithValue(context.Background(), _traceIdContextKey, traceId)
	}
	if value, ok := ctx.Value(_traceIdContextKey).(string); ok && value == traceId {
		return ctx
	}
	return context.WithValue(ctx, _traceIdContextKey, traceId)
}

func FromContext(ctx context.Context) (traceId string, ok bool) {
	if ctx == nil {
		return "", false
	}
	traceId, ok = ctx.Value(_traceIdContextKey).(string)
	return
}

func FromRequest(req *http.Request) (traceId string, ok bool) {
	traceId, ok = FromContext(req.Context())
	if ok {
		return traceId, true
	}
	return FromHeader(req.Header)
}

const TraceIdHeaderKey = "X-Request-Id"

func FromHeader(header http.Header) (traceId string, ok bool) {
	traceId = header.Get(TraceIdHeaderKey)
	if traceId != "" {
		return traceId, true
	}
	return "", false
}
