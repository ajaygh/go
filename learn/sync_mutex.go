package main

import (
	"fmt"
	"sync"
	//"time"
)

type SafeCounter struct {
	count map[string]int
	mux   sync.Mutex
}

func (c *SafeCounter) Inc(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	c.mux.Lock()
	c.count[s]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(s string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.count[s]
}

func main() {
	var wg sync.WaitGroup
	c := SafeCounter{count: make(map[string]int)}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go c.Inc("hello", &wg)
	}
	wg.Wait()
	//	time.Sleep(time.Second)
	defer fmt.Println(c.Value("hello"))
}
