package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := 44.44
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	fmt.Println(t, v)
}
