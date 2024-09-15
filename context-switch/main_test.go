package main

import (
	"runtime"
	"sync"
	"testing"

	"golang.org/x/sys/unix"
)

// 测试 Goroutine 切换的开销，使用 runtime.Gosched()
func BenchmarkGoroutineSwitchWithGosched(b *testing.B) {
	// 运行 b.N 次，模拟 Goroutine 切换
	for i := 0; i < b.N; i++ {
		runtime.Gosched() // 主动触发 Goroutine 切换
	}
}

func BenchmarkThreadSwitchWithGosched(b *testing.B) {
	// 运行 b.N 次，模拟 Goroutine 切换
	runtime.GOMAXPROCS(2)
	wg := sync.WaitGroup{}
	run := func() {
		tid := unix.Gettid() // 获取线程 ID
		b.Logf("Goroutine running on thread ID: %d", tid)
		runtime.LockOSThread()
		for i := 0; i < b.N; i++ {
			runtime.Gosched() // 主动触发 Goroutine 切换
		}
		wg.Done()
	}
	wg.Add(1)
	go run()
	wg.Add(1)
	go run()

	wg.Wait()
	b.Logf("Goroutine running done")
}
