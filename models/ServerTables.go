package models

import (
	"time"
)

// Server 服务列表
type Server struct {
	ID int `orm:"column(id);pk;auto" json:"id"`
	//base
	Name   string `orm:"column(name);size(30)" json:"name"`               //app name
	Host   string `orm:"column(host);size(20)" json:"host"`               //host
	Port   string `orm:"column(port);size(20)" json:"port"`               //port
	Path   string `orm:"column(path);size(255)" json:"path"`              //项目路径
	Status uint32 `orm:"column(status);size(4);default(0)" json:"status"` //状态: 0停止,1运行,2开放访问,3关闭访问
	//config
	ProjectID int    `orm:"column(project_id);size(4)" json:"project_id"`       //ProjectID
	JwtSecret string `orm:"column(jwt_secret);size(43)" json:"jwt_secret"`      //jwt secret
	Token     string `orm:"column(token);size(43);null" json:"token"`           //header token
	Appid     string `orm:"column(appid);size(18);null" json:"appid"`           //wx appid
	AppSecret string `orm:"column(app_secret);size(32);null" json:"app_secret"` //wx app secret
	Group     string `orm:"column(group);size(20);null" json:"group"`           //router group
	Prefixs   string `orm:"column(prefixs);size(255);null" json:"prefixs"`      //prefixs (json)
	//databases
	DefaultDB     string `orm:"column(default_db);size(255);null" json:"default_db"`        //默认数据库连接
	DefaultPrefix string `orm:"column(default_prefix);size(20);null" json:"default_prefix"` //db prefix
	LoggerDB      string `orm:"column(logger_db);size(255);null" json:"logger_db"`          //日志数据库连接
	LoggerPrefix  string `orm:"column(logger_prefix);size(20);null" json:"logger_prefix"`   //db prefix
	// time
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`     // 更新时间
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *Server) TableName() string {
	return "system_server"
}

// TableEngine 设置引擎为 INNODB
func (u *Server) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *Server) TableUnique() [][]string {
	return [][]string{
		[]string{"Name", "Host"},
	}
}
