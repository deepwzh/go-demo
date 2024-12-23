# 原子操作性能基准测试
## 测试过程
```
# go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: go-demo/atomic
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
BenchmarkAtomicAdd-24           368387625                3.222 ns/op           0 B/op          0 allocs/op
BenchmarkNonAtomicAdd-24        1000000000               0.1006 ns/op          0 B/op          0 allocs/op
BenchmarkAtomicLoad-24          1000000000               0.1827 ns/op          0 B/op          0 allocs/op
BenchmarkNonAtomicLoad-24       1000000000               0.1196 ns/op          0 B/op          0 allocs/op
BenchmarkAtomicStore-24         394983414                3.097 ns/op           0 B/op          0 allocs/op
BenchmarkNonAtomicStore-24      1000000000               0.09617 ns/op         0 B/op          0 allocs/op
BenchmarkAtomicUint32-24        373946335                3.234 ns/op           0 B/op          0 allocs/op
BenchmarkNonAtomicUint32-24     1000000000               0.1112 ns/op          0 B/op          0 allocs/op
BenchmarkAtomicTime-24          25965378                45.86 ns/op           24 B/op          1 allocs/op
BenchmarkNonAtomicTime-24       46106446                25.88 ns/op            0 B/op          0 allocs/op
PASS
ok      go-demo/atomic  7.732s
```

## 结果分析
- 原子操作的性能比非原子操作慢了大概30倍左右
- 基本类型原子操作的内存分配和非原子操作一样，都是0
- time.Time类型的原子操作性能比非原子操作慢了大概2倍，且有额外的内存分配
