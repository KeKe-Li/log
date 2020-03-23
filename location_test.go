package log

import (
	"strings"
	"testing"
)

func testCallerLocation() string {
	return callerLocation(0)
}

func TestCallerLocation(t *testing.T) {
	location := testCallerLocation()
	switch {
	case location == "log.testCallerLocation(github.com/chanxuehong/log/location_test.go:9)":
	case strings.HasPrefix(location, "log.testCallerLocation(") && strings.HasSuffix(location, "/log/location_test.go:9)"):
	default:
		t.Errorf("not expected location: %s", location)
		return
	}
}

func TestTrimFileName(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{
			"/a/b/c.go",
			"/a/b/c.go",
		},
		{
			"/a/src/b/c.go",
			"b/c.go",
		},
		{
			"/a/src/d/vendor/b/c.go",
			"b/c.go",
		},
		{
			"/a/vendor/b/c.go",
			"b/c.go",
		},
	}

	for _, v := range tests {
		have := trimFileName(v.str)
		if have != v.want {
			t.Errorf("have:%s, want:%s", have, v.want)
			return
		}
	}
}
