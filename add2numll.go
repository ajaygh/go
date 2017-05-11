package main

import "fmt"

type Node struct {
	left, right *Node
	val         int
}

func (node *Node) create(val int) {
	node.val = val
}

func main() {
	nd := Node{nil, nil, 55}
	nd.create(555)
	fmt.Println("vim-go ", nd, nd.val)
}
