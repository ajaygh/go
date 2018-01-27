package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const url = "https://foaas.herokuapp.com/shakespeare/Ali/Hilda"

type Foo struct {
	Message  string
	Subtitle string
}

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Accept", "Application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	v := Foo{}
	err = decoder.Decode(&v)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(v)
}
