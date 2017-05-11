package main

import (
	"fmt"
	"testing"
)

type TestPair struct {
	Input  []float64
	Output float64
}

func Average(in ...float64) float64 {
	if len(in) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, val := range in {
		sum += val
	}
	return sum / float64(len(in))
}

func TestAverage(t *testing.T) {
	v := Average(1, 2)
	if v != 1.5 {
		t.Error("Expected 1.5 got ", v)
	}
	fmt.Println("Testing float slices")
	tests := []TestPair{
		{[]float64{1, 2, 3}, 2},
		{[]float64{1, 2}, 1.5},
		{[]float64{1, 1, 1}, 1},
		{[]float64{}, 0},
	}
	//start testing
	for _, pair := range tests {
		v := Average(pair.Input...)
		if v != pair.Output {
			t.Error("Expected ", pair.Output,
				"got ", v)
		}
	}
}
