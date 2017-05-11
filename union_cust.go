package main

import "fmt"

type Person struct {
	age    [2]uint8
	weight [3]uint8
}

type Rcv struct {
	data []uint8
	Person
}

func main() {
	rcv := Rcv{data: []uint8{1, 1, 1, 1, 1}}
	fmt.Println(rcv.data[0], rcv.age[1])
}
