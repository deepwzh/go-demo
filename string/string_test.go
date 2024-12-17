package string

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func stringSprintf() string {
	var s string
	v := "benmark"
	for i := 0; i < 10; i++ {
		s = fmt.Sprintf("%s[%s]", s, v)
	}
	return s
}

func stringPlus() string {
	var s string
	v := "benmark"
	for i := 0; i < 10; i++ {
		s = s + "[" + v + "]"
	}
	return s
}

func stringBuffer() string {
	var buffer bytes.Buffer
	v := "benmark"
	for i := 0; i < 10; i++ {
		buffer.WriteString("[")
		buffer.WriteString(v)
		buffer.WriteString("]")
	}
	data := buffer.Bytes()
	return string(data)
}

func stringBuilder() string {
	var builder strings.Builder
	v := "benmark"
	for i := 0; i < 10; i++ {
		builder.WriteString("[")
		builder.WriteString(v)
		builder.WriteString("]")
	}
	return builder.String()
}

func BenchmarkStringSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringSprintf()
	}
}

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringPlus()
	}
}

func BenchmarkStringBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuffer()
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilder()
	}
}
