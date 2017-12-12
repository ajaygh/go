package main

import (
	"fmt"
)

func main() {
	var N int64
	fmt.Scanln(&N)
	res := make(chan int64, 1)
	fib(N, res)
	fmt.Println(<-res)
}

func fib(N int64, res chan int64) {
	if N < 2 {
		res <- N
		return
	}
	tmp := make(chan int64, 1)
	go fib(N-1, tmp)
	go fib(N-2, tmp)
	total := <-tmp
	total += <-tmp
	res <- total
}
