package main

import "fmt"
import "math"

type Point struct {
	x, y int
}
type Centre struct {
	Point
}

type shape interface {
	area() float64
}

type Circle struct {
	centre Centre
	radius float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	P := new(Point)
	c := Centre{*P}
	c.x = 4
	c.y = 8
	cir := Circle{c, 5}
	fmt.Printf("Point coordinate are x = %v, y = %v\n", c.x, c.y)
	fmt.Println("Circle area is ", cir.area())
}
