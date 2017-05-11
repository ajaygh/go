package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	select {
	case c <- x:
		x, y = y, x+y
	case <-quit:
		fmt.Println("quitting")
		return
	}

}

func main() {
	c, quit := make(chan int), make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
