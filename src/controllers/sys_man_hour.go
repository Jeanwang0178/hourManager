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
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/beego/bee/logger"
	"hourManager/src/common"
	"hourManager/src/models"
	"hourManager/src/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	sheetName = "Sheet1"
)

type ManHourController struct {
	BaseController
}

func (this *ManHourController) List() {

	sourceList := sourceLists(this.userId)
	this.Data["sourceList"] = sourceList
	this.Data["pageTitle"] = "项目工时录入"
	this.Data["isFilter"] = "y"
	this.Data["realName"] = this.GetString("realName")
	this.display()
}

func (this *ManHourController) ListAll() {

	sourceList := sourceLists(this.userId)
	this.Data["sourceList"] = sourceList
	this.Data["pageTitle"] = "项目工时管理"
	this.Data["isFilter"] = "n"
	this.display("manhour/list")
}

func (this *ManHourController) ListUser() {
	this.Data["pageTitle"] = "工时信息"
	tplName := "project/list_user"
	this.Data["isExists"] = this.GetString("isExists")
	this.Data["projectId"] = this.GetString("projectId")
	this.display(tplName)
}

type sourceList struct {
	ProjectId   int
	ProjectName string
}

func (this *ManHourController) Add() {
	manHour := new(models.SysManHour)
	this.Data["manHour"] = manHour
	this.Data["workDateStr"] = ""
	//获取项目列表
	sourceList := sourceLists(this.userId)
	this.Data["sourceList"] = sourceList
	this.Data["isView"] = "n"
	this.Data["pageTitle"] = "新增项目工时"
	this.display("manhour/edit")

}

func (this *ManHourController) Edit() {
	this.Data["pageTitle"] = "编辑项目工时"
	sourceList := sourceLists(this.userId)
	this.Data["sourceList"] = sourceList

	id, _ := this.GetInt64("id")
	isView := this.GetString("isView")
	manHour, _ := models.GetSysManHourById(id)
	this.Data["manHour"] = manHour
	this.Data["isView"] = isView
	this.Data["workDateStr"] = beego.Date(manHour.WorkDate, "Y-m-d")
	this.display()

}

func sourceLists(userId int64) (sl []sourceList) {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	projectList := models.GetSysProjectListByUserId(userId)
	for _, sv := range projectList {
		sourceRow := sourceList{}
		sourceRow.ProjectId = int(sv.Id)
		sourceRow.ProjectName = sv.ProjectName
		sl = append(sl, sourceRow)
	}
	return sl
}

func (this *ManHourController) AjaxSave() {
	manHourId, _ := this.GetInt64("id")
	projectId, _ := this.GetInt64("project_id")
	userId := this.userId
	work_date := this.GetString("work_date")
	workDate, err := beego.DateParse(work_date, "Y-m-d")
	if err != nil {
		beeLogger.Log.Errorf("parse date failed :"+work_date, err)
	}
	taskTarget := this.GetString("task_target")
	taskProgress := this.GetString("task_progress")
	manHours, _ := this.GetFloat("man_hour")

	if manHourId == 0 {
		manHour := new(models.SysManHour)
		manHour.ProjectId = projectId
		manHour.UserId = userId
		manHour.WorkDate = workDate
		manHour.TaskTarget = taskTarget
		manHour.TaskProgress = taskProgress
		manHour.ManHour = manHours
		manHour.CreateTime = time.Now()
		manHour.UpdateTime = time.Now()
		manHour.Status, _ = strconv.Atoi(common.Enable)
		//新增
		manHour.CreateTime = time.Now()
		manHour.UpdateTime = time.Now()
		manHour.CreateId = this.userId
		manHour.UpdateId = this.userId

		_, err := models.GetSysManHourByWorkDate(-1)

		if err == nil {
			this.ajaxMsg("当前日期已经录入，请勿重复录入", MSG_ERR)
		}

		if _, err := models.AddSysManHour(manHour); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
		this.ajaxMsg("", MSG_OK)
	}

	update, _ := models.GetSysManHourById(manHourId)
	// 修改
	update.ProjectId = projectId
	update.WorkDate = workDate
	update.TaskTarget = taskTarget
	update.TaskProgress = taskProgress
	update.ManHour = manHours
	update.UpdateId = this.userId
	update.UpdateTime = time.Now()
	update.Status = 1

	if err := update.UpdateSysManHour(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ManHourController) AjaxDel() {

	man_hour_id, _ := this.GetInt64("id")
	manHour, _ := models.GetSysManHourById(man_hour_id)
	manHour.Status = 0
	manHour.Id = man_hour_id
	manHour.UserId = this.userId
	manHour.UpdateTime = time.Now()

	if err := manHour.UpdateSysManHour(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	this.ajaxMsg("", MSG_OK)
}

func (this *ManHourController) Table() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 300
	}

	projectId, _ := this.GetInt64("projectId")
	realName := strings.TrimSpace(this.GetString("realName"))
	isFilter := strings.TrimSpace(this.GetString("isFilter"))
	dateRange := strings.TrimSpace(this.GetString("dateRange"))
	this.pageSize = limit
	//查询条件
	filters := make(map[string]interface{})
	beginDate := ""
	endDate := ""
	if dateRange != "" {
		date := strings.Split(dateRange, " - ")
		beginDate = date[0]
		endDate = date[1]
	}

	beeLogger.Log.Infof("quey projectId :{}", projectId)
	filters["projectId"] = projectId
	filters["realName"] = realName
	filters["beginDate"] = beginDate
	filters["endDate"] = endDate
	if isFilter == "" || isFilter != "n" {
		filters["userId"] = this.userId
	} else {
		filters["userId"] = int64(0)
	}

	result, count := models.GetSysManHourInfoByParam(page, this.pageSize, filters)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["project_id"] = v.ProjectId
		row["project_name"] = v.ProjectName
		row["user_id"] = v.UserId
		row["real_name"] = v.RealName
		row["work_date"] = beego.Date(v.WorkDate, "Y-m-d")
		row["task_target"] = v.TaskTarget
		row["task_progress"] = v.TaskProgress
		row["man_hour"] = v.ManHour
		row["status"] = v.Status

		row["create_time"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["update_time"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		list[k] = row
	}
	this.ajaxList("查询成功", MSG_OK, count, list)
}

func (this *ManHourController) Excel() {
	//列表
	page, err := this.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = 300
	}

	projectId, _ := this.GetInt64("projectId")
	realName := strings.TrimSpace(this.GetString("realName"))
	isFilter := strings.TrimSpace(this.GetString("isFilter"))
	dateRange := strings.TrimSpace(this.GetString("dateRange"))
	this.pageSize = limit
	//查询条件
	filters := make(map[string]interface{})
	beginDate := ""
	endDate := ""
	if dateRange != "" {
		date := strings.Split(dateRange, " - ")
		beginDate = date[0]
		endDate = date[1]
	}

	beeLogger.Log.Infof("quey projectId :{}", projectId)
	filters["projectId"] = projectId
	filters["realName"] = realName
	filters["beginDate"] = beginDate
	filters["endDate"] = endDate
	if isFilter == "" || isFilter != "n" {
		filters["userId"] = this.userId
	} else {
		filters["userId"] = int64(0)
	}

	result, _ := models.GetSysManHourInfoByParam(page, this.pageSize, filters)
	details := make([]interface{}, 0)
	others := make(map[string]interface{})
	beeLogger.Log.Infof("export total count ", len(result))
	start := ""
	end := ""
	cellSize := 0
	for k, v := range result {
		row := make([]interface{}, 0)
		row = append(row, beego.Date(v.WorkDate, "Y-m-d"))
		row = append(row, v.TaskTarget)
		row = append(row, v.TaskProgress)
		row = append(row, v.ManHour)
		cellSize = 4

		if k == 0 {
			start = beego.Date(v.WorkDate, "Y-m-d")
			others["${workTarget}"] = "本周工作目标"
			others["${companyName}"] = v.CompanyName
			others["${realName}"] = v.RealName
		} else if k == len(result)-1 {
			end = beego.Date(v.WorkDate, "Y-m-d")
		}
		details = append(details, &row)
	}

	if dateRange == "" {
		dateRange = start + " - " + end
	}
	repstr := strings.Replace(dateRange, "-", "/", -1)
	others["${rangDate}"] = strings.Replace(repstr, " / ", " - ", -1)
	beeLogger.Log.Infof("replace :", others["${rangDate}"])

	openFile, err := excelize.OpenFile("static/excel/template_man_hour.xlsx")

	if err != nil {
		beeLogger.Log.Errorf("open excel template failed ", err)
	}

	fileName := utils.ExportExcel(openFile, sheetName, details,cellSize, others)

	defer func() {
		beeLogger.Log.Infof("delete excel %s", fileName)
		err := os.Remove(fileName)
		if err != nil {
			beeLogger.Log.Errorf("delete tmp excel failed %s ", fileName)
		}
	}()

	err = openFile.SaveAs(fileName)
	if err != nil {
		beeLogger.Log.Errorf("save excel failed %s ", fileName)
	}
	if others["${realName}"] == nil {
		others["${realName}"] = "XXX"
	}
	this.Ctx.Output.Download(fileName, "能源科技周工作计划单-XXX项目-"+others["${realName}"].(string)+"-"+time.Now().Format("20060102150405")+".xlsx")
}
