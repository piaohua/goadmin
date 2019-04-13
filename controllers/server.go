package controllers

import (
	"encoding/json"
	"errors"
	"time"

	"goadmin/models"

	"github.com/astaxie/beego"
)

// ServerController ...
type ServerController struct {
	baseAdminController
}

// Index ...
func (c *ServerController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	name := c.GetString("appname", "")
	projectID, _ := c.GetInt("project_id", 0)
	report, _ := c.GetInt("report", 0)
	if report != 0 {
		projectID = report //report service
	}
	status, _ := c.GetInt("status", 0)
	beego.Info("server index startTime: ", startTime)
	beego.Info("server index endTime: ", endTime)
	if !c.isPost() {
		status = -1 //默认全部
	}
	b := new(models.Server)
	lists, pagination, _ := b.List(name, startTime, endTime,
		status, projectID, page, c.pageSize, c.isPost())
	//beego.Info("server lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["BaseData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("ServerController.Index")
	beego.Info("server pagination: ", pagination)
	c.Data["Status"] = status
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	c.Data["BaseData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("server/index.html")
}

// BaseEdit ...
func (c *ServerController) BaseEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	report, _ := c.GetInt("report", 0)
	if report != 0 {
		projectID = report //report service
	}
	id, _ := c.GetInt("id", 0)
	Name := c.GetString("name", "")
	Host := c.GetString("host", "")
	Port := c.GetString("port", "")
	Path := c.GetString("path", "")
	JwtSecret := c.GetString("jwt_secret", "")
	Token := c.GetString("token", "")
	Appid := c.GetString("appid", "")
	AppSecret := c.GetString("app_secret", "")
	Group := c.GetString("group", "")
	Prefixs := c.GetString("prefixs", "")
	DefaultDB := c.GetString("default_db", "")
	DefaultPrefix := c.GetString("default_prefix", "")
	LoggerDB := c.GetString("logger_db", "")
	LoggerPrefix := c.GetString("logger_prefix", "")
	if Name == "" || Host == "" || Port == "" || JwtSecret == "" ||
		Token == "" || Appid == "" || AppSecret == "" ||
		DefaultDB == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	if Prefixs != "" {
		m := make(map[int]string)
		if err := json.Unmarshal([]byte(Prefixs), &m); err != nil {
			c.RenderJson(300, "参数错误", nil)
		}
	}
	beego.Info("server baseEdit name: ", Name)
	b := new(models.Server)
	b.ID = id
	b.Name = Name
	b.Host = Host
	b.Port = Port
	b.Path = Path
	b.ProjectID = projectID
	b.JwtSecret = JwtSecret
	b.Token = Token
	b.Appid = Appid
	b.AppSecret = AppSecret
	b.Group = Group
	b.Prefixs = Prefixs
	b.DefaultDB = DefaultDB
	b.DefaultPrefix = DefaultPrefix
	b.LoggerDB = LoggerDB
	b.LoggerPrefix = LoggerPrefix
	err := b.Upsert()
	if err != nil {
		beego.Error("server baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// BaseAct 1start,2stop,3open,4close,5create,6update
func (c *ServerController) BaseAct() {
	id, _ := c.GetInt("id", 0)
	actStatus, _ := c.GetInt("status", 0)
	if actStatus < 1 || actStatus > 6 {
		c.RenderJson(300, "参数错误", nil)
	}
	b := new(models.Server)
	b.ID = id
	err := b.Get()
	if err != nil {
		beego.Error("server err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	switch actStatus {
	case 1, 5, 6:
		switch b.Status {
		case 0: //停止状态才能操作
		default:
			err = errors.New("停止状态才能操作")
		}
	case 2:
		switch b.Status {
		case 3: //关闭访问才能操作
		default:
			err = errors.New("关闭访问才能操作")
		}
	case 3:
		switch b.Status {
		case 1: //运行状态才能操作
		default:
			err = errors.New("运行状态才能操作")
		}
	case 4:
		switch b.Status {
		case 2: //开放访问时才能操作
		default:
			err = errors.New("开放访问时才能操作")
		}
	}
	if err != nil {
		beego.Error("server err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	val := &models.SystemChat{
		UserId:    c.userId,
		UserName:  c.getUserName(),
		CreatedAt: time.Now(),
	}
	err = b.Act(actStatus, val)
	val.Save()
	if err != nil {
		beego.Error("server err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	b.UpdateStatus(actStatus)
	c.RenderJson(200, "处理成功", nil)
}
