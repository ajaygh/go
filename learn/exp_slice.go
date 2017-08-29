package main

import (
	"fmt"
)

func sl(data *[]byte) {
	(*data)[1] -= 5
	*data = append(*data, 1)
	*data = append(*data, 2)
	*data = append(*data, 3)
	*data = append(*data, 4)
}

func main() {
	data := make([]byte, 0)
	data = append(data, byte(11))
	data = append(data, byte(12))
	fmt.Println("data ", data)
	sl(&data)
	fmt.Println("data ", data)
}
