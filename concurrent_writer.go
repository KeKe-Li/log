package log

import (
	"io"
	"os"
	"sync"
)

var (
	ConcurrentStdout io.Writer = ConcurrentWriter(os.Stdout)
	ConcurrentStderr io.Writer = ConcurrentWriter(os.Stderr)
)

// ConcurrentWriter wraps an io.Writer and returns a concurrent io.Writer.
func ConcurrentWriter(w io.Writer) io.Writer {
	if w == nil {
		return nil
	}
	return &concurrentWriter{
		w: w,
	}
}

type concurrentWriter struct {
	mu sync.Mutex
	w  io.Writer
}

func (w *concurrentWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.w.Write(p)
}
