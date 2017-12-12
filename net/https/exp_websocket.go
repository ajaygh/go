package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(wc *websocket.Conn) {
	//continuously receive and send data to socket
	for {
		var reply string
		if err := websocket.Message.Receive(wc, &reply); err != nil {
			log.Fatalln("can't receive ", err)
			break
		}
		fmt.Println("received message is ", reply)

		if err := websocket.Message.Send(wc, "Received: "+reply); err != nil {
			log.Println("can't send ", err)
			break
		}

	}
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("working")
	w.Write([]byte("nice working"))
}
func main() {
	fmt.Println("Learning websocket in go")

	http.Handle("/ws", websocket.Handler(Echo))
	http.HandleFunc("/test", test)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	if err := http.ListenAndServeTLS(":1234", "server.crt", "server.key", nil); err != nil {
		log.Println("error is ", err)
	}
}
