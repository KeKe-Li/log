// <-- to adjust the line number
package log

import (
	"bytes"
	"context"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func testFatalContextLocationWithoutLogger() {
	FatalContext(testWithoutLoggerContext, "msg")
}
func testErrorContextLocationWithoutLogger() {
	ErrorContext(testWithoutLoggerContext, "msg")
}
func testWarnContextLocationWithoutLogger() {
	WarnContext(testWithoutLoggerContext, "msg")
}
func testInfoContextLocationWithoutLogger() {
	InfoContext(testWithoutLoggerContext, "msg")
}
func testDebugContextLocationWithoutLogger() {
	DebugContext(testWithoutLoggerContext, "msg")
}
func testOutputContextLocationWithoutLogger() {
	OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "msg")
}

var testWithoutLoggerContext = context.Background()

func setWithoutLoggerContextOptionsToDefault() {
	SetFormatter(TextFormatter)
	SetOutput(ConcurrentStdout)
	SetLevel(DebugLevel)
}

//type locationFormat struct{}
//
//func (locationFormat) Format(entry *Entry) ([]byte, error) {
//	return []byte(entry.Location), nil
//}

func TestLocationContextWithoutLoggerWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	// fatal
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(locationFormat{})

		testFatalContextLocationWithoutLogger()

		location := buf.String()
		switch {
		case location == "log.testFatalContextLocationWithoutLogger(github.com/chanxuehong/log/shortcut_without_logger_test.go:14)":
		case strings.HasPrefix(location, "log.testFatalContextLocationWithoutLogger(") && strings.HasSuffix(location, "/log/shortcut_without_logger_test.go:14)"):
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

		testErrorContextLocationWithoutLogger()

		location := buf.String()
		switch {
		case location == "log.testErrorContextLocationWithoutLogger(github.com/chanxuehong/log/shortcut_without_logger_test.go:17)":
		case strings.HasPrefix(location, "log.testErrorContextLocationWithoutLogger(") && strings.HasSuffix(location, "/log/shortcut_without_logger_test.go:17)"):
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

		testWarnContextLocationWithoutLogger()

		location := buf.String()
		switch {
		case location == "log.testWarnContextLocationWithoutLogger(github.com/chanxuehong/log/shortcut_without_logger_test.go:20)":
		case strings.HasPrefix(location, "log.testWarnContextLocationWithoutLogger(") && strings.HasSuffix(location, "/log/shortcut_without_logger_test.go:20)"):
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

		testInfoContextLocationWithoutLogger()

		location := buf.String()
		switch {
		case location == "log.testInfoContextLocationWithoutLogger(github.com/chanxuehong/log/shortcut_without_logger_test.go:23)":
		case strings.HasPrefix(location, "log.testInfoContextLocationWithoutLogger(") && strings.HasSuffix(location, "/log/shortcut_without_logger_test.go:23)"):
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

		testDebugContextLocationWithoutLogger()

		location := buf.String()
		switch {
		case location == "log.testDebugContextLocationWithoutLogger(github.com/chanxuehong/log/shortcut_without_logger_test.go:26)":
		case strings.HasPrefix(location, "log.testDebugContextLocationWithoutLogger(") && strings.HasSuffix(location, "/log/shortcut_without_logger_test.go:26)"):
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

		testOutputContextLocationWithoutLogger()

		location := buf.String()
		switch {
		case location == "log.testOutputContextLocationWithoutLogger(github.com/chanxuehong/log/shortcut_without_logger_test.go:29)":
		case strings.HasPrefix(location, "log.testOutputContextLocationWithoutLogger(") && strings.HasSuffix(location, "/log/shortcut_without_logger_test.go:29)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}
}

//type testJsonFormatter struct{}
//
//func (testJsonFormatter) Format(entry *Entry) ([]byte, error) {
//	// ignored Entry.Location
//
//	// check entry.Time
//	t := time.Now()
//	d := t.Sub(entry.Time)
//	if d > time.Millisecond || d < -time.Millisecond {
//		return nil, fmt.Errorf("time mismatch, have:%v, want:%v", entry.Time, t)
//	}
//
//	// ignored entry.Time
//	m := make(map[string]interface{})
//	prefixFieldClashes(entry.Fields)
//	m[fieldKeyTraceId] = entry.TraceId
//	m[fieldKeyLevel] = entry.Level.String()
//	m[fieldKeyMessage] = entry.Message
//	for k, v := range entry.Fields {
//		m[k] = v
//	}
//	return json.Marshal(m)
//}

func TestFatalContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	FatalContext(testWithoutLoggerContext, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestErrorContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestWarnContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	WarnContext(testWithoutLoggerContext, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestInfoContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	InfoContext(testWithoutLoggerContext, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestDebugContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	DebugContext(testWithoutLoggerContext, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestOutputContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	// fatal
	{
		var buf bytes.Buffer
		SetOutput(ConcurrentWriter(&buf))
		SetFormatter(testJsonFormatter{})

		OutputContext(testWithoutLoggerContext, 0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

		OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

		OutputContext(testWithoutLoggerContext, 0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

		OutputContext(testWithoutLoggerContext, 0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

		OutputContext(testWithoutLoggerContext, 0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

		OutputContext(testWithoutLoggerContext, 0, FatalLevel-1, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

		OutputContext(testWithoutLoggerContext, 0, DebugLevel+1, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()
		if len(data) != 0 {
			t.Errorf("want empty, but now is: %s", data)
			return
		}
	}
}

func TestWithFieldContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	{
		l := WithFieldContext(testWithoutLoggerContext, "field100-key", "field100-value")
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
		ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestWithFieldsContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	var buf bytes.Buffer
	SetOutput(ConcurrentWriter(&buf))
	SetFormatter(testJsonFormatter{})

	{
		l := WithFieldsContext(testWithoutLoggerContext, "field100-key", "field100-value", "field200-key", "field200-value")
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
		ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLeveledOutputContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	// fatal
	{
		SetLevel(FatalLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			OutputContext(testWithoutLoggerContext, 0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			OutputContext(testWithoutLoggerContext, 0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLeveledPrintContextWithoutLogger(t *testing.T) {
	defer setWithoutLoggerContextOptionsToDefault()

	// fatal
	{
		SetLevel(FatalLevel)

		// fatal
		{
			var buf bytes.Buffer
			SetOutput(ConcurrentWriter(&buf))
			SetFormatter(testJsonFormatter{})

			FatalContext(testWithoutLoggerContext, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			WarnContext(testWithoutLoggerContext, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			InfoContext(testWithoutLoggerContext, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			DebugContext(testWithoutLoggerContext, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			FatalContext(testWithoutLoggerContext, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			WarnContext(testWithoutLoggerContext, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			InfoContext(testWithoutLoggerContext, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			DebugContext(testWithoutLoggerContext, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			FatalContext(testWithoutLoggerContext, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			WarnContext(testWithoutLoggerContext, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			InfoContext(testWithoutLoggerContext, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			DebugContext(testWithoutLoggerContext, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			FatalContext(testWithoutLoggerContext, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			WarnContext(testWithoutLoggerContext, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			InfoContext(testWithoutLoggerContext, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			DebugContext(testWithoutLoggerContext, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			FatalContext(testWithoutLoggerContext, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			ErrorContext(testWithoutLoggerContext, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			WarnContext(testWithoutLoggerContext, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			InfoContext(testWithoutLoggerContext, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

			DebugContext(testWithoutLoggerContext, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
