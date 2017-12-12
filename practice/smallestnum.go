package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	id := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&id[i])
	}

	var m int
	fmt.Scanln(&m)
	//find the minimum
	res := make(map[int]int)
	for _, d := range id {
		if _, ok := res[d]; ok {
			res[d] += 1
		} else {
			res[d] = 1
		}
	}

	min := 1 << 32
	for k, v := range res {
		if v == m && k < min {
			min = k
		}
	}
	fmt.Println(min)
}
