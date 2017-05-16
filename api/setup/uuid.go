package main

import (
	"fmt"
	"os/exec"
	"log"
	"net"
)

func main(){
	uuid, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
		return 
	}
	fmt.Printf("uuid is %s len is %v\n", uuid, len(uuid))
	udpAddr, err := net.ResolveUDPAddr("udp4", "localhost:8080")
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	fmt.Printf("udpAddr %T\n", udpAddr)
	fmt.Printf("udpConn type %T\n", udpConn)
	
}