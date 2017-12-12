package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	t := make([]byte, 0)
	t = append(t, 254)
	t = append(t, 254)
	d := binary.BigEndian.Uint16(t)

	fmt.Println(d)
}
