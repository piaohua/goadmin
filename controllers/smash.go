package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// SmashController ...
type SmashController struct {
	baseAdminController
}

// Index ...
func (c *SmashController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gold := c.GetString("gold", "")
	diamond, _ := c.GetInt("diamond", 0)
	beego.Info("smash index startTime: ", startTime)
	beego.Info("smash index endTime: ", endTime)
	beego.Info("smash index userid: ", userid)
	beego.Info("smash index gold: ", gold)
	beego.Info("smash index diamond: ", diamond)
	b := new(models.SmBaseData)
	lists, pagination, _ := b.GetDataList(userid, startTime, endTime,
		gold, diamond, page, c.pageSize, c.isPost())
	//beego.Info("smash lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("SmashController.Index")
	beego.Info("smash pagination: ", pagination)
	c.Data["Gold"] = gold
	if diamond != 0 {
		c.Data["Diamond"] = diamond
	}
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	c.Data["SmData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("smash/index.html")
}

// BaseEdit ...
func (c *SmashController) BaseEdit() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	gold := c.GetString("gold", "")
	diamond, err := c.GetUint32("diamond", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "钻石参数错误", nil)
	}
	step, err := c.GetInt("guide_step", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	beego.Info("smash baseEdit step: ", step)
	beego.Info("smash baseEdit gold: ", gold)
	beego.Info("smash baseEdit diamond: ", diamond)
	beego.Info("smash baseEdit userid: ", userid)
	b := new(models.SmBaseData)
	b.ID = userid
	b.Gold = gold
	b.Diamond = diamond
	b.GuideStep = step
	err = b.UpdateBase()
	if err != nil {
		beego.Error("smash baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
