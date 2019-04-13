package api

import (
	"github.com/astaxie/beego"
)

// StatisController ...
type StatisController struct {
	baseApiController
}

// Stat ...
func (c *StatisController) Stat() {
	timestamp, _ := c.GetInt("timestamp")
	beego.Info("stat timestamp: ", timestamp)
	c.RenderJson(200, "处理成功", nil)
}
