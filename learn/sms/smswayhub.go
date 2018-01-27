package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	sendSMS("9997703718", "DIM810032")
}

func sendSMS(to, msg string) {
	resp, err := http.Get("http://login.smsgatewayhub.com/api/mt/SendSMS?APIKey=" +
		"a58upNs8MEOmC9FQsyxezQ&senderid=WEBSMS&channel=2&DCS=0&flashsms=0&number=" +
		to + "&text=" + msg + "&route=1")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err != nil {
		log.Fatal(err)
	} else {
		log.Println("response is ", string(body))
	}
}
