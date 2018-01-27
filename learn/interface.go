package main

import "fmt"

func main() {
	type Person struct {
		Weight [2]uint8
	}
	b := []byte{4, 5}
	var i interface{}
	i = b
	var p Person
	p = i.(Person)
	fmt.Printf("Person Weight O : %v 1: %v\n", p.Weight[0], p.Weight[1])
}
