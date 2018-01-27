package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Id, Salary int
}

func main() {
	person := Person{"tom", 40}
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	json.NewEncoder(b1).Encode(person)
	fmt.Println("encoded person is ")
	io.Copy(os.Stdout, b1)
	employee := Employee{person, 55, 4000}
	json.NewEncoder(b2).Encode(employee)
	fmt.Println("encoded employee is ")
	io.Copy(os.Stdout, b2)
}
