package main

import (
	"fmt"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type (
	message []byte
	buffer  [windowSize]message
)

var worst time.Duration

func makeMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(b *buffer, n int) {
	start := time.Now()
	m := makeMessage(n)
	(*b)[n%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	var b buffer
	for i := 0; i < msgCount; i++ {
		go pushMsg(&b, i)
	}
	fmt.Println("Worst push time ", worst)
}
