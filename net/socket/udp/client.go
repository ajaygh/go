package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

var Scan = []byte{
	0xff, 0xff, 0xfd, //header
	0x0,        //version
	0x1,        //reserved
	0x01, 0x09, //sender
	0x01, 0x02, //receiver
	0x02, 0x02, //length
	0x32, 0x24, 0x02, 0x3, //timestamp
	0x01, 0x02, //packet type
	0x3,                 //icrid
	0x01, 0x0, 0x0, 0x0, //jobid
	0x02, 0x0, //casket
	0x05, 0x0, //width
	0x09, 0x0, //length
	0xa, 0x0, //height
	0x2, 0x1, 0x4, //box volume
	0x3, 0x1, 0x4, //real vol
	0x0,       //vol status
	0x01,      //weight status
	0x45, 0x3, //weight
	0x23, //input number
	//uuid
	0x02, //number of barcodes
	//barcode 1
	//barcode 2
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s  host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	var buf [512]byte
	for {
		//_, err = conn.Write([]byte("testing udp client"))
		_, err = conn.Write(Scan)
		checkError(err)
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:n]))
		time.Sleep(time.Millisecond * 100)
	}
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s\n", err.Error())
		os.Exit(1)
	}
}
