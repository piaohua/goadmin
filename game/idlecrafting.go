package controllers

import (
	"goadmin/libs"
	"goadmin/models"

	"github.com/astaxie/beego"
)

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
	pagination.BaseUrl = beego.URLFor("StatController.Index")
	beego.Info("stat pagination: ", pagination)
	if version != 0 {
		c.Data["Version"] = version
	}
	if cover != 0 {
		c.Data["Cover"] = cover
	}
	projectList := models.GetProjectList()
	c.Data["Projects"] = &projectList
	c.Data["ProjectID"] = projectID
	c.Data["Timestamp"] = timestamp
	c.Data["StatData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/index.html")
}
