package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// BanziController ...
type BanziController struct {
	baseAdminController
}

// Index ...
func (c *BanziController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gold, _ := c.GetInt("gold", 0)
	beego.Info("banzi index startTime: ", startTime)
	beego.Info("banzi index endTime: ", endTime)
	beego.Info("banzi index userid: ", userid)
	beego.Info("banzi index gold: ", gold)
	b := new(models.BzBaseData)
	lists, pagination, _ := b.GetDataList(userid, startTime, endTime,
		gold, page, c.pageSize, c.isPost())
	//beego.Info("banzi lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("BanziController.Index")
	beego.Info("banzi pagination: ", pagination)
	if gold != 0 {
		c.Data["Gold"] = gold
	}
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	c.Data["BzData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("banzi/index.html")
}

// BaseEdit ...
func (c *BanziController) BaseEdit() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	gold, err := c.GetUint32("gold", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "金币参数错误", nil)
	}
	skinLock := c.GetString("skin_lock", "")
	petCount, err := c.GetUint32("pet_count", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "宠物参数错误", nil)
	}
	topGate, err := c.GetUint32("top_gate", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	perfectPass, err := c.GetUint32("perfect_pass", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	beego.Info("banzi baseEdit gold: ", gold)
	beego.Info("banzi baseEdit userid: ", userid)
	b := new(models.BzBaseData)
	b.ID = userid
	b.Gold = gold
	b.SkinLock = skinLock
	b.PetCount = petCount
	b.TopGate = topGate
	b.PerfectPass = perfectPass
	err = b.UpdateBase()
	if err != nil {
		beego.Error("banzi baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
