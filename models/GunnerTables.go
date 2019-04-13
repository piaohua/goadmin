package models

import (
	"time"
)

// GnBaseData 用户基础数据表
type GnBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold     uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"`               //金币数量
	Gate     uint32 `orm:"column(gate);default(0);" xorm:"INT(4) default(0) 'gate'" json:"gate"`                //关卡数量
	TopScore uint32 `orm:"column(top_score);default(0)" xorm:"INT(11) default(0) 'top_score'" json:"top_score"` //最高分数
	AtkProp  uint32 `orm:"column(atk_prop);default(0)" xorm:"INT(11) default(0) 'atk_prop'" json:"atk_prop"`    //连击道具数量
	BoomProp uint32 `orm:"column(boom_prop);default(0)" xorm:"INT(11) default(0) 'boom_prop'" json:"boom_prop"` //手雷道具数量
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	//
	Add     bool  `orm:"column(add);bool" xorm:"Bool default(0) 'add'" json:"add"`                         // 添加到小程序
	OutTime int64 `orm:"column(out_time);default(0)" xorm:"INT(11) default(0) 'out_time'" json:"out_time"` //退出时间
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *GnBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *GnBaseData) TableEngine() string {
	return "INNODB"
}
