package main

import (
	"encoding/gob"

	"goadmin/models"
	"goadmin/routers"

	"github.com/astaxie/beego"
)

func init() {
	//session 值类型注册
	gob.Register([]models.ParentMenu{})
	gob.Register([]models.Project{})
	gob.Register(map[string]bool{})
}

func main() {

	//logs
	if beego.AppConfig.String("runmode") == "dev" {
		beego.SetLevel(beego.LevelDebug)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log_file")+`"}`)
	} else {
		beego.SetLevel(beego.LevelDebug)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log_file")+`"}`)
		beego.BeeLogger.DelLogger("console")
	}

	//orm
	models.InitOrm()

	//router
	routers.SetupRouter()

	beego.Run()
}
