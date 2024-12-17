package string

import (
	"testing"
	"unsafe"
)

var (
	byteData = []byte("hello world")
	strData  = "hello world"
)

func BenchmarkByte2String(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = string(byteData)
	}
	_ = res
}

func BenchmarkString2Byte(b *testing.B) {
	var res []byte

	for i := 0; i < b.N; i++ {
		res = []byte(strData)
	}
	_ = res
}

func BenchmarkString2ByteNoCopy(b *testing.B) {
	var res []byte
	for i := 0; i < b.N; i++ {
		res = unsafe.Slice(unsafe.StringData(strData), len(strData))
	}
	_ = res
}
func BenchmarkByte2StringNoCopy(b *testing.B) {
	var res string
	for i := 0; i < b.N; i++ {
		res = unsafe.String(unsafe.SliceData(byteData), len(byteData))
	}
	_ = res
}
