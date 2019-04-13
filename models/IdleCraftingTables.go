package models

import (
	"time"
)

// IcBaseData 用户基础数据表
type IcBaseData struct {
	ID         int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	OpenID     string `orm:"column(openid);size(28);unique" xorm:"CHAR(28) notnull unique 'openid'" json:"openid"` // wechat openid
	SessionKey string `orm:"column(session_key);size(24)" xorm:"varchar(24) null 'session_key'" json:"-"`          // wechat session key
	// base
	Level  uint32 `orm:"column(level);" xorm:"INT(11) default(1) 'level'" json:"level"` // 主角等级
	Gate   uint32 `orm:"column(gate);default(0)" xorm:"INT(11) default(0) 'gate'" json:"gate"`
	Gold   string `orm:"column(gold);default('0');size(255)" xorm:"varchar(255) default('0') 'gold'" json:"gold"`
	Jewel  uint32 `orm:"column(jewel);default(0)" xorm:"INT(11) default(0) 'jewel'" json:"jewel"`
	KeyNum uint32 `orm:"column(key_num);default(0)" xorm:"INT(11) default(0) 'key_num'" json:"key_num"`
	//
	TopGold   string `orm:"column(top_gold);default('0');size(255)" xorm:"varchar(255) default('0') 'top_gold'" json:"top_gold"` //累计金币
	TopJewel  uint32 `orm:"column(top_jewel);default(0)" xorm:"INT(11) default(0) 'top_jewel'" json:"top_jewel"`                 //累计宝石
	TopKeyNum uint32 `orm:"column(top_key_num);default(0)" xorm:"INT(11) default(0) 'top_key_num'" json:"top_key_num"`           //累计钥匙
	//
	Score        uint32 `orm:"column(score);default(0)" xorm:"INT(11) default(0) 'score'" json:"score"`                         //成就积分
	BrickNum     uint32 `orm:"column(brick_num);default(0)" xorm:"INT(11) default(0) 'brick_num'" json:"brick_num"`             //砖块数量
	LotteryTimes uint32 `orm:"column(lottery_times);default(0)" xorm:"INT(11) default(0) 'lottery_times'" json:"lottery_times"` //累计抽奖次数
	WeaponTimes  uint32 `orm:"column(weapon_times);default(0)" xorm:"INT(11) default(0) 'weapon_times'" json:"weapon_times"`    //使用黄金武器次数
	BeatBoss     uint32 `orm:"column(beat_boss);default(0)" xorm:"INT(11) default(0) 'beat_boss'" json:"beat_boss"`             //累计击败精英boss次数
	AutoAtk      uint32 `orm:"column(auto_atk);default(0)" xorm:"INT(11) default(0) 'auto_atk'" json:"auto_atk"`                //自动攻击次数
	//
	LeaveTime int64  `orm:"column(leave_time);default(0)" xorm:"INT(11) default(0) 'leave_time'" json:"leave_time"`    //离开时间截
	GuideStep string `orm:"column(guide_step);null;size(255)" xorm:"varchar(255) null 'guide_step'" json:"guide_step"` //新手引导
	// sign
	Signin   uint32    `orm:"column(signin);default(0)" xorm:"INT(11) default(0) 'signin'" json:"signin"` //签到0-27天 (1 << i)
	SignTime time.Time `orm:"column(sign_time);datetime" xorm:"datetime 'sign_time'" json:"sign_time"`    //签到时间
	//
	SourceID int    `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	Weapon   string `orm:"column(weapon);null;size(255)" xorm:"varchar(255) null 'weapon'" json:"weapon"`      //
	Gift     uint32 `orm:"column(gift);default(0)" xorm:"TINYINT default(0) 'gift'" json:"gift"`
	//
	WeaponLev uint32 `orm:"column(weapon_lev);default(1)" xorm:"INT(4) default(1) 'weapon_lev'" json:"weapon_lev"` //武器等级
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseData) TableName() string {
	return "base_data"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseData) TableEngine() string {
	return "INNODB"
}

// IcBaseEffect 用户额外效果加成数据表
type IcBaseEffect struct {
	ID       int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID   int    `orm:"column(userid)" xorm:"INT notnull unique(userid_effect) 'userid'" json:"userid"`          // userid
	EffectID int    `orm:"column(effect_id)" xorm:"INT notnull unique(userid_effect) 'effect_id'" json:"effect_id"` // effect id
	Number   uint32 `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"`                  // 数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseEffect) TableName() string {
	return "base_effect"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseEffect) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseEffect) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "EffectID"},
	}
}

// IcBaseMaterials 用户材料数据表
type IcBaseMaterials struct {
	ID          int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID      int    `orm:"column(userid)" xorm:"INT notnull unique(userid_materials) 'userid'" json:"userid"`                    // userid
	MaterialsID int    `orm:"column(materials_id);" xorm:"INT notnull unique(userid_materials) 'materials_id'" json:"materials_id"` // 材料id
	Number      uint32 `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"`                               // 数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseMaterials) TableName() string {
	return "base_materials"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseMaterials) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseMaterials) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "MaterialsID"},
	}
}

// IcBasePet 用户宠物数据表
type IcBasePet struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(userid_pet) 'userid'" json:"userid"`  // userid
	PetID  int `orm:"column(pet_id);" xorm:"INT notnull unique(userid_pet) 'pet_id'" json:"pet_id"` // 宠物id
	// level
	Level uint32 `orm:"column(level);default(1)" xorm:"INT default(1) 'level'" json:"level"` // 宠物初始等级(默认1级)
	Star  uint32 `orm:"column(star);default(0)" xorm:"INT default(0) 'star'" json:"star"`    // 宠物对应进化等级
	Extra uint32 `orm:"column(extra);default(0)" xorm:"INT default(0) 'extra'" json:"extra"` // 宠物对应进化等级的进度
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBasePet) TableName() string {
	return "base_pet"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBasePet) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBasePet) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "PetID"},
	}
}

// IcBaseSkill 用户技能数据表
type IcBaseSkill struct {
	ID      int    `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID  int    `orm:"column(userid)" xorm:"INT notnull unique(userid_skill) 'userid'" json:"userid"`        // userid
	SkillID int    `orm:"column(skill_id);" xorm:"INT notnull unique(userid_skill) 'skill_id'" json:"skill_id"` // 技能id
	Number  uint32 `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"`               // 数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseSkill) TableName() string {
	return "base_skill"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseSkill) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseSkill) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "SkillID"},
	}
}

// IcBaseAchiement 用户已完成成就数据表
type IcBaseAchiement struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(ach) 'userid'" json:"userid"` // userid
	//
	AchID  uint32 `orm:"column(achid)" xorm:"INT(4) notnull unique(ach) 'achid'" json:"achid"`      //成就ID
	Type   int32  `orm:"column(type)" xorm:"INT(4) notnull unique(ach) 'type'" json:"type"`         //成就类型
	Status uint32 `orm:"column(status);default(0)" xorm:"INT(4) default(0) 'status'" json:"status"` //成就状态0:初始,1:完成,2全部完成领取
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseAchiement) TableName() string {
	return "base_achiement"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseAchiement) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseAchiement) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "AchID", "Type"},
	}
}

// IcBaseTask 用户已完成任务数据表
type IcBaseTask struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(userid_task) 'userid'" json:"userid"` // userid
	//
	TaskID uint32 `orm:"column(taskid)" xorm:"INT(4) notnull unique(userid_task) 'taskid'" json:"taskid"` //任务ID
	Status uint32 `orm:"column(status);default(0)" xorm:"INT(4) default(0) 'status'" json:"status"`       //任务状态0:初始,1:完成,2全部完成领取
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseTask) TableName() string {
	return "base_task"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseTask) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseTask) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "TaskID"},
	}
}

// IcBaseBuyTimes 用户购买次数数据表
type IcBaseBuyTimes struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(a) 'userid'" json:"userid"` // userid
	//
	BuyID uint32 `orm:"column(buy_id)" xorm:"INT notnull unique(a) 'buy_id'" json:"buy_id"`  // id
	Times uint32 `orm:"column(times);default(0)" xorm:"INT default(0) 'times'" json:"times"` // 次数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseBuyTimes) TableName() string {
	return "base_buy_times"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseBuyTimes) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseBuyTimes) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "BuyID"},
	}
}

// IcBaseChest 用户宝箱数据表
type IcBaseChest struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(a) 'userid'" json:"userid"` // userid
	//
	ChestID   int    `orm:"column(chest_id)" xorm:"INT notnull unique(userid_chest) 'chest_id'" json:"chest_id"` // 宝箱id
	Number    uint32 `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"`              // 数量
	OpenTimes uint32 `orm:"column(open_times);default(0)" xorm:"INT default(0) 'open_times'" json:"open_times"`  // 打开次数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *IcBaseChest) TableName() string {
	return "base_chest"
}

// TableEngine 设置引擎为 INNODB
func (u *IcBaseChest) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *IcBaseChest) TableUnique() [][]string {
	return [][]string{
		[]string{"UserID", "ChestID"},
	}
}
