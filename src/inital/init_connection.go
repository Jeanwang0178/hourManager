package inital

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	blog "github.com/beego/bee/logger"
	_ "github.com/go-sql-driver/mysql"
	"hourManager/src/common"
	"net/url"
)

func Init() {

	//默认数据库
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		panic("conect default database failed " + err.Error())
	}

	if beego.AppConfig.String("runmode") == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		orm.Debug = true
	}

	//redis 初始化
	common.InitCache()

	data := make([]interface{}, 0)
	data = append(data, "default")
	err = common.SetCache(common.AliasName, data, 6000000)
	if err != nil {
		blog.Log.Errorf("conect redis failed ", err.Error())
	}

}
