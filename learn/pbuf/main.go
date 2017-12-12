package main

import (
	"io/ioutil"
	"log"

	pbuf "github.com/ajaygh/learn/pbuf/pb"
	proto "github.com/golang/protobuf/proto"
)

func main() {
	ab := &pbuf.AddressBook{
		People: []*pbuf.Person{
			{
				Name:  "sachin",
				Id:    1,
				Email: "sachin@tt.in",
				Phones: []*pbuf.Person_PhoneNumber{
					{"2113456", pbuf.Person_HOME},
					{"6753456", pbuf.Person_WORK},
				},
			},
			{
				Name:  "saurav",
				Id:    2,
				Email: "saurav@bengal.in",
				Phones: []*pbuf.Person_PhoneNumber{
					{"0113879", pbuf.Person_MOBILE},
					{"1753456", pbuf.Person_HOME},
				},
			},
		},
	}

	out, err := proto.Marshal(ab)
	if err != nil {
		log.Fatalln("failed to marshal:", err)
	}
	if err = ioutil.WriteFile("firstpfout", out, 0644); err != nil {
		log.Fatalln("failed to write:", err)
	}

	tmp := new(pbuf.AddressBook)
	in, err := ioutil.ReadFile("firstpfout")
	if err != nil {
		log.Fatalln("failed to read:", err)
	}

	if err := proto.Unmarshal(in, tmp); err != nil {
		log.Fatalln("failed to unmarshal:", err)
	}
	log.Println("successfully read.", tmp)
}
