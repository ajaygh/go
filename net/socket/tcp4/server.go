package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":10003"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println("tcpAddr : ", tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	fmt.Println("listener : ", listener)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "invalid conn received : %s", err.Error())
			continue
		}
		fmt.Println("conn : ", conn)
		go handleClient(conn)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error :%s", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
