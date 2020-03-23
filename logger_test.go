package log

import "testing"

func TestLogger_New(t *testing.T) {
	lg1 := _New([]Option{
		WithFormatter(JsonFormatter),
		WithOutput(ConcurrentStderr),
		WithLevel(InfoLevel),
	})

	lg2 := _New(nil)
	lg2.SetFormatter(JsonFormatter)
	lg2.SetOutput(ConcurrentStderr)
	lg2.SetLevel(InfoLevel)

	opts1 := lg1.getOptions()
	opts2 := lg2.getOptions()

	if *opts1 != *opts2 {
		t.Errorf("have:%v, want:%v", opts1, opts2)
		return
	}
}

func TestLogger_SetOptions_GetOptions(t *testing.T) {
	lg := _New(nil)

	// default
	{
		opts := lg.getOptions()
		want := options{
			traceId:   "",
			formatter: TextFormatter,
			output:    ConcurrentStdout,
			level:     DebugLevel,
		}
		if have := *opts; have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// set and get
	{
		opts := &options{
			traceId:   "123456789",
			formatter: JsonFormatter,
			output:    ConcurrentStderr,
			level:     InfoLevel,
		}
		lg.setOptions(opts)
		have := lg.getOptions()
		if *have != *opts {
			t.Errorf("have:%v, want:%v", have, opts)
			return
		}
	}
}
