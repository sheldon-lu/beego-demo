package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AccessController{})
	beego.AutoRouter(&controllers.ApisController{})
}
