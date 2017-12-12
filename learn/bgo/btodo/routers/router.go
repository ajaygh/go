package routers

import (
	"github.com/ajaygh/learn/bgo/btodo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
