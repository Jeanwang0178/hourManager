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

type SysManHour struct {
	Id           int64   `orm:"column(id);auto"`
	ProjectId    int64   `orm:"column(project_id)" description:"项目ID"`
	UserId       int64   `orm:"column(user_id)" description:"用户ID"`
	WorkDate     int64   `orm:"column(work_date)" description:"日期"`
	TaskTarget   string  `orm:"column(task_target);size(1024)" description:"当日工作目标"`
	TaskProgress string  `orm:"column(task_progress);size(20)" description:"任务进展情况"`
	ManHour      float64 `orm:"column(man_hour);null;digits(15);decimals(5)" description:"本日用时"`
	Status       int     `orm:"column(status)" description:"状态，1-正常 0禁用"`
	CreateId     int64   `orm:"column(create_id)" description:"创建者ID"`
	UpdateId     int64   `orm:"column(update_id)" description:"修改者ID"`
	CreateTime   int64   `orm:"column(create_time)" description:"创建时间"`
	UpdateTime   int64   `orm:"column(update_time)" description:"修改时间"`
}

func (t *SysManHour) TableName() string {
	return "sys_man_hour"
}

func init() {
	orm.RegisterModel(new(SysManHour))
}

// AddSysManHour insert a new SysManHour into database and returns
// last inserted Id on success.
func AddSysManHour(m *SysManHour) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysManHourById retrieves SysManHour by Id. Returns error if
// Id doesn't exist
func GetSysManHourById(id int64) (v *SysManHour, err error) {
	o := orm.NewOrm()
	v = &SysManHour{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//根据工作日期查找
func GetSysManHourByWorkDate(workDate int64) (v *SysManHour, err error) {
	o := orm.NewOrm()
	v = &SysManHour{Id: workDate}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func (p *SysManHour) UpdateSysManHour(fields ...string) error {
	if _, err := orm.NewOrm().Update(p, fields...); err != nil {
		return err
	}
	return nil
}

func GetSysManHourList(page, pageSize int, filters ...interface{}) ([]*SysManHour, int64) {
	offset := (page - 1) * pageSize
	list := make([]*SysManHour, 0)
	query := orm.NewOrm().QueryTable("sys_man_hour")
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

func GetSysManHourInfoByParam(page, pageSize int, filters map[string]interface{}) ([]SysManHourInfo, int64) {
	offset := (page - 1) * pageSize
	res := make(orm.Params)
	var list []SysManHourInfo
	totalSql := new(bytes.Buffer)
	querySql := new(bytes.Buffer)

	totalSql.WriteString("select 'total' total, count(*) count from sys_man_hour a where exists (select 1 from sys_project b where b.id = a.project_id and (b.id = ?  or 0 = ? )  ) and (a.user_id = ? or 0 = ?)  and a.status  = 1 ")
	totalSql.WriteString(" and  exists (select 1 from sys_user c where c.id = a.user_id and c.real_name like ?  )  ")
	querySql.WriteString("select b.project_name,c.real_name,a.`id`, a.`project_id`, a.`user_id`,c.`company_name`, a.`work_date`, a.`task_target`, a.`task_progress`, a.`man_hour`, a.`status`, a.`create_id`, a.`update_id`, a.`create_time`, a.`update_time` from sys_man_hour a ,sys_project b ,sys_user c where a.project_id = b.id and a.user_id = c.id  and (  b.id = ? or 0 = ? ) and (a.user_id = ? or 0 = ?) and a.status = 1 and c.real_name like ? ")

	beginDate := filters["beginDate"].(string)
	endDate := filters["endDate"].(string)
	projectId := filters["projectId"].(int64)
	userId := filters["userId"].(int64)
	realName := filters["realName"].(string)
	realName = "%" + realName + "%"
	params := make([]interface{}, 0)
	params = append(params, projectId)
	params = append(params, projectId)
	params = append(params, userId)
	params = append(params, userId)
	params = append(params, realName)
	if beginDate != "" && endDate != "" {
		totalSql.WriteString(" and a.work_date between ? and ? ")
		querySql.WriteString(" and a.work_date between ? and ? ")
		start, _ := time.Parse("2006-01-02 15:04:05", beginDate+" 00:00:00")
		end, _ := time.Parse("2006-01-02 15:04:05", endDate+" 00:00:00")
		params = append(params, start.Unix())
		params = append(params, end.Unix())
	}

	total, err := orm.NewOrm().Raw(totalSql.String(), params).RowsToMap(&res, "total", "count")

	if err != nil {
		beeLogger.Log.Errorf("GetSysManHourInfoByParam total falied ", err)
	} else {
		total, err = strconv.ParseInt(res["total"].(string), 10, 64)
		if err != nil {
			beeLogger.Log.Errorf("strconv ParseInt failed :", res["total"])
		}
	}
	querySql.WriteString(" order by a.work_date,a.update_time  ")
	querySql.WriteString("  limit ? offset ?")
	params = append(params, pageSize)
	params = append(params, offset)

	orm.NewOrm().Raw(querySql.String(), params).QueryRows(&list)

	return list, total
}

// GetAllSysManHour retrieves all SysManHour matches certain condition. Returns empty list if
// no records exist
func GetAllSysManHour(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysManHour))
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

	var l []SysManHour
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

// UpdateSysManHour updates SysManHour by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysManHourById(m *SysManHour) (err error) {
	o := orm.NewOrm()
	v := SysManHour{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysManHour deletes SysManHour by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysManHour(id int64) (err error) {
	o := orm.NewOrm()
	v := SysManHour{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysManHour{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
