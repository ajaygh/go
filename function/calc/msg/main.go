package main

import (
	"fmt"
	"io"
	"net"
)

type mux struct {
	ops chan func(map[net.Addr]net.Conn)
}

func (m *mux) add(conn net.Conn) {
	m.ops <- func(m map[net.Addr]net.Conn) {
		m[conn.RemoteAddr()] = conn
	}
}

func (m *mux) remove(conn net.Conn) {
	m.ops <- func(m map[net.Addr]net.Conn) {
		delete(m, conn.RemoteAddr())
	}
}

func (m *mux) sendMsg(msg string) error {
	result := make(chan error, 1)
	m.ops <- func(m map[net.Addr]net.Conn) {
		for _, conn := range m {
			_, err := io.WriteString(conn, msg)
			if err != nil {
				result <- err
				return
			}
			result <- nil
		}
	}
	return <-result
}

func (m *mux) sendPvtMsg(msg string, addr net.Addr) error {
	result := make(chan net.Conn, 1)
	m.ops <- func(m map[net.Addr]net.Conn) {
		result <- m[addr]
	}
	conn := <-result
	if conn == nil {
		return fmt.Errorf("client %v not registered", addr)
	}
	_, err := io.WriteString(conn, msg)
	return err
}

func (m *mux) loop() {
	conns := make(map[net.Addr]net.Conn)
	for {
		op := <-m.ops
		op(conns)
	}
}
func main() {
	m := mux{ops: make(chan func(map[net.Addr]net.Conn))}
	go m.loop()
}
