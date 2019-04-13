package controllers

import (
	"goadmin/libs"
	"goadmin/models"

	"github.com/astaxie/beego"
)

type SysuserController struct {
	baseAdminController
}

func (c *SysuserController) Index() {
	page, _ := c.GetInt("page", 1)
	user_name := c.GetString("uname", "")
	user_list, pagination := SystemUserService.GetPage(page, c.pageSize, user_name)
	beego.Info("user_list: ", user_list)
	var ids []int32
	for _, v := range user_list {
		ids = append(ids, v.RoleId)
	}
	//role_list, _ := UserRoleService.GetRoleListBy(ids)
	role_list, _, _ := UserRoleService.GetRoleList(1, 1000) //取全部
	beego.Info("role_list: ", role_list)
	if c.isPost() {
		var ls = struct {
			RoleList map[int32]models.UserRole `json:"RoleList"`
			UserList []models.SystemUser       `json:"UserList"`
		}{
			RoleList: role_list,
			UserList: user_list,
		}
		c.RenderJson(200, "处理成功", ls)
		return
	}
	pagination.BaseUrl = beego.URLFor("SysuserController.Index")
	c.Data["RoleList"] = &role_list
	c.Data["UserList"] = &user_list
	c.Data["Pagination"] = &pagination
	c.display("user/index.html")
}

func (c *SysuserController) Save() {
	var user_name = c.GetString("username")
	var password = c.GetString("password")
	var pwdcheck = c.GetString("pwdcheck")
	var email = c.GetString("email")
	var realname = c.GetString("realname")
	var role_id, _ = c.GetInt32("role_id", int32(0))
	var status, _ = c.GetInt8("status", int8(0))
	var uid, _ = c.GetInt32("user_id", 0)
	if uid == 0 && user_name == "" || email == "" || realname == "" || role_id == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	password_md5 := libs.MD5(password)
	pwdcheck_md5 := libs.MD5(pwdcheck)
	if password_md5 != pwdcheck_md5 {
		c.RenderJson(300, "密码有误", nil)
	}
	user := models.SystemUser{UserId: uid, UserName: user_name, Password: password_md5,
		Email: email, RoleId: role_id, Status: status, RealName: realname}
	if ok, err := user.Save(); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

func (c *SysuserController) UpdateStatus() {
	var status, _ = c.GetInt8("status", int8(0))
	var uid, _ = c.GetInt32("user_id", 0)
	if uid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	user := models.SystemUser{UserId: uid, Status: status}
	if ok, err := user.Update("Status"); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

func (c *SysuserController) ResetPasswd() {
	var uid, _ = c.GetInt32("user_id", 0)
	var password = c.GetString("password")
	var pwdcheck = c.GetString("pwdcheck")
	if uid == 0 || password == "" || pwdcheck == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	password_md5 := libs.MD5(password)
	pwdcheck_md5 := libs.MD5(pwdcheck)
	if password_md5 != pwdcheck_md5 {
		c.RenderJson(300, "密码有误", nil)
	}
	user := models.SystemUser{UserId: uid, Password: password_md5}
	if ok, err := user.Update("Password"); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
