/**
* @Project: hourManager
* @Package models
* @Description: 用户管理
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

	"github.com/astaxie/beego/orm"
	"time"
)

type SysUser struct {
	Id          int64     `orm:"column(id);auto"`
	CompanyName string    `orm:"column(company_name);size(64)" description:"公司名称"`
	LoginName   string    `orm:"column(login_name);size(20)" description:"用户名"`
	RealName    string    `orm:"column(real_name);size(32)" description:"真实姓名"`
	Password    string    `orm:"column(password);size(32)" description:"密码"`
	RoleIds     string    `orm:"column(role_ids);size(255)" description:"角色id字符串，如：2,3,4"`
	Phone       string    `orm:"column(phone);size(20)" description:"手机号码"`
	Email       string    `orm:"column(email);size(50)" description:"邮箱"`
	Salt        string    `orm:"column(salt);size(10)" description:"密码盐"`
	LastLogin   int64     `orm:"column(last_login)" description:"最后登录时间"`
	LastIp      string    `orm:"column(last_ip);size(15)" description:"最后登录IP"`
	Status      int       `orm:"column(status)" description:"状态，1-正常 0禁用"`
	CreateId    int64     `orm:"column(create_id)" description:"创建者ID"`
	UpdateId    int64     `orm:"column(update_id)" description:"修改者ID"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null" description:"更新时间"`
}

func (t *SysUser) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(SysUser))
}

// AddSysUser insert a new SysUser into database and returns
// last inserted Id on success.
func AddSysUser(m *SysUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSysUserById retrieves SysUser by Id. Returns error if
// Id doesn't exist
func GetSysUserById(id int64) (v *SysUser, err error) {
	o := orm.NewOrm()
	v = &SysUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetSysUserByName retrieves User by loginName. Returns error if
// no records exist
func GetSysUserByName(loginName string) (*SysUser, error) {
	u := new(SysUser)
	err := orm.NewOrm().QueryTable("sys_user").Filter("login_name", loginName).Filter("status", "1").One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func GetSysUserList(page, pageSize int, filters ...interface{}) ([]*SysUser, int64) {
	offset := (page - 1) * pageSize
	list := make([]*SysUser, 0)
	query := orm.NewOrm().QueryTable("sys_user")
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

func (u *SysUser) UpdateSysUser(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

// GetAllSysUser retrieves all SysUser matches certain condition. Returns empty list if
// no records exist
func GetAllSysUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SysUser))
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

	var l []SysUser
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

// UpdateSysUser updates SysUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateSysUserById(m *SysUser) (err error) {
	o := orm.NewOrm()
	v := SysUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteSysUser deletes SysUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteSysUser(id int64) (err error) {
	o := orm.NewOrm()
	v := SysUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&SysUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
