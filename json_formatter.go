package log

import (
	"bytes"
	"encoding/json"
)

var JsonFormatter Formatter = jsonFormatter{}

type jsonFormatter struct{}

func (jsonFormatter) Format(entry *Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = bytes.NewBuffer(make([]byte, 0, 16<<10))
	}
	var fields map[string]interface{}
	if fields = entry.Fields; len(fields) > 0 {
		prefixFieldClashes(fields)
		for k, v := range fields {
			if vv, ok := v.(error); ok {
				fields[k] = vv.Error()
			}
		}
	} else {
		fields = make(map[string]interface{}, 8)
	}
	fields[fieldKeyTime] = FormatTimeString(entry.Time.In(_beijingLocation))
	fields[fieldKeyLevel] = entry.Level.String()
	fields[fieldKeyTraceId] = entry.TraceId
	fields[fieldKeyLocation] = entry.Location
	fields[fieldKeyMessage] = entry.Message
	if err := json.NewEncoder(buffer).Encode(fields); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
