/**
* @Project: hourManager
* @Package models
* @Description: 角色权限
* @author : wj
* @date Date : 2018/12/06/ 16:00
* @version V1.0
 */
package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"bytes"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type SysRoleAuth struct {
	Id     int64 `orm:"column(id);auto" description:"主键ID"`
	RoleId int64 `orm:"column(role_id)" description:"角色ID"`
	AuthId int64 `orm:"column(auth_id)" description:"权限ID"`
}

func (t *SysRoleAuth) TableName() string {
	return "sys_role_auth"
}

func init() {
	orm.RegisterModel(new(SysRoleAuth))
}

// AddSysRoleAuth insert a new SysRoleAuth into database and returns
// last inserted Id on success.
func AddSysRoleAuth(m *SysRoleAuth) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysRoleAuthById retrieves SysRoleAuth by Id. Returns error if
// Id doesn't exist
func GetSysRoleAuthById(id int64) (v *SysRoleAuth, err error) {
	o := orm.NewOrm()
	v = &SysRoleAuth{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetSysRoleAuthByRoleId(roleId int64) ([]*SysRoleAuth, error) {
	list := make([]*SysRoleAuth, 0)
	query := orm.NewOrm().QueryTable("sys_role_auth")
	_, err := query.Filter("role_id", roleId).All(&list, "AuthId")
	if err != nil {
		return nil, err
	}
	return list, nil
}

//获取多个
func GetSysRoleAuthByIds(RoleIds string) (Authids string, err error) {
	list := make([]*SysRoleAuth, 0)
	query := orm.NewOrm().QueryTable("sys_role_auth")
	ids := strings.Split(RoleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.FormatInt(v.AuthId, 10))
			b.WriteString(",")
		}
	}
	Authids = strings.TrimRight(b.String(), ",")
	return Authids, nil
}

// GetAllSysRoleAuth retrieves all SysRoleAuth matches certain condition. Returns empty list if
// no records exist
func GetAllSysRoleAuth(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysRoleAuth))
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

	var l []SysRoleAuth
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

// UpdateSysRoleAuth updates SysRoleAuth by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysRoleAuthById(m *SysRoleAuth) (err error) {
	o := orm.NewOrm()
	v := SysRoleAuth{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysRoleAuth deletes SysRoleAuth by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysRoleAuth(id int64) (err error) {
	o := orm.NewOrm()
	v := SysRoleAuth{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysRoleAuth{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
