package atomic_test

import (
	"sync/atomic"
	"testing"
	"time"
)

func BenchmarkAtomicAdd(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		atomic.AddInt64(&counter, 1)
	}
}

func BenchmarkNonAtomicAdd(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		counter++
	}
}

func BenchmarkAtomicLoad(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		atomic.LoadInt64(&counter)
	}
}

func BenchmarkNonAtomicLoad(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		_ = counter
	}
}

func BenchmarkAtomicStore(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		atomic.StoreInt64(&counter, 1)
	}
}

func BenchmarkNonAtomicStore(b *testing.B) {
	var counter int64
	for i := 0; i < b.N; i++ {
		counter = 1
	}
	_ = counter
}

func BenchmarkAtomicUint32(b *testing.B) {
	var counter uint32
	for i := 0; i < b.N; i++ {
		atomic.AddUint32(&counter, 1)
	}
}

func BenchmarkNonAtomicUint32(b *testing.B) {
	var counter uint32
	for i := 0; i < b.N; i++ {
		counter++
	}
}

func BenchmarkAtomicTime(b *testing.B) {
	var t atomic.Value
	t.Store(time.Now())
	for i := 0; i < b.N; i++ {
		t.Store(time.Now())
	}
}

func BenchmarkNonAtomicTime(b *testing.B) {
	var t time.Time
	for i := 0; i < b.N; i++ {
		t = time.Now()
	}
	_ = t
}
