package main

import (
	"testing"

	"golang.org/x/sys/unix"
)

// 测试获取当前进程 ID 的开销
func BenchmarkGetpid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = unix.Getpid() // 获取当前进程 ID
	}
}

// 测试获取文件状态的开销 (Fstat)
func BenchmarkFstat(b *testing.B) {
	file, err := unix.Open("/dev/null", unix.O_RDONLY, 0)
	if err != nil {
		b.Fatal(err)
	}
	defer unix.Close(file)

	var stat unix.Stat_t
	for i := 0; i < b.N; i++ {
		// 获取文件状态
		if err := unix.Fstat(file, &stat); err != nil {
			b.Fatal(err)
		}
	}
}

// 测试获取当前时间的开销 (ClockGettime)
func BenchmarkClockGettime(b *testing.B) {
	var ts unix.Timespec
	for i := 0; i < b.N; i++ {
		_ = unix.ClockGettime(unix.CLOCK_REALTIME, &ts) // 获取当前时间
	}
}
