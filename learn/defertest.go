package main

import "fmt"

func trace(s string) string {
	fmt.Println("Entering ", s)
	return s
}

func un(s string) {
	fmt.Println("Leaving ", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("In a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("In b")
	a()
}

func main() {
	b()
	s := "test2"
	defer fmt.Println(trace(s))
	s = "test1"
	fmt.Println("test 1")
}
