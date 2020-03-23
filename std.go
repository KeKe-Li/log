package log

import "io"

var _std = _New(nil)

// Fatal logs a message at FatalLevel on the standard logger.
// For more information see the Logger interface.
func Fatal(msg string, fields ...interface{}) {
	_std.output(1, FatalLevel, msg, fields)
}

// Error logs a message at ErrorLevel on the standard logger.
// For more information see the Logger interface.
func Error(msg string, fields ...interface{}) {
	_std.output(1, ErrorLevel, msg, fields)
}

// Warn logs a message at WarnLevel on the standard logger.
// For more information see the Logger interface.
func Warn(msg string, fields ...interface{}) {
	_std.output(1, WarnLevel, msg, fields)
}

// Info logs a message at InfoLevel on the standard logger.
// For more information see the Logger interface.
func Info(msg string, fields ...interface{}) {
	_std.output(1, InfoLevel, msg, fields)
}

// Debug logs a message at DebugLevel on the standard logger.
// For more information see the Logger interface.
func Debug(msg string, fields ...interface{}) {
	_std.output(1, DebugLevel, msg, fields)
}

// Output logs a message at specified level on the standard logger.
// For more information see the Logger interface.
func Output(calldepth int, level Level, msg string, fields ...interface{}) {
	_std.Output(calldepth+1, level, msg, fields...)
}

// WithField creates a new Logger from the standard Logger and adds a field to it.
// For more information see the Logger interface.
func WithField(key string, value interface{}) Logger {
	return _std.WithField(key, value)
}

// WithFields creates a new Logger from the standard Logger and adds multiple fields to it.
// For more information see the Logger interface.
func WithFields(fields ...interface{}) Logger {
	return _std.WithFields(fields...)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter Formatter) {
	_std.SetFormatter(formatter)
}

// SetOutput sets the standard logger output.
//  NOTE: output must be thread-safe, see ConcurrentWriter.
func SetOutput(output io.Writer) {
	_std.SetOutput(output)
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) error {
	return _std.SetLevel(level)
}

// SetLevelString sets the standard logger level.
func SetLevelString(str string) error {
	return _std.SetLevelString(str)
}
