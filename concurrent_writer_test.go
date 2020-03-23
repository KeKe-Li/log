package log

import (
	"bytes"
	"testing"
)

func TestConcurrentWriter(t *testing.T) {
	// nil
	{
		w := ConcurrentWriter(nil)
		if w != nil {
			t.Error("want nil")
			return
		}
	}
	// non-nil
	{
		w := bytes.NewBuffer(make([]byte, 0, 64))

		cw := ConcurrentWriter(w)
		cw.Write([]byte("123456789"))

		have := w.String()
		want := "123456789"
		if have != want {
			t.Errorf("have:%s, want:%s", have, want)
			return
		}
	}
}
