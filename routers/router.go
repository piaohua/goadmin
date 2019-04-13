package routers

import (
	"goadmin/api"
	"goadmin/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

// SetupRouter setup router
func SetupRouter() {
	//namespace
	namespace := beego.AppConfig.String("namespace")
	if len(namespace) != 0 {

		domain := beego.AppConfig.String("domain")
		ns :=
			beego.NewNamespace(namespace,
				beego.NSCond(func(ctx *context.Context) bool {
					beego.Info("domain: ", ctx.Input.Domain())
					if ctx.Input.Domain() == domain {
						return true
					}
					return false
				}),
				beego.NSRouter("/", &controllers.MainController{}),
				beego.NSRouter("/login", &controllers.MainController{}, "post:Login"),
				beego.NSAutoRouter(&controllers.AdminController{}),
				beego.NSAutoRouter(&controllers.SystemController{}),
				beego.NSAutoRouter(&controllers.SysuserController{}),
				beego.NSAutoRouter(&controllers.ModuleController{}),
				beego.NSAutoRouter(&controllers.RoleController{}),
				//api
				beego.NSRouter("/api/stat", &api.StatisController{}, "post:Stat"),
				//stat
				beego.NSAutoRouter(&controllers.StatController{}),
				beego.NSAutoRouter(&controllers.ReportController{}),
				//IdleCrafting
				beego.NSAutoRouter(&controllers.IdleCraftingController{}),
				//Smash
				beego.NSAutoRouter(&controllers.SmashController{}),
				//Banzi
				beego.NSAutoRouter(&controllers.BanziController{}),
				//Qise
				beego.NSAutoRouter(&controllers.QiseController{}),
				//Gunner
				beego.NSAutoRouter(&controllers.GunnerController{}),
				//Triangle
				beego.NSAutoRouter(&controllers.TriangleController{}),
				//Project
				beego.NSAutoRouter(&controllers.ProjectController{}),
				//ApiDoc
				beego.NSAutoRouter(&controllers.ApiDocController{}),
				//Server
				beego.NSAutoRouter(&controllers.ServerController{}),
				//Log
				beego.NSAutoRouter(&controllers.LogController{}),
				//Electro
				beego.NSAutoRouter(&controllers.ElectroController{}),
				//Fudai
				beego.NSAutoRouter(&controllers.FuDaiController{}),
				//Card2048
				beego.NSAutoRouter(&controllers.Card2048Controller{}),
				//Pintu
				beego.NSAutoRouter(&controllers.PintuController{}),
			)
		beego.AddNamespace(ns)

	} else {

		beego.Router("/", &controllers.MainController{})
		beego.Router("/login", &controllers.MainController{}, "post:Login")
		beego.AutoRouter(&controllers.AdminController{})
		beego.AutoRouter(&controllers.SystemController{})
		beego.AutoRouter(&controllers.SysuserController{})
		beego.AutoRouter(&controllers.ModuleController{})
		beego.AutoRouter(&controllers.RoleController{})
		//api
		beego.Router("/api/stat", &api.StatisController{}, "post:Stat")
		//stat
		beego.AutoRouter(&controllers.StatController{})
		beego.AutoRouter(&controllers.ReportController{})
		//IdleCrafting
		beego.AutoRouter(&controllers.IdleCraftingController{})
		//Smash
		beego.AutoRouter(&controllers.SmashController{})
		//Banzi
		beego.AutoRouter(&controllers.BanziController{})
		//Qise
		beego.AutoRouter(&controllers.QiseController{})
		//Gunner
		beego.AutoRouter(&controllers.GunnerController{})
		//Triangle
		beego.AutoRouter(&controllers.TriangleController{})
		//Project
		beego.AutoRouter(&controllers.ProjectController{})
		//ApiDoc
		beego.AutoRouter(&controllers.ApiDocController{})
		//Server
		beego.AutoRouter(&controllers.ServerController{})
		//Log
		beego.AutoRouter(&controllers.LogController{})
		//Electro
		beego.AutoRouter(&controllers.ElectroController{})
		//FuDai
		beego.AutoRouter(&controllers.FuDaiController{})
		//Card2048
		beego.AutoRouter(&controllers.Card2048Controller{})
		//Pintu
		beego.AutoRouter(&controllers.PintuController{})

	}

	beego.SetStaticPath("/"+namespace+"/static", "static")
}
