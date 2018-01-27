package main

import "fmt"

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(2 * i)
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(2*i + 1)
		}
	}()
}
