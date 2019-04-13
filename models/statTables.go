package models

import (
	"time"
)

// StatGateRetimes 关卡复活/重新挑战次数统计数据日志
type StatGateRetimes struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	//
	Gate uint32 `orm:"column(gate);default(0)" xorm:"INT(4) default(0) 'gate'" json:"gate"` //关卡
	Type uint32 `orm:"column(type);default(0)" xorm:"INT(4) default(0) 'type'" json:"type"` //类型1:复活,2:重新挑战
	//
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` //
	LogData string `orm:"column(log_data);null;type(text)" xorm:"Text null json 'log_data'" json:"log_data"` //
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//RegData GateMap `xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册数据
	//LogData GateMap `xorm:"Text null json 'log_data'" json:"log_data"` // 登录数据
}

// TableName 表名
func (u *StatGateRetimes) TableName() string {
	return "stat_gate_retimes"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateRetimes) TableEngine() string {
	return "INNODB"
}

// StatGateScore 关卡得分情况统计数据日志
type StatGateScore struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	//
	Gate uint32 `orm:"column(gate);default(0)" xorm:"INT(4) default(0) 'gate'" json:"gate"` //关卡
	Type uint32 `orm:"column(type);default(0)" xorm:"INT(4) default(0) 'type'" json:"type"` //类型1:挑战,2:奖励
	//
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` //
	LogData string `orm:"column(log_data);null;type(text)" xorm:"Text null json 'log_data'" json:"log_data"` //
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//RegData GateMap `xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册数据
	//LogData GateMap `xorm:"Text null json 'log_data'" json:"log_data"` // 登录数据
}

// TableName 表名
func (u *StatGateScore) TableName() string {
	return "stat_gate_score"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateScore) TableEngine() string {
	return "INNODB"
}

// StatGateLoss 新手关卡流失统计数据日志
type StatGateLoss struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Gate      int `orm:"column(gate);default(0)" xorm:"default(0) 'gate'" json:"gate"` //关卡
	Num       int `orm:"column(num);default(0)" xorm:"num" json:"num"`                 //人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatGateLoss) TableName() string {
	return "stat_gate_loss"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateLoss) TableEngine() string {
	return "INNODB"
}

// StatGateWeapon 合成的武器等级统计数据日志
type StatGateWeapon struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	Count int `orm:"column(count);default(0)" xorm:"count" json:"count"`
	Gate  int `orm:"column(gate);default(0)" xorm:"gate" json:"gate"`
	Level int `orm:"column(level);default(0)" xorm:"level" json:"level"`
	Num   int `orm:"column(num);default(0)" xorm:"num" json:"num"` //人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatGateWeapon) TableName() string {
	return "stat_gate_weapon"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateWeapon) TableEngine() string {
	return "INNODB"
}

// StatPetLevel 宠物主角升级分布统计数据日志
type StatPetLevel struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	Regman int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman int `orm:"column(logman)" xorm:"logman" json:"logman"`
	//
	RegRoleData string `orm:"column(reg_role_data);null;type(text)" orm:"Text null json 'reg_role_data'" json:"reg_role_data"` // 数据
	RegPetData  string `orm:"column(reg_pet_data);null;type(text)" orm:"Text null json 'reg_pet_data'" json:"reg_pet_data"`    // 数据
	LogRoleData string `orm:"column(log_role_data);null;type(text)" orm:"Text null json 'log_role_data'" json:"log_role_data"` // 数据
	LogPetData  string `orm:"column(log_pet_data);null;type(text)" orm:"Text null json 'log_pet_data'" json:"log_pet_data"`    // 数据
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatPetLevel) TableName() string {
	return "stat_pet_level"
}

// TableEngine 设置引擎为 INNODB
func (u *StatPetLevel) TableEngine() string {
	return "INNODB"
}

//   - electro
// StatGateGoldDis 关卡分布金币数据统计
type StatGateGoldDis struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	Gate   int    `orm:"column(gate);default(0)" xorm:"INT default(0) 'gate'" json:"gate"`       //关卡
	Regman int    `orm:"column(regman);default(0)" xorm:"INT default(0) 'regman'" json:"regman"` //注册人数
	RAdd   string `orm:"column(radd);null;type(text)" xorm:"Text null json 'radd'" json:"radd"`
	RCost  string `orm:"column(rcost);null;type(text)" xorm:"Text null json 'rcost'" json:"rcost"`
	//
	Logman int    `orm:"column(logman);default(0)" xorm:"INT default(0) 'logman'" json:"logman"` //登录人数
	LAdd   string `orm:"column(ladd);null;type(text)" xorm:"Text null json 'ladd'" json:"ladd"`
	LCost  string `orm:"column(lcost);null;type(text)" xorm:"Text null json 'lcost'" json:"lcost"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	Data    string `orm:"column(data);null;type(text)" xorm:"Text null json 'data'" json:"data"`
	LogData string `orm:"column(logdata);null;type(text)" xorm:"Text null json 'logdata'" json:"logdata"`
}

//// GoldDis 关卡分布金币数据统计
//type GoldDis struct {
//	Number int `json:"number,omitempty"` //人数
//	Times  int `json:"times,omitempty"`  //次数
//	Count  int `json:"count,omitempty"`  //金币数量
//}

// TableName 表名
func (u *StatGateGoldDis) TableName() string {
	return "stat_gate_gold_dis"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateGoldDis) TableEngine() string {
	return "INNODB"
}

//   - electro
// StatGateGoldDep 金币存量数据统计
type StatGateGoldDep struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	Gate int `orm:"column(gate);default(0)" xorm:"INT default(0) 'gate'" json:"gate"` //关卡
	//
	Regman  int `orm:"column(regman);default(0)" xorm:"INT default(0) 'regman'" json:"regman"`    //注册人数
	RNumber int `orm:"column(rnumber);default(0)" xorm:"INT default(0) 'rnumber'" json:"rnumber"` //人数
	RCount  int `orm:"column(rcount);default(0)" xorm:"INT default(0) 'rcount'" json:"rcount"`    //金币数量
	//
	Logman  int `orm:"column(logman);default(0)" xorm:"INT default(0) 'logman'" json:"logman"` //登录人数
	LNumber int `orm:"column(lnumber);default(0)" xorm:"INT default(0) 'lnumber'" json:"lnumber"`
	LCount  int `orm:"column(lcount);default(0)" xorm:"INT default(0) 'lcount'" json:"lcount"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatGateGoldDep) TableName() string {
	return "stat_gate_gold_dep"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateGoldDep) TableEngine() string {
	return "INNODB"
}

//   - IdleCrafting
// StatWeaponLev 武器等级停留统计数据日志
type StatWeaponLev struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	Regman  int    `orm:"column(regman);default(0)" xorm:"INT default(0) 'regman'" json:"regman"`            //注册人数
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` //
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatWeaponLev) TableName() string {
	return "stat_weapon_lev"
}

// TableEngine 设置引擎为 INNODB
func (u *StatWeaponLev) TableEngine() string {
	return "INNODB"
}

// StatProp bomb道具数据统计日志
type StatProp struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	//1:玩家开局数分布,2:分数分布情况
	//3:火箭道具使用关卡分布情况,4:弃牌道具使用关卡分布情况
	//5:火箭道具使用次数情况,6:弃牌道具使用次数情况
	Type uint32 `orm:"column(type)" xorm:"type int(4)" json:"type"`
	//
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册数据
	LogData string `orm:"column(log_data);null;type(text)" xorm:"Text null json 'log_data'" json:"log_data"` // 登录数据
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatProp) TableName() string {
	return "stat_prop"
}

// TableEngine 设置引擎为 INNODB
func (u *StatProp) TableEngine() string {
	return "INNODB"
}
