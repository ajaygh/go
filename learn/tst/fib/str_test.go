package fib

import (
	"strings"
	"testing"
)

func BenchmarkStringjoin1(b *testing.B) {
	b.ReportAllocs()
	input := []string{"Hello", "World"}
	for n := 0; n < b.N; n++ {
		output := strings.Join(input, " ")
		if output != "Hello World" {
			b.Error("Unexpected result:", output)
		}
	}
}
