package main

import (
	"fmt"
	"reflect"
	"sync"
)

func main() {
	var mp sync.Map
	mp.Store("32dea", map[int]string{
		1234: "swdf",
		3214: "qazx",
	})
	mp.Store("america", 67.45)
	res1, loaded := mp.Load("32dea")
	if loaded {
		tmp, _ := res1.(map[int]string)
		tmp[5678] = "ppppppp"
	}
	res2, _ := mp.Load("america")
	fmt.Println("res ", res1, reflect.TypeOf(res1))
	fmt.Println("res ", res2, reflect.TypeOf(res2))
}
