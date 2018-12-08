/**********************************************
** @Des: login
** @Author: haodaquan
** @Date:   2017-09-07 16:30:10
** @Last Modified by:   haodaquan
** @Last Modified time: 2017-09-17 11:55:21
***********************************************/
package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	cache "github.com/patrickmn/go-cache"
	"hourManager/src/models"
	"hourManager/src/utils"
)

type LoginController struct {
	BaseController
}

//登录 TODO:XSRF过滤
// @router / [get,post]
func (this *LoginController) LoginIn() {
	if this.userId > 0 {
		this.redirect(beego.URLFor("HomeController.Index"))
	}
	beego.ReadFromRequest(&this.Controller)
	if this.isPost() {

		username := strings.TrimSpace(this.GetString("username"))
		password := strings.TrimSpace(this.GetString("password"))

		if username != "" && password != "" {
			user, err := models.GetSysUserByName(username)
			fmt.Println(user)
			flash := beego.NewFlash()
			errorMsg := ""
			if err != nil || user.Password != utils.Md5([]byte(password+user.Salt)) {
				errorMsg = "帐号或密码错误"
			} else if user.Status == 0 {
				errorMsg = "该帐号已禁用"
			} else {
				user.LastIp = this.getClientIp()
				user.LastLogin = time.Now().Unix()
				user.UpdateSysUser()
				utils.Che.Set("uid"+strconv.FormatInt(user.Id, 10), user, cache.DefaultExpiration)
				authkey := utils.Md5([]byte(this.getClientIp() + "|" + user.Password + user.Salt))
				this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 7*86400)

				this.redirect(beego.URLFor("HomeController.Index"))
			}
			flash.Error(errorMsg)
			flash.Store(&this.Controller)
			this.redirect(beego.URLFor("LoginController.LoginIn"))
		}
	}
	this.TplName = "login/login.html"
}

//登出
// @router / [get,post]
func (this *LoginController) LoginOut() {
	this.Ctx.SetCookie("auth", "")
	this.redirect(beego.URLFor("LoginController.LoginIn"))
}

func (this *LoginController) NoAuth() {
	this.Ctx.WriteString("没有权限")
}
