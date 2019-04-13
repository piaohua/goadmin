package models

import (
	"time"
)

// ApiDoc 接口文档
type ApiDoc struct {
	ID        int `orm:"column(id);pk;auto" json:"id"`
	ProjectId int `orm:"column(project_id);default(0)" json:"project_id"` // 项目id
	//
	//自定义类型,如:
	//1-基础模板,2-接口模板,
	//3-基础说明(接口简介,接入须知,回调说明,环境说明,状态码说明),
	//4-接口说明(商城(购买,出售)...),5-其它
	//TODO 分组
	Type   int `orm:"column(type);default(0)" json:"type"`     // 自定义类型
	Status int `orm:"column(status);default(0)" json:"status"` // 状态:1-正在开发,2-正在审核,3-正常,4-未通过,5-暂停使用
	Sort   int `orm:"column(sort);default(0)" json:"sort"`     // 自定义排序编号
	//
	URL    string `orm:"column(url);size(100)" json:"url"`              // 接口地址
	Method int    `orm:"column(method);default(1)" json:"method"`       // 请求方式:1-GET 2-POST 3-PUT 4-PATCH 5-DELETE
	Name   string `orm:"column(name);size(100)" json:"name"`            // 接口名称或标题
	Detail string `orm:"column(detail);null;type(text);" json:"detail"` // 接口详细内容
	//
	DocID     int       `orm:"column(doc_id);default(0)" json:"doc_id"`                   // 接口ID(主键ID,修改记录)
	CreateID  int       `orm:"column(create_id);default(0)" json:"create_id"`             // 创建者ID
	UpdateID  int       `orm:"column(update_id);default(0)" json:"update_id"`             // 修改者ID
	AuditID   int       `orm:"column(audit_id);default(0)" json:"audit_id"`               // 审核人
	AuditTime time.Time `orm:"column(audit_time);null;type(datetime);" json:"audit_time"` // 审核时间
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *ApiDoc) TableName() string {
	return "system_apidoc"
}

// TableEngine 设置引擎为 INNODB
func (u *ApiDoc) TableEngine() string {
	return "INNODB"
}
