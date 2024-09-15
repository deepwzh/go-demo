# 测试系统调用的性能
## 测试结果
```
goos: linux
goarch: amd64
pkg: go-test/syscall
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
BenchmarkGetpid-24              21115286                57.94 ns/op
BenchmarkFstat-24                9094418               133.5 ns/op
BenchmarkClockGettime-24        11050359               108.8 ns/op
PASS
ok      go-test/syscall 3.948s
```
## 测试分析
getpid 系统调用的开销: 57.94ns

fstat 系统调用的开销: 133.5ns

clock_gettime 系统调用的开销: 108.8ns