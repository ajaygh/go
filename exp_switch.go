package main

import (
	"fmt"
)

func main() {
	i := 5
	switch i {
	case 4:
		fmt.Println("matched to 4")
	case 5:
		fmt.Println("matched to 5")
	default:
		fmt.Println("matched not")

	}
}
