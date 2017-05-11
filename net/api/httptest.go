package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // Parse arguments
	fmt.Println(r.Form) //Print form's information
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	rand.Seed(14)
	time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	fmt.Println("Sleeping for random times")

	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("Values: ", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello golang")
}

func main() {
	http.HandleFunc("/", sayHelloName) // setup router
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
