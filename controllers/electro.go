package controllers

import (
	"time"

	"goadmin/models"

	"github.com/astaxie/beego"
)

// ElectroController ...
type ElectroController struct {
	baseAdminController
}

// Index ...
func (c *ElectroController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gold, _ := c.GetInt("gold", 0)
	Type, _ := c.GetInt("type", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	projectID, _ := c.GetInt("project_id", 0)
	if !c.isPost() {
		sourceID = -1 //默认显示全部
	}
	beego.Info("electro index startTime: ", startTime)
	beego.Info("electro index endTime: ", endTime)
	beego.Info("electro index userid: ", userid)
	beego.Info("electro index gold: ", gold)
	var pagination models.Pagination
	var lists []models.EmBaseData
	var lists1 []models.EmBaseFu
	if Type == 1 {
		b := new(models.EmBaseFu)
		lists1, pagination, _ = b.GetDataList(userid, startTime, endTime,
			gold, page, c.pageSize, c.isPost())
	} else {
		b := new(models.EmBaseData)
		lists, pagination, _ = b.GetDataList(userid, startTime, endTime,
			gold, page, c.pageSize, c.isPost())
	}
	//beego.Info("electro lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		if Type == 1 {
			m["BaseData"] = lists1
		} else {
			m["BaseData"] = lists
		}
		c.RenderJson(200, "处理成功", m)
		return
	}
	sourceList := models.GetSourceList(projectID)
	pagination.BaseUrl = beego.URLFor("ElectroController.Index")
	beego.Info("electro pagination: ", pagination)
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
	if Type == 1 {
		c.Data["BaseData"] = &lists1
	} else {
		c.Data["BaseData"] = &lists
	}
	c.Data["Pagination"] = &pagination
	c.display("electro/index.html")
}

// BaseEdit ...
func (c *ElectroController) BaseEdit() {
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
	signin, _ := c.GetUint32("signin", 0)
	signTime := c.GetString("sign_time", "")
	sT, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", signTime, time.Local)
	beego.Info("electro signTime: ", signTime, sT)
	beego.Info("electro gold: ", gold)
	beego.Info("electro userid: ", userid)
	b := new(models.EmBaseData)
	b.ID = userid
	b.Gold = gold
	b.Gate = gate
	b.GuideStep = guideStep
	b.UnlockGate = unlock_gate
	b.Add = add
	b.Signin = signin
	b.SignTime = sT
	err = b.UpdateBase()
	if err != nil {
		beego.Error("electro err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// BaseDelete ...
func (c *ElectroController) BaseDelete() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	beego.Info("userid: ", userid)
	u := new(models.EmBaseData)
	u.ID = userid
	err := u.Remove()
	if err != nil {
		beego.Error("err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// FuDai ...
func (c *ElectroController) FuDai() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	b := new(models.EmBaseFu)
	b.UserID = userid
	err := b.GetData()
	if err != nil {
		beego.Error("err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", b)
}

// FuDaiEdit ...
func (c *ElectroController) FuDaiEdit() {
	id, _ := c.GetInt("id", 0)
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 || id == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	today, err := c.GetInt("today", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	days, err := c.GetUint32("days", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	level, err := c.GetUint32("level", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	bag, err := c.GetUint32("bag", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	bag_num, err := c.GetUint32("bag_num", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	number, err := c.GetUint32("number", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	ticket, err := c.GetUint32("ticket", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	b := new(models.EmBaseFu)
	b.ID = id
	b.UserID = userid
	b.Today = today
	b.Days = days
	b.Level = level
	b.Bag = bag
	b.BagNum = bag_num
	b.Number = number
	b.Ticket = ticket
	err = b.UpdateBase()
	if err != nil {
		beego.Error("electro err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// RankDelete ...
func (c *ElectroController) RankDelete() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	beego.Info("userid: ", userid)
	u := new(models.EmBaseData)
	u.ID = userid
	err := u.RankRemove()
	if err != nil {
		beego.Error("err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
