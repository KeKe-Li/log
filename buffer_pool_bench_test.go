package log

import "testing"

func BenchmarkGetBytesBufferPool(b *testing.B) {
	b.ReportAllocs()
	b.SetParallelism(64)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			getBytesBufferPool()
		}
	})
}
