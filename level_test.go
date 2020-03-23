package log

import (
	"testing"
)

func TestIsValidLevel(t *testing.T) {
	tests := []struct {
		level Level
		want  bool
	}{
		{
			FatalLevel,
			true,
		},
		{
			ErrorLevel,
			true,
		},
		{
			WarnLevel,
			true,
		},
		{
			InfoLevel,
			true,
		},
		{
			DebugLevel,
			true,
		},
		{
			invalidLevel,
			false,
		},
		{
			0,
			false,
		},
		{
			100,
			false,
		},
		{
			6,
			false,
		},
	}
	for _, v := range tests {
		have := isValidLevel(v.level)
		if have != v.want {
			t.Errorf("level:%v, have:%t, want:%t", v.level, have, v.want)
			return
		}
	}
}

func TestIsLevelEnabled(t *testing.T) {
	tests := []struct {
		level       Level
		loggerLevel Level
		want        bool
	}{
		// level is fatal
		{
			FatalLevel,
			FatalLevel,
			true,
		},
		{
			FatalLevel,
			ErrorLevel,
			true,
		},
		{
			FatalLevel,
			WarnLevel,
			true,
		},
		{
			FatalLevel,
			InfoLevel,
			true,
		},
		{
			FatalLevel,
			DebugLevel,
			true,
		},

		// level is error
		{
			ErrorLevel,
			FatalLevel,
			false,
		},
		{
			ErrorLevel,
			ErrorLevel,
			true,
		},
		{
			ErrorLevel,
			WarnLevel,
			true,
		},
		{
			ErrorLevel,
			InfoLevel,
			true,
		},
		{
			ErrorLevel,
			DebugLevel,
			true,
		},

		// level is warning
		{
			WarnLevel,
			FatalLevel,
			false,
		},
		{
			WarnLevel,
			ErrorLevel,
			false,
		},
		{
			WarnLevel,
			WarnLevel,
			true,
		},
		{
			WarnLevel,
			InfoLevel,
			true,
		},
		{
			WarnLevel,
			DebugLevel,
			true,
		},

		// level is info
		{
			InfoLevel,
			FatalLevel,
			false,
		},
		{
			InfoLevel,
			ErrorLevel,
			false,
		},
		{
			InfoLevel,
			WarnLevel,
			false,
		},
		{
			InfoLevel,
			InfoLevel,
			true,
		},
		{
			InfoLevel,
			DebugLevel,
			true,
		},

		// level is debug
		{
			DebugLevel,
			FatalLevel,
			false,
		},
		{
			DebugLevel,
			ErrorLevel,
			false,
		},
		{
			DebugLevel,
			WarnLevel,
			false,
		},
		{
			DebugLevel,
			InfoLevel,
			false,
		},
		{
			DebugLevel,
			DebugLevel,
			true,
		},
	}
	for _, v := range tests {
		have := isLevelEnabled(v.level, v.loggerLevel)
		if have != v.want {
			t.Errorf("level:%v, loggerLevel:%v, have:%t, want:%t", v.level, v.loggerLevel, have, v.want)
			return
		}
	}
}

func TestParseLevelString(t *testing.T) {
	tests := []struct {
		str   string
		level Level
		ok    bool
	}{
		{
			"fatal",
			FatalLevel,
			true,
		},
		{
			"error",
			ErrorLevel,
			true,
		},
		{
			"warning",
			WarnLevel,
			true,
		},
		{
			"info",
			InfoLevel,
			true,
		},
		{
			"debug",
			DebugLevel,
			true,
		},
		{
			"trace",
			invalidLevel,
			false,
		},
		{
			"panic",
			invalidLevel,
			false,
		},
		{
			"",
			invalidLevel,
			false,
		},
	}
	for _, v := range tests {
		level, ok := parseLevelString(v.str)
		if level != v.level || ok != v.ok {
			t.Errorf("str:%s, have:(%d, %t), want:(%d, %t)", v.str, level, ok, v.level, v.ok)
			return
		}
	}
}

func TestLevel_String(t *testing.T) {
	tests := []struct {
		level Level
		str   string
	}{
		{
			FatalLevel,
			"fatal",
		},
		{
			ErrorLevel,
			"error",
		},
		{
			WarnLevel,
			"warning",
		},
		{
			InfoLevel,
			"info",
		},
		{
			DebugLevel,
			"debug",
		},
	}
	for _, v := range tests {
		str := v.level.String()
		if str != v.str {
			t.Errorf("level:%v, have:%s, want:%s", v.level, str, v.str)
			return
		}
	}
}
