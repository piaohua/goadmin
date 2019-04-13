package controllers

import (
	"sort"
	"time"

	"goadmin/libs"
	"goadmin/models"

	"github.com/astaxie/beego"
)

// AdminController ...
type AdminController struct {
	baseAdminController
}

// Main 主界面
func (c *AdminController) Main() {
	c.display("main.html")
}

// Control 控制台
func (c *AdminController) Control() {
	lists := models.GetChatList()
	sort.Slice(lists, func(i, j int) bool {
		return lists[i].ID < lists[j].ID
	})
	c.Data["MsgData"] = &lists
	c.display("control.html")
}

// Logout ...
func (c *AdminController) Logout() {
	c.DestroySession()
	c.redirect(beego.URLFor("MainController.Get"))
}

// Send ...
func (c *AdminController) Send() {
	message := c.GetString("message", "")
	beego.Info("message: ", message)
	if message == "" {
		c.RenderJson(300, "输入不能为空", nil)
	}
	val := models.SystemChat{
		UserId:    c.userId,
		UserName:  c.getUserName(),
		Message:   message,
		CreatedAt: time.Now(),
	}
	switch message {
	case "pm2 list",
		"~/ctrl list",
		"~/ctrl flush",
		"~/ctrl nginx",
		"~/ctrl nginx-restart",
		"pm2 show banzi",
		"pm2 show banzi2",
		"pm2 show goadmin",
		"pm2 show idlecrafting",
		"pm2 show idlecrafting2",
		"pm2 show qise",
		"pm2 show qise2",
		"pm2 show report",
		"pm2 show report2",
		"pm2 show smash2":
		b, err := libs.ExecCmd(message)
		if err != nil {
			val.Content = err.Error()
		} else {
			val.Content = string(b)
		}
	}
	ok, err := val.Save()
	beego.Info("val: ", val, ", ok: ", ok, ", err: ", err)
	c.RenderJson(200, "处理成功", val)
}

// Sync ...
func (c *AdminController) Sync() {
	id, _ := c.GetInt("lastid", 0)
	//beego.Info("sync id: ", id, ", err: ", err)
	lists := models.GetChatLast(id)
	sort.Slice(lists, func(i, j int) bool {
		return lists[i].ID < lists[j].ID
	})
	c.RenderJson(200, "处理成功", lists)
}

// SetProject ...
func (c *AdminController) SetProject() {
	id, _ := c.GetInt("project_id", 0)
	beego.Info("SetProject id: ", id)
	if id != c.projectId {
		c.setProjectId(id)
	}
	c.RenderJson(200, "处理成功", nil)
}
