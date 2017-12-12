package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":62000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println("tcpAddr : ", tcpAddr)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	defer listener.Close()
	fmt.Println("listener : ", listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection : %s", err.Error())
			os.Exit(1)
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
	conn.Write([]byte("sent data "))
	buf := make([]byte, 1024)
	len, err := conn.Read(buf)
	fmt.Println("Read len: ", len, "Content: ", buf)
	checkError(err)
	time.Sleep(time.Second * 4)
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
