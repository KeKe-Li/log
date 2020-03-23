package log

import (
	"context"
)

// FatalContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	lg.Output(1, FatalLevel, msg, fields...)
//  	return
//  }
//  Output(1, FatalLevel, msg, fields...)
func FatalContext(ctx context.Context, msg string, fields ...interface{}) {
	lg, ok := FromContext(ctx)
	if ok {
		lg.Output(1, FatalLevel, msg, fields...)
		return
	}
	Output(1, FatalLevel, msg, fields...)
}

// ErrorContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	lg.Output(1, ErrorLevel, msg, fields...)
//  	return
//  }
//  Output(1, ErrorLevel, msg, fields...)
func ErrorContext(ctx context.Context, msg string, fields ...interface{}) {
	lg, ok := FromContext(ctx)
	if ok {
		lg.Output(1, ErrorLevel, msg, fields...)
		return
	}
	Output(1, ErrorLevel, msg, fields...)
}

// WarnContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	lg.Output(1, WarnLevel, msg, fields...)
//  	return
//  }
//  Output(1, WarnLevel, msg, fields...)
func WarnContext(ctx context.Context, msg string, fields ...interface{}) {
	lg, ok := FromContext(ctx)
	if ok {
		lg.Output(1, WarnLevel, msg, fields...)
		return
	}
	Output(1, WarnLevel, msg, fields...)
}

// InfoContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	lg.Output(1, InfoLevel, msg, fields...)
//  	return
//  }
//  Output(1, InfoLevel, msg, fields...)
func InfoContext(ctx context.Context, msg string, fields ...interface{}) {
	lg, ok := FromContext(ctx)
	if ok {
		lg.Output(1, InfoLevel, msg, fields...)
		return
	}
	Output(1, InfoLevel, msg, fields...)
}

// DebugContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	lg.Output(1, DebugLevel, msg, fields...)
//  	return
//  }
//  Output(1, DebugLevel, msg, fields...)
func DebugContext(ctx context.Context, msg string, fields ...interface{}) {
	lg, ok := FromContext(ctx)
	if ok {
		lg.Output(1, DebugLevel, msg, fields...)
		return
	}
	Output(1, DebugLevel, msg, fields...)
}

// OutputContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	lg.Output(calldepth+1, level, msg, fields...)
//  	return
//  }
//  Output(calldepth+1, level, msg, fields...)
func OutputContext(ctx context.Context, calldepth int, level Level, msg string, fields ...interface{}) {
	lg, ok := FromContext(ctx)
	if ok {
		lg.Output(calldepth+1, level, msg, fields...)
		return
	}
	Output(calldepth+1, level, msg, fields...)
}

// WithFieldContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	return lg.WithField(key, value)
//  }
//  return WithField(key, value)
func WithFieldContext(ctx context.Context, key string, value interface{}) Logger {
	lg, ok := FromContext(ctx)
	if ok {
		return lg.WithField(key, value)
	}
	return WithField(key, value)
}

// WithFieldsContext is a shortcut to the following code:
//  lg, ok := FromContext(ctx)
//  if ok {
//  	return lg.WithFields(fields...)
//  }
//  return WithFields(fields...)
func WithFieldsContext(ctx context.Context, fields ...interface{}) Logger {
	lg, ok := FromContext(ctx)
	if ok {
		return lg.WithFields(fields...)
	}
	return WithFields(fields...)
}
