package main

import (
	_ "beego-demo/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true




	beego.BConfig.WebConfig.TemplateLeft = "[[%"
	beego.BConfig.WebConfig.TemplateRight = "%]]"
	beego.Run()
}

