package log

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestWithTraceId(t *testing.T) {
	// empty traceId
	{
		opt := WithTraceId("")

		var o = options{
			traceId: "987654321",
		}
		opt(&o)
		want := options{
			traceId: "",
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-empty traceId
	{
		opt := WithTraceId("123456789")

		var o = options{
			traceId: "987654321",
		}
		opt(&o)
		want := options{
			traceId: "123456789",
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestWithTraceIdFunc(t *testing.T) {
	// nil function
	{
		var fn func() string
		opt := WithTraceIdFunc(fn)

		var o = options{
			traceId: "987654321",
		}
		opt(&o)
		want := options{
			traceId: "987654321",
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-nil function with empty returns
	{
		fn := func() string { return "" }
		opt := WithTraceIdFunc(fn)

		var o = options{
			traceId: "987654321",
		}
		opt(&o)
		want := options{
			traceId: "",
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-nil function with non-empty returns
	{
		fn := func() string { return "123456789" }
		opt := WithTraceIdFunc(fn)

		var o = options{
			traceId: "987654321",
		}
		opt(&o)
		want := options{
			traceId: "123456789",
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestWithFormatter(t *testing.T) {
	// nil Formatter
	{
		var formatter Formatter
		opt := WithFormatter(formatter)

		var o = options{
			formatter: JsonFormatter,
		}
		opt(&o)
		want := options{
			formatter: JsonFormatter,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-nil Formatter
	{
		opt := WithFormatter(TextFormatter)

		var o = options{
			formatter: JsonFormatter,
		}
		opt(&o)
		want := options{
			formatter: TextFormatter,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestWithOutput(t *testing.T) {
	// nil output
	{
		var output io.Writer
		opt := WithOutput(output)

		var o = options{
			output: ConcurrentStderr,
		}
		opt(&o)
		want := options{
			output: ConcurrentStderr,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-nil output
	{
		opt := WithOutput(ConcurrentStdout)

		var o = options{
			output: ConcurrentStderr,
		}
		opt(&o)
		want := options{
			output: ConcurrentStdout,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestWithLevel(t *testing.T) {
	// fatal-1
	{
		opt := WithLevel(FatalLevel - 1)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: DebugLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// fatal
	{
		opt := WithLevel(FatalLevel)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: FatalLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// error
	{
		opt := WithLevel(ErrorLevel)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: ErrorLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// warning
	{
		opt := WithLevel(WarnLevel)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: WarnLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// info
	{
		opt := WithLevel(InfoLevel)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: InfoLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// debug
	{
		opt := WithLevel(DebugLevel)

		var o = options{
			level: FatalLevel,
		}
		opt(&o)
		want := options{
			level: DebugLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// debug+1
	{
		opt := WithLevel(DebugLevel + 1)

		var o = options{
			level: FatalLevel,
		}
		opt(&o)
		want := options{
			level: FatalLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestWithLevelString(t *testing.T) {
	// panic
	{
		opt := WithLevelString("panic")

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: DebugLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// fatal
	{
		opt := WithLevelString(FatalLevelString)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: FatalLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// error
	{
		opt := WithLevelString(ErrorLevelString)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: ErrorLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// warning
	{
		opt := WithLevelString(WarnLevelString)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: WarnLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// info
	{
		opt := WithLevelString(InfoLevelString)

		var o = options{
			level: DebugLevel,
		}
		opt(&o)
		want := options{
			level: InfoLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// debug
	{
		opt := WithLevelString(DebugLevelString)

		var o = options{
			level: FatalLevel,
		}
		opt(&o)
		want := options{
			level: DebugLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// trace
	{
		opt := WithLevelString("trace")

		var o = options{
			level: FatalLevel,
		}
		opt(&o)
		want := options{
			level: FatalLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestOptions_SetFormatter(t *testing.T) {
	// nil Formatter
	{
		var o = options{
			formatter: JsonFormatter,
		}
		var formatter Formatter
		o.SetFormatter(formatter)

		want := options{
			formatter: JsonFormatter,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-nil Formatter
	{
		var o = options{
			formatter: JsonFormatter,
		}
		o.SetFormatter(TextFormatter)

		want := options{
			formatter: TextFormatter,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestOptions_SetOutput(t *testing.T) {
	// nil output
	{
		var o = options{
			output: ConcurrentStderr,
		}
		var output io.Writer
		o.SetOutput(output)

		want := options{
			output: ConcurrentStderr,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// non-nil output
	{
		var o = options{
			output: ConcurrentStderr,
		}
		o.SetOutput(ConcurrentStdout)

		want := options{
			output: ConcurrentStdout,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestOptions_SetLevel(t *testing.T) {
	// fatal-1
	{
		var o = options{
			level: DebugLevel,
		}
		o.SetLevel(FatalLevel - 1)

		want := options{
			level: DebugLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// fatal
	{
		var o = options{
			level: DebugLevel,
		}
		o.SetLevel(FatalLevel)

		want := options{
			level: FatalLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// error
	{
		var o = options{
			level: DebugLevel,
		}
		o.SetLevel(ErrorLevel)

		want := options{
			level: ErrorLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// warning
	{
		var o = options{
			level: DebugLevel,
		}
		o.SetLevel(WarnLevel)

		want := options{
			level: WarnLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// info
	{
		var o = options{
			level: DebugLevel,
		}
		o.SetLevel(InfoLevel)

		want := options{
			level: InfoLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// debug
	{
		var o = options{
			level: FatalLevel,
		}
		o.SetLevel(DebugLevel)

		want := options{
			level: DebugLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
	// debug+1
	{
		var o = options{
			level: FatalLevel,
		}
		o.SetLevel(DebugLevel + 1)

		want := options{
			level: FatalLevel,
		}
		if o != want {
			t.Errorf("have:%+v, want:%+v", o, want)
			return
		}
	}
}

func TestNewOptions(t *testing.T) {
	// nil == getDefaultOptions()
	{
		SetDefaultOptions(nil)

		defaultWant := options{
			traceId:   "",
			formatter: TextFormatter,
			output:    ConcurrentStdout,
			level:     DebugLevel,
		}

		// with none
		{
			opts := newOptions(nil)
			want := defaultWant
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithTraceId
		{
			opts := newOptions([]Option{WithTraceId("123456789")})
			want := defaultWant
			want.traceId = "123456789"
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithTraceIdFunc
		{
			fn := func() string { return "123456789" }
			opts := newOptions([]Option{WithTraceIdFunc(fn)})
			want := defaultWant
			want.traceId = "123456789"
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithFormatter
		{
			opts := newOptions([]Option{WithFormatter(JsonFormatter)})
			want := defaultWant
			want.formatter = JsonFormatter
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithOutput
		{
			opts := newOptions([]Option{WithOutput(ConcurrentStderr)})
			want := defaultWant
			want.output = ConcurrentStderr
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithLevel
		{
			opts := newOptions([]Option{WithLevel(InfoLevel)})
			want := defaultWant
			want.level = InfoLevel
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithLevelString
		{
			opts := newOptions([]Option{WithLevelString(WarnLevelString)})
			want := defaultWant
			want.level = WarnLevel
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
	}

	// nil != getDefaultOptions()
	{
		formatter := testJsonFormatter{}
		output := ConcurrentWriter(os.Stdout)
		SetDefaultOptions([]Option{WithTraceId("987654321"), WithFormatter(formatter), WithOutput(output), WithLevel(FatalLevel)})
		defer SetDefaultOptions(nil)

		defaultWant := options{
			traceId:   "987654321",
			formatter: formatter,
			output:    output,
			level:     FatalLevel,
		}

		// with none
		{
			opts := newOptions(nil)
			want := defaultWant
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithTraceId
		{
			opts := newOptions([]Option{WithTraceId("123456789")})
			want := defaultWant
			want.traceId = "123456789"
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithTraceIdFunc
		{
			fn := func() string { return "123456789" }
			opts := newOptions([]Option{WithTraceIdFunc(fn)})
			want := defaultWant
			want.traceId = "123456789"
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithFormatter
		{
			opts := newOptions([]Option{WithFormatter(TextFormatter)})
			want := defaultWant
			want.formatter = TextFormatter
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithOutput
		{
			opts := newOptions([]Option{WithOutput(ConcurrentStdout)})
			want := defaultWant
			want.output = ConcurrentStdout
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithLevel
		{
			opts := newOptions([]Option{WithLevel(InfoLevel)})
			want := defaultWant
			want.level = InfoLevel
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
		// WithLevelString
		{
			opts := newOptions([]Option{WithLevelString(WarnLevelString)})
			want := defaultWant
			want.level = WarnLevel
			if have := *opts; have != want {
				t.Errorf("have:%+v, want:%+v", have, want)
				return
			}
		}
	}
}

func Test_SetDefaultOptions_GetDefaultOptions(t *testing.T) {
	defer SetDefaultOptions(nil)

	// no called SetDefaultOptions yet
	opts := getDefaultOptions()
	if opts != nil {
		t.Errorf("have:%v, want:nil", opts)
		return
	}

	// call SetDefaultOptions with non-nil []Option
	opt := WithFormatter(JsonFormatter)
	opts = []Option{
		opt,
	}
	SetDefaultOptions(opts)
	have := getDefaultOptions()
	if !reflect.DeepEqual(have, opts) {
		t.Errorf("have:%v, want:%v", have, opts)
		return
	}

	// call SetDefaultOptions with nil []Option
	SetDefaultOptions(nil)
	opts = getDefaultOptions()
	if opts != nil {
		t.Errorf("have:%v, want:nil", opts)
		return
	}

	// call SetDefaultOptions with non-nil []Option
	opt = WithFormatter(TextFormatter)
	opts = []Option{
		opt,
	}
	SetDefaultOptions(opts)
	have = getDefaultOptions()
	if !reflect.DeepEqual(have, opts) {
		t.Errorf("have:%v, want:%v", have, opts)
		return
	}
}
