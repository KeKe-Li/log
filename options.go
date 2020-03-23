package log

import (
	"io"
	"sync/atomic"
	"unsafe"
)

type Option func(*options)

func WithTraceId(traceId string) Option {
	return func(o *options) {
		o.traceId = traceId
	}
}

func WithTraceIdFunc(fn func() string) Option {
	return func(o *options) {
		if fn == nil {
			return
		}
		o.traceId = fn()
	}
}

// WithOutput sets the logger output.
//  NOTE: output must be thread-safe, see ConcurrentWriter.
func WithOutput(output io.Writer) Option {
	return func(o *options) {
		if output == nil {
			return
		}
		o.output = output
	}
}

func WithFormatter(formatter Formatter) Option {
	return func(o *options) {
		if formatter == nil {
			return
		}
		o.formatter = formatter
	}
}

func WithLevel(level Level) Option {
	return func(o *options) {
		if !isValidLevel(level) {
			return
		}
		o.level = level
	}
}

func WithLevelString(str string) Option {
	return func(o *options) {
		level, ok := parseLevelString(str)
		if !ok {
			return
		}
		o.level = level
	}
}

type options struct {
	traceId   string
	formatter Formatter
	output    io.Writer
	level     Level
}

func (opts *options) SetFormatter(formatter Formatter) {
	if formatter == nil {
		return
	}
	opts.formatter = formatter
}
func (opts *options) SetOutput(output io.Writer) {
	if output == nil {
		return
	}
	opts.output = output
}
func (opts *options) SetLevel(level Level) {
	if !isValidLevel(level) {
		return
	}
	opts.level = level
}

func newOptions(opts []Option) *options {
	var o options
	for _, opt := range getDefaultOptions() {
		if opt == nil {
			continue
		}
		opt(&o)
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&o)
	}
	if o.formatter == nil {
		o.formatter = TextFormatter
	}
	if o.output == nil {
		o.output = ConcurrentStdout
	}
	if o.level == invalidLevel {
		o.level = DebugLevel
	}
	return &o
}

var _defaultOptionsPtr unsafe.Pointer // *[]Option

func SetDefaultOptions(opts []Option) {
	if opts == nil {
		atomic.StorePointer(&_defaultOptionsPtr, nil)
		return
	}
	atomic.StorePointer(&_defaultOptionsPtr, unsafe.Pointer(&opts))
}

func getDefaultOptions() []Option {
	ptr := (*[]Option)(atomic.LoadPointer(&_defaultOptionsPtr))
	if ptr == nil {
		return nil
	}
	return *ptr
}
