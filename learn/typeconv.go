package main

import "fmt"
import "strconv"

func main() {
	var scu uint64 = 000100
	sscu := strconv.FormatUint(scu, 10)

	fmt.Println(sscu)
}
