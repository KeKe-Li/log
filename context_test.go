package log

import (
	"context"
	"net/http"
	"testing"
)

func Test_NewContext_FromContext(t *testing.T) {
	// no called NewContext yet
	{
		ctx := context.Background()

		lg, ok := FromContext(ctx)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// called NewContext with nil Logger
	{
		ctx := context.Background()

		_lg := Logger(nil)
		ctx = NewContext(ctx, _lg)

		lg, ok := FromContext(ctx)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// called NewContext with non-nil Logger, nil context.Context
	{
		var ctx context.Context

		_lg := New()
		ctx = NewContext(ctx, _lg)

		lg, ok := FromContext(ctx)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// called NewContext with non-nil Logger, non-nil context.Context
	{
		ctx := context.Background()

		_lg := New()
		ctx = NewContext(ctx, _lg)

		lg, ok := FromContext(ctx)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}
}

func TestNewContext(t *testing.T) {
	// nil Logger
	{
		ctx := context.Background()

		ctx2 := NewContext(ctx, Logger(nil))
		if ctx != ctx2 {
			t.Error("want equal")
			return
		}
	}

	// nil context.Context
	{
		var ctx context.Context

		_lg := New()
		ctx2 := NewContext(ctx, _lg)
		if ctx == ctx2 {
			t.Error("want not equal")
			return
		}

		lg, ok := FromContext(ctx2)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// serially NewContext with same Logger
	{
		ctx := context.Background()

		_lg := New()
		ctx2 := NewContext(ctx, _lg)
		ctx3 := NewContext(ctx2, _lg)
		if ctx2 != ctx3 {
			t.Error("want equal")
			return
		}
	}

	// parallel NewContext with same Logger
	{
		ctx := context.Background()

		_lg := New()
		ctx2 := NewContext(ctx, _lg)
		ctx3 := NewContext(ctx, _lg)
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

		lg, ok := FromContext(ctx)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// non-nil context that does not contain Logger
	{
		ctx := context.Background()

		lg, ok := FromContext(ctx)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// non-nil context that contains nil Logger
	// this will never happen, see NewContext

	// non-nil context that contains non-nil Logger
	{
		_lg := New()
		ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)

		lg, ok := FromContext(ctx)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}
}

func TestMustFromContext(t *testing.T) {
	// nil context
	func() {
		var lg Logger

		defer func() {
			if err := recover(); err != nil {
				if lg != nil {
					t.Error("lg must nil")
				}
			} else {
				t.Error("must panic")
			}
		}()

		var ctx context.Context
		lg = MustFromContext(ctx)
	}()

	// non-nil context that does not contain Logger
	func() {
		var lg Logger

		defer func() {
			if err := recover(); err != nil {
				if lg != nil {
					t.Error("lg must nil")
				}
			} else {
				t.Error("must panic")
			}
		}()

		ctx := context.Background()
		lg = MustFromContext(ctx)
	}()

	// non-nil context that contains nil Logger
	// this will never happen, see NewContext

	// non-nil context that contains non-nil Logger
	func() {
		var lg Logger

		defer func() {
			if err := recover(); err != nil {
				t.Error("must not panic")
			} else {
				if lg == nil {
					t.Error("lg must not nil")
				}
			}
		}()

		_lg := New()
		ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)

		lg = MustFromContext(ctx)
		if lg != _lg {
			t.Error("want equal")
			return
		}
	}()
}

func TestFromContextOrNew(t *testing.T) {
	// nil new function
	{
		// nil context
		{
			var ctx context.Context

			lg, ctx2, ok := FromContextOrNew(ctx, nil)
			if lg == nil {
				t.Error("want non-nil")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if ctx2 == ctx {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromContext(ctx2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that does not contain Logger
		{
			ctx := context.Background()

			lg, ctx2, ok := FromContextOrNew(ctx, nil)
			if lg == nil {
				t.Error("want non-nil")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if ctx2 == ctx {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromContext(ctx2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that contains nil Logger
		// this will never happen, see NewContext

		// non-nil context that contains non-nil Logger
		{
			_lg := New()
			ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)

			lg, ctx2, ok := FromContextOrNew(ctx, nil)
			if lg != _lg {
				t.Error("want equal")
				return
			}
			if ok {
				t.Error("want false")
				return
			}
			if ctx2 != ctx {
				t.Error("want equal")
				return
			}
			lg2, ok := FromContext(ctx2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}
	}

	// non-nil new func
	{
		var _newLogger = New()
		var _new = func() Logger { return _newLogger }

		// nil context
		{
			var ctx context.Context

			lg, ctx2, ok := FromContextOrNew(ctx, _new)
			if lg != _newLogger {
				t.Error("want equal")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if ctx2 == ctx {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromContext(ctx2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that does not contain Logger
		{
			ctx := context.Background()

			lg, ctx2, ok := FromContextOrNew(ctx, _new)
			if lg != _newLogger {
				t.Error("want equal")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if ctx2 == ctx {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromContext(ctx2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that contains nil Logger
		// this will never happen, see NewContext

		// non-nil context that contains non-nil Logger
		{
			_lg := New()
			ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)

			lg, ctx2, ok := FromContextOrNew(ctx, _new)
			if lg != _lg {
				t.Error("want equal")
				return
			}
			if ok {
				t.Error("want false")
				return
			}
			if ctx2 != ctx {
				t.Error("want equal")
				return
			}
			lg2, ok := FromContext(ctx2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}
	}
}

func Test_NewRequest_FromRequest(t *testing.T) {
	// no called NewContext yet
	{
		req := &http.Request{}
		req = req.WithContext(context.Background())

		lg, ok := FromRequest(req)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// called NewContext with nil Logger
	{
		req := &http.Request{}
		req = req.WithContext(context.Background())

		_lg := Logger(nil)
		req = NewRequest(req, _lg)

		lg, ok := FromRequest(req)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// called NewContext with non-nil Logger, nil context.Context
	{
		req := &http.Request{}

		_lg := New()
		req = NewRequest(req, _lg)

		lg, ok := FromRequest(req)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// called NewContext with non-nil Logger, non-nil context.Context
	{
		req := &http.Request{}
		req = req.WithContext(context.Background())

		_lg := New()
		req = NewRequest(req, _lg)

		lg, ok := FromRequest(req)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}
}

func TestNewRequest(t *testing.T) {
	// nil Logger
	{
		req := &http.Request{}

		req2 := NewRequest(req, Logger(nil))
		if req != req2 {
			t.Error("want equal")
			return
		}
	}

	// nil context.Context
	{
		req := &http.Request{}

		_lg := New()
		req2 := NewRequest(req, _lg)
		if req == req2 {
			t.Error("want not equal")
			return
		}

		lg, ok := FromRequest(req2)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// serially NewContext with same Logger
	{
		req := &http.Request{}
		req = req.WithContext(context.Background())

		_lg := New()
		req2 := NewRequest(req, _lg)
		req3 := NewRequest(req2, _lg)
		if req2 != req3 {
			t.Error("want equal")
			return
		}
	}

	// parallel NewContext with same Logger
	{
		req := &http.Request{}
		req = req.WithContext(context.Background())

		_lg := New()
		req2 := NewRequest(req, _lg)
		req3 := NewRequest(req, _lg)
		if req2 == req3 {
			t.Error("want not equal")
			return
		}
	}
}

func TestFromRequest(t *testing.T) {
	//// nil Request
	//{
	//	var req *http.Request
	//
	//	lg, ok := FromRequest(req)
	//	want, wantOk := Logger(nil), false
	//	if lg != want || ok != wantOk {
	//		t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
	//		return
	//	}
	//}

	// nil context
	{
		req := &http.Request{}

		lg, ok := FromRequest(req)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// non-nil context that does not contain Logger
	{
		req := &http.Request{}

		ctx := context.Background()
		req = req.WithContext(ctx)

		lg, ok := FromRequest(req)
		want, wantOk := Logger(nil), false
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}

	// non-nil context that contains nil Logger
	// this will never happen, see NewContext

	// non-nil context that contains non-nil Logger
	{
		req := &http.Request{}

		_lg := New()
		ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)
		req = req.WithContext(ctx)

		lg, ok := FromRequest(req)
		want, wantOk := _lg, true
		if lg != want || ok != wantOk {
			t.Errorf("have:(%#v, %t), want:(%#v, %t)", lg, ok, want, wantOk)
			return
		}
	}
}

func TestMustFromRequest(t *testing.T) {
	//// nil Request
	//func() {
	//	var lg Logger
	//
	//	defer func() {
	//		if err := recover(); err != nil {
	//			if lg != nil {
	//				t.Error("lg must nil")
	//			}
	//		} else {
	//			t.Error("must panic")
	//		}
	//	}()
	//
	//	var req *http.Request
	//
	//	lg = MustFromRequest(req)
	//}()

	// nil context
	func() {
		var lg Logger

		defer func() {
			if err := recover(); err != nil {
				if lg != nil {
					t.Error("lg must nil")
				}
			} else {
				t.Error("must panic")
			}
		}()

		req := &http.Request{}

		lg = MustFromRequest(req)
	}()

	// non-nil context that does not contain Logger
	func() {
		var lg Logger

		defer func() {
			if err := recover(); err != nil {
				if lg != nil {
					t.Error("lg must nil")
				}
			} else {
				t.Error("must panic")
			}
		}()

		req := &http.Request{}

		ctx := context.Background()
		req = req.WithContext(ctx)

		lg = MustFromRequest(req)
	}()

	// non-nil context that contains nil Logger
	// this will never happen, see NewContext

	// non-nil context that contains non-nil Logger
	func() {
		var lg Logger

		defer func() {
			if err := recover(); err != nil {
				t.Error("must not panic")
			} else {
				if lg == nil {
					t.Error("lg must not nil")
				}
			}
		}()

		req := &http.Request{}

		_lg := New()
		ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)
		req = req.WithContext(ctx)

		lg = MustFromRequest(req)
		if lg != _lg {
			t.Error("want equal")
			return
		}
	}()
}

func TestFromRequestOrNew(t *testing.T) {
	// nil new function
	{
		//// nil Request
		//{
		//	var req *http.Request
		//
		//	lg, req2, ok := FromRequestOrNew(req, nil)
		//	if lg == nil {
		//		t.Error("want non-nil")
		//		return
		//	}
		//	if !ok {
		//		t.Error("want true")
		//		return
		//	}
		//	if req2 == req {
		//		t.Error("want not equal")
		//		return
		//	}
		//	lg2, ok := FromRequest(req2)
		//	if !ok {
		//		t.Error("want true")
		//		return
		//	}
		//	if lg2 != lg {
		//		t.Error("want equal")
		//		return
		//	}
		//}

		// nil context
		{
			req := &http.Request{}

			lg, req2, ok := FromRequestOrNew(req, nil)
			if lg == nil {
				t.Error("want non-nil")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if req2 == req {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromRequest(req2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that does not contain Logger
		{
			req := &http.Request{}

			ctx := context.Background()
			req = req.WithContext(ctx)

			lg, req2, ok := FromRequestOrNew(req, nil)
			if lg == nil {
				t.Error("want non-nil")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if req2 == req {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromRequest(req2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that contains nil Logger
		// this will never happen, see NewContext

		// non-nil context that contains non-nil Logger
		{
			req := &http.Request{}

			_lg := New()
			ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)
			req = req.WithContext(ctx)

			lg, req2, ok := FromRequestOrNew(req, nil)
			if lg != _lg {
				t.Error("want equal")
				return
			}
			if ok {
				t.Error("want false")
				return
			}
			if req2 != req {
				t.Error("want equal")
				return
			}
			lg2, ok := FromRequest(req2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}
	}

	// non-nil new function
	{
		var _newLogger = New()
		var _new = func() Logger { return _newLogger }

		//// nil Request
		//{
		//	var req *http.Request
		//
		//	lg, req2, ok := FromRequestOrNew(req, _new)
		//	if lg != _newLogger {
		//		t.Error("want equal")
		//		return
		//	}
		//	if !ok {
		//		t.Error("want true")
		//		return
		//	}
		//	if req2 == req {
		//		t.Error("want not equal")
		//		return
		//	}
		//	lg2, ok := FromRequest(req2)
		//	if !ok {
		//		t.Error("want true")
		//		return
		//	}
		//	if lg2 != lg {
		//		t.Error("want equal")
		//		return
		//	}
		//}

		// nil context
		{
			req := &http.Request{}

			lg, req2, ok := FromRequestOrNew(req, _new)
			if lg != _newLogger {
				t.Error("want equal")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if req2 == req {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromRequest(req2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that does not contain Logger
		{
			req := &http.Request{}

			ctx := context.Background()
			req = req.WithContext(ctx)

			lg, req2, ok := FromRequestOrNew(req, _new)
			if lg != _newLogger {
				t.Error("want equal")
				return
			}
			if !ok {
				t.Error("want true")
				return
			}
			if req2 == req {
				t.Error("want not equal")
				return
			}
			lg2, ok := FromRequest(req2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}

		// non-nil context that contains nil Logger
		// this will never happen, see NewContext

		// non-nil context that contains non-nil Logger
		{
			req := &http.Request{}

			_lg := New()
			ctx := context.WithValue(context.Background(), _loggerContextKey, _lg)
			req = req.WithContext(ctx)

			lg, req2, ok := FromRequestOrNew(req, _new)
			if lg != _lg {
				t.Error("want equal")
				return
			}
			if ok {
				t.Error("want false")
				return
			}
			if req2 != req {
				t.Error("want equal")
				return
			}
			lg2, ok := FromRequest(req2)
			if !ok {
				t.Error("want true")
				return
			}
			if lg2 != lg {
				t.Error("want equal")
				return
			}
		}
	}
}
