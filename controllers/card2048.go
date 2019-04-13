package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// Card2048Controller ...
type Card2048Controller struct {
	baseAdminController
}

// Index ...
func (c *Card2048Controller) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gold, _ := c.GetInt("gold", 0)
	score, _ := c.GetUint32("score", 0)
	bomb, _ := c.GetUint32("bomb", 0)
	throw, _ := c.GetUint32("throw", 0)
	refresh, _ := c.GetUint32("refresh", 0)
	gate, _ := c.GetUint32("gate", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	projectID, _ := c.GetInt("project_id", 0)
	if !c.isPost() {
		sourceID = -1 //默认显示全部
	}
	beego.Info("index startTime: ", startTime)
	beego.Info("index endTime: ", endTime)
	beego.Info("index userid: ", userid)
	beego.Info("index gold: ", gold)
	b := new(models.CaBaseData)
	b.Score = score
	b.Bomb = bomb
	b.Throw = throw
	b.Refresh = refresh
	b.Gate = gate
	lists, pagination, _ := b.GetDataList(userid, startTime, endTime,
		gold, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["BaseData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	sourceList := models.GetSourceList(projectID)
	pagination.BaseUrl = beego.URLFor("Card2048Controller.Index")
	beego.Info("pagination: ", pagination)
	if gold != 0 {
		c.Data["Gold"] = gold
	}
	if score != 0 {
		c.Data["Score"] = score
	}
	if bomb != 0 {
		c.Data["Bomb"] = bomb
	}
	if throw != 0 {
		c.Data["Throw"] = throw
	}
	if refresh != 0 {
		c.Data["Refresh"] = refresh
	}
	if gate != 0 {
		c.Data["Gate"] = gate
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
	c.display("card2048/index.html")
}

// BaseEdit ...
func (c *Card2048Controller) BaseEdit() {
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
	score, err := c.GetUint32("score", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	bomb, err := c.GetUint32("bomb", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	throw, err := c.GetUint32("throw", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	refresh, _ := c.GetUint32("refresh", 0)
	prize, _ := c.GetInt("prize", 0)
	playday, _ := c.GetInt("playday", 0)
	playtime, _ := c.GetUint32("playtime", 0)
	beego.Info("baseEdit gold: ", gold)
	beego.Info("baseEdit userid: ", userid)
	b := new(models.CaBaseData)
	b.ID = userid
	b.Gold = gold
	b.Score = score
	b.Bomb = bomb
	b.Throw = throw
	b.Refresh = refresh
	b.Gate = gate
	b.GuideStep = guideStep
	b.UnlockGate = unlock_gate
	b.Add = add
	b.Prize = prize
	b.Playday = playday
	b.Playtime = playtime
	err = b.UpdateBase()
	if err != nil {
		beego.Error("baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// BaseDelete ...
func (c *Card2048Controller) BaseDelete() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	beego.Info("userid: ", userid)
	u := new(models.CaBaseData)
	u.ID = userid
	err := u.Remove()
	if err != nil {
		beego.Error("err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
