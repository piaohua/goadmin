package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// GunnerController ...
type GunnerController struct {
	baseAdminController
}

// Index ...
func (c *GunnerController) Index() {
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
	beego.Info("gunner index startTime: ", startTime)
	beego.Info("gunner index endTime: ", endTime)
	beego.Info("gunner index userid: ", userid)
	beego.Info("gunner index gold: ", gold)
	b := new(models.GnBaseData)
	lists, pagination, _ := b.GetDataList(userid, startTime, endTime,
		gold, page, c.pageSize, c.isPost())
	//beego.Info("gunner lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["BaseData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	sourceList := models.GetSourceList(projectID)
	pagination.BaseUrl = beego.URLFor("GunnerController.Index")
	beego.Info("gunner pagination: ", pagination)
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
	c.display("gunner/index.html")
}

// BaseEdit ...
func (c *GunnerController) BaseEdit() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	gold, err := c.GetUint32("gold", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "金币参数错误", nil)
	}
	TopScore, err := c.GetUint32("top_score", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	AtkProp, err := c.GetUint32("atk_prop", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	BoomProp, err := c.GetUint32("boom_prop", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	OutTime, err := c.GetInt64("out_time", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	beego.Info("gunner baseEdit gold: ", gold)
	beego.Info("gunner baseEdit userid: ", userid)
	b := new(models.GnBaseData)
	b.ID = userid
	b.Gold = gold
	b.TopScore = TopScore
	b.AtkProp = AtkProp
	b.BoomProp = BoomProp
	b.OutTime = OutTime
	err = b.UpdateBase()
	if err != nil {
		beego.Error("gunner baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
