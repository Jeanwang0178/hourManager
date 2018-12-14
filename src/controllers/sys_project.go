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
	"hourManager/src/models"
	"strings"
	"time"
)

type ProjectController struct {
	BaseController
}

func (this *ProjectController) List() {
	this.Data["pageTitle"] = "项目设置"
	this.display()
}

func (this *ProjectController) Add() {
	this.Data["pageTitle"] = "新增项目"
	this.display()
}

func (this *ProjectController) Edit() {
	this.Data["pageTitle"] = "编辑项目"

	id, _ := this.GetInt64("id")
	project, _ := models.GetSysProjectById(id)
	row := make(map[string]interface{})
	row["id"] = project.Id
	row["project_name"] = project.ProjectName
	row["detail"] = project.Detail
	this.Data["project"] = row
	this.display()
}

func (this *ProjectController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 30
	}
	projectName := strings.TrimSpace(this.GetString("projectName"))

	this.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if projectName != "" {
		filters = append(filters, "project_name__icontains", projectName)
	}
	result, count := models.GetSysProjectList(page, this.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["project_name"] = v.ProjectName
		row["detail"] = v.Detail
		row["create_time"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["update_time"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("成功", MSG_OK, count, list)
}

func (this *ProjectController) AjaxSave() {
	projectId, _ := this.GetInt64("id")
	if projectId == 0 {
		project := new(models.SysProject)

		project.ProjectName = strings.TrimSpace(this.GetString("project_name"))
		project.Detail = strings.TrimSpace(this.GetString("detail"))
		project.CreateId = this.userId
		project.UpdateId = this.userId
		project.CreateTime = time.Now()
		project.UpdateTime = time.Now()
		project.Status = 1

		_, err := models.GetSysProjectByName(project.ProjectName)

		if err == nil {
			this.ajaxMsg("项目名称已经存在", MSG_ERR)
		}

		if _, err := models.AddSysProject(project); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	proUpdate, _ := models.GetSysProjectById(projectId)
	// 修改
	proUpdate.ProjectName = strings.TrimSpace(this.GetString("project_name"))
	proUpdate.Detail = strings.TrimSpace(this.GetString("detail"))
	proUpdate.UpdateId = this.userId
	proUpdate.UpdateTime = time.Now()
	proUpdate.Status = 1

	if err := proUpdate.UpdateSysProject(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ProjectController) AjaxDel() {

	project_id, _ := this.GetInt64("id")
	pro, _ := models.GetSysProjectById(project_id)
	pro.UpdateTime = time.Now()
	pro.UpdateId = this.userId
	pro.Status = 0
	pro.Id = project_id

	if err := pro.UpdateSysProject(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}
