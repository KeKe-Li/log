package log

import (
	"testing"
	"time"
)

func BenchmarkFormatTime(b *testing.B) {
	now := time.Now()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormatTime(now)
	}
}

func BenchmarkFormatTimeString(b *testing.B) {
	now := time.Now()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FormatTimeString(now)
	}
}

func BenchmarkStdTimeFormat(b *testing.B) {
	now := time.Now()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		now.Format(TimeFormatLayout)
	}
}
