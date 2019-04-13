package controllers

import (
	"goadmin/libs"
	"goadmin/models"

	"github.com/astaxie/beego"
)

// StatController ...
type StatController struct {
	baseAdminController
}

// Index ...
func (c *StatController) Index() {
	page, _ := c.GetInt("page", 1)
	timestamp := c.GetString("timestamp", "")
	version, _ := c.GetInt("version", 0)
	cover, _ := c.GetInt("cover", 0)
	var t int64
	if timestamp != "" {
		t, _ = libs.Str2Local(timestamp)
	}
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("stat index t: ", t)
	beego.Info("stat index timestamp: ", timestamp)
	beego.Info("stat index version: ", version)
	beego.Info("stat index cover: ", cover)
	stat := new(models.StatSendClick)
	lists, pagination, _ := stat.GetStatList(int(t), projectID, version, cover, page, c.pageSize)
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.Index", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	if version != 0 {
		c.Data["Version"] = version
	}
	if cover != 0 {
		c.Data["Cover"] = cover
	}
	c.Data["Timestamp"] = timestamp
	c.Data["StatData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/index.html")
}

// GateScore 关卡挑战/奖励得分
func (c *StatController) GateScore() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	Type, _ := c.GetInt("type", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	if !c.isPost() {
		Type = 1 //默认
	}
	stat := new(models.StatGateScore)
	lists, pagination, _ := stat.GetGateScoreList(dateat, projectID, Type, page, c.pageSize, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.GateScore", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["GateScoreData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/gatescore.html")
}

// GateRetimes 关卡挑战/奖励得分
func (c *StatController) GateRetimes() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	Type, _ := c.GetInt("type", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	if !c.isPost() {
		Type = 1 //默认
	}
	stat := new(models.StatGateRetimes)
	lists, pagination, _ := stat.GetGateRetimesList(dateat, projectID, Type, page, c.pageSize, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.GateRetimes", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["GateRetimesData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/gateretimes.html")
}

// GateLoss 新手关卡流失
func (c *StatController) GateLoss() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	if dateat == "" {
		dateat = libs.TodayStr(-1)
	}
	stat := new(models.StatGateLoss)
	lists, pagination, _ := stat.GetGateLossList(dateat, projectID, page, 100, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.GateLoss", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["GateLossData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/gateloss.html")
}

// PetLevel 宠物主角升级
func (c *StatController) PetLevel() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	stat := new(models.StatPetLevel)
	lists, pagination, _ := stat.GetPetLevelList(dateat, projectID, page, 1, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.PetLevel", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["PetLevelData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/petlevel.html")
}

// GateWeapon 合成的武器等级
func (c *StatController) GateWeapon() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	if dateat == "" {
		dateat = libs.TodayStr(-1)
	}
	stat := new(models.StatGateWeapon)
	lists, pagination, _ := stat.GetGateWeaponList(dateat, projectID, page, 1200, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.GateWeapon", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["GateWeaponData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/gateweapon.html")
}

// GateGoldDis 金币产出消耗
func (c *StatController) GateGoldDis() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	if dateat == "" {
		dateat = libs.TodayStr(-1)
	}
	stat := new(models.StatGateGoldDis)
	lists, pagination, _ := stat.GetGateGoldDisList(dateat, projectID, page, 3200, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.GateGoldDis", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["DisData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/gategolddis.html")
}

// GateGoldDep 金币存量
func (c *StatController) GateGoldDep() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	if dateat == "" {
		dateat = libs.TodayStr(-1)
	}
	stat := new(models.StatGateGoldDep)
	lists, pagination, _ := stat.GetGateGoldDepList(dateat, projectID, page, 1200, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.GateGoldDep", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["DepData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/gategolddep.html")
}

// WeaponLev 武器等级停留
func (c *StatController) WeaponLev() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	stat := new(models.StatWeaponLev)
	lists, pagination, _ := stat.GetWeaponLevList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.WeaponLev", "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["WeaponLevData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/weaponlev.html")
}

// PropList 道具
func (c *StatController) PropList() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	Type, _ := c.GetUint32("type", 1)
	beego.Info("stat index projectID: ", projectID)
	beego.Info("stat index dateat: ", dateat)
	stat := new(models.StatProp)
	lists, pagination, _ := stat.GetPropList(Type, dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("stat lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("StatController.PropList", "type", Type, "project_id", projectID)
	beego.Info("stat pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["Type"] = Type
	c.Data["PropData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/proplist.html")
}
