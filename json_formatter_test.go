package log

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

func TestJsonFormatter_Format(t *testing.T) {
	entry := &Entry{
		Location: "function(file:line)",
		Time:     time.Date(2018, time.May, 20, 8, 20, 30, 666000000, time.UTC),
		Level:    InfoLevel,
		TraceId:  "trace_id_123456789",
		Message:  "message_123456789",
		Fields: map[string]interface{}{
			"key1":           "fields_value1",
			"key2":           "fields_value2",
			"key3":           &testError{X: "123456789"}, // error
			fieldKeyTime:     "time",
			fieldKeyLevel:    "level",
			fieldKeyTraceId:  "request_id",
			fieldKeyLocation: "location",
			fieldKeyMessage:  "msg",
		},
		Buffer: nil,
	}
	data, err := JsonFormatter.Format(entry)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if len(data) == 0 || data[len(data)-1] != '\n' {
		t.Error("want end with '\n'")
		return
	}
	var have map[string]string
	if err = json.Unmarshal(data, &have); err != nil {
		t.Error(err.Error())
		return
	}
	want := map[string]string{
		"time":              "2018-05-20 16:20:30.666",
		"level":             "info",
		"request_id":        "trace_id_123456789",
		"location":          "function(file:line)",
		"msg":               "message_123456789",
		"fields.level":      "level",
		"fields.location":   "location",
		"fields.msg":        "msg",
		"fields.request_id": "request_id",
		"fields.time":       "time",
		"key1":              "fields_value1",
		"key2":              "fields_value2",
		"key3":              "test_error_123456789", // error
	}
	if !reflect.DeepEqual(have, want) {
		t.Errorf("\nhave:%v\nwant:%v", have, want)
		return
	}
}

type testError struct {
	X string `json:"x"`
}

func (e *testError) Error() string {
	return "test_error_123456789"
}
