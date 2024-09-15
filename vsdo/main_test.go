package main

import (
	"testing"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

func BenchmarkVsdoClockGettime(b *testing.B) {
	// 使用 __vdso_clock_gettime 获取时间的基准测试
	b.ResetTimer() // 重置计时器，以免包括准备时间
	for i := 0; i < b.N; i++ {
		_ = time.Now() // 最终会调用到vDSO clock_gettime
	}
}

// 	golang没有调用 __vdso_getcpu 的方法
// func BenchmarkVsdoGetcpu(b *testing.B) {
// 	// 使用 __vdso_getcpu 获取 CPU 信息的基准测试
// 	b.ResetTimer() // 重置计时器，以免包括准备时间
// 	for i := 0; i < b.N; i++ {
// 		_ = runtime.NumCPU() // vDSO getcpu
// 	}
// }

func BenchmarkVsdoGettimeofday(b *testing.B) {
	var tv unix.Timeval
	// 使用 __vdso_gettimeofday 获取时间的基准测试
	b.ResetTimer() // 重置计时器，以免包括准备时间
	for i := 0; i < b.N; i++ {
		_ = unix.Gettimeofday(&tv) // 最终会调用到vDSO gettimeofday
	}
}

// 	golang没有调用 __vdso_time 的方法
// func BenchmarkVsdoTime(b *testing.B) {
// }

func BenchmarkSyscallClockGettime(b *testing.B) {
	var ts unix.Timespec
	// 使用系统调用获取时间的基准测试
	b.ResetTimer() // 重置计时器，以免包括准备时间
	for i := 0; i < b.N; i++ {
		_, _, _ = unix.Syscall(unix.SYS_CLOCK_GETTIME, unix.CLOCK_REALTIME, uintptr(unsafe.Pointer(&ts)), 0) // 系统调用 clock_gettime
	}
}

func BenchmarkSyscallGetcpu(b *testing.B) {
	var cpu int
	// 使用系统调用获取 CPU 信息的基准测试
	b.ResetTimer() // 重置计时器，以免包括准备时间
	for i := 0; i < b.N; i++ {
		_, _, _ = unix.Syscall(unix.SYS_GETCPU, uintptr(unsafe.Pointer(&cpu)), 0, 0) // 系统调用 getcpu
	}
}

func BenchmarkSyscallGettimeofday(b *testing.B) {
	var tv unix.Timeval
	// 使用系统调用获取时间的基准测试
	b.ResetTimer() // 重置计时器，以免包括准备时间
	for i := 0; i < b.N; i++ {
		_, _, _ = unix.Syscall(unix.SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(&tv)), 0, 0) // 系统调用 gettimeofday
	}
}

func BenchmarkSyscallTime(b *testing.B) {
	// 使用系统调用获取当前时间戳的基准测试
	b.ResetTimer() // 重置计时器，以免包括准备时间
	for i := 0; i < b.N; i++ {
		_, _, _ = unix.Syscall(unix.SYS_TIME, 0, 0, 0) // 系统调用 time
	}
}
