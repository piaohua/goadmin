package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// TriangleController ...
type TriangleController struct {
	baseAdminController
}

// Index ...
func (c *TriangleController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gold, _ := c.GetInt("gold", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	projectID, _ := c.GetInt("project_id", 0)
	if !c.isPost() {
		sourceID = -1 //默认显示全部
	}
	beego.Info("triangle index startTime: ", startTime)
	beego.Info("triangle index endTime: ", endTime)
	beego.Info("triangle index userid: ", userid)
	beego.Info("triangle index gold: ", gold)
	b := new(models.TgBaseData)
	lists, pagination, _ := b.GetDataList(userid, startTime, endTime,
		gold, page, c.pageSize, c.isPost())
	//beego.Info("triangle lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["BaseData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	sourceList := models.GetSourceList(projectID)
	pagination.BaseUrl = beego.URLFor("TriangleController.Index")
	beego.Info("triangle pagination: ", pagination)
	if gold != 0 {
		c.Data["Gold"] = gold
	}
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Sources"] = &sourceList
	c.Data["SourceID"] = sourceID
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	c.Data["BaseData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("triangle/index.html")
}

// BaseEdit ...
func (c *TriangleController) BaseEdit() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	gold, err := c.GetUint32("gold", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "金币参数错误", nil)
	}
	gate, err := c.GetUint32("gate", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	guideStep, err := c.GetUint32("guide_step", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	add, err := c.GetBool("add", false)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	unlock_gate, err := c.GetUint32("unlock_gate", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	beego.Info("triangle baseEdit gold: ", gold)
	beego.Info("triangle baseEdit userid: ", userid)
	b := new(models.TgBaseData)
	b.ID = userid
	b.Gold = gold
	b.Gate = gate
	b.GuideStep = guideStep
	b.UnlockGate = unlock_gate
	b.Add = add
	err = b.UpdateBase()
	if err != nil {
		beego.Error("triangle baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
