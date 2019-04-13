package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// QiseController ...
type QiseController struct {
	baseAdminController
}

// Index ...
func (c *QiseController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gold, _ := c.GetInt("gold", 0)
	gate, _ := c.GetInt("gate", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	if !c.isPost() {
		sourceID = -1 //默认显示全部
	}
	beego.Info("qise index startTime: ", startTime)
	beego.Info("qise index endTime: ", endTime)
	beego.Info("qise index userid: ", userid)
	beego.Info("qise index gold: ", gold)
	beego.Info("qise index gate: ", gate)
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("qise index projectID: ", projectID)
	var lists interface{}
	var pagination models.Pagination
	switch projectID {
	case 18: //体验服数据
		b := new(models.Qs99BaseData)
		lists, pagination, _ = b.GetDataList(sourceID, userid, startTime, endTime,
			gold, gate, page, c.pageSize, c.isPost())
	default:
		b := new(models.QsBaseData)
		lists, pagination, _ = b.GetDataList(sourceID, userid, startTime, endTime,
			gold, gate, page, c.pageSize, c.isPost())
	}
	//beego.Info("qise lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["QsData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	sourceList := models.GetSourceList(projectID)
	pagination.BaseUrl = beego.URLFor("QiseController.Index")
	beego.Info("qise pagination: ", pagination)
	if gold != 0 {
		c.Data["Gold"] = gold
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
	c.Data["QsData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("qise/index.html")
}

// BaseEdit ...
func (c *QiseController) BaseEdit() {
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
		c.RenderJson(300, "关卡参数错误", nil)
	}
	guide_step, err := c.GetUint32("guide_step", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "指引参数错误", nil)
	}
	add, err := c.GetBool("add", false)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	version, _ := c.GetUint32("version", 0)
	beego.Info("qise baseEdit gold: ", gold)
	beego.Info("qise baseEdit gate: ", gate)
	beego.Info("qise baseEdit guide_step: ", guide_step)
	beego.Info("qise baseEdit userid: ", userid)
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("qise baseEdit projectID: ", projectID)
	switch projectID {
	case 18:
		b := new(models.Qs99BaseData)
		b.ID = userid
		b.Gold = gold
		b.Gate = gate
		b.GuideStep = guide_step
		b.Version = version
		err = b.UpdateBase()
		if err != nil {
			beego.Error("qise baseEdit err: ", err)
			c.RenderJson(300, err.Error(), nil)
		}
	default:
		b := new(models.QsBaseData)
		b.ID = userid
		b.Gold = gold
		b.Gate = gate
		b.GuideStep = guide_step
		b.Version = version
		b.Add = add
		err = b.UpdateBase()
		if err != nil {
			beego.Error("qise baseEdit err: ", err)
			c.RenderJson(300, err.Error(), nil)
		}
	}
	c.RenderJson(200, "处理成功", nil)
}
