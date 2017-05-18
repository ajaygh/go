package main

import (
	"fmt"
)

type Strategy interface{
	doOperation(a, b int) int
}

type AddOperation struct{}
type SubOperation struct{}
type MulOperation struct{}

func (AddOperation) doOperation(a, b int) int {
	return a + b
}

func (SubOperation) doOperation(a, b int) int {
	return a - b
}

func (MulOperation) doOperation(a, b int) int {
	return a * b
}

type Context struct{
	strategy Strategy
}

func (ctx *Context) executeOperation(a, b int ) int {
	return ctx.strategy.doOperation(a, b)
}

func main(){
	context := &Context{AddOperation{}}
	fmt.Println("100 + 50 = ", context.executeOperation(100, 50))

	context = &Context{&SubOperation{}}
	fmt.Println("100 - 50 = ", context.executeOperation(100, 50))

	context = &Context{&MulOperation{}}
	fmt.Println("100 * 50 = ", context.executeOperation(100, 50))
}