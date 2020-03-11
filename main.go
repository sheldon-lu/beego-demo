package main

import (
	_ "beego-demo/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.InsertFilter("/*", beego.BeforeRouter, func(ctx *context.Context) {
	// 	cookie, err := ctx.Request.Cookie("Authorization")
	//	if err != nil || !hjwt.CheckToken(cookie.Value) {
	//		http.Redirect(ctx.ResponseWriter, ctx.Request, "/", http.StatusMovedPermanently)
	//	}
	//})

	beego.BConfig.WebConfig.TemplateLeft = "[[%"
	beego.BConfig.WebConfig.TemplateRight = "%]]"
	beego.Run()
}
