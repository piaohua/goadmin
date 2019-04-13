package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// LogController ...
type LogController struct {
	baseAdminController
}

// Login 登录记录日志
func (c *LogController) Login() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogLoginList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Login", "project_id", projectID)
	//beego.Info("pagination: ", pagination)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/login.html")
}

// Register 注册记录日志
func (c *LogController) Register() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogRegisterList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Register", "project_id", projectID)
	//beego.Info("pagination: ", pagination)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/register.html")
}

// SendClick 记录日志
func (c *LogController) SendClick() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogSendClickList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.SendClick", "project_id", projectID)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/sendclick.html")
}

// Progress 记录日志
func (c *LogController) Progress() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogProgressList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Progress", "project_id", projectID)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/progress.html")
}

// SceneAct 记录日志
func (c *LogController) SceneAct() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogSceneActList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.SceneAct", "project_id", projectID)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/sceneact.html")
}

// Playtime 记录日志
func (c *LogController) Playtime() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogPlaytimeList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Playtime", "project_id", projectID)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/playtime.html")
}

// Tip 记录日志
func (c *LogController) Tip() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogTipList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Tip", "project_id", projectID)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/tip.html")
}

// Wall 记录日志
func (c *LogController) Wall() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogWallList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Wall", "project_id", projectID)
	//beego.Info("pagination: ", pagination)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/wall.html")
}

// GateGrow 记录日志
func (c *LogController) GateGrow() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogGateGrowList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.GateGrow", "project_id", projectID)
	//beego.Info("pagination: ", pagination)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/gategrow.html")
}

// Gold 记录日志
func (c *LogController) Gold() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogGoldList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Gold", "project_id", projectID)
	//beego.Info("pagination: ", pagination)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/gold.html")
}

// Gate 记录日志
func (c *LogController) Gate() {
	page, _ := c.GetInt("page", 1)
	datetime := c.GetString("datetime", "")
	userid, _ := c.GetInt("userid", 0)
	projectID, _ := c.GetInt("project_id", 0)
	lists, pagination, _ := models.GetLogGateList(datetime,
		projectID, userid, page, c.pageSize, c.isPost())
	//beego.Info("lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("LogController.Gate", "project_id", projectID)
	//beego.Info("pagination: ", pagination)
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Datetime"] = datetime
	c.Data["LogData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("logger/gate.html")
}
