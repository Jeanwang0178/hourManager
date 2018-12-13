/**
* @Project: hourManager
* @Package
* @Description: 用户管理
* @author : wj
* @date Date : 2018/12/06/ 16:06
* @version V1.0
 */
package controllers

import (
	"PPGo_ApiAdmin/libs"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/patrickmn/go-cache"
	"hourManager/src/models"
	"hourManager/src/utils"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	BaseController
}

func (self *UserController) List() {
	self.Data["pageTitle"] = "用户管理"
	self.display()
}

func (self *UserController) Add() {
	self.Data["pageTitle"] = "新增用户"

	// 角色
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.GetSysRoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}

	self.Data["role"] = list

	self.display()
}

func (self *UserController) Edit() {
	self.Data["pageTitle"] = "编辑用户"

	id, _ := self.GetInt64("id")
	sysUser, _ := models.GetSysUserById(id)
	row := make(map[string]interface{})
	row["id"] = sysUser.Id

	row["company_name"] = sysUser.CompanyName
	row["login_name"] = sysUser.LoginName
	row["real_name"] = sysUser.RealName
	row["phone"] = sysUser.Phone
	row["email"] = sysUser.Email
	row["role_ids"] = sysUser.RoleIds
	self.Data["user"] = row

	role_ids := strings.Split(sysUser.RoleIds, ",")

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.GetSysRoleList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["checked"] = 0
		for i := 0; i < len(role_ids); i++ {
			role_id, _ := strconv.ParseInt(role_ids[i], 10, 64)
			if role_id == v.Id {
				row["checked"] = 1
			}
			fmt.Println(role_ids[i])
		}
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}
	self.Data["role"] = list
	self.display()
}

func (self *UserController) AjaxSave() {
	userId, _ := self.GetInt64("id")
	if userId == 0 {
		user := new(models.SysUser)
		user.CompanyName = strings.TrimSpace(self.GetString("company_name"))
		user.LoginName = strings.TrimSpace(self.GetString("login_name"))
		user.RealName = strings.TrimSpace(self.GetString("real_name"))
		user.Phone = strings.TrimSpace(self.GetString("phone"))
		user.Email = strings.TrimSpace(self.GetString("email"))
		user.RoleIds = strings.TrimSpace(self.GetString("roleids"))
		user.UpdateTime = time.Now().Unix()
		user.UpdateId = self.userId
		user.Status = 1

		// 检查登录名是否已经存在
		_, err := models.GetSysUserByName(user.LoginName)

		if err == nil {
			self.ajaxMsg("登录名已经存在", MSG_ERR)
		}
		//新增
		pwd, salt := libs.Password(4, "")
		user.Password = pwd
		user.Salt = salt
		user.CreateTime = time.Now().Unix()
		user.CreateId = self.userId
		if _, err := models.AddSysUser(user); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	newUser, _ := models.GetSysUserById(userId)
	//修改
	newUser.Id = userId
	newUser.UpdateTime = time.Now().Unix()
	newUser.UpdateId = self.userId
	newUser.CompanyName = strings.TrimSpace(self.GetString("company_name"))
	newUser.LoginName = strings.TrimSpace(self.GetString("login_name"))
	newUser.RealName = strings.TrimSpace(self.GetString("real_name"))
	newUser.Phone = strings.TrimSpace(self.GetString("phone"))
	newUser.Email = strings.TrimSpace(self.GetString("email"))
	newUser.RoleIds = strings.TrimSpace(self.GetString("roleids"))
	newUser.UpdateTime = time.Now().Unix()
	newUser.UpdateId = self.userId
	newUser.Status = 1

	resetPwd, _ := self.GetInt("reset_pwd")
	if resetPwd == 1 {
		pwd, salt := libs.Password(4, "")
		newUser.Password = pwd
		newUser.Salt = salt
	}
	if err := newUser.UpdateSysUser(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg(strconv.Itoa(resetPwd), MSG_OK)
}

func (self *UserController) AjaxDel() {

	userId, _ := self.GetInt64("id")
	status := strings.TrimSpace(self.GetString("status"))
	if userId == 1 {
		self.ajaxMsg("超级管理员不允许操作", MSG_ERR)
	}

	userStatus := 0
	if status == "enable" {
		userStatus = 1
	}
	user, _ := models.GetSysUserById(userId)
	user.UpdateTime = time.Now().Unix()
	user.Status = userStatus
	user.Id = userId

	if err := user.UpdateSysUser(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *UserController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	realName := strings.TrimSpace(self.GetString("realName"))

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	//
	if realName != "" {
		filters = append(filters, "real_name__icontains", realName)
	}
	result, count := models.GetSysUserList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["company_name"] = v.CompanyName
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

func (this *UserController) Modify() {
	this.Data["pageTitle"] = "资料修改"
	id := this.userId
	user, _ := models.GetSysUserById(id)
	row := make(map[string]interface{})
	row["id"] = user.Id
	row["company_name"] = user.CompanyName
	row["login_name"] = user.LoginName
	row["real_name"] = user.RealName
	row["phone"] = user.Phone
	row["email"] = user.Email
	this.Data["user"] = row
	utils.Che.Set("uid"+strconv.FormatInt(this.user.Id, 10), nil, cache.DefaultExpiration)
	this.display()
}

func (this *UserController) AjaxModify() {
	userId, _ := this.GetInt64("id")
	user, _ := models.GetSysUserById(userId)
	//修改
	user.Id = userId
	user.UpdateTime = time.Now().Unix()
	user.UpdateId = this.userId
	user.CompanyName = strings.TrimSpace(this.GetString("company_name"))
	user.LoginName = strings.TrimSpace(this.GetString("login_name"))
	user.RealName = strings.TrimSpace(this.GetString("real_name"))
	user.Phone = strings.TrimSpace(this.GetString("phone"))
	user.Email = strings.TrimSpace(this.GetString("email"))

	resetPwd := this.GetString("reset_pwd")
	if resetPwd == "1" {
		pwdOld := strings.TrimSpace(this.GetString("password_old"))
		pwdOldMd5 := utils.Md5([]byte(pwdOld + user.Salt))
		if user.Password != pwdOldMd5 {
			this.ajaxMsg("旧密码错误", MSG_ERR)
		}

		pwdNew1 := strings.TrimSpace(this.GetString("password_new1"))
		pwdNew2 := strings.TrimSpace(this.GetString("password_new2"))

		if pwdNew1 != pwdNew2 {
			this.ajaxMsg("两次密码不一致", MSG_ERR)
		}

		pwd, salt := utils.Password(4, pwdNew1)
		user.Password = pwd
		user.Salt = salt
	}
	user.UpdateTime = time.Now().Unix()
	user.UpdateId = this.userId
	user.Status = 1

	if err := user.UpdateSysUser(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
