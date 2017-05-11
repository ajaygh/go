package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second * time.Duration(rand.Intn(20)))
		fmt.Println("pseudo- random num is ", rand.Intn(20))
	}
}
