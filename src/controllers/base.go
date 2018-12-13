package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/bee/logger"
	"github.com/patrickmn/go-cache"
	"hourManager/src/models"
	"hourManager/src/utils"
	"strconv"
	"strings"
	"webcron/app/libs"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.SysUser
	userId         int64
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
}

func (this *BaseController) Prepare() {
	this.pageSize = 20
	controllerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.Data["version"] = beego.AppConfig.String("version")
	this.Data["siteName"] = beego.AppConfig.String("site.name")
	this.Data["curRoute"] = this.controllerName + "." + this.actionName
	this.Data["curController"] = this.controllerName
	this.Data["curAction"] = this.actionName
	if (strings.Compare(this.controllerName, "apidoc")) != 0 {
		this.auth()
	}

	this.Data["loginUserId"] = this.userId
	this.Data["loginUserName"] = this.userName
}

//登录权限验证
func (this *BaseController) auth() {
	arr := strings.Split(this.Ctx.GetCookie("auth"), "|")
	this.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.ParseInt(idstr, 10, 64)
		if userId > 0 {
			var err error

			cheUser, found := utils.Che.Get("uid" + idstr)
			user := &models.SysUser{}
			if found && cheUser != nil { //从缓存取用户
				user = cheUser.(*models.SysUser)
			} else {
				user, err = models.GetSysUserById(userId)
				utils.Che.Set("uid"+idstr, user, cache.DefaultExpiration)
			}
			if err == nil && password == libs.Md5([]byte(this.getClientIp()+"|"+user.Password+user.Salt)) {
				this.userId = user.Id

				this.loginName = user.LoginName
				this.userName = user.RealName
				this.user = user
				this.AdminAuth()
			}

			isHasAuth := strings.Contains(this.allowUrl, this.controllerName+"/"+this.actionName)
			//不需要权限检查
			noAuth := "ajaxmodify/ajaxsave/ajaxdel/table/tableuser/listuser/prouserlist/list/loginin/loginout/getnodes/start/show/ajaxapisave/index/group/public/env/code/apidetail"
			isNoAuth := strings.Contains(noAuth, this.actionName)
			if isHasAuth == false && isNoAuth == false {
				//this.Ctx.WriteString("没有权限")
				beeLogger.Log.Error("没有权限")
				this.ajaxMsg("没有权限", MSG_ERR)
				return
			}
		}
	}

	if this.userId == 0 && (this.controllerName != "login" && this.actionName != "loginin") {
		this.redirect(beego.URLFor("LoginController.LoginIn"))
	}
}

func (this *BaseController) AdminAuth() {
	cheMen, found := utils.Che.Get("menu" + strconv.FormatInt(this.user.Id, 10))
	if found && cheMen != nil { //从缓存取菜单
		menu := cheMen.(*CheMenu)
		//fmt.Println("调用显示菜单")
		this.Data["SideMenu1"] = menu.FirstMenu  //一级菜单
		this.Data["SideMenu2"] = menu.SecondMenu //二级菜单
		this.allowUrl = menu.AllowUrl
	} else {
		// 左侧导航栏
		filters := make([]interface{}, 0)
		filters = append(filters, "status", 1)
		if this.userId != 1 {
			//普通管理员
			adminAuthIds, _ := models.GetSysRoleAuthByIds(this.user.RoleIds)
			adminAuthIdArr := strings.Split(adminAuthIds, ",")
			filters = append(filters, "id__in", adminAuthIdArr)
		}
		result, _ := models.GetSysAuthList(1, 1000, filters...)
		levelFirst := make([]map[string]interface{}, len(result))
		levelSecond := make([]map[string]interface{}, len(result))
		allow_url := ""
		i, j := 0, 0
		for _, v := range result {
			if v.AuthUrl != " " || v.AuthUrl != "/" {
				allow_url += v.AuthUrl
			}
			row := make(map[string]interface{})
			if v.Pid == 1 && v.IsShow == 1 {
				row["Id"] = int(v.Id)
				row["Sort"] = v.Sort
				row["AuthName"] = v.AuthName
				row["AuthUrl"] = v.AuthUrl
				row["Icon"] = v.Icon
				row["Pid"] = int(v.Pid)
				levelFirst[i] = row
				i++
			}
			if v.Pid != 1 && v.IsShow == 1 {
				row["Id"] = int(v.Id)
				row["Sort"] = v.Sort
				row["AuthName"] = v.AuthName
				row["AuthUrl"] = v.AuthUrl
				row["Icon"] = v.Icon
				row["Pid"] = int(v.Pid)
				levelSecond[j] = row
				j++
			}
		}
		this.Data["SideMenu1"] = levelFirst[:i]  //一级菜单
		this.Data["SideMenu2"] = levelSecond[:j] //二级菜单

		this.allowUrl = allow_url + "/home/index"
		cheM := &CheMenu{}
		cheM.AllowUrl = this.allowUrl
		cheM.FirstMenu = this.Data["SideMenu1"].([]map[string]interface{})
		cheM.SecondMenu = this.Data["SideMenu2"].([]map[string]interface{})
		utils.Che.Set("menu"+strconv.FormatInt(this.user.Id, 10), cheM, cache.DefaultExpiration)
	}

}

type CheMenu struct {
	FirstMenu  []map[string]interface{}
	SecondMenu []map[string]interface{}
	AllowUrl   string
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (this *BaseController) getClientIp() string {
	s := this.Ctx.Request.RemoteAddr
	l := strings.LastIndex(s, ":")
	return s[0:l]
}

// 重定向
func (this *BaseController) redirect(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//渲染模板
func (this *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = this.controllerName + "/" + this.actionName + ".html"
	}
	this.Layout = "layout/layout.html"
	this.TplName = tplname
}

//ajax返回
func (this *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//ajax返回 列表
func (this *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}

//显示错误信息
func (this *BaseController) showMsg(args ...string) {
	this.Data["message"] = args[0]
	redirect := this.Ctx.Request.Referer()
	if len(args) > 1 {
		redirect = args[1]
	}

	this.Data["redirect"] = redirect
	this.Data["pageTitle"] = "系统提示"
	this.display("error/message")
	this.Render()
	this.StopRun()
}

//输出JSON
func (this *BaseController) jsonResult(out interface{}) {
	this.Data["json"] = out
	this.ServeJSON()
	this.StopRun()
}
