package log

import (
	"bytes"
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	now := time.Now()
	have := FormatTime(now)
	want := now.Format(TimeFormatLayout)
	if have := string(have[:]); have != want {
		t.Errorf("have:%s, want:%s", have, want)
		return
	}
}

func TestFormatTimeString(t *testing.T) {
	now := time.Now()
	have := FormatTimeString(now)
	want := now.Format(TimeFormatLayout)
	if have != want {
		t.Errorf("have:%s, want:%s", have, want)
		return
	}
}

func TestItoa(t *testing.T) {
	tests := []struct {
		buf  []byte
		n    int
		want []byte
	}{
		{
			make([]byte, 2),
			0,
			[]byte{'0', '0'},
		},
		{
			make([]byte, 2),
			1,
			[]byte{'0', '1'},
		},
		{
			make([]byte, 2),
			2,
			[]byte{'0', '2'},
		},
		{
			make([]byte, 2),
			11,
			[]byte{'1', '1'},
		},
		{
			make([]byte, 2),
			98,
			[]byte{'9', '8'},
		},

		{
			make([]byte, 3),
			0,
			[]byte{'0', '0', '0'},
		},
		{
			make([]byte, 3),
			6,
			[]byte{'0', '0', '6'},
		},
		{
			make([]byte, 3),
			67,
			[]byte{'0', '6', '7'},
		},
		{
			make([]byte, 3),
			567,
			[]byte{'5', '6', '7'},
		},

		{
			make([]byte, 4),
			0,
			[]byte{'0', '0', '0', '0'},
		},
		{
			make([]byte, 4),
			6,
			[]byte{'0', '0', '0', '6'},
		},
		{
			make([]byte, 4),
			67,
			[]byte{'0', '0', '6', '7'},
		},
		{
			make([]byte, 4),
			567,
			[]byte{'0', '5', '6', '7'},
		},
		{
			make([]byte, 4),
			4567,
			[]byte{'4', '5', '6', '7'},
		},
	}

	for _, v := range tests {
		itoa(v.buf, v.n)
		if !bytes.Equal(v.buf, v.want) {
			t.Errorf("n=%d, have=%v, want=%v", v.n, v.buf, v.want)
			return
		}
	}
}
