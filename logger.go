package log

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type Logger interface {
	// Fatal logs a message at FatalLevel.
	//
	// Unlike other golang log libraries (for example, the golang standard log library),
	// Fatal just logs a message and does not call os.Exit, so you need to explicitly call os.Exit if necessary.
	//
	// For fields, the following conditions must be satisfied
	//  1. the len(fields) must be an even number, that is to say len(fields)%2==0
	//  2. the even index element of fields must be non-empty string
	Fatal(msg string, fields ...interface{})

	// Error logs a message at ErrorLevel.
	// The requirements for fields can see the comments of Fatal.
	Error(msg string, fields ...interface{})

	// Warn logs a message at WarnLevel.
	// The requirements for fields can see the comments of Fatal.
	Warn(msg string, fields ...interface{})

	// Info logs a message at InfoLevel.
	// The requirements for fields can see the comments of Fatal.
	Info(msg string, fields ...interface{})

	// Debug logs a message at DebugLevel.
	// The requirements for fields can see the comments of Fatal.
	Debug(msg string, fields ...interface{})

	// Output logs a message at specified level.
	//
	// For level==FatalLevel, unlike other golang log libraries (for example, the golang standard log library),
	// Output just logs a message and does not call os.Exit, so you need to explicitly call os.Exit if necessary.
	//
	// The requirements for fields can see the comments of Fatal.
	Output(calldepth int, level Level, msg string, fields ...interface{})

	// WithField creates a new Logger from the current Logger and adds a field to it.
	WithField(key string, value interface{}) Logger

	// WithFields creates a new Logger from the current Logger and adds multiple fields to it.
	// The requirements for fields can see the comments of Fatal.
	WithFields(fields ...interface{}) Logger

	// SetFormatter sets the logger formatter.
	SetFormatter(Formatter)

	// SetOutput sets the logger output.
	SetOutput(io.Writer)

	// SetLevel sets the logger level.
	SetLevel(Level) error

	// SetLevelString sets the logger level.
	SetLevelString(string) error
}

type Formatter interface {
	Format(*Entry) ([]byte, error)
}

type Entry struct {
	Location string // function(file:line)
	Time     time.Time
	Level    Level
	TraceId  string
	Message  string
	Fields   map[string]interface{}
	Buffer   *bytes.Buffer
}

func New(opts ...Option) Logger { return _New(opts) }

func _New(opts []Option) *logger {
	l := &logger{}
	l.setOptions(newOptions(opts))
	return l
}

type logger struct {
	mu      sync.Mutex     // protects the following options field
	options unsafe.Pointer // *options

	fields map[string]interface{}
}

func (l *logger) getOptions() (opts *options) {
	return (*options)(atomic.LoadPointer(&l.options))
}
func (l *logger) setOptions(opts *options) {
	atomic.StorePointer(&l.options, unsafe.Pointer(opts))
}

func (l *logger) SetFormatter(formatter Formatter) {
	if formatter == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	opts := *l.getOptions()
	opts.SetFormatter(formatter)
	l.setOptions(&opts)
}
func (l *logger) SetOutput(output io.Writer) {
	if output == nil {
		return
	}
	l.mu.Lock()
	defer l.mu.Unlock()

	opts := *l.getOptions()
	opts.SetOutput(output)
	l.setOptions(&opts)
}
func (l *logger) SetLevel(level Level) error {
	if !isValidLevel(level) {
		return fmt.Errorf("invalid level: %d", level)
	}
	l.setLevel(level)
	return nil
}
func (l *logger) SetLevelString(str string) error {
	level, ok := parseLevelString(str)
	if !ok {
		return fmt.Errorf("invalid level string: %q", str)
	}
	l.setLevel(level)
	return nil
}
func (l *logger) setLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()

	opts := *l.getOptions()
	opts.SetLevel(level)
	l.setOptions(&opts)
}

func (l *logger) Fatal(msg string, fields ...interface{}) {
	l.output(1, FatalLevel, msg, fields)
}
func (l *logger) Error(msg string, fields ...interface{}) {
	l.output(1, ErrorLevel, msg, fields)
}
func (l *logger) Warn(msg string, fields ...interface{}) {
	l.output(1, WarnLevel, msg, fields)
}
func (l *logger) Info(msg string, fields ...interface{}) {
	l.output(1, InfoLevel, msg, fields)
}
func (l *logger) Debug(msg string, fields ...interface{}) {
	l.output(1, DebugLevel, msg, fields)
}

func (l *logger) Output(calldepth int, level Level, msg string, fields ...interface{}) {
	if !isValidLevel(level) {
		return
	}
	if calldepth < 0 {
		calldepth = 0
	}
	l.output(calldepth+1, level, msg, fields)
}

func (l *logger) output(calldepth int, level Level, msg string, fields []interface{}) {
	opts := l.getOptions()
	if !isLevelEnabled(level, opts.level) {
		return
	}
	location := callerLocation(calldepth + 1)

	combinedFields, err := combineFields(l.fields, fields)
	if err != nil {
		fmt.Fprintf(ConcurrentStderr, "log: failed to combine fields, error=%v, location=%s\n", err, location)
	}

	pool := getBytesBufferPool()
	buffer := pool.Get()
	defer pool.Put(buffer)
	buffer.Reset()

	data, err := opts.formatter.Format(&Entry{
		Location: location,
		Time:     time.Now(),
		Level:    level,
		TraceId:  opts.traceId,
		Message:  msg,
		Fields:   combinedFields,
		Buffer:   buffer,
	})
	if err != nil {
		fmt.Fprintf(ConcurrentStderr, "log: failed to format Entry, error=%v, location=%s\n", err, location)
		return
	}
	if _, err = opts.output.Write(data); err != nil {
		fmt.Fprintf(ConcurrentStderr, "log: failed to write to log, error=%v, location=%s\n", err, location)
		return
	}
}

func (l *logger) WithField(key string, value interface{}) Logger {
	if key == "" {
		return l
	}
	if len(l.fields) == 0 {
		nl := &logger{
			fields: map[string]interface{}{key: value},
		}
		nl.setOptions(l.getOptions())
		return nl
	}
	m := make(map[string]interface{}, len(l.fields)+1)
	for k, v := range l.fields {
		m[k] = v
	}
	m[key] = value
	nl := &logger{
		fields: m,
	}
	nl.setOptions(l.getOptions())
	return nl
}

func (l *logger) WithFields(fields ...interface{}) Logger {
	if len(fields) == 0 {
		return l
	}
	m, err := combineFields(l.fields, fields)
	if err != nil {
		fmt.Fprintf(ConcurrentStderr, "log: failed to combine fields, error=%v, location=%s\n", err, callerLocation(1))
	}
	nl := &logger{
		fields: m,
	}
	nl.setOptions(l.getOptions())
	return nl
}
