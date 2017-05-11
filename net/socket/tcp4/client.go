package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
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
	fmt.Println("client conn : ", conn)

	_, err = conn.Write([]byte("HEAD from client"))
	checkError(err)
	fmt.Println("Write done")

	result, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println("Read done")
	fmt.Println(string(result))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, " Fatal Error: %s\n", err.Error())
		os.Exit(1)
	}
}
