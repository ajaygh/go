package main

import (
	"fmt"
)

func sltest(data []byte) {
	data[0] = 111
}

func main() {
	data := make([]byte, 2)
	data[0] = 11
	data[1] = 22
	fmt.Println("before ", data)
	sltest(data)
	fmt.Println("after ", data)
}
