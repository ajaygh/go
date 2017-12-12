package main

import (
	"fmt"
)
const sz = 7
func prt(d [sz]int, s [sz]bool) {
	for i := 0; i < len(d); i++ {
		if s[i] {
			fmt.Print(d[i], " ")
		}
	}
	fmt.Println()
}

func printAll(d [sz]int, s [sz]bool, pos int) {
	if pos == len(d)-1 {
		prt(d, s)
		s[pos] = true
		prt(d, s)
		s[pos] = false
		return
	}
	printAll(d, s, pos+1)
	s[pos] = true
	printAll(d, s, pos+1)
	s[pos] = false
}

func main() {
	d := [sz]int{1, 2, 3, 4, 5, 6, 7}
	status := make([sz]bool, len(d))
	printAll(d, status, 0)
}
