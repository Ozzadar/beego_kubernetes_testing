package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/jmcvetta/napping"
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
	c.Data["appname"] = c.GetString("appname")

	c.TplNames = "applaunch.tpl"
}
