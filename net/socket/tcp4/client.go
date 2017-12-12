package main

import (
	"fmt"
	"net"
	"os"
	//	"time"
)

func main() {
	//Read server address
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage %s host:port", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	fmt.Println("tcpAddr : ", tcpAddr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	defer conn.Close()

	fmt.Println("client conn : ", conn)

	buf := make([]byte, 1000)
	for {
		_, err = conn.Read(buf)
		checkError(err)
		fmt.Println("received : ", buf)
		//	_, err = conn.Write([]byte(fmt.Sprintln("HEAD from client %s", os.Args[2])))
		//	checkError(err)
		//	fmt.Println("Write done")
		//		time.Sleep(time.Millisecond * 3)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, " Fatal Error: %s\n", err.Error())
		os.Exit(1)
	}
}
