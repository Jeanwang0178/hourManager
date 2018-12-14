package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"bytes"
	"github.com/astaxie/beego/orm"
	"github.com/beego/bee/logger"
	"strconv"
	"time"
)

type SysProjectUser struct {
	Id         int64     `orm:"column(id);auto" description:"主键ID"`
	ProjectId  int64     `orm:"column(project_id)" description:"项目ID"`
	UserId     int64     `orm:"column(user_id)" description:"用户ID"`
	CreateId   int64     `orm:"column(create_id)" description:"创建者ID"`
	UpdateId   int64     `orm:"column(update_id)" description:"修改者ID"`
	Status     int8      `orm:"column(status)" description:"状态1-正常，0-删除"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
}

func (t *SysProjectUser) TableName() string {
	return "sys_project_user"
}

func init() {
	orm.RegisterModel(new(SysProjectUser))
}

// AddSysProjectUser insert a new SysProjectUser into database and returns
// last inserted Id on success.
func AddSysProjectUser(m *SysProjectUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func AddSysProjectUserList(mlist []SysProjectUser) (count int64, err error) {
	o := orm.NewOrm()
	count, err = o.InsertMulti(len(mlist), mlist)
	if err != nil {
		return 0, nil
	}
	return count, nil
}

// GetSysProjectUserById retrieves SysProjectUser by Id. Returns error if
// Id doesn't exist
func GetSysProjectUserById(id int64) (v *SysProjectUser, err error) {
	o := orm.NewOrm()
	v = &SysProjectUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (p *SysProjectUser) UpdateSysProjectUser(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func GetSysProjectUserList(page, pageSize int, filters ...interface{}) ([]*SysProjectUser, int64) {
	offset := (page - 1) * pageSize
	list := make([]*SysProjectUser, 0)
	query := orm.NewOrm().QueryTable("sys_project_user")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}

func GetSysProjectListByUserId(userId int64) []SysProject {
	var list []SysProject
	sql := new(bytes.Buffer)
	sql.WriteString("select * from sys_project a where exists (select 1 from sys_project_user b where a.id = b.project_id and b.status = 1  ")
	if userId != 0 {
		sql.WriteString(" and b.user_id = ?  ")
	}
	sql.WriteString(") and a.status = 1 order by a.project_name ")
	beeLogger.Log.Infof(sql.String(), userId)

	if userId != 0 {
		orm.NewOrm().Raw(sql.String(), userId).QueryRows(&list)
	} else {
		orm.NewOrm().Raw(sql.String()).QueryRows(&list)
	}
	return list
}

func GetSysUserListByParam(page, pageSize int, filters map[string]interface{}) ([]SysUser, int64) {
	offset := (page - 1) * pageSize
	res := make(orm.Params)
	var list []SysUser
	totalSql := ""
	querySql := ""
	isExists := filters["isExists"].(string)
	projectId := filters["projectId"].(int64)
	if isExists == "y" {
		totalSql = "select 'total' total, count(*) count from sys_user a where exists (select 1 from sys_project_user b where a.id = b.user_id and b.status = 1 and b.project_id = ? ) and a.status = 1"
		querySql = "select * from sys_user a where exists (select 1 from sys_project_user b where a.id = b.user_id and b.status = 1 and b.project_id = ? ) and a.status = 1  limit ? offset ? "
	} else {
		totalSql = "select 'total' total, count(*) count from sys_user a where not exists (select 1 from sys_project_user b where a.id = b.user_id and b.status = 1 and b.project_id = ?) and a.status = 1"
		querySql = "select * from sys_user a where not exists (select 1 from sys_project_user b where a.id = b.user_id and b.status = 1 and b.project_id = ? ) and a.status = 1 limit ? offset ? "
	}
	total, err := orm.NewOrm().Raw(totalSql, projectId).RowsToMap(&res, "total", "count")
	if err != nil {
		beeLogger.Log.Errorf("GetSysUserListByParam total falied ", err)
	} else {
		total, err = strconv.ParseInt(res["total"].(string), 10, 64)
		if err != nil {
			beeLogger.Log.Errorf("strconv ParseInt failed :", res["total"])
		}
	}
	orm.NewOrm().Raw(querySql, projectId, pageSize, offset).QueryRows(&list)

	return list, total
}

func GetSysProjectUserInfoByParam(page, pageSize int, filters ...interface{}) ([]SysProjectUserInfo, int64) {
	offset := (page - 1) * pageSize
	res := make(orm.Params)
	projectId := filters
	var list []SysProjectUserInfo
	totalSql := "select 'total' total, count(*) count from sys_user a where exists (select 1 from sys_project_user b where a.id = b.user_id and b.status = 1 and b.project_id = ? ) and a.status = 1"
	querySql := "select b.id ,b.project_id ,c.project_name,b.user_id,a.company_name, a.`login_name`, a.`real_name`, a.`role_ids`, a.`phone`, a.`email`, a.`last_login`, a.`last_ip`, a.`status`, b.`create_id`, d.real_name create_user, b.`update_id`, b.`create_time`, b.`update_time` from sys_user a ,sys_project_user b ,sys_project c ,sys_user d where a.id = b.user_id and b.project_id = c.id and b.create_id = d.id and b.project_id = ? and b.status  = 1 limit ? offset ? "

	total, err := orm.NewOrm().Raw(totalSql, projectId).RowsToMap(&res, "total", "count")
	if err != nil {
		beeLogger.Log.Errorf("GetSysUserListByParam total falied ", err)
	} else {
		total, err = strconv.ParseInt(res["total"].(string), 10, 64)
		if err != nil {
			beeLogger.Log.Errorf("strconv ParseInt failed :", res["total"])
		}
	}
	orm.NewOrm().Raw(querySql, projectId, pageSize, offset).QueryRows(&list)

	return list, total
}

// GetAllSysProjectUser retrieves all SysProjectUser matches certain condition. Returns empty list if
// no records exist
func GetAllSysProjectUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysProjectUser))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []SysProjectUser
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateSysProjectUser updates SysProjectUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysProjectUserById(m *SysProjectUser) (err error) {
	o := orm.NewOrm()
	v := SysProjectUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysProjectUser deletes SysProjectUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysProjectUser(id int64) (err error) {
	o := orm.NewOrm()
	v := SysProjectUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysProjectUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
