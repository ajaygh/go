package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 57600}
	s, err := serial.OpenPort(c)
	chkErr(err)
	defer s.Close()

	log.Println("opened device")

	buf := make([]byte, 1024)
	for {
		//n, err := s.Write([]byte{0x08, 0x01, 0x02, 0x03, 0x00, 0x00})
		n, err := s.Write([]byte{0x05, 0x00, 0x00, 0x00, 0x00, 0x01})
		chkErr(err)
		log.Println("written data")

		n, err = s.Read(buf)
		chkErr(err)
		log.Println("read data")
		log.Printf("% x\n", buf[:n])
	}
}

func chkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
