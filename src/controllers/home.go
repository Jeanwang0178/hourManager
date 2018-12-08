/**
* @Project: hourManager
* @Package controllers
* @Description: 首页管理
* @author : wj
* @date Date : 2018/12/06/ 17:00
* @version V1.0
 */
package controllers

type HomeController struct {
	BaseController
}

// @router / [get]
func (self *HomeController) Index() {
	self.Data["pageTitle"] = "系统首页"
	self.TplName = "layout/main.html"
}
