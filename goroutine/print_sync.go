package main

import "fmt"

const N = 100
const num = 2

func printnum(i int, c1, c2 chan int) {
	for curr := i; curr <= N; {
		<-c1
		fmt.Println(curr, "from thread ", i)
		curr += num
		c2 <- num - i + 1

	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go printnum(1, c1, c2)
	c1 <- 1
	go printnum(2, c2, c1)
}
