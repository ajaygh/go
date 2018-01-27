package perf

import "testing"

func BenchmarkUpdate(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Update(i)
	}
}

func BenchmarkUpdateRef(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Update(i)
	}
}

func BenchmarkUpdateRef1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Update(i)
	}
}
