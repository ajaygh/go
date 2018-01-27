package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

type Student struct {
	Id   int
	Name string
	Cnt  []byte
}

func main() {
	var stu1 bytes.Buffer
	enc := gob.NewEncoder(&stu1)
	//	dec := gob.NewDecoder(&stu1)
	err := enc.Encode(Student{1, "ajay", []byte("01234")})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(stu1)
}
