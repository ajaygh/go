package main

import (
	"fmt"

	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("vim-go")
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello beego"))
	})
	beego.Run()
}
