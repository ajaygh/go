package main

import (
	"fmt"
)

func main() {
	var data []int
	data = append(data, 1)
	fmt.Println(data)
	data = append(data, 2, 3)
	fmt.Println(data)
	t := []int{10, 20, 30, 40}
	data = append(data, t...)
	fmt.Println(data)
}
