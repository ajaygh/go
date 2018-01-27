package main

import "fmt"

type Name struct {
	First, Second string
}

func (name Name) fullName() string {
	fmt.Println(name)
	return name.First + name.Second
}

func (name *Name) firstName() string {
	return name.First
}

func main() {
	n1 := &Name{"Sachin", "Tendulkar"}
	n2 := Name{"Tom", "Hanks"}
	fmt.Println(n1.fullName())
	fmt.Println(n2.fullName())
	fmt.Println(n1.firstName())
}
