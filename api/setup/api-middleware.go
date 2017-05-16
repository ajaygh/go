package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
	"math/rand"
)

const (
	Port = "5555"
	host = "127.0.0.1"
	dsi = 17
)


func main() {
	go scanTest()
	go sortTest()
	eventTest()
	os.Exit(0)
}

func getConnection() *net.UDPConn {
	service := host + ":" + Port
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	checkError(err)
	return conn
}

func initCommonData(buf []byte){
	buf[0] ,buf[1], buf[2] = 0xff, 0xff, 0xfd //header
	buf[3] = 0x00       //version
	buf[4] = 0x01       //reserved
	buf[5], buf[6] = 0x01, 0x09 //receiver
	buf[7], buf[8] = 0x01, 0x02 //sender
	buf[11],  buf[12], buf[13], buf[14] = 0x32, 0x24, 0x02, 0x3 //timestamp
}

func makeSortData(buf []byte) []byte {
	initCommonData(buf)		
	buf[9],  buf[10]= 0x02, 0x02 //length
	buf[15],  buf[16] = 0x02, 0x02 //packet type
	buf[dsi],  buf[dsi+1], buf[dsi+2], buf[dsi + 3] = 0x01, 0x02, 0x03, 0x00 //jobid
	buf[dsi+4], buf[dsi+5] = 0x02, 0x01 //casketID
	buf[dsi+6],  buf[dsi+7] = 0x05, 0x01 //chuteID
	buf[dsi+8] = 0x01 //sort status

	uuid , _ := exec.Command("uuidgen").Output()	
	buf = buf[:dsi+9]	
	buf = append(buf, uuid...)
	sorterId := byte(rand.Intn(255)) //sorter id = dsi+46
	buf = append(buf, sorterId)
	return buf
}

func makeScanData(buf []byte) []byte {
	initCommonData(buf)			
	//buf[9],  buf[10]= 0x02, 0x02 //length
	buf[15],  buf[16] = 0x01, 0x02 //packet type
	buf[dsi] = 0x03                 //icrid
	buf[dsi +1],  buf[dsi+2], buf[dsi+3], buf[dsi + 4] = 0x01, 0x02, 0x03, 0x00 //jobid	
	buf[dsi +5],  buf[dsi+6] = 0x02, 0x00 //casket
	buf[dsi +7],  buf[dsi+8] = 0x05, 0x00 //width
	buf[dsi +9],  buf[dsi+10] = 0x09, 0x00 //length
	buf[dsi +11],  buf[dsi+12] = 0xa, 0x00 //height
	buf[dsi +13],  buf[dsi+14], buf[dsi + 15] = 0x2, 0x01, 0x04 //box volume
	buf[dsi +16],  buf[dsi+17], buf[dsi + 18] = 0x3, 0x01, 0x04 //real vol
	buf[dsi +19],  buf[dsi+20] = 0x00, 0x04 //vol status
	buf[dsi +21] = 0x01       //weight status
	buf[dsi +22], buf[dsi +23] = 0x45, 0x03 //weight
	buf[dsi +24] = 0x23       //input number
	buf[dsi +25] = 0x09       //image day
	buf[dsi +26] = 0x05      //image month
	buf[dsi +27], buf[dsi +28] = 0x04, 0x04 //image year
	buf[dsi +29] = 0x02       //image hour
	buf[dsi +30] = 0x05       //image minutes
	buf[dsi +31] = 0x22       //image seconds
	buf[dsi +32], buf[dsi +33] = 0x01, 0x44 //image milliseconds
	buf[dsi +34], buf[dsi +35] = 0x01, 0x02 //image unique number

	uuid , _ := exec.Command("uuidgen").Output()	
	buf = buf[:dsi+36]	
	buf = append(buf, uuid...)
	numBarcode := byte(0x03) 
	buf = append(buf, numBarcode)
	barcodes := []byte("barcode1\nbarcode2\nbarcode3")
	lenBC := len(barcodes)
	lenTotal := dsi + 74 + lenBC
	buf[9],  buf[10]= byte(lenTotal%256), byte(lenTotal/256 )//length
	fmt.Println("scan len", lenTotal)
	buf = append(buf, barcodes...)
	return buf
}

func makeEventData(buf []byte) []byte {
	initCommonData(buf)			
	buf[15],  buf[16] = 0x04, 0x02 //packet type
	buf = buf[:dsi]
	eventStr := "event number " //event
	eventStr += strconv.Itoa(rand.Intn(500))
	buf = append(buf, []byte(eventStr)...)
	buf[9],  buf[10]= byte(len(eventStr)%256),  byte(len(eventStr)/256)//length
	fmt.Printf("event data : %s, len = %v", eventStr, len(eventStr))
	return buf
}

func scanTest() {
	conn := getConnection()

	for {
		buf := make([]byte,100)
		buf = makeScanData(buf)
		_, err := conn.Write(buf)
		checkError(err)
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:n]))
		time.Sleep(time.Millisecond * 200)
	}
}

func sortTest() {
	conn := getConnection()
	for {
		buf := make([]byte, 40, 80)
		buf = makeSortData(buf)	

		_, err := conn.Write(buf)
		checkError(err)
		n, err := conn.Read(buf[0:])
		checkError(err)
		fmt.Println(string(buf[0:n]))
		time.Sleep(time.Millisecond * 100)
	}
}

func eventTest() {
	conn := getConnection()

	for {
		buf := make([]byte, 40, 100)
		buf = makeEventData(buf)	

		_, err := conn.Write(buf)
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
