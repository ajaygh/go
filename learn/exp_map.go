package main

import (
	"fmt"
)

var (
	fruit = make(map[int]int)
)

func init() {
	fruit[44] = 6
	fruit[33] = 9
}

func main() {
	for k, v := range fruit {
		fmt.Println("k=", k, "v=", v)
	}
	tmp := 9007306628923392
	fr, ok := fruit[tmp]
	if !ok {
		fr = 5
	}
	fmt.Println(fr)

}
