/**
* @Project: hourManager
* @Package
* @Description: TODO
* @author : wj
* @date Date : 2018/12/08/ 17:20
* @version V1.0
 */
package models

type SysManHourInfo struct {
	SysManHour
	ProjectName string //`orm:"column(project_name);size(20)" description:"项目名称"`
	CompanyName string //`orm:"column(company_name);size(20)" description:"公司名称"`
	RealName    string //`orm:"column(real_name);size(32)" description:"真实姓名"`

}
