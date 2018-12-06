package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:ComUserController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"],
		beego.ControllerComments{
			Method:           "Index",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"],
		beego.ControllerComments{
			Method:           "GetTime",
			Router:           `/gettime`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"],
		beego.ControllerComments{
			Method:           "PostLogin",
			Router:           `/postLogin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"] = append(beego.GlobalControllerRouter["hourManager/src/controllers:DefaultController"],
		beego.ControllerComments{
			Method:           "Profile",
			Router:           `/profile`,
			AllowHTTPMethods: []string{"get", "post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
