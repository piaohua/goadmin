package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	if c.GetSession("user_id") != nil {
		c.redirect(beego.URLFor("AdminController.Main"))
	}
	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		// 显示设置成功
		c.Data["ErrorMsg"] = n
	} else if n, ok = flash.Data["error"]; ok {
		// 显示错误
		c.Data["ErrorMsg"] = n
	} else {
		// 不然默认显示设置页面
		c.Data["ErrorMsg"] = ""
	}
	c.Data["SystemName"] = beego.AppConfig.String("systenName")
	c.Data["SystemAlias"] = beego.AppConfig.String("systenAlias")
	c.display("login.html")
}

func (c *MainController) Login() {
	flash := beego.NewFlash()
	email := c.GetString("email", "")
	passwd := c.GetString("passwd", "")
	if len(email) == 0 || len(passwd) == 0 {
		flash.Error("请输入用户名和密码")
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor("MainController.Get"))
		return
	}
	user := new(models.SystemUser)
	if ok, err := user.CheckLogin(email, passwd); !ok {
		flash.Error("登录失败: " + err.Error())
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor("MainController.Get"))
		return
	}
	if user.Status != 1 {
		flash.Error("账户被禁用")
		flash.Store(&c.Controller)
		c.redirect(beego.URLFor("MainController.Get"))
		return
	}
	c.SetSession("username", string(user.UserName))
	c.SetSession("email", string(user.Email))
	c.SetSession("user_id", int32(user.UserId))
	c.SetSession("role_id", int32(user.RoleId))
	c.redirect(beego.URLFor("AdminController.Main"))
}
