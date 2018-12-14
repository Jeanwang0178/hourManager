/**
* @Project: hourManager
* @Package
* @Description: TODO
* @author : wj
* @date Date : 2018/12/08/ 11:09
* @version V1.0
 */
package controllers

import (
	"fmt"
	"github.com/patrickmn/go-cache"
	"hourManager/src/models"
	"hourManager/src/utils"
	"strconv"
	"strings"
	"time"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Index() {

	this.Data["pageTitle"] = "权限因子"
	this.display()
}

func (this *AuthController) List() {
	this.Data["zTree"] = true //引入ztreecss
	this.Data["pageTitle"] = "权限因子"
	this.display()
}

//获取全部节点
func (this *AuthController) GetNodes() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, count := models.GetSysAuthList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pId"] = v.Pid
		row["name"] = v.AuthName
		row["open"] = true
		list[k] = row
	}

	this.ajaxList("成功", MSG_OK, count, list)
}

//获取一个节点
func (this *AuthController) GetNode() {
	id, _ := this.GetInt64("id")
	result, _ := models.GetSysAuthById(id)
	// if err == nil {
	// 	this.ajaxMsg(err.Error(), MSG_ERR)
	// }
	row := make(map[string]interface{})
	row["id"] = result.Id
	row["pid"] = result.Pid
	row["auth_name"] = result.AuthName
	row["auth_url"] = result.AuthUrl
	row["sort"] = result.Sort
	row["is_show"] = result.IsShow
	row["icon"] = result.Icon

	fmt.Println(row)

	this.ajaxList("成功", MSG_OK, 0, row)
}

//新增或修改
func (this *AuthController) AjaxSave() {
	auth := new(models.SysAuth)
	auth.UserId = this.userId
	auth.Pid, _ = this.GetInt64("pid")
	auth.AuthName = strings.TrimSpace(this.GetString("auth_name"))
	auth.AuthUrl = strings.TrimSpace(this.GetString("auth_url"))
	auth.Sort, _ = this.GetInt("sort")
	auth.IsShow, _ = this.GetInt("is_show")
	auth.Icon = strings.TrimSpace(this.GetString("icon"))
	auth.UpdateTime = time.Now()

	auth.Status = 1

	id, _ := this.GetInt64("id")
	if id == 0 {
		//新增
		auth.CreateTime = time.Now()
		auth.CreateId = this.userId
		auth.UpdateId = this.userId
		if _, err := models.AddSysAuth(auth); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	} else {
		auth.Id = id
		auth.UpdateId = this.userId
		if err := auth.UpdateSysAuth(); err != nil {
			this.ajaxMsg(err.Error(), MSG_ERR)
		}
	}
	utils.Che.Set("menu"+strconv.FormatInt(this.user.Id, 10), nil, cache.DefaultExpiration)
	this.ajaxMsg("", MSG_OK)
}

//删除
func (this *AuthController) AjaxDel() {
	id, _ := this.GetInt64("id")
	auth, _ := models.GetSysAuthById(id)
	auth.Id = id
	auth.Status = 0
	if err := auth.UpdateSysAuth(); err != nil {
		this.ajaxMsg(err.Error(), MSG_ERR)
	}
	utils.Che.Set("menu"+strconv.FormatInt(this.user.Id, 10), nil, cache.DefaultExpiration)
	this.ajaxMsg("", MSG_OK)
}
