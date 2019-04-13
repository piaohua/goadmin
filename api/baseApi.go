package api

import (
	. "goadmin/controllers"
	"goadmin/libs"

	"github.com/astaxie/beego"
)

type baseApiController struct {
	BaseController
}

// Prepare ...
func (c *baseApiController) Prepare() {
	sign := c.GetString("sign")
	if sign == "" {
		c.RenderJson(400, "缺少必要验签参数", nil)
	}
	formData := c.Ctx.Request.Form
	beego.Info("url: ", c.Ctx.Input.URL(),
		", formData: ", formData)
	var param = make(map[string]string)
	for k, v := range formData {
		param[k] = v[len(v)-1]
	}
	apikey := beego.AppConfig.String("apikey")
	//if apikey == "" {
	//	c.RenderJson(400, "apikey无效", nil)
	//}
	if sign != libs.Sign(param, apikey) {
		c.RenderJson(400, "签名不正确", nil)
	}
}
