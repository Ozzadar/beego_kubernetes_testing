package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/jmcvetta/napping"
	"fmt"
)

type MainController struct {
	beego.Controller
}
type AppLaunch struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "Klouds.io"
	c.Data["Email"] = "faddat@gmail.com"
	c.TplNames = "index.tpl"
}

func (c *AppLaunch) Get() {
	c.Data["Website"] = "Klouds.io"
	c.Data["Email"] = "faddat@gmail.com"
	c.TplNames = "applaunch.tpl"
}

func (c *AppLaunch) Post() {
	//c.ParseForm()

	appname := c.GetString("appname")

	c.Data["Website"] = "Klouds.io"
	c.Data["Email"] = "faddat@gmail.com"

	fmt.Println("App Name: " ,appname)
	c.TplNames = "applaunch.tpl"
	c.Ctx.Output.Body([]byte("Launching " + appname + "!"))
}
