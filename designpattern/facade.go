package main

import "fmt"

type Shape interface {
	draw()
}

type Circle struct{}
type Square struct{}
type Rectangle struct{}

func (Circle) draw() {
	fmt.Println("Circle :: draw method called.")
}

func (Square) draw() {
	fmt.Println("Square :: draw method called.")
}
func (Rectangle) draw() {
	fmt.Println("Rectangle :: draw method called.")
}

type ShapeMaker struct {
	circle, square, rectangle Shape
}

func (shapeMaker *ShapeMaker) drawCircle() {
	shapeMaker.circle.draw()
}

func (shapeMaker *ShapeMaker) drawSquare() {
	shapeMaker.square.draw()
}

func (shapeMaker *ShapeMaker) drawRectangle() {
	shapeMaker.rectangle.draw()
}

func main() {
	shapeMaker := ShapeMaker{Circle{}, Square{}, Rectangle{}}
	shapeMaker.drawCircle()
	shapeMaker.drawSquare()
	shapeMaker.drawRectangle()

	fmt.Println(`Facade, a wrapper to hide something,
	 basically hides the system complexity and provides an
	  interface to client to access the whole system.`)
}
