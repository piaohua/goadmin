package models

import (
	"time"
)

// PtBaseData 用户基础数据表
type PtBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"` //金币数量
	Gate uint32 `orm:"column(gate);default(0);" xorm:"INT(4) default(0) 'gate'" json:"gate"`  //关卡数量
	//
	SourceID int  `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	Add      bool `orm:"column(add);bool" xorm:"Bool default(0) 'add'" json:"add"`                           // 添加到小程序
	//
	GuideStep uint32 `xorm:"INT(4) default(0) 'guide_step'" json:"guide_step"` //新手引导
	//
	UnlockGate uint32 `orm:"column(unlock_gate);default(0)" xorm:"INT(4) default(0) 'unlock_gate'" json:"unlock_gate"` //解锁关卡数量
	// sign
	Signin   uint32    `orm:"column(signin);default(0)" xorm:"INT(4) default(0) 'signin'" json:"signin"` //签到0-6天 (1 << i)
	SignTime time.Time `orm:"column(sign_time);datetime" xorm:"datetime 'sign_time'" json:"sign_time"`   //签到时间
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *PtBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *PtBaseData) TableEngine() string {
	return "INNODB"
}

// PtBaseFu 用户福袋数据表
type PtBaseFu struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid);unique" xorm:"INT notnull unique(fu) 'userid'" json:"userid,omitempty"` // userid
	//当天状态
	Today  int    `orm:"column(today);default(0);size(4)" xorm:"INT(4) default(0) 'today'" json:"today"`       //当天日期yyyymmdd
	Days   uint32 `orm:"column(days);default(0);size(4)" xorm:"INT(4) default(0) 'days'" json:"days"`          //第几天
	Level  uint32 `orm:"column(level);default(0);size(4)" xorm:"INT(4) default(0) 'level'" json:"level"`       //当前福袋通关数量
	Bag    uint32 `orm:"column(bag);default(0);size(4)" xorm:"INT(4) default(0) 'bag'" json:"bag"`             //当前进行第几个福袋
	BagNum uint32 `orm:"column(bag_num);default(0);size(4)" xorm:"INT(4) default(0) 'bag_num'" json:"bag_num"` //当天领取第几个福袋
	//累计数量
	Number uint32 `orm:"column(number);default(0);size(4)" xorm:"INT(4) default(0) 'number'" json:"number"` //福袋获得总数量
	Ticket uint32 `orm:"column(ticket);default(0);size(4)" xorm:"INT(4) default(0) 'ticket'" json:"ticket"` //福券获得总数量
	//
	Give uint32 `orm:"column(give);default(0)" xorm:"TINYINT default(0) 'give'" json:"give"` //福袋赠送0无,1可领取,2已领取,3跳过
	//
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *PtBaseFu) TableName() string {
	return "base_fu"
}

// TableEngine 设置引擎为 INNODB
func (u *PtBaseFu) TableEngine() string {
	return "INNODB"
}
