package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func doRequest(url string) {
	for {
		go func() {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(time.Now(), " res body: ", string(body))
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		}()
	}
}

func main() {
	doRequest("http://127.0.0.1:9090")
}
