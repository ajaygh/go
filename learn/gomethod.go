package main

import (
	"fmt"
	"time"
)

type Count struct {
	count int
}

func (c Count) countNum() {
	ticker := time.NewTicker(time.Millisecond * 200)
	for {
		<-ticker.C
		fmt.Println("val = ", c.count)
		c.count++
		fmt.Println("recievd from chan ", <-ch)
	}
}

var ch = make(chan int)

func send() {
	ticker := time.NewTicker(time.Millisecond * 300)
	ii := 100
	for {
		select {
		case <-ticker.C:
			ch <- ii
			fmt.Println("Sent value", ii)
			ii += 5
		}
	}
}

func main() {
	ct := &Count{1}
	go ct.countNum()
	go send()
	time.Sleep(time.Second * 10)
}
