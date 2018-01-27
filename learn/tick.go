package main

import (
	"fmt"
	"time"
)

var quit = make(chan struct{})

func do() {
	tick := time.NewTicker(time.Millisecond * 500)
	i := 0
	for {
		select {
		case <-tick.C:
			fmt.Println("receievd ", i)
			i++
		case <-quit:
			tick.Stop()
		}
	}
}

func main() {
	go do()
	time.Sleep(time.Second * 6)
}
