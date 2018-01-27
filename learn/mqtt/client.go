package main

import (
	"log"
	"sync"

	"github.com/eclipse/paho.mqtt.golang"
)

var wg sync.WaitGroup

func main() {
	o := mqtt.NewClientOptions()
	o = o.AddBroker("tcp://localhost:1883")
	o = o.SetUsername("senra")
	o = o.SetPassword("sc2havellssenra")
	cl := mqtt.NewClient(o)

	if token := cl.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	wg.Add(1)
	if token := cl.Subscribe("test", 0, handler); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}
	wg.Wait()
}

func handler(cl mqtt.Client, msg mqtt.Message) {
	log.Printf("message received %s\n", msg.Payload())
	wg.Done()
}
