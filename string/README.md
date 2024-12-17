## 测试
go test . -bench=. -benchmem
goos: linux
goarch: amd64
pkg: go-demo/string
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
BenchmarkByte2String-24                 130288062                9.089 ns/op          16 B/op          1 allocs/op
BenchmarkString2Byte-24                 95289331                11.32 ns/op           16 B/op          1 allocs/op
BenchmarkString2ByteNoCopy-24           1000000000               0.2795 ns/op          0 B/op          0 allocs/op
BenchmarkByte2StringNoCopy-24           1000000000               0.3024 ns/op          0 B/op          0 allocs/op
BenchmarkStringSprintf-24                1601800               749.0 ns/op           872 B/op         29 allocs/op
BenchmarkStringPlus-24                   5036642               239.9 ns/op           568 B/op         10 allocs/op
BenchmarkStringBuffer-24                 9433900               126.6 ns/op           288 B/op          3 allocs/op
BenchmarkStringBuilder-24               11412094               107.8 ns/op           248 B/op          5 allocs/op
PASS
ok      go-demo/string  9.928s

## 结论
- 使用unsafe的方式进行string和[]byte的转换，性能最好，大概快40倍左右，且没有额外的内存分配
- 使用Buffer和Builder进行字符串拼接性能较好，使用Sprintf性能最差，最好的Builder要比最差的Sprintf快7倍左右，内存分配也少很多