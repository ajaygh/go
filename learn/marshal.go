package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type Person struct {
	Name  string
	Phone string ",omitempty"
}

type Scan struct {
	Height, Weight [2]uint8
}

func main() {
	data, err := bson.Marshal(&Person{Name: "Bob"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", data)
	var p Person
	err = bson.Unmarshal(data, &p)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(p.Name)

	var scan Scan
	b := []byte{4, 5, 6, 7}
	err = bson.Unmarshal(b, &scan)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(scan)
}
