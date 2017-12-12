package main

import "fmt"
import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Ctx.WriteString("hello beego")
}
func main() {
	fmt.Println("beego learn")
	beego.Router("/hello", &MainController{})
	beego.Run()
}
