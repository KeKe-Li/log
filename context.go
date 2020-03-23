package log

import (
	"context"
	"net/http"
)

type loggerContextKey struct{}

var _loggerContextKey loggerContextKey

func NewContext(ctx context.Context, logger Logger) context.Context {
	if logger == nil {
		return ctx
	}
	if ctx == nil {
		return context.WithValue(context.Background(), _loggerContextKey, logger)
	}
	if value, ok := ctx.Value(_loggerContextKey).(Logger); ok && value == logger {
		return ctx
	}
	return context.WithValue(ctx, _loggerContextKey, logger)
}

func FromContext(ctx context.Context) (lg Logger, ok bool) {
	if ctx == nil {
		return nil, false
	}
	lg, ok = ctx.Value(_loggerContextKey).(Logger)
	return
}

func MustFromContext(ctx context.Context) Logger {
	lg, ok := FromContext(ctx)
	if !ok {
		panic("log: failed to get from context.Context")
	}
	return lg
}

func FromContextOrNew(ctx context.Context, new func() Logger) (lg Logger, ctx2 context.Context, isNew bool) {
	lg, ok := FromContext(ctx)
	if ok {
		return lg, ctx, false
	}
	if new != nil {
		lg = new()
		ctx2 = NewContext(ctx, lg)
		isNew = true
		return
	}
	lg = New()
	ctx2 = NewContext(ctx, lg)
	isNew = true
	return
}

func NewRequest(req *http.Request, logger Logger) *http.Request {
	if logger == nil {
		return req
	}
	ctx := req.Context()
	ctx2 := NewContext(ctx, logger)
	if ctx2 == ctx {
		return req
	}
	return req.WithContext(ctx2)
}

func FromRequest(req *http.Request) (lg Logger, ok bool) {
	return FromContext(req.Context())
}

func MustFromRequest(req *http.Request) Logger {
	lg, ok := FromRequest(req)
	if !ok {
		panic("log: failed to get from http.Request")
	}
	return lg
}

func FromRequestOrNew(req *http.Request, new func() Logger) (lg Logger, req2 *http.Request, isNew bool) {
	lg, ok := FromRequest(req)
	if ok {
		return lg, req, false
	}
	if new != nil {
		lg = new()
		req2 = NewRequest(req, lg)
		isNew = true
		return
	}
	lg = New()
	req2 = NewRequest(req, lg)
	isNew = true
	return
}
