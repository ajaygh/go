package main

import (
	"fmt"
	"time"
)

const max = 1000000

var data = make(chan int, max)
var rcv = make(chan struct{}, max)

func parallel() {
	start := time.Now()

	for i := 0; i < max; i++ {
		data <- i
	}
	for i := 0; i < max; i++ {
		go prt()
	}
	for i := 0; i < max; i++ {
		<-rcv
	}
	fmt.Println("total time ", start.After(time.Now()))
}

func prt() {
	fmt.Println("write by goroutine ", <-data)
	rcv <- struct{}{}
}

func main() {
	parallel()
}
