package controllers

import (
	"encoding/json"
	"strconv"

	"goadmin/models"

	"github.com/astaxie/beego"
)

type RoleController struct {
	baseAdminController
}

func (c *RoleController) Index() {
	page, _ := c.GetInt("page", 1)
	role_lists, pagination, _ := UserRoleService.GetRoleList(page, c.pageSize)
	if c.isPost() {
		c.RenderJson(200, "处理成功", role_lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("RoleController.Index")
	c.Data["RoleData"] = &role_lists
	c.Data["Pagination"] = &pagination
	c.display("role/index.html")
}

// Access 权限设置
func (c *RoleController) Access() {
	var role_id, _ = c.GetInt32("role_id", 0)
	access_lists := UserPermService.GetAccessList(role_id)
	c.Data["AccessData"] = &access_lists
	c.Data["RoleId"] = role_id
	module_lists, _ := SystemModuleService.GetModuleListAll()
	c.Data["ModuleData"] = &module_lists
	c.display("role/access.html")
}

func (c *RoleController) SetAccess() {
	var item = c.GetStrings("item")
	var id, _ = c.GetInt32("role_id", 0)
	if id == 0 || item == nil {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	var permissions []models.UserPermission
	for _, value := range item {
		module_id, _ := strconv.ParseInt(value, 10, 32)
		permissions = append(permissions, models.UserPermission{RoleId: id,
			ModuleId: int32(module_id)})
	}
	if ok := UserPermService.SetPermission(id, permissions); !ok {
		c.RenderJson(300, "设置失败", nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

func (c *RoleController) ViewRole() {
	var id, _ = c.GetInt32("role_id", 0)
	c.Data["RoleName"] = UserRoleService.GetRoleName(id)
	c.Data["UserList"] = SystemUserService.GetRoleUsers(id)
	c.displayTpl("role/viewRole.html")
}

func (c *RoleController) Save() {
	var name = c.GetString("role_name")
	var desc = c.GetString("description")
	var status, _ = c.GetInt8("status", int8(0))
	var id, _ = c.GetInt32("role_id", 0)
	if name == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	// 项目选择
	projectIds := make([]int, 0)
	for _, v := range c.GetStrings("project_ids") {
		if projectId, _ := strconv.Atoi(v); projectId > 0 {
			projectIds = append(projectIds, projectId)
		}
	}
	//beego.Info("projectIds: ", projectIds)
	//
	module := models.UserRole{
		RoleId:      id,
		RoleName:    name,
		Description: desc,
		Status:      status,
	}
	//
	b, err := json.Marshal(projectIds)
	if err != nil {
		beego.Error("projectIds error ", projectIds)
	} else {
		module.ProjectIds = string(b)
	}
	if ok, err := module.Save(); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
