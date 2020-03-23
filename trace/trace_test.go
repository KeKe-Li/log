package trace

import (
	"context"
	"net/http"
	"testing"
)

func Test_NewContext_FromContext(t *testing.T) {
	// no called NewContext yet
	{
		ctx := context.Background()

		id, ok := FromContext(ctx)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// called NewContext with empty traceId
	{
		ctx := context.Background()

		traceId := ""
		ctx = NewContext(ctx, traceId)
		id, ok := FromContext(ctx)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// called NewContext with non-empty traceId
	{
		ctx := context.Background()

		traceId := "123456789"
		ctx = NewContext(ctx, traceId)
		id, ok := FromContext(ctx)
		wantId, wantOk := traceId, true
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}
}

func TestNewContext(t *testing.T) {
	// empty traceId
	{
		ctx := context.Background()

		ctx2 := NewContext(ctx, "")
		if ctx != ctx2 {
			t.Error("want equal")
			return
		}
	}

	// nil context.Context
	{
		var ctx context.Context

		ctx2 := NewContext(ctx, "123456789")
		if ctx == ctx2 {
			t.Error("want not equal")
			return
		}

		id, ok := FromContext(ctx2)
		wantId, wantOk := "123456789", true
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// serially NewContext with same traceId
	{
		ctx := context.Background()

		traceId := "123456789"
		ctx2 := NewContext(ctx, traceId)
		ctx3 := NewContext(ctx2, traceId)
		if ctx2 != ctx3 {
			t.Error("want equal")
			return
		}
	}

	// parallel NewContext with same traceId
	{
		ctx := context.Background()

		traceId := "123456789"
		ctx2 := NewContext(ctx, traceId)
		ctx3 := NewContext(ctx, traceId)
		if ctx2 == ctx3 {
			t.Error("want not equal")
			return
		}
	}
}

func TestFromContext(t *testing.T) {
	// nil context
	{
		var ctx context.Context
		id, ok := FromContext(ctx)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// non-nil context that does not contain traceId
	{
		ctx := context.Background()
		id, ok := FromContext(ctx)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// non-nil context that contains empty traceId
	// this will never happen, see NewContext

	// non-nil context that contains non-empty traceId
	{
		ctx := context.WithValue(context.Background(), _traceIdContextKey, "123456789")
		id, ok := FromContext(ctx)
		wantId, wantOk := "123456789", true
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}
}

func TestFromRequest(t *testing.T) {
	//// nil *http.Request
	//{
	//	var req *http.Request
	//	id, ok := FromRequest(req)
	//	wantId, wantOk := "", false
	//	if id != wantId || ok != wantOk {
	//		t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
	//		return
	//	}
	//}

	// non-nil *http.Request

	// nil Request.Context()
	{
		// nil Request.Header
		{
			req := &http.Request{}

			id, ok := FromRequest(req)
			wantId, wantOk := "", false
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}

		// non-nil Request.Header without valid traceId
		{
			req := &http.Request{}

			header := make(http.Header)
			req.Header = header

			id, ok := FromRequest(req)
			wantId, wantOk := "", false
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}

		// non-nil Request.Header with valid traceId
		{
			req := &http.Request{}

			header := make(http.Header)
			header.Set(TraceIdHeaderKey, "123456789")
			req.Header = header

			id, ok := FromRequest(req)
			wantId, wantOk := "123456789", true
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}
	}

	// non-nil Request.Context() without valid traceId
	{
		// nil Request.Header
		{
			req := &http.Request{}
			req = req.WithContext(context.Background())

			id, ok := FromRequest(req)
			wantId, wantOk := "", false
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}

		// non-nil Request.Header without valid traceId
		{
			req := &http.Request{}
			req = req.WithContext(context.Background())

			header := make(http.Header)
			req.Header = header

			id, ok := FromRequest(req)
			wantId, wantOk := "", false
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}

		// non-nil Request.Header with valid traceId
		{
			req := &http.Request{}
			req = req.WithContext(context.Background())

			header := make(http.Header)
			header.Set(TraceIdHeaderKey, "123456789")
			req.Header = header

			id, ok := FromRequest(req)
			wantId, wantOk := "123456789", true
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}
	}

	// non-nil Request.Context() with valid traceId
	{
		// nil Request.Header
		{
			req := &http.Request{}

			traceId := "987654321"
			ctx := NewContext(context.Background(), traceId)
			req = req.WithContext(ctx)

			id, ok := FromRequest(req)
			wantId, wantOk := traceId, true
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}

		// non-nil Request.Header without valid traceId
		{
			req := &http.Request{}

			traceId := "987654321"
			ctx := NewContext(context.Background(), traceId)
			req = req.WithContext(ctx)

			header := make(http.Header)
			req.Header = header

			id, ok := FromRequest(req)
			wantId, wantOk := traceId, true
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}

		// non-nil Request.Header with valid traceId
		{
			req := &http.Request{}

			traceId := "987654321"
			ctx := NewContext(context.Background(), traceId)
			req = req.WithContext(ctx)

			header := make(http.Header)
			header.Set(TraceIdHeaderKey, "123456789")
			req.Header = header

			id, ok := FromRequest(req)
			wantId, wantOk := traceId, true
			if id != wantId || ok != wantOk {
				t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
				return
			}
		}
	}
}

func TestFromHeader(t *testing.T) {
	// nil header
	{
		var header http.Header
		id, ok := FromHeader(header)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// non-nil header without TraceIdHeaderKey
	{
		header := make(http.Header)
		id, ok := FromHeader(header)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// non-nil header with empty value for TraceIdHeaderKey
	{
		header := make(http.Header)
		header.Set(TraceIdHeaderKey, "")
		id, ok := FromHeader(header)
		wantId, wantOk := "", false
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}

	// non-nil header with valid value for TraceIdHeaderKey
	{
		header := make(http.Header)
		header.Set(TraceIdHeaderKey, "123456789")
		id, ok := FromHeader(header)
		wantId, wantOk := "123456789", true
		if id != wantId || ok != wantOk {
			t.Errorf("have:(%s, %t), want:(%s, %t)", id, ok, wantId, wantOk)
			return
		}
	}
}
