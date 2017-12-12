package fib

import (
	"log"
	"testing"
)

func TestFib(t *testing.T) {
	tt := []struct {
		n, exp int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 5},
		{4, 3},
		{5, 5},
		{6, 8},
	}

	for _, tc := range tt {
		if actual := Fib(tc.n); actual != tc.exp {
			t.Fatalf("Fib(%d) expected: %d, actual %d\n", tc.n, tc.exp, actual)
		}
		log.Printf("works ")
	}
}
