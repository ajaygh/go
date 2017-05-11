package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	Port = "5555"
	host = "127.0.0.1"
)

var scan = []byte{
	0xff, 0xff, 0xfd, //header
	0x00,       //version
	0x01,       //reserved
	0x01, 0x09, //receiver
	0x01, 0x02, //sender
	0x02, 0x02, //length
	0x32, 0x24, 0x02, 0x3, //timestamp
	0x01, 0x02, //packet type
	0x3,                    //icrid
	0x01, 0x00, 0x00, 0x00, //jobid
	0x02, 0x00, //casket
	0x05, 0x00, //width
	0x09, 0x00, //length
	0xa, 0x00, //height
	0x2, 0x01, 0x04, //box volume
	0x3, 0x01, 0x04, //real vol
	0x00, 0x04, //vol status
	0x01,       //weight status
	0x45, 0x03, //weight
	0x23,       //input number
	0x09,       //image day
	0x05,       //image month
	0x04, 0x04, //image year
	0x02,       //image hour
	0x05,       //image minutes
	0x22,       //image seconds
	0x01, 0x44, //image milliseconds
	0x01, 0x02, //image unique number
	//@uuid
	0x02, //number of barcodes
	//barcode 1
	//barcode 2
}

var sorterSort = []byte{
	0xff, 0xff, 0xfd, //header
	0x00,       //version
	0x01,       //reserved
	0x01, 0x09, //receiver
	0x01, 0x02, //sender
	0x02, 0x02, //length
	0x32, 0x24, 0x02, 0x3, //timestamp
	0x02, 0x02, //packet type
	0x01, 0x02, 0x03, 0x00, //jobid
	0x02, 0x01, //casketID
	0x05, 0x01, //chuteID
	0x01,                                                                                                          //sort status
	1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, //uuid
	0x04, //sorter id
}

func main() {
	//scanTest()
	sortTest()
	os.Exit(0)
}

func scanTest() {
	service := host + ":" + Port
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	var buf [512]byte
	for {
		_, err = conn.Write(scan)
		checkError(err)
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:n]))
		time.Sleep(time.Millisecond * 100)
	}
}

func sortTest() {
	service := host + ":" + Port
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)

	var buf [512]byte
	for {
		_, err = conn.Write(sorterSort)
		checkError(err)
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:n]))
		time.Sleep(time.Millisecond * 100)
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s\n", err.Error())
		os.Exit(1)
	}
}
