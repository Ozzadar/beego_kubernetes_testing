package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
	_ "github.com/jmcvetta/napping"
)

func main() {
	beego.Run()
}

