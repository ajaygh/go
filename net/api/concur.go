package main

import (
	"fmt"
	"sync"
	"time"
)

func do_some() {
	fmt.Println("Sleeping for 1 second")
	time.Sleep(time.Second)
}

func serial_test() {
	fmt.Println("testing serial")
	start := time.Now()
	for i := 0; i < 10; i++ {
		do_some()
	}
	end := time.Since(start)
	fmt.Println("Total time taken serial", end)
}

func concur_test(wg1 *sync.WaitGroup) {
	defer wg1.Done()
	fmt.Println("testing concur")
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Sleeping for 1 second")
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
	end := time.Since(start)
	fmt.Println("Total time taken concur", end)
}

func main() {
	//serial_test()
	var wg sync.WaitGroup
	wg.Add(1)
	go concur_test(&wg)
	wg.Wait()
}
