package log

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"strings"
	"testing"
)

func testLogger_FatalLocation(buf *bytes.Buffer) {
	lg := _New(nil)
	lg.SetOutput(ConcurrentWriter(buf))
	lg.SetFormatter(locationFormat{})
	lg.Fatal("msg")
}
func testLogger_ErrorLocation(buf *bytes.Buffer) {
	lg := _New(nil)
	lg.SetOutput(ConcurrentWriter(buf))
	lg.SetFormatter(locationFormat{})
	lg.Error("msg")
}
func testLogger_WarnLocation(buf *bytes.Buffer) {
	lg := _New(nil)
	lg.SetOutput(ConcurrentWriter(buf))
	lg.SetFormatter(locationFormat{})
	lg.Warn("msg")
}
func testLogger_InfoLocation(buf *bytes.Buffer) {
	lg := _New(nil)
	lg.SetOutput(ConcurrentWriter(buf))
	lg.SetFormatter(locationFormat{})
	lg.Info("msg")
}
func testLogger_DebugLocation(buf *bytes.Buffer) {
	lg := _New(nil)
	lg.SetOutput(ConcurrentWriter(buf))
	lg.SetFormatter(locationFormat{})
	lg.Debug("msg")
}
func testLogger_OutputLocation(buf *bytes.Buffer) {
	lg := _New(nil)
	lg.SetOutput(ConcurrentWriter(buf))
	lg.SetFormatter(locationFormat{})
	lg.Output(0, ErrorLevel, "msg")
}

func TestLogger_Location(t *testing.T) {
	// fatal
	{
		var buf bytes.Buffer
		testLogger_FatalLocation(&buf)

		location := buf.String()
		switch {
		case location == "log.testLogger_FatalLocation(github.com/chanxuehong/log/logger_api_test.go:16)":
		case strings.HasPrefix(location, "log.testLogger_FatalLocation(") && strings.HasSuffix(location, "/log/logger_api_test.go:16)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// error
	{
		var buf bytes.Buffer
		testLogger_ErrorLocation(&buf)

		location := buf.String()
		switch {
		case location == "log.testLogger_ErrorLocation(github.com/chanxuehong/log/logger_api_test.go:22)":
		case strings.HasPrefix(location, "log.testLogger_ErrorLocation(") && strings.HasSuffix(location, "/log/logger_api_test.go:22)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// warning
	{
		var buf bytes.Buffer
		testLogger_WarnLocation(&buf)

		location := buf.String()
		switch {
		case location == "log.testLogger_WarnLocation(github.com/chanxuehong/log/logger_api_test.go:28)":
		case strings.HasPrefix(location, "log.testLogger_WarnLocation(") && strings.HasSuffix(location, "/log/logger_api_test.go:28)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// info
	{
		var buf bytes.Buffer
		testLogger_InfoLocation(&buf)

		location := buf.String()
		switch {
		case location == "log.testLogger_InfoLocation(github.com/chanxuehong/log/logger_api_test.go:34)":
		case strings.HasPrefix(location, "log.testLogger_InfoLocation(") && strings.HasSuffix(location, "/log/logger_api_test.go:34)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// debug
	{
		var buf bytes.Buffer
		testLogger_DebugLocation(&buf)

		location := buf.String()
		switch {
		case location == "log.testLogger_DebugLocation(github.com/chanxuehong/log/logger_api_test.go:40)":
		case strings.HasPrefix(location, "log.testLogger_DebugLocation(") && strings.HasSuffix(location, "/log/logger_api_test.go:40)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}

	// output
	{
		var buf bytes.Buffer
		testLogger_OutputLocation(&buf)

		location := buf.String()
		switch {
		case location == "log.testLogger_OutputLocation(github.com/chanxuehong/log/logger_api_test.go:46)":
		case strings.HasPrefix(location, "log.testLogger_OutputLocation(") && strings.HasSuffix(location, "/log/logger_api_test.go:46)"):
		default:
			t.Errorf("not expected location: %s", location)
			return
		}
	}
}

func TestLogger_Fatal(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	lg.Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_Error(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_Warn(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	lg.Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_Info(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	lg.Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_Debug(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	lg.Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_Output(t *testing.T) {
	lg := _New(nil)

	// fatal
	{
		var buf bytes.Buffer
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, FatalLevel-1, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()
		if len(data) != 0 {
			t.Errorf("want empty, but now is: %s", data)
			return
		}
	}
	// debug + 1
	{
		var buf bytes.Buffer
		lg.SetOutput(ConcurrentWriter(&buf))
		lg.SetFormatter(testJsonFormatter{})

		lg.Output(0, DebugLevel+1, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
		data := buf.Bytes()
		if len(data) != 0 {
			t.Errorf("want empty, but now is: %s", data)
			return
		}
	}
}

func TestLogger_WithField(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	{
		l := lg.WithField("field100-key", "field100-value")
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
		lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_WithFields(t *testing.T) {
	lg := _New(nil)

	var buf bytes.Buffer
	lg.SetOutput(ConcurrentWriter(&buf))
	lg.SetFormatter(testJsonFormatter{})

	{
		l := lg.WithFields("field100-key", "field100-value", "field200-key", "field200-value")
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
		lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_SetFormatter(t *testing.T) {
	lg := _New(nil)

	// default
	{
		have := lg.getOptions().formatter
		want := TextFormatter
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// nil Formatter
	{
		var formatter Formatter
		lg.SetFormatter(formatter)

		have := lg.getOptions().formatter
		want := TextFormatter
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// non-nil Formatter
	{
		var formatter = testJsonFormatter{}
		lg.SetFormatter(formatter)

		have := lg.getOptions().formatter
		want := formatter
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestLogger_SetOutput(t *testing.T) {
	lg := _New(nil)

	// default
	{
		have := lg.getOptions().output
		want := ConcurrentStdout
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	// nil output
	{
		var output io.Writer
		lg.SetOutput(output)

		have := lg.getOptions().output
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
		lg.SetOutput(output)

		have := lg.getOptions().output
		want := output
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestLogger_SetLevel(t *testing.T) {
	lg := _New(nil)

	{
		lg.SetLevel(FatalLevel)

		have := lg.getOptions().level
		want := FatalLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(ErrorLevel)

		have := lg.getOptions().level
		want := ErrorLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(WarnLevel)

		have := lg.getOptions().level
		want := WarnLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(InfoLevel)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(DebugLevel)

		have := lg.getOptions().level
		want := DebugLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(InfoLevel)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(FatalLevel - 1)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(DebugLevel + 1)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevel(InfoLevel)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestLogger_SetLevelString(t *testing.T) {
	lg := _New(nil)

	{
		lg.SetLevelString(FatalLevelString)

		have := lg.getOptions().level
		want := FatalLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString(ErrorLevelString)

		have := lg.getOptions().level
		want := ErrorLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString(WarnLevelString)

		have := lg.getOptions().level
		want := WarnLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString(InfoLevelString)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString(DebugLevelString)

		have := lg.getOptions().level
		want := DebugLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString(InfoLevelString)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString("panic")

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString("trace")

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
	{
		lg.SetLevelString(InfoLevelString)

		have := lg.getOptions().level
		want := InfoLevel
		if have != want {
			t.Errorf("have:%v, want:%v", have, want)
			return
		}
	}
}

func TestLogger_LeveledOutput(t *testing.T) {
	lg := _New(nil)

	// fatal
	{
		lg.SetLevel(FatalLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// error
	{
		lg.SetLevel(ErrorLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// warning
	{
		lg.SetLevel(WarnLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// info
	{
		lg.SetLevel(InfoLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// debug
	{
		lg.SetLevel(DebugLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, FatalLevel, "fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, ErrorLevel, "error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, WarnLevel, "warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, InfoLevel, "info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Output(0, DebugLevel, "debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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

func TestLogger_LeveledPrint(t *testing.T) {
	lg := _New(nil)

	// fatal
	{
		lg.SetLevel(FatalLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// warning
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// error
	{
		lg.SetLevel(ErrorLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// info
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// warning
	{
		lg.SetLevel(WarnLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
		// debug
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// info
	{
		lg.SetLevel(InfoLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
			data := buf.Bytes()
			if len(data) != 0 {
				t.Errorf("want empty, but now is: %s", data)
				return
			}
		}
	}

	// debug
	{
		lg.SetLevel(DebugLevel)

		// fatal
		{
			var buf bytes.Buffer
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Fatal("fatal-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Error("error-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Warn("warning-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Info("info-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
			lg.SetOutput(ConcurrentWriter(&buf))
			lg.SetFormatter(testJsonFormatter{})

			lg.Debug("debug-msg", "field1-key", "field1-value", "field2-key", "field2-value")
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
