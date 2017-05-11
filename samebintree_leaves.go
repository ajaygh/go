package main

import "fmt"

type Tree struct {
	left  *Tree
	val   int
	right *Tree
}

func getLeaves(t *Tree) l {
	l := make([]int)
	//create stack
}
func sameLeaf(t1, t2 *Tree) bool {
	l1 = getLeaves(t1)
	l2 = getLeaves(t2)
	if l1 != l2 {
		fmt.Println("Leaves not same ", len(l1), len(l2))
		return false
	} else {
		for i := 0; i < len(l1); i++ {
			if l1[i] != l2[i] {
				fmt.Println("leaves mismatch at index ", i)
				return false
			}
		}
		return true
	}
}

func main() {
	//create tree t1
	//create tree t2
	res := sameLeaf(t1, t2)
	if res {
		fmt.Println("leaves are same.")
	} else {
		fmt.Println("leaves are not same.")
	}
}
