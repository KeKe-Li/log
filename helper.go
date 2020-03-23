package log

import (
	"encoding/json"
	"encoding/xml"
)

// JSON is a helper function, following is its function code.
//
//  data, _ := json.Marshal(v)
//  return string(data)
func JSON(v interface{}) string {
	pool := getBytesBufferPool()
	buffer := pool.Get()
	defer pool.Put(buffer)
	buffer.Reset()

	if err := json.NewEncoder(buffer).Encode(v); err != nil {
		return ""
	}
	data := buffer.Bytes()

	// remove the trailing newline
	i := len(data) - 1
	if i < 0 || i >= len(data) /* BCE */ {
		return ""
	}
	if data[i] == '\n' {
		data = data[:i]
	}
	return string(data)
}

// XML is a helper function, following is its function code.
//
//  data, _ := xml.Marshal(v)
//  return string(data)
func XML(v interface{}) string {
	pool := getBytesBufferPool()
	buffer := pool.Get()
	defer pool.Put(buffer)
	buffer.Reset()

	if err := xml.NewEncoder(buffer).Encode(v); err != nil {
		return ""
	}
	return string(buffer.Bytes())
}
