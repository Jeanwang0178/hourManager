package routers

import (
	"github.com/astaxie/beego"

	"hourManager/src/controllers"
)

func init() {

	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")

}
