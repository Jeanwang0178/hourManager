package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/beego/bee/logger"
	"hourManager/src/common"
	"hourManager/src/services"
	"strings"
	"time"
	"webcron/app/libs"
)

type DefaultController struct {
	BaseController
}

// @router / [get]
func (ctl *DefaultController) Index() {
	ctl.Data["Website"] = "beego.me"
	ctl.Data["Email"] = "wjian0124@163.com"
	ctl.TplName = "index.html"
	ctl.display()
}

// 获取系统时间
// @router /gettime [get]
func (this *DefaultController) GetTime() {
	out := make(map[string]interface{})
	out["time"] = time.Now().UnixNano() / int64(time.Millisecond)
	this.jsonResult(out)
}

// 登录
// @router /login [get]
func (this *DefaultController) Login() {
	/*if len(this.userId) > 0 {
		this.redirect("../index.html")
	}*/
	beego.ReadFromRequest(&this.Controller)
	this.TplName = "default/login.html"
}

// 登录
// @router /postLogin [post]
func (this *DefaultController) PostLogin() {

	response := make(map[string]interface{})

	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))
		remember := this.GetString("remember")
		if username != "" && password != "" {
			user, err := services.UserServiceUserGetByName(username)
			errorMsg := ""
			ss := libs.Md5([]byte(password + user.Salt))
			beeLogger.Log.Info(ss)
			if err != nil || user.Password != libs.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == -1 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = int(time.Now().Unix())
				services.UserServiceUserUpdate(user)

				authkey := libs.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				token := ""
				token = user.Id + "|" + authkey
				if remember == "yes" {
					this.Ctx.SetCookie("token", token, 7*86400)
				} else {
					this.Ctx.SetCookie("token", token)
				}

				this.redirect(beego.URLFor("DefaultController.Index"))
			}
			if errorMsg != "" {
				response["code"] = common.FailedCode
				response["msg"] = errorMsg
			} else {
				response["code"] = common.SuccessCode
				response["msg"] = common.SuccessMsg
			}
		}
	}

	this.Data["json"] = response
	this.ServeJSON()

}

// 退出登录
// @router /logout [get]
func (this *DefaultController) Logout() {
	this.Ctx.SetCookie("token", "")
	this.redirect(beego.URLFor("DefaultController.Login"))
}

// 个人信息
//@router /profile [get,post]
func (this *DefaultController) Profile() {
	beego.ReadFromRequest(&this.Controller)
	user, _ := services.UserServiceUserGetById("")

	if this.isPost() {
		flash := beego.NewFlash()
		user.Email = this.GetString("email")
		services.UserServiceUserUpdate(user)
		password1 := this.GetString("password1")
		password2 := this.GetString("password2")
		if password1 != "" {
			if len(password1) < 6 {
				flash.Error("密码长度必须大于6位")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else if password2 != password1 {
				flash.Error("两次输入的密码不一致")
				flash.Store(&this.Controller)
				this.redirect(beego.URLFor(".Profile"))
			} else {
				user.Salt = string(utils.RandomCreateBytes(10))
				user.Password = libs.Md5([]byte(password1 + user.Salt))
				services.UserServiceUserUpdate(user)
			}
		}
		flash.Success("修改成功！")
		flash.Store(&this.Controller)
		this.redirect(beego.URLFor(".Profile"))
	}

	this.Data["pageTitle"] = "个人信息"
	this.Data["user"] = user
	this.display()
}
