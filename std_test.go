package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
	"time"
)

func testFatalLocation() {
	Fatal("msg")
}
func testErrorLocation() {
	Error("msg")
}
func testWarnLocation() {
	Warn("msg")
}
func testInfoLocation() {
	Info("msg")
}
func testDebugLocation() {
	Debug("msg")
}
func testOutputLocation() {
	Output(0, ErrorLevel, "msg")
}

func setStdOptionsToDefault() {
	SetFormatter(TextFormatter)
	SetOutput(ConcurrentStdout)
	SetLevel(DebugLevel)
}

type locationFormat struct{}

func (locationFormat) Format(entry *Entry) ([]byte, error) {
	return []byte(entry.Location), nil
}

func TestLocation(t *testing.T) {
	defer setStdOptionsToDefault()

	// fatal
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testFatalLocation()

		location := buf.String()
		switch {
		case location == "log.testFatalLocation(github.com/chanxuehong/log/std_test.go:15)":
		case strings.HasPrefix(location, "log.testFatalLocation(") && strings.HasSuffix(location, "/log/std_test.go:15)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// error
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testErrorLocation()

		location := buf.String()
		switch {
		case location == "log.testErrorLocation(github.com/chanxuehong/log/std_test.go:18)":
		case strings.HasPrefix(location, "log.testErrorLocation(") && strings.HasSuffix(location, "/log/std_test.go:18)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// warning
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testWarnLocation()

		location := buf.String()
		switch {
		case location == "log.testWarnLocation(github.com/chanxuehong/log/std_test.go:21)":
		case strings.HasPrefix(location, "log.testWarnLocation(") && strings.HasSuffix(location, "/log/std_test.go:21)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// info
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testInfoLocation()

		location := buf.String()
		switch {
		case location == "log.testInfoLocation(github.com/chanxuehong/log/std_test.go:24)":
		case strings.HasPrefix(location, "log.testInfoLocation(") && strings.HasSuffix(location, "/log/std_test.go:24)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// debug
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testDebugLocation()

		location := buf.String()
		switch {
		case location == "log.testDebugLocation(github.com/chanxuehong/log/std_test.go:27)":
		case strings.HasPrefix(location, "log.testDebugLocation(") && strings.HasSuffix(location, "/log/std_test.go:27)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// output
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testOutputLocation()

		location := buf.String()
		switch {
		case location == "log.testOutputLocation(github.com/chanxuehong/log/std_test.go:30)":
		case strings.HasPrefix(location, "log.testOutputLocation(") && strings.HasSuffix(location, "/log/std_test.go:30)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}
}

type testJsonFormatter struct{}

func (testJsonFormatter) Format(entry *Entry) ([]byte, error) {
	// ignored Entry.Location

	// check entry.Time
	t := time.Now()
	d := t.Sub(entry.Time)
	if d > time.Millisecond || d < -time.Millisecond {
		return nil, fmt.Errorf("time mismatch, have:%v, want:%v", entry.Time, t)
	}

	// ignored entry.Time
	m := make(map[string]interface{})
	prefixFieldClashes(entry.Fields)
	m[fieldKeyTraceId] = entry.TraceId
	m[fieldKeyLevel] = entry.Level.String()
	m[fieldKeyMessage] = entry.Message
	for k, v := range entry.Fields {
		m[k] = v
	}
	return json.Marshal(m)
}

func TestFatal(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
	data := buf.Bytes()

	var have map[string]interface{}
	if err := json.Unmarshal(data, &have); err != nil {
		t.Error(err.Error())
		return
	}
	want := map[string]interface{}{
		fieldKeyTraceId: "",
		fieldKeyLevel:   FatalLevelString,
		fieldKeyMessage: "fatal-msg",
		"field1-key":    "field1-value",
		"field2-key":    "field2-value",
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave:%v\nwant:%v", have, want)
		return
	}
}

func TestError(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
	data := buf.Bytes()

	var have map[string]interface{}
	if err := json.Unmarshal(data, &have); err != nil {
		t.Error(err.Error())
		return
	}
	want := map[string]interface{}{
		fieldKeyTraceId: "",
		fieldKeyLevel:   ErrorLevelString,
		fieldKeyMessage: "error-msg",
		"field1-key":    "field1-value",
		"field2-key":    "field2-value",
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave:%v\nwant:%v", have, want)
		return
	}
}

func TestWarn(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
	data := buf.Bytes()

	var have map[string]interface{}
	if err := json.Unmarshal(data, &have); err != nil {
		t.Error(err.Error())
		return
	}
	want := map[string]interface{}{
		fieldKeyTraceId: "",
		fieldKeyLevel:   WarnLevelString,
		fieldKeyMessage: "warning-msg",
		"field1-key":    "field1-value",
		"field2-key":    "field2-value",
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave:%v\nwant:%v", have, want)
		return
	}
}

func TestInfo(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
	data := buf.Bytes()

	var have map[string]interface{}
	if err := json.Unmarshal(data, &have); err != nil {
		t.Error(err.Error())
		return
	}
	want := map[string]interface{}{
		fieldKeyTraceId: "",
		fieldKeyLevel:   InfoLevelString,
		fieldKeyMessage: "info-msg",
		"field1-key":    "field1-value",
		"field2-key":    "field2-value",
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave:%v\nwant:%v", have, want)
		return
	}
}

func TestDebug(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
	data := buf.Bytes()

	var have map[string]interface{}
	if err := json.Unmarshal(data, &have); err != nil {
		t.Error(err.Error())
		return
	}
	want := map[string]interface{}{
		fieldKeyTraceId: "",
		fieldKeyLevel:   DebugLevelString,
		fieldKeyMessage: "debug-msg",
		"field1-key":    "field1-value",
		"field2-key":    "field2-value",
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave:%v\nwant:%v", have, want)
		return
	}
}

func TestOutput(t *testing.T) {
	defer setStdOptionsToDefault()

	// fatal
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   FatalLevelString,
			fieldKeyMessage: "fatal-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
	// error
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   ErrorLevelString,
			fieldKeyMessage: "error-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
	// warning
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   WarnLevelString,
			fieldKeyMessage: "warning-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
	// info
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   InfoLevelString,
			fieldKeyMessage: "info-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
	// debug
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   DebugLevelString,
			fieldKeyMessage: "debug-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
	// fatal - 1
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, FatalLevel-1, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()
		if len(data) != 0 {
			t.Errorf("want empty, but now is: %s", data)
			return
		}
	}
	// debug + 1
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		Output(0, DebugLevel+1, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()
		if len(data) != 0 {
			t.Errorf("want empty, but now is: %s", data)
			return
		}
	}
}

func TestWithField(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	{
		l := WithField("field100-key", "field100-value")
		l.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   ErrorLevelString,
			fieldKeyMessage: "error-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
			"field100-key":  "field100-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}

	buf.Reset()

	// WithField cannot affect original logger
	{
		Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   ErrorLevelString,
			fieldKeyMessage: "error-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
}

func TestWithFields(t *testing.T) {
	defer setStdOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	{
		l := WithFields("field100-key", "field100-value", "field200-key", "field200-value")
		l.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   ErrorLevelString,
			fieldKeyMessage: "error-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
			"field100-key":  "field100-value",
			"field200-key":  "field200-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}

	buf.Reset()

	// WithFields cannot affect original logger
	{
		Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()

		var have map[string]interface{}
		if err := json.Unmarshal(data, &have); err != nil {
			t.Error(err.Error())
			return
		}
		want := map[string]interface{}{
			fieldKeyTraceId: "",
			fieldKeyLevel:   ErrorLevelString,
			fieldKeyMessage: "error-msg",
			"field1-key":    "field1-value",
			"field2-key":    "field2-value",
		}
		if !reflect.DeepEqual(have, want) {
			t.Errorf("\nhave:%v\nwant:%v", have, want)
			return
		}
	}
}

func TestSetFormatter(t *testing.T) {
	defer setStdOptionsToDefault()

	// default
	{
		have := _std.getOptions().formatter
		want := TextFormatter
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// nil Formatter
	{
		var formatter Formatter
		SetFormatter(formatter)

		have := _std.getOptions().formatter
		want := TextFormatter
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// non-nil Formatter
	{
		var formatter = testJsonFormatter{}
		SetFormatter(formatter)

		have := _std.getOptions().formatter
		want := formatter
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestSetOutput(t *testing.T) {
	defer setStdOptionsToDefault()

	// default
	{
		have := _std.getOptions().output
		want := ConcurrentStdout
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// nil output
	{
		var output io.Writer
		SetOutput(output)

		have := _std.getOptions().output
		want := ConcurrentStdout
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// non-nil output
	{
		var buf bytes.Buffer
		var output io.Writer = ConcurrentWriter(&buf)
		SetOutput(output)

		have := _std.getOptions().output
		want := output
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestSetLevel(t *testing.T) {
	defer setStdOptionsToDefault()

	{
		SetLevel(FatalLevel)

		have := _std.getOptions().level
		want := FatalLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(ErrorLevel)

		have := _std.getOptions().level
		want := ErrorLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(WarnLevel)

		have := _std.getOptions().level
		want := WarnLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(InfoLevel)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(DebugLevel)

		have := _std.getOptions().level
		want := DebugLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(InfoLevel)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(FatalLevel - 1)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(DebugLevel + 1)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevel(InfoLevel)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestSetLevelString(t *testing.T) {
	defer setStdOptionsToDefault()

	{
		SetLevelString(FatalLevelString)

		have := _std.getOptions().level
		want := FatalLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString(ErrorLevelString)

		have := _std.getOptions().level
		want := ErrorLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString(WarnLevelString)

		have := _std.getOptions().level
		want := WarnLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString(InfoLevelString)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString(DebugLevelString)

		have := _std.getOptions().level
		want := DebugLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString(InfoLevelString)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString("panic")

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString("trace")

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		SetLevelString(InfoLevelString)

		have := _std.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestLeveledOutput(t *testing.T) {
	defer setStdOptionsToDefault()

	// fatal
	{
		SetLevel(FatalLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// error
	{
		SetLevel(ErrorLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// warning
	{
		SetLevel(WarnLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   WarnLevelString,
				fieldKeyMessage: "warning-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// info
	{
		SetLevel(InfoLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   WarnLevelString,
				fieldKeyMessage: "warning-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   InfoLevelString,
				fieldKeyMessage: "info-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// debug
	{
		SetLevel(DebugLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   WarnLevelString,
				fieldKeyMessage: "warning-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   InfoLevelString,
				fieldKeyMessage: "info-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   DebugLevelString,
				fieldKeyMessage: "debug-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
	}
}

func TestLeveledPrint(t *testing.T) {
	defer setStdOptionsToDefault()

	// fatal
	{
		SetLevel(FatalLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// error
	{
		SetLevel(ErrorLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// warning
	{
		SetLevel(WarnLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   WarnLevelString,
				fieldKeyMessage: "warning-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// info
	{
		SetLevel(InfoLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   WarnLevelString,
				fieldKeyMessage: "warning-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   InfoLevelString,
				fieldKeyMessage: "info-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// debug
	{
		SetLevel(DebugLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   FatalLevelString,
				fieldKeyMessage: "fatal-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// error
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   ErrorLevelString,
				fieldKeyMessage: "error-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   WarnLevelString,
				fieldKeyMessage: "warning-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   InfoLevelString,
				fieldKeyMessage: "info-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()

			var have map[string]interface{}
			if err := json.Unmarshal(data, &have); err != nil {
				t.Error(err.Error())
				return
			}
			want := map[string]interface{}{
				fieldKeyTraceId: "",
				fieldKeyLevel:   DebugLevelString,
				fieldKeyMessage: "debug-msg",
				"field1-key":    "field1-value",
				"field2-key":    "field2-value",
			}
			if !reflect.DeepEqual(have, want) {
				t.Errorf("\nhave:%v\nwant:%v", have, want)
				return
			}
		}
	}
}
