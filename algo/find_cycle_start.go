package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func findCycleStart(head *Node) int {
	slow, fast := head.next, head.next.next
	for slow != fast {
		slow = slow.next
		fast = fast.next.next
	}
	slow = head
	for slow != fast {
		slow = slow.next
		fast = fast.next
	}
	return slow.val
}

func main() {
	head := &Node{val: 1}
	h1 := &Node{val: 2}
	h2 := &Node{val: 3}
	h3 := &Node{val: 4}
	h4 := &Node{val: 5}
	h5 := &Node{val: 6}
	h6 := &Node{val: 7}
	h7 := &Node{val: 8}

	head.next = h1
	h1.next = h2
	h2.next = h3
	h3.next = h4
	h4.next = h5
	h5.next = h6
	h6.next = h7
	h7.next = h3

	val := findCycleStart(head)
	fmt.Println(val)
}
