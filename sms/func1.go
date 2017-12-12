// Function as first class citizen.
package main

import (
	"fmt"
	"math"
)

type calculator struct {
	accum float64
}

type opfunc func(float64) float64

func (c *calculator) Do(of opfunc) {
	c.accum = of(c.accum)
}

func add(n float64) opfunc {
	return func(accum float64) float64 {
		return accum + n
	}
}

func mul(n float64) opfunc {
	return func(accum float64) float64 {
		return accum * n
	}
}

func main() {
	var c calculator
	c.Do(add(5))
	fmt.Println(c.accum)
	c.Do(mul(5))
	fmt.Println(c.accum)
	c.Do(math.Sqrt)
	fmt.Println(c.accum)
}
