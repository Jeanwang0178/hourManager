package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/bee/generate/swaggergen"
	beeUtils "github.com/beego/bee/utils"
	"github.com/patrickmn/go-cache"
	_ "hourManager/routers"
	"hourManager/src/inital"
	"hourManager/src/utils"
	"html/template"
	"net/http"
	"os"
	"time"
)

const VERSION = "1.0.1"

var (
	workspace = os.Getenv("BeeWorkspace")
)

func main() {

	inital.Init()

	utils.Che = cache.New(60*time.Minute, 120*time.Minute)
	//设置默认404页面
	beego.ErrorHandler("404", func(writer http.ResponseWriter, request *http.Request) {
		t, _ := template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath + "/error/404.html")
		data := make(map[string]interface{})
		data["content"] = "page not found"
		t.Execute(writer, data)
	})

	beego.BConfig.WebConfig.Session.SessionOn = true

	//是否异常恢复，默认值为 true
	beego.BConfig.RecoverPanic = true
	beego.BConfig.WebConfig.EnableDocs = true

	//初始化swagger
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.SetStaticPath("/swagger", "swagger")
	}
	beego.SetStaticPath("/module", "module")
	beego.BConfig.Log.AccessLogs = true

	currentpath, _ := os.Getwd()
	if workspace != "" {
		currentpath = workspace
	}
	if beeUtils.IsInGOPATH(currentpath) {
		swaggergen.GenerateDocs(currentpath)
	}else{
		swaggergen.GenerateDocs("")
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	beego.Run()

}
