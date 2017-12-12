package main

import (
	"fmt"
)

func mp(p map[int]string) {
	p[4] = "four"
	p[5] = "five"
}

func main() {
	p := make(map[int]string)
	p[1] = "one"
	fmt.Println("before ", p)
	mp(p)
	fmt.Println("after ", p)

}
