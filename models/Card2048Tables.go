package models

import (
	"time"
)

// CaBaseData 用户基础数据表
type CaBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"` //金币数量
	Gate uint32 `orm:"column(gate);default(0);" xorm:"INT(4) default(0) 'gate'" json:"gate"`  //关卡数量
	//
	Score   uint32 `orm:"column(score);default(0)" xorm:"INT(4) default(0) 'score'" json:"score"`       //分数
	Bomb    uint32 `orm:"column(bomb);default(0)" xorm:"INT(4) default(0) 'bomb'" json:"bomb"`          //炸弹数量
	Throw   uint32 `orm:"column(throw);default(0)" xorm:"INT(4) default(0) 'throw'" json:"throw"`       //throw
	Refresh uint32 `orm:"column(refresh);default(0)" xorm:"INT(4) default(0) 'refresh'" json:"refresh"` //refresh
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	//Wall     string `orm:"column(wall);size(255);null" xorm:"varchar(255) null 'wall'" json:"wall"`            // wall
	Add bool `orm:"column(add);bool" xorm:"Bool default(0) 'add'" json:"add"` // 添加到小程序
	//
	GuideStep uint32 `orm:"column(guide_step);default(0)" xorm:"INT(4) default(0) 'guide_step'" json:"guide_step"` //新手引导
	//
	UnlockGate uint32 `orm:"column(unlock_gate);default(0)" xorm:"INT(4) default(0) 'unlock_gate'" json:"unlock_gate"` //解锁关卡数量
	Prize      int    `orm:"column(prize);default(0)" xorm:"INT(4) default(0) 'prize'" json:"prize"`                   //prize
	//
	Playday  int    `orm:"column(playday);default(0)" xorm:"INT(4) default(0) 'playday'" json:"playday"`    //playday
	Playtime uint32 `orm:"column(playtime);default(0)" xorm:"INT(4) default(0) 'playtime'" json:"playtime"` //playtime
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *CaBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *CaBaseData) TableEngine() string {
	return "INNODB"
}
