package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ComUser struct {
	Id        string `orm:"column(id);pk"`
	UserName  string `orm:"column(user_name);size(20)" description:"用户名"`
	Email     string `orm:"column(email);size(50)" description:"邮箱"`
	Password  string `orm:"column(password);size(32)" description:"密码"`
	Salt      string `orm:"column(salt);size(10)" description:"密码盐"`
	LastLogin int    `orm:"column(last_login)" description:"最后登录时间"`
	LastIp    string `orm:"column(last_ip);size(15)" description:"最后登录IP"`
	Status    int8   `orm:"column(status)" description:"状态，0正常 -1禁用"`
}

func (t *ComUser) TableName() string {
	return "com_user"
}

func init() {
	orm.RegisterModel(new(ComUser))
}

func (u *ComUser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

// AddComUser insert a new ComUser into database and returns
// last inserted Id on success.
func AddComUser(m *ComUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetComUserById retrieves ComUser by Id. Returns error if
// Id doesn't exist
func GetComUserById(id string) (v *ComUser, err error) {
	o := orm.NewOrm()
	v = &ComUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllComUser retrieves all ComUser matches certain condition. Returns empty list if
// no records exist
func GetAllComUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ComUser))
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

	var l []ComUser
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

// UpdateComUser updates ComUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateComUserById(m *ComUser) (err error) {
	o := orm.NewOrm()
	v := ComUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// UserGetByName retrieves User by userName. Returns error if
// no records exist
func GetComUserByName(userName string) (*ComUser, error) {
	u := new(ComUser)
	err := orm.NewOrm().QueryTable("com_user").Filter("user_name", userName).One(u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteComUser deletes ComUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteComUser(id string) (err error) {
	o := orm.NewOrm()
	v := ComUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ComUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
