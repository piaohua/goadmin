package models

import (
	"time"
)

// SmBaseData 用户基础数据表
type SmBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold    string `orm:"column(gold);size(32);null" xorm:"varchar(32) null 'gold'" json:"gold"`         //金币数量
	Diamond uint32 `orm:"column(diamond);default(0)" xorm:"INT(11) default(0) 'diamond'" json:"diamond"` //钻石数量
	//
	FruitInfo       string `orm:"column(fruit_info);size(800);null" xorm:"varchar(800) null 'fruit_info'" json:"fruit_info"`                   //水果信息
	TechnologyInfo  string `orm:"column(technology_info);size(255);null" xorm:"varchar(255) null 'technology_info'" json:"technology_info"`    //武器信息
	AchievementInfo string `orm:"column(achievement_info);size(500);null" xorm:"varchar(500) null 'achievement_info'" json:"achievement_info"` //成就信息
	HideTime        int64  `orm:"column(hide_time);default(0)" xorm:"INT(11) default(0) 'hide_time'" json:"hide_time"`                         //用户退出时间戳
	//
	OfflineLv int `orm:"column(offline_lv);default(0)" xorm:"INT default(0) 'offline_lv'" json:"offline_lv"` //离线数据
	GuideStep int `orm:"column(guide_step);default(0)" xorm:"INT default(0) 'guide_step'" json:"guide_step"` //新手引导
	//
	//SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *SmBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *SmBaseData) TableEngine() string {
	return "INNODB"
}
