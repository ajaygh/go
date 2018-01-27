package strjoin

import "testing"
import "strings"

func BenchmarkStringJoin1(b *testing.B) {
	b.ReportAllocs()
	in := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		res := strings.Join(in, " ")
		if res != "Hello World" {
			b.Error("Unexpected result: ", res)
		}
	}
}

func BenchmarkStringJoin2(b *testing.B) {
	b.ReportAllocs()

	join := func(sarr []string, delim string) string {
		if len(sarr) == 2 {
			return sarr[0] + delim + sarr[1]
		}
		return ""
	}
	in := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		res := join(in, " ")
		if res != "Hello World" {
			b.Error("Unexpected result: ", res)
		}
	}
}

func BenchmarkStringJoin3(b *testing.B) {
	b.ReportAllocs()
	in := []string{"Hello", "World"}
	for i := 0; i < b.N; i++ {
		res := in[0] + " " + in[1]
		if res != "Hello World" {
			b.Error("Unexpected result: ", res)
		}
	}
}
