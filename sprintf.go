package main

import "fmt"

func main() {
	j := 55
	i := 33
	k := fmt.Sprintf("%04d-%03d", i, j)
	fmt.Println("k value is ", k)
}
