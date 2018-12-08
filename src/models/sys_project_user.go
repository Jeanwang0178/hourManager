package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type SysProjectUser struct {
	Id         int64 `orm:"column(id);auto" description:"主键ID"`
	ProjectId  int64 `orm:"column(project_id)" description:"项目ID"`
	UserId     int64 `orm:"column(user_id)" description:"用户ID"`
	CreateId   int64 `orm:"column(create_id)" description:"创建者ID"`
	UpdateId   int64 `orm:"column(update_id)" description:"修改者ID"`
	Status     uint8 `orm:"column(status)" description:"状态1-正常，0-删除"`
	CreateTime int64 `orm:"column(create_time)" description:"添加时间"`
	UpdateTime int64 `orm:"column(update_time)" description:"修改时间"`
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
