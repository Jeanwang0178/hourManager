/**
* @Project: hourManager
* @Package
* @Description: TODO
* @author : wj
* @date Date : 2018/12/08/ 13:46
* @version V1.0
 */
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/bee/logger"
	"hourManager/src/models"
	"strconv"
	"strings"
	"time"
)

type ProjectUserController struct {
	BaseController
}

func (this *ProjectUserController) List() {
	this.Data["pageTitle"] = "项目用户管理"
	this.display()
}

func (this *ProjectUserController) ListUser() {
	this.Data["pageTitle"] = "用户信息"
	tplName := "project/list_user"
	this.Data["isExists"] = this.GetString("isExists")
	this.Data["projectId"] = this.GetString("projectId")
	this.display(tplName)
}

func (this *ProjectUserController) Add() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "新增项目用户"
	this.display()
}

func (this *ProjectUserController) AjaxSave() {

	projectId, _ := this.GetInt64("project_id")
	userIds := this.GetString("user_ids")
	userIdSlice := strings.Split(userIds, ",")

	proUsers := make([]models.SysProjectUser, 0)

	for _, v := range userIdSlice {

		userId, _ := strconv.ParseInt(v, 10, 64)
		proUser := new(models.SysProjectUser)
		proUser.ProjectId = projectId
		proUser.UserId = userId
		proUser.CreateTime = time.Now()
		proUser.UpdateTime = time.Now()
		proUser.Status = 1
		//新增
		proUser.CreateTime = time.Now()
		proUser.UpdateTime = time.Now()
		proUser.CreateId = this.userId
		proUser.UpdateId = this.userId
		proUsers = append(proUsers, *proUser)
	}

	if id, err := models.AddSysProjectUserList(proUsers); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	} else {
		beeLogger.Log.Infof("add porject user :", id)
	}
	this.ajaxMsg("", MSG_OK)

}

func (this *ProjectUserController) AjaxDel() {

	role_id, _ := this.GetInt64("id")
	proUser, _ := models.GetSysProjectUserById(role_id)
	proUser.Status = 0
	proUser.Id = role_id
	proUser.UpdateTime = time.Now()

	if err := proUser.UpdateSysProjectUser(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ProjectUserController) TableUser() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	isExists := strings.TrimSpace(this.GetString("isExists"))
	projectId, _ := this.GetInt64("projectId")
	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	this.pageSize = limit
	//查询条件
	filters := make(map[string]interface{})
	//
	if isExists != "" {
		filters["isExists"] = isExists
	}
	filters["projectId"] = projectId
	result, count := models.GetSysUserListByParam(page, this.pageSize, filters)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.LoginName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_time"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["update_time"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ProjectUserController) ProUserList() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}

	projectId, _ := this.GetInt64("projectId")

	StatusText := make(map[int]string)
	StatusText[0] = "<font color='red'>禁用</font>"
	StatusText[1] = "正常"

	this.pageSize = limit
	//查询条件
	result, count := models.GetSysProjectUserInfoByParam(page, this.pageSize, projectId)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["login_name"] = v.LoginName
		row["company_name"] = v.CompanyName
		row["real_name"] = v.RealName
		row["phone"] = v.Phone
		row["email"] = v.Email
		row["role_ids"] = v.RoleIds
		row["create_user"] = v.CreateUser
		row["create_time"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["update_time"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		row["status"] = v.Status
		row["status_text"] = StatusText[v.Status]
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ProjectUserController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 300
	}

	projectId := strings.TrimSpace(this.GetString("projectId"))
	this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if projectId != "" {
		beeLogger.Log.Infof("quey projectId :{}", projectId)
		filters = append(filters, "projectId", projectId)
	}
	result, count := models.GetSysProjectUserList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["project_id"] = v.ProjectId
		row["user_id"] = v.UserId
		row["create_time"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["update_time"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}
