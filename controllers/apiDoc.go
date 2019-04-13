package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// ApiDocController ...
type ApiDocController struct {
	baseAdminController
}

// Show ...
func (c *ApiDocController) Show() {
	projectID, _ := c.GetInt("project_id", 0)
	id, _ := c.GetInt("id", 0)
	beego.Info("apidoc show projectID: ", projectID)
	g := &models.ApiDoc{ID: id}
	err := g.One(projectID)
	if err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.Data["ApiDocData"] = g
	c.display("apidoc/show.html")
}
