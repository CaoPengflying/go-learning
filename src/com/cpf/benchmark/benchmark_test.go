package benchmark

import (
	"bytes"
	"testing"
)

func BenchmarkConcatStringByBuffer(b *testing.B) {
	//使用resetTimer 和 stopTimer将代码隔开
	elems := []string{"", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		for _, elem := range elems {
			buf.WriteString(elem)
		}
	}
	b.StopTimer()
}

func BenchmarkConcatStringByAdd(b *testing.B) {
	//使用resetTimer 和 stopTimer将代码隔开
	elems := []string{"", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := ""
		for _, elem := range elems {
			result += elem
		}
	}
	b.StopTimer()
}
