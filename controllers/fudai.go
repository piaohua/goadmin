package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// FuDaiController ...
type FuDaiController struct {
	baseAdminController
}

// FuDai ...
func (c *FuDaiController) FuDai() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.GameFuDai
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.GameFuDai)
		lists, pagination, _ = p.GetGameFuDaiList(projectID, page, c.pageSize, c.isPost())
		beego.Info("game lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("FuDaiController.FuDai", "project_id", projectID)
	beego.Info("pagination: ", pagination)
	c.Data["FuDaiData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("fudai/gamefudai.html")
}

// FuDaiEdit ...
func (c *FuDaiController) FuDaiEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var bag, _ = c.GetUint32("bag", 0)
	var num, _ = c.GetUint32("num", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || bag == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameFuDai{
		ID:  id,
		Bag: bag,
		Num: num,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// FuGate ...
func (c *FuDaiController) FuGate() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.GameFuGate
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.GameFuGate)
		lists, pagination, _ = p.GetGameFuGateList(projectID, page, c.pageSize, c.isPost())
		beego.Info("game lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("FuDaiController.FuGate", "project_id", projectID)
	beego.Info("pagination: ", pagination)
	c.Data["FuGateData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("fudai/gamefugate.html")
}

// FuGateEdit ...
func (c *FuDaiController) FuGateEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var bag, _ = c.GetUint32("bag", 0)
	var num, _ = c.GetUint32("num", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || bag == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameFuGate{
		ID:  id,
		Bag: bag,
		Num: num,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// FuDay ...
func (c *FuDaiController) FuDay() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.GameFuDay
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.GameFuDay)
		lists, pagination, _ = p.GetGameFuDayList(projectID, page, c.pageSize, c.isPost())
		beego.Info("game lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("FuDaiController.FuDay", "project_id", projectID)
	beego.Info("pagination: ", pagination)
	c.Data["FuDayData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("fudai/gamefuday.html")
}

// FuDayEdit ...
func (c *FuDaiController) FuDayEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var day, _ = c.GetUint32("day", 0)
	var num, _ = c.GetUint32("num", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || day == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameFuDay{
		ID:  id,
		Day: day,
		Num: num,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
