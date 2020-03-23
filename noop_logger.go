package log

import "io"

// NoopLogger no operation Logger
type NoopLogger struct{}

// Fatal impl Logger Fatal
func (NoopLogger) Fatal(msg string, fields ...interface{}) {
}

// Error impl Logger Error
func (NoopLogger) Error(msg string, fields ...interface{}) {
}

// Warn impl Logger Warn
func (NoopLogger) Warn(msg string, fields ...interface{}) {
}

// Info impl Logger Info
func (NoopLogger) Info(msg string, fields ...interface{}) {
}

// Debug impl Logger Debug
func (NoopLogger) Debug(msg string, fields ...interface{}) {
}

// Output impl Logger Output
func (NoopLogger) Output(calldepth int, level Level, msg string, fields ...interface{}) {
}

// WithField impl Logger WithField
func (NoopLogger) WithField(key string, value interface{}) Logger {
	return NoopLogger{}
}

// WithFields impl Logger WithFields
func (NoopLogger) WithFields(fields ...interface{}) Logger {
	return NoopLogger{}
}

// SetFormatter impl Logger SetFormatter
func (NoopLogger) SetFormatter(Formatter) {
}

// SetOutput impl Logger SetOutput
func (NoopLogger) SetOutput(io.Writer) {
}

// SetLevel impl Logger SetLevel
func (NoopLogger) SetLevel(Level) error {
	return nil
}

// SetLevelString impl Logger SetLevelString
func (NoopLogger) SetLevelString(string) error {
	return nil
}
