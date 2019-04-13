package models

import (
	"time"
)

// BzBaseData 用户基础数据表
type BzBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Gold      uint32 `orm:"column(gold);default(0);" xorm:"INT(11) default(0) 'gold'" json:"gold"`                   //金币数量
	TopGold   uint32 `orm:"column(top_gold);default(0);" xorm:"INT(11) default(0) 'top_gold'" json:"top_gold"`       //累计金币数量
	TopGate   uint32 `orm:"column(top_gate);default(0)" xorm:"INT(11) default(0) 'top_gate'" json:"top_gate"`        //最大关卡数
	TopAttack uint32 `orm:"column(top_attack);default(0);" xorm:"INT(11) default(0) 'top_attack'" json:"top_attack"` //最大连击数
	SkinID    uint32 `orm:"column(skinid);default(0);" xorm:"INT(11) default(0) 'skinid'" json:"skinid"`             //当前皮肤
	SkinLock  string `orm:"column(skin_lock);size(255);null;" xorm:"varchar(255) null 'skin_lock'" json:"skin_lock"` //皮肤解锁
	HighScore uint32 `orm:"column(high_score);default(0);" xorm:"INT(11) default(0) 'high_score'" json:"high_score"` //最高积分数量
	TopScore  uint32 `orm:"column(top_score);default(0);" xorm:"INT(11) default(0) 'top_score'" json:"top_score"`    //累计积分数量
	PetCount  uint32 `orm:"column(pet_count);default(0);" xorm:"INT(11) default(0) 'pet_count'" json:"pet_count"`    //宠物数值
	// pass
	PerfectPass uint32 `orm:"column(perfect_pass);default(0)"xorm:"INT(11) default(0) 'perfect_pass'" json:"perfect_pass"` //完美通关次数
	NicePass    uint32 `orm:"column(nice_pass);default(0)"xorm:"INT(11) default(0) 'nice_pass'" json:"nice_pass"`          //优秀通关次数
	GreatPass   uint32 `orm:"column(great_pass);default(0)"xorm:"INT(11) default(0) 'great_pass'" json:"great_pass"`       //很棒通关次数
	NormalPass  uint32 `orm:"column(normal_pass);default(0)"xorm:"INT(11) default(0) 'normal_pass'" json:"normal_pass"`    //通过通关次数
	//login
	TopLoginTimes uint32    `orm:"column(top_login_times);default(0)" xorm:"INT(11) default(0) 'top_login_times'" json:"top_login_times"` //最长连续登录天数
	LoginTimes    uint32    `orm:"column(login_times);default(0)" xorm:"INT(11) default(0) 'login_times'" json:"login_times"`             //连续登录天数
	LoginTime     time.Time `orm:"column(login_time);null" xorm:"DateTime null 'login_time'" json:"login_time"`                           //登录时间
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *BzBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *BzBaseData) TableEngine() string {
	return "INNODB"
}

// BzBaseAction 用户操作动作数据表
type BzBaseAction struct {
	ID     int `xorm:"INT(11) notnull pk autoincr 'id'" json:"-"`                // id
	UserID int `xorm:"INT notnull unique(userid_action) 'userid'" json:"userid"` // userid
	//
	Type   int32  `xorm:"INT notnull unique(userid_action) 'type'" json:"type"` // 类型
	Number uint32 `xorm:"INT default(0) 'number'" json:"number"`                // 数量
	//
	CreatedAt time.Time `xorm:"created" json:"created_at,omitempty"`
	UpdatedAt time.Time `xorm:"updated" json:"updated_at,omitempty"`
}

// TableName 表名
func (u *BzBaseAction) TableName() string {
	return "base_action"
}

// TableEngine 设置引擎为 INNODB
func (u *BzBaseAction) TableEngine() string {
	return "INNODB"
}
