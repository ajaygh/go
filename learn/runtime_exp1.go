package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("max cores ", runtime.NumCPU())
}
