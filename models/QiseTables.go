package models

import (
	"time"
)

// QsBaseData 用户基础数据表
type QsBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"` //金币数量
	Gate uint32 `orm:"column(gate);default(0)" xorm:"INT(11) default(0) 'gate'" json:"gate"`  //关卡数值
	Box  uint32 `orm:"column(box);default(0)" xorm:"INT(11) default(0) 'box'" json:"box"`     //宝箱数量
	//
	GuideStep uint32 `orm:"column(guide_step);default(0)" xorm:"INT(4) default(0) 'guide_step'" json:"guide_step"` //新手引导
	Version   uint32 `orm:"column(version);default(0)" xorm:"INT(4) default(0) 'version'" json:"version"`          //关卡版本标识
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	//Wall     []string `orm:"column(wall);size(255);null"xorm:"varchar(255) null 'wall'" json:"wall"`             // wall
	Wall string `orm:"column(wall);size(255);null" xorm:"varchar(255) null 'wall'" json:"wall"` // wall
	Add  bool   `orm:"column(add);bool" xorm:"Bool default(0) 'add'" json:"add"`                // 添加到小程序
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *QsBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *QsBaseData) TableEngine() string {
	return "INNODB"
}

// Qs99BaseData 用户基础数据表
type Qs99BaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"` //金币数量
	Gate uint32 `orm:"column(gate);default(0)" xorm:"INT(11) default(0) 'gate'" json:"gate"`  //关卡数值
	Box  uint32 `orm:"column(box);default(0)" xorm:"INT(11) default(0) 'box'" json:"box"`     //宝箱数量
	//
	GuideStep uint32 `orm:"column(guide_step);default(0)" xorm:"INT(4) default(0) 'guide_step'" json:"guide_step"` //新手引导
	Version   uint32 `orm:"column(version);default(0)" xorm:"INT(4) default(0) 'version'" json:"version"`          //关卡版本标识
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	//Wall     []string `orm:"column(wall);size(255);null"xorm:"varchar(255) null 'wall'" json:"wall"`             // wall
	Wall string `orm:"column(wall);size(255);null"xorm:"varchar(255) null 'wall'" json:"wall"` // wall
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *Qs99BaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *Qs99BaseData) TableEngine() string {
	return "INNODB"
}
