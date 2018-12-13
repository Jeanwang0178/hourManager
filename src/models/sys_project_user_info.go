/**
* @Project: hourManager
* @Package
* @Description: TODO
* @author : wj
* @date Date : 2018/12/08/ 17:20
* @version V1.0
 */
package models

type SysProjectUserInfo struct {
	SysProjectUser
	ProjectName string //`orm:"column(project_name);size(20)" description:"项目名称"`
	CompanyId   string //`orm:"column(company_id);size(20)" description:"公司ID"`
	CompanyName string //`orm:"column(company_name);size(20)" description:"公司名称"`
	LoginName   string //`orm:"column(login_name);size(20)" description:"用户名"`
	RealName    string //`orm:"column(real_name);size(32)" description:"真实姓名"`
	Password    string //`orm:"column(password);size(32)" description:"密码"`
	RoleIds     string //`orm:"column(role_ids);size(255)" description:"角色id字符串，如：2,3,4"`
	Phone       string //`orm:"column(phone);size(20)" description:"手机号码"`
	Email       string //`orm:"column(email);size(50)" description:"邮箱"`
	Salt        string //`orm:"column(salt);size(10)" description:"密码盐"`
	LastLogin   int64  //`orm:"column(last_login)" description:"最后登录时间"`
	LastIp      string //`orm:"column(last_ip);size(15)" description:"最后登录IP"`
	Status      int    //`orm:"column(status)" description:"状态，1-正常 0禁用"`
	CreateUser  string //`orm:"column(login_name);size(20)" description:"用户名"`
}
