package controllers

import "goadmin/models"

var (
	UserRoleService     *models.UserRole       // 系统角色
	UserPermService     *models.UserPermission // 用户权限
	SystemUserService   *models.SystemUser     // 系统用户
	SystemConfigService *models.SystemConfig   // 系统设置
	SystemModuleService *models.SystemModule   // 系统模块
)

func init() {
	UserRoleService = &models.UserRole{}
	UserPermService = &models.UserPermission{}
	SystemUserService = &models.SystemUser{}
	SystemConfigService = &models.SystemConfig{}
	SystemModuleService = &models.SystemModule{}
}
