# 测试协程切换和线程切换的性能
## 测试结果
```
goos: linux
goarch: amd64
pkg: go-test/context-switch
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
BenchmarkGoroutineSwitchWithGosched-24          12298116                97.66 ns/op
BenchmarkThreadSwitchWithGosched-24                47210             25405 ns/op
--- BENCH: BenchmarkThreadSwitchWithGosched-24
    main_test.go:25: Goroutine running on thread ID: 1397207
    main_test.go:25: Goroutine running on thread ID: 1397208
    main_test.go:38: Goroutine running done
    main_test.go:25: Goroutine running on thread ID: 1397204
    main_test.go:25: Goroutine running on thread ID: 1397202
    main_test.go:38: Goroutine running done
    main_test.go:25: Goroutine running on thread ID: 1397209
    main_test.go:25: Goroutine running on thread ID: 1397206
    main_test.go:38: Goroutine running done
    main_test.go:25: Goroutine running on thread ID: 1397256
        ... [output truncated]
testing: BenchmarkThreadSwitchWithGosched-24 left GOMAXPROCS set to 2
PASS
ok      go-test/context-switch        2.761s
```
## 测试分析

线程上下文开销: 25.405us

协程上下文开销: 97.66ns