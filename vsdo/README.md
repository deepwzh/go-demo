## 测试vsdo与系统调用的性能
### vsdo提供的符号
#### __vdso_clock_gettime
| Symbol                 | Version   | 描述 |
|------------------------|-----------| ----|
| __vdso_clock_gettime    | LINUX_2.6 | 该符号是 clock_gettime 函数的 VDSO 实现，它用于获取指定时钟的时间。通过 VDSO 实现，它可以避免用户空间到内核空间的上下文切换，提供高效的时间获取操作。
| __vdso_getcpu           | LINUX_2.6 | 该符号提供了获取当前 CPU 号的功能，通常用于了解正在执行进程的 CPU 位置。这对于 CPU 亲和性（affinity）相关的优化非常有用。
| __vdso_gettimeofday     | LINUX_2.6 | 这是 gettimeofday 函数的 VDSO 实现，用于获取当前的日期和时间（以秒和微秒为单位）。通过 VDSO 实现，可以快速获取时间而无需调用内核，减少系统调用开销。
| __vdso_time             | LINUX_2.6 | 这是 time 函数的 VDSO 实现，用于获取自 Unix 纪元以来的秒数。该符号实现可以在用户空间快速访问时间信息。

### 测试结果

```
goos: linux
goarch: amd64
pkg: go-test/vsdo
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
BenchmarkVsdoClockGettime-24            46890218                25.32 ns/op
BenchmarkVsdoGettimeofday-24            94176427                12.71 ns/op
BenchmarkSyscallClockGettime-24         10810088               111.5 ns/op
BenchmarkSyscallGetcpu-24               12904531                95.64 ns/op
BenchmarkSyscallGettimeofday-24         10332813               123.6 ns/op
BenchmarkSyscallTime-24                 14027950                91.69 ns/op
PASS
ok      go-test/vsdo    7.849s
```

### 测试分析

golang中vsdo只用了__vdso_clock_gettime和__vdso_gettimeofday

__vdso_clock_gettime的开销: 25.32ns, 系统调用的开销: 111.5ns，是他的4倍
__vdso_gettimeofday的开销: 12.71ns, 系统调用的开销: 123.6ns，是他的10倍