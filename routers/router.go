package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/appLaunch/:id([0-9]+)", &controllers.AppLaunch{})
}
