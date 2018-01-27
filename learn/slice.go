package main

import (
	"fmt"
)

func test(data *[]byte) {
	*data = append(*data, byte(8))
}

func main() {
	data := make([]byte, 0)
	data = append(data, byte(3))
	for _, i := range data {
		fmt.Println("data content ", i)
	}
	test(&data)
	fmt.Println("After")
	for _, i := range data {
		fmt.Println("data content ", i)
	}
}
