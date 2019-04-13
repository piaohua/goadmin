package models

import (
	"time"
)

// TgBaseData 用户基础数据表
type TgBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"` //金币数量
	Gate uint32 `orm:"column(gate);default(0);" xorm:"INT(4) default(0) 'gate'" json:"gate"`  //关卡数量
	//
	SourceID int    `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	Wall     string `orm:"column(wall);size(255);null" xorm:"varchar(255) null 'wall'" json:"wall"`            // wall
	Add      bool   `orm:"column(add);bool" xorm:"Bool default(0) 'add'" json:"add"`                           // 添加到小程序
	//
	GuideStep uint32 `orm:"column(guide_step);default(0)" xorm:"INT(4) default(0) 'guide_step'" json:"guide_step"` //新手引导
	//
	UnlockGate uint32 `orm:"column(unlock_gate);default(0)" xorm:"INT(4) default(0) 'unlock_gate'" json:"unlock_gate"` //解锁关卡数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *TgBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *TgBaseData) TableEngine() string {
	return "INNODB"
}
