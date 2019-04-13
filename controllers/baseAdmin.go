package controllers

import (
	"strconv"
	"strings"

	"goadmin/models"

	"github.com/astaxie/beego"
)

type baseAdminController struct {
	BaseController
	controllerName string // 控制器名
	actionName     string // 动作名
	roleId         int32  // 当前登录的角色id
	userId         int32  // 当前登录的用户id
	pageSize       int    // 默认分页大小
	lang           string // 当前语言环境
	projectId      int    // 当前选择的项目
}

// Prepare ...
func (c *baseAdminController) Prepare() {
	// header set
	c.Ctx.Output.Header("X-Powered-By", beego.AppConfig.String("appname"))
	c.Ctx.Output.Header("X-Author-By", "piaohua")
	// baseAdmin set (大小写敏感)
	control, action := c.GetControllerAndAction()
	c.controllerName = beego.Substr(control, 0, len(control)-10)
	c.actionName = action
	c.pageSize = 20
	//userId
	c.getUserId()
	//判断登录
	if c.userId == 0 {
		c.redirect(beego.URLFor("MainController.Get"))
		return
	}
	c.Data["LoginUserName"] = c.GetSession("username")
	c.Data["LoginUserEmail"] = c.GetSession("email")
	//角色id
	c.getRoleId()
	//beego.Info("base userid roleid ", c.userId, c.roleId)
	//判断url权限
	c.accessPerm()
	//menu list
	roleMenus := c.getMenuList()
	c.Data["UserMenu"] = &roleMenus
	c.Data["Module"] = c.controllerName
	c.Data["Action"] = c.actionName
	//project list
	projectList := c.getProjectList()
	c.Data["Projects"] = &projectList
	c.getProjectId()
	//beego.Info("projectId : ", c.projectId)
	if c.projectId == 0 && len(projectList) != 0 {
		c.projectId = projectList[0].ProjectID //默认第一个
	}
	c.Data["projectId"] = c.projectId
	//set project_id
	projectID, _ := c.GetInt("project_id", 0)
	if projectID == 0 && c.projectId != 0 {
		c.Ctx.Input.SetParam("project_id", strconv.Itoa(c.projectId))
	}
	//c.Data["ProjectID"] = projectID
}

// 渲染模版
func (c *baseAdminController) display(tplname string) {
	if tplname == "" {
		tplname = c.controllerName + "/" + c.actionName + ".html"
		tplname = strings.ToLower(tplname)
	}
	//输出layout需要的数据
	c.Data["SystemName"] = beego.AppConfig.String("systenName")
	c.Data["SystemAlias"] = beego.AppConfig.String("systenAlias")
	c.Data["namespace"] = beego.AppConfig.String("namespace")
	//layout
	c.Layout = "common/layout.html"
	c.TplName = tplname
}

func (c *baseAdminController) displayTpl(tplname string) {
	if tplname == "" {
		tplname = c.controllerName + "/" + c.actionName + ".html"
		tplname = strings.ToLower(tplname)
	}
	c.TplName = tplname
}

// 提示消息
//func (c *baseAdminController) showMsg(msg string, msgno int, redirect ...string) {
//}

// 获取玩家id
func (c *baseAdminController) getUserId() {
	if v, ok := c.GetSession("user_id").(int32); ok {
		c.userId = v
	}
}

// 获取角色id
func (c *baseAdminController) getRoleId() {
	if v, ok := c.GetSession("role_id").(int32); ok {
		c.roleId = v
	}
}

// 获取玩家名字
func (c *baseAdminController) getUserName() string {
	if v, ok := c.GetSession("username").(string); ok {
		return v
	}
	return ""
}

// 检查是否有某个权限
func (c *baseAdminController) HasAccessPerm(module, action string) bool {
	if c.userId == 0 || c.roleId == 0 {
		return false
	}
	if v, ok := c.GetSession("UserPermission").(map[string]bool); ok {
		if val, ok := v[(module + "." + action)]; ok && val {
			return true
		}
	}
	return false
}

// 登录权限验证
func (c *baseAdminController) noPerm() {
	if c.IsAjax() {
		c.RenderJson(400, "您没有操作权限", nil)
	} else {
		c.RenderScript("您没有操作权限", true)
	}
}

// 设置权限
func (c *baseAdminController) accessPerm() {
	key := c.controllerName + "." + c.actionName
	//beego.Info("accessPerm key ", key)
	if v, ok := c.GetSession("UserPermission").(map[string]bool); ok {
		//beego.Info("accessPerm v ", v)
		if val, ok := v[key]; !(ok && val) {
			c.noPerm()
		}
		return
	}
	c.accessPerm2(key)
}

func (c *baseAdminController) accessPerm2(key string) {
	//key := c.controllerName + "." + c.actionName
	access, err := SystemModuleService.GetAccessModuleList(c.roleId)
	//beego.Info("accessPerm ", access, err)
	if err != nil {
		beego.Error("accessPerm err ", err)
	}
	if len(access) == 0 || err != nil {
		//设置默认页面
		access = map[string]bool{
			"Admin.Main":   true,
			"Admin.Logout": true,
		}
	}
	c.SetSession("UserPermission", access)
	if val, ok := access[key]; !(ok && val) {
		c.noPerm()
	}
}

// 功能菜单
func (c *baseAdminController) getMenuList() []models.ParentMenu {
	if v, ok := c.GetSession("UserMenu").([]models.ParentMenu); ok {
		return v
	}
	return c.getMenuList2()
}

func (c *baseAdminController) getMenuList2() []models.ParentMenu {
	roleMenus, err := UserRoleService.GetMenus(c.roleId)
	if err != nil {
		beego.Error("getMenuList err ", err)
		return nil
	}
	c.SetSession("UserMenu", roleMenus)
	return roleMenus
}

// 项目列表
func (c *baseAdminController) getProjectList() []models.Project {
	if v, ok := c.GetSession("Projects").([]models.Project); ok {
		return v
	}
	return c.getProjectList2()
}

func (c *baseAdminController) getProjectList2() []models.Project {
	projectList := models.GetProjectListBy(c.roleId)
	c.SetSession("Projects", projectList)
	return projectList
}

// 当前项目
func (c *baseAdminController) getProjectId() {
	if v, ok := c.GetSession("project_id").(int); ok {
		c.projectId = v
	}
}

// 项目选择
func (c *baseAdminController) setProjectId(projectId int) {
	c.SetSession("project_id", projectId)
}
