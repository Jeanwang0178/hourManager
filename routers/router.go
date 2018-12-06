package routers

import (
	"github.com/astaxie/beego"
	"hourManager/src/controllers"
)

func init() {
	ns := beego.NewNamespace("/hour",
		beego.NSNamespace("/default",
			beego.NSInclude(
				&controllers.DefaultController{},
			),
		),
	)

	beego.AddNamespace(ns)

}
