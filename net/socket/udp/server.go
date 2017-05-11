package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":12000"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	for {
		handleClient(conn)
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s\n", err.Error())
		os.Exit(1)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte
	n, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		fmt.Println("error in read ", err.Error())
		return
	}
	fmt.Printf("Addr: %s | Data : %v \n", addr, buf[0:n])

	fmt.Printf("value %d\n", uint(buf[0])+uint(buf[1])*256)
	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
