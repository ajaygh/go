//find the minimum average wait time
package main

import (
	"fmt"
)

type pair struct {
	first, second int
}

func main() {
	var n int
	fmt.Scan(n)
	var orders []pair
	for i := 0; i < n; i++ {
		var f, s int
		fmt.Scan(f, s)
		orders = append(orders, pair{f, s})
	}
	res := calcMinAvgWaitTime(orders)
}

func calcMinAvgWaitTime(orders []pair) int64 {
	heap := heap.New()

}
