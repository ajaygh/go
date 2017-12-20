package serial

import (
	"errors"
	"log"

	"github.com/tarm/serial"
)

var s *serial.Port

func init() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 57600}
	tmp, err := serial.OpenPort(c)
	chkErr(err)
	s = tmp
	go rcv()
}

// Send sends data to serial
func Send(data []byte) (err error) {
	log.Printf("sending data % x\n", data)
	if s == nil {
		return
	}
	n, err := s.Write(data)

	if err != nil {
		return
	}
	if n != len(data) {
		return errors.New("incomplete  send")
	}
	return
}

func rcv() {
	buf := make([]byte, 1024)
	for {
		if s == nil {
			break
		}
		n, err := s.Read(buf)
		if err == nil {
			log.Printf("received data % x\n", buf[:n])
		}
	}
}

func chkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
