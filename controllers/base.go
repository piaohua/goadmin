package controllers

import (
	"github.com/astaxie/beego"
)

// BaseController ...
type BaseController struct {
	beego.Controller
}

// 响应json数据格式
type ResponseJson struct {
	Code    int32       `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

// RenderJson ...
func (c *BaseController) RenderJson(code int32, message string, data interface{}) {
	mystruct := ResponseJson{Code: code, Data: data, Message: message}
	c.Data["json"] = &mystruct
	c.ServeJSON()
	c.StopRun()
}

// RenderScript ...
func (c *BaseController) RenderScript(message string, back bool) {
	ext_str := ""
	if back {
		ext_str += "window.history.back();"
	}
	content := "<script>alert(\"" + message + "\");" + ext_str + "</script>"
	c.Ctx.Output.Body([]byte(content))
	c.StopRun()
}

// 是否POST提交
func (c *BaseController) isPost() bool {
	return c.Ctx.Request.Method == "POST"
}

// 重定向
func (c *BaseController) redirect(url string) {
	c.Redirect(url, 302)
	c.StopRun()
}

// 获取用户IP地址
func (c *BaseController) getClientIp() string {
	if p := c.Ctx.Input.Proxy(); len(p) > 0 {
		return p[0]
	}
	return c.Ctx.Input.IP()
}

// 渲染模版
func (c *BaseController) display(tplname string) {
	c.TplName = tplname
}
