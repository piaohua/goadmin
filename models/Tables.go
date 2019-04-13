package models

import (
	"time"
)

// SystemUser 系统用户
type SystemUser struct {
	UserId   int32     `orm:"column(user_id);pk;auto"`
	UserName string    `orm:"column(username)"`
	Password string    `orm:"column(password);size(32)"`
	Email    string    `orm:"column(email);unique"`
	RealName string    `orm:"column(realname)"`
	RoleId   int32     `orm:"column(role_id)"`
	Status   int8      `orm:"column(status)"`
	Created  time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

// TableName 表名
func (u *SystemUser) TableName() string {
	return "system_user"
}

// TableEngine 设置引擎为 INNODB
func (u *SystemUser) TableEngine() string {
	return "INNODB"
}

// SystemModule 系统模块
type SystemModule struct {
	ModuleId   int32          `orm:"column(module_id);pk;auto"`
	ModuleName string         `orm:"column(module_name)"`
	MethodName string         `orm:"column(method_name)"`
	Title      string         `orm:"column(title)"`
	Icon       string         `orm:"column(icon)"`
	ParentId   int32          `orm:"column(parent_id)"`
	Sort       int32          `orm:"column(sort)"`
	Show       int8           `orm:"column(show)"`
	Type       int8           `orm:"column(type)"`
	Status     int8           `orm:"column(status)"`
	Children   []SystemModule `orm:"-"`
}

// TableName 表名
func (u *SystemModule) TableName() string {
	return "system_module"
}

// TableEngine 设置引擎为 INNODB
func (u *SystemModule) TableEngine() string {
	return "INNODB"
}

// UserRole 系统角色
type UserRole struct {
	RoleId      int32  `orm:"column(role_id);pk;auto"`
	RoleName    string `orm:"column(role_name)"`
	Description string `orm:"column(description)"`
	Count       int64  `orm:"-"`
	Status      int8   `orm:"column(status)"`
	ProjectIds  string `orm:"column(project_ids)"`
	//
	Projects    []int     `orm:"-"`
	ProjectList []Project `orm:"-"`
}

// TableName 表名
func (u *UserRole) TableName() string {
	return "system_role"
}

// TableEngine 设置引擎为 INNODB
func (u *UserRole) TableEngine() string {
	return "INNODB"
}

// UserPermission 用户权限
type UserPermission struct {
	Id       int32 `orm:"column(id);pk"`
	RoleId   int32 `orm:"column(role_id)"`
	ModuleId int32 `orm:"column(module_id)"`
}

// TableName 表名
func (u *UserPermission) TableName() string {
	return "system_access"
}

// TableEngine 设置引擎为 INNODB
func (u *UserPermission) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *UserPermission) TableUnique() [][]string {
	return [][]string{
		[]string{"RoleId", "ModuleId"},
	}
}

// SystemConfig 系统设置
type SystemConfig struct {
	Key         string `orm:"column(k);size(32);pk"`
	Value       string `orm:"column(v);type(text)"`
	Description string `orm:"column(d);size(100)"`
}

// DicValue 字典值
type DicValue struct {
	Key         string `json:"k"`
	Description string `json:"d"`
	Time        int64  `json:"at"`
	Operate     string `json:"opt"`
}

// TableName 表名
func (u *SystemConfig) TableName() string {
	return "system_setting"
}

// TableEngine 设置引擎为 INNODB
func (u *SystemConfig) TableEngine() string {
	return "INNODB"
}

// Pagination 分页对象
type Pagination struct {
	CurrentPage int32  `json:"currentPage"`
	TotalCounts int32  `json:"totalCounts"`
	BaseUrl     string `json:"baseUrl"`
}
