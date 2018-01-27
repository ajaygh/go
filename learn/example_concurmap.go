package main

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
)

var (
	mfruit = cmap.New()
)

func init() {
	mfruit.Set("mango", 5)
	mfruit.Set("apple", 10)
}
func main() {
	fmt.Println("fruit ", mfruit)
	for k, v := range mfruit {
		fmt.Println("key = ", k, "value = ", v)
	}
}
