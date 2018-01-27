package main

import (
	"fmt"
	"math"
)

type calculator struct {
	acc int
}

type op func(int) int

func (c *calculator) add(num int) op {
	return func(val int) int {
		return val + num
	}
}

func (c *calculator) sub(num int) op {
	return func(val int) int {
		return val - num
	}
}

func (c *calculator) mul(num int) op {
	return func(val int) int {
		return val * num
	}
}

func (c *calculator) div(num int) op {
	return func(val int) int {
		return val / num
	}
}

func (c *calculator) sqrt() op {
	return func(val int) int {
		return int(math.Sqrt(float64(val)))
	}
}

func (c *calculator) do(op op) {
	c.acc = op(c.acc)
}
func main() {
	c := calculator{}
	fmt.Println("accum", c.acc)
	c.do(c.add(5))
	fmt.Println("accum", c.acc)
	c.do(c.sub(2))
	fmt.Println("accum", c.acc)
	c.do(c.mul(3))
	fmt.Println("accum", c.acc)
	c.do(c.div(4))
	fmt.Println("accum", c.acc)
	c.do(c.sqrt())
	fmt.Println("accum", c.acc)
}
