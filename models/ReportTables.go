package models

import "time"

// StatActive 活跃留存数据
type StatActive struct {
	Id        int `orm:"column(id);pk;auto" xorm:"pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	SourceID  int `orm:"column(source_id)" xorm:"INT default(0) 'source_id'" json:"source_id"` // 来源
	Dau       int `orm:"column(dau)" xorm:"dau" json:"dau"`
	RegMan    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	LogMan    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	BackMan   int `orm:"column(backman)" xorm:"backman" json:"backman"`
	Left2     int `orm:"column(left2)" xorm:"left2" json:"left2"`
	Left3     int `orm:"column(left3)" xorm:"left3" json:"left3"`
	Left4     int `orm:"column(left4)" xorm:"left4" json:"left4"`
	Left5     int `orm:"column(left5)" xorm:"left5" json:"left5"`
	Left6     int `orm:"column(left6)" xorm:"left6" json:"left6"`
	Left7     int `orm:"column(left7)" xorm:"left7" json:"left7"`
	Left10    int `orm:"column(left10)" xorm:"left10" json:"left10"`
	Left15    int `orm:"column(left15)" xorm:"left15" json:"left15"`
	Log2      int `orm:"column(log2)" xorm:"log2" json:"log2"`
	Log3      int `orm:"column(log3)" xorm:"log3" json:"log3"`
	Log4      int `orm:"column(log4)" xorm:"log4" json:"log4"`
	Log5      int `orm:"column(log5)" xorm:"log5" json:"log5"`
	Log6      int `orm:"column(log6)" xorm:"log6" json:"log6"`
	Log7      int `orm:"column(log7)" xorm:"log7" json:"log7"`
	Log10     int `orm:"column(log10)" xorm:"log10" json:"log10"`
	Log15     int `orm:"column(log15)" xorm:"log15" json:"log15"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatActive) TableName() string {
	return "stat_active"
}

// TableEngine 设置引擎为 INNODB
func (u *StatActive) TableEngine() string {
	return "INNODB"
}

// StatPlaytime 在线时长数据
type StatPlaytime struct {
	Id        int `orm:"column(id);pk;auto" xorm:"pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime0    int `orm:"column(rtime0)" xorm:"rtime0" json:"rtime0"`
	Rtime1    int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"`
	Rtime5    int `orm:"column(rtime5)" xorm:"rtime5" json:"rtime5"`
	Rtime10   int `orm:"column(rtime10)" xorm:"rtime10" json:"rtime10"`
	Rtime15   int `orm:"column(rtime15)" xorm:"rtime15" json:"rtime15"`
	Rtime30   int `orm:"column(rtime30)" xorm:"rtime30" json:"rtime30"`
	Rtime60   int `orm:"column(rtime60)" xorm:"rtime60" json:"rtime60"`
	Ravgtime  int `orm:"column(ravgtime)" xorm:"ravgtime" json:"ravgtime"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime0    int `orm:"column(ltime0)" xorm:"ltime0" json:"ltime0"`
	Ltime1    int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"`
	Ltime5    int `orm:"column(ltime5)" xorm:"ltime5" json:"ltime5"`
	Ltime10   int `orm:"column(ltime10)" xorm:"ltime10" json:"ltime10"`
	Ltime15   int `orm:"column(ltime15)" xorm:"ltime15" json:"ltime15"`
	Ltime30   int `orm:"column(ltime30)" xorm:"ltime30" json:"ltime30"`
	Ltime60   int `orm:"column(ltime60)" xorm:"ltime60" json:"ltime60"`
	Lravgtime int `orm:"column(lravgtime)" xorm:"lravgtime" json:"lravgtime"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatPlaytime) TableName() string {
	return "stat_playtime"
}

// TableEngine 设置引擎为 INNODB
func (u *StatPlaytime) TableEngine() string {
	return "INNODB"
}

// StatStart 启动次数数据
type StatStart struct {
	Id        int `orm:"column(id);pk;auto" xorm:"pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1    int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"`
	Rtime2    int `orm:"column(rtime2)" xorm:"rtime2" json:"rtime2"`
	Rtime3    int `orm:"column(rtime3)" xorm:"rtime3" json:"rtime3"`
	Rtime5    int `orm:"column(rtime5)" xorm:"rtime5" json:"rtime5"`
	Rtime8    int `orm:"column(rtime8)" xorm:"rtime8" json:"rtime8"`
	Rtime10   int `orm:"column(rtime10)" xorm:"rtime10" json:"rtime10"`
	Rtime15   int `orm:"column(rtime15)" xorm:"rtime15" json:"rtime15"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1    int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"`
	Ltime2    int `orm:"column(ltime2)" xorm:"ltime2" json:"ltime2"`
	Ltime3    int `orm:"column(ltime3)" xorm:"ltime3" json:"ltime3"`
	Ltime5    int `orm:"column(ltime5)" xorm:"ltime5" json:"ltime5"`
	Ltime8    int `orm:"column(ltime8)" xorm:"ltime8" json:"ltime8"`
	Ltime10   int `orm:"column(ltime10)" xorm:"ltime10" json:"ltime10"`
	Ltime15   int `orm:"column(ltime15)" xorm:"ltime15" json:"ltime15"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatStart) TableName() string {
	return "stat_start"
}

// TableEngine 设置引擎为 INNODB
func (u *StatStart) TableEngine() string {
	return "INNODB"
}

// StatStart2 启动次数数据
type StatStart2 struct {
	Id        int `orm:"column(id);pk;auto" xorm:"pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
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
func (u *StatStart2) TableName() string {
	return "stat_start2"
}

// TableEngine 设置引擎为 INNODB
func (u *StatStart2) TableEngine() string {
	return "INNODB"
}

//Type:记录类型
//1:loading节点1完成人数
//2:loading节点2完成人数
//3:进入新手教程人数
//4:完成新手教程节点1人数
//5:完成新手教程节点2人数
//6:完成新手教程人数
//7:达到新增点
// StatProgress 点击进入转化统计数据日志
type StatProgress struct {
	ID        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull unique(stat_progress) 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today)" xorm:"INT notnull unique(stat_progress) 'today'" json:"today"`                // 当天时间yyyymmdd
	//
	Data string `orm:"column(data);null;type(text)" xorm:"Text null 'data'" json:"data"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//Data Progress `xorm:"Text null json 'data'" json:"data"`
}

// TableName 表名
func (u *StatProgress) TableName() string {
	return "stat_progress"
}

// TableEngine 设置引擎为 INNODB
func (u *StatProgress) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *StatProgress) TableUnique() [][]string {
	return [][]string{
		[]string{"ProjectId", "Today"},
	}
}

//Type:记录类型             - 粉碎水果
//1:loading节点1完成人数   1  新手第1步
//2:loading节点2完成人数   2  新手第2步
//3:进入新手教程人数       3  新手第3步
//4:完成新手教程节点1人数  4  新手第4步
//5:完成新手教程节点2人数  5  配置开始加载
//6:完成新手教程人数       6  配置加载完成
//7:达到新增点             7  进入游戏界面
// Progress 步骤分布
type Progress struct {
	Type1 int `json:"type1,omitempty"` // type1
	Type2 int `json:"type2,omitempty"` // type2
	Type3 int `json:"type3,omitempty"` // type3
	Type4 int `json:"type4,omitempty"` // type4
	Type5 int `json:"type5,omitempty"` // type5
	Type6 int `json:"type6,omitempty"` // type6
	Type7 int `json:"type7,omitempty"` // type7
}

// StatGateData 关卡数据分布统计日志
type StatGateData struct {
	ID        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull unique(stat_progress) 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today)" xorm:"INT notnull unique(stat_progress) 'today'" json:"today"`                // 当天时间yyyymmdd
	// 注册 类型1:通关时间分布,2:使用步数分布,3:金币提示次数分布,4:免费提示次数分布
	// 登录 类型10：通关时间分布,20:使用步数分布,30:金币提示次数分布,40:免费提示次数分布
	Type   int `orm:"column(type);default(0)" xorm:"INT default(0) 'type'" json:"type"`
	Number int `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"` // 注册人数，登录人数
	Gate   int `orm:"column(gate);default(0)" xorm:"INT default(0) 'gate'" json:"gate"`       // 关卡
	//
	Data string `orm:"column(data);null;type(text)" xorm:"Text null 'data'" json:"data"` // 分布统计数据
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//Data GateData `xorm:"Text null json 'data'" json:"data"` // 分布统计数据
}

// TableName 表名
func (u *StatGateData) TableName() string {
	return "stat_gate_data"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateData) TableEngine() string {
	return "INNODB"
}

/*
type   通关时间分布:    使用步数分布:    金币提示分布:   免费提示分布:
type1  少于10秒         少于10步         少于2次         少于2次
type2  10~20            10~15            2~4             2~4
type3  20~40            15~20            4~6             4~6
type4  40~60            20~30            6~8             6~8
type5  60~9999          30~999           8~99            8~99
*/
// GateData 关卡数据分布
type GateData struct {
	Type1 int `json:"type1,omitempty"` // type1
	Type2 int `json:"type2,omitempty"` // type2
	Type3 int `json:"type3,omitempty"` // type3
	Type4 int `json:"type4,omitempty"` // type4
	Type5 int `json:"type5,omitempty"` // type5
	Type6 int `json:"type6,omitempty"` // type6 //通关人数,提示人数
}

//1:分享触发,2:分享点击,3:分享进入
// StatScene 场景操作类型统计数据日志
type StatScene struct {
	ID        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull unique(stat_scene) 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today)" xorm:"INT notnull unique(stat_scene) 'today'" json:"today"`                // 时间yyyymmdd
	Scene     int `orm:"column(scene)" xorm:"INT notnull unique(stat_scene) 'scene'" json:"scene"`                // 场景id
	Type      int `orm:"column(type)" xorm:"INT notnull unique(stat_scene) 'type'" json:"type"`                   // 操作类型
	SourceID  int `orm:"column(source_id)" xorm:"INT default(0) 'source_id'" json:"source_id"`                    // 来源
	Number    int `orm:"column(number)" xorm:"INT default(0) 'number'" json:"number"`                             // 人数
	Times     int `orm:"column(times)" xorm:"INT default(0) 'times'" json:"times"`                                // 次数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	////
	//Number1 int `orm:"-" xorm:"-" json:"number1"` // 人数
	//Times1  int `orm:"-" xorm:"-" json:"times1"`  // 次数
	//Number2 int `orm:"-" xorm:"-" json:"number2"` // 人数
	//Times2  int `orm:"-" xorm:"-" json:"times2"`  // 次数
	//Number3 int `orm:"-" xorm:"-" json:"number3"` // 人数
	//Times3  int `orm:"-" xorm:"-" json:"times3"`  // 次数
}

// TableName 表名
func (u *StatScene) TableName() string {
	return "stat_scene"
}

// TableEngine 设置引擎为 INNODB
func (u *StatScene) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *StatScene) TableUnique() [][]string {
	return [][]string{
		[]string{"ProjectId", "Today", "Scene", "Type"},
	}
}

// StatGate 关卡分布统计数据日志
type StatGate struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1    int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"`    //1
	Rtime2    int `orm:"column(rtime2)" xorm:"rtime2" json:"rtime2"`    //2~5
	Rtime3    int `orm:"column(rtime3)" xorm:"rtime3" json:"rtime3"`    //6~10
	Rtime4    int `orm:"column(rtime4)" xorm:"rtime4" json:"rtime4"`    //11~30
	Rtime5    int `orm:"column(rtime5)" xorm:"rtime5" json:"rtime5"`    //31~50
	Rtime6    int `orm:"column(rtime6)" xorm:"rtime6" json:"rtime6"`    //51~
	Rtime7    int `orm:"column(rtime7)" xorm:"rtime7" json:"rtime7"`    //71~100
	Rtime8    int `orm:"column(rtime8)" xorm:"rtime8" json:"rtime8"`    //101~150
	Rtime9    int `orm:"column(rtime9)" xorm:"rtime9" json:"rtime9"`    //151~200
	Rtime10   int `orm:"column(rtime10)" xorm:"rtime10" json:"rtime10"` //201~
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1    int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"` //1
	Ltime2    int `orm:"column(ltime2)" xorm:"ltime2" json:"ltime2"` //2~10
	Ltime3    int `orm:"column(ltime3)" xorm:"ltime3" json:"ltime3"` //11~30
	Ltime4    int `orm:"column(ltime4)" xorm:"ltime4" json:"ltime4"` //31~50
	Ltime5    int `orm:"column(ltime5)" xorm:"ltime5" json:"ltime5"` //51~80
	Ltime6    int `orm:"column(ltime6)" xorm:"ltime6" json:"ltime6"` //81~
	Ltime7    int `orm:"column(ltime7)" xorm:"ltime7" json:"ltime7"` //101~150
	Ltime8    int `orm:"column(ltime8)" xorm:"ltime8" json:"ltime8"` //151~200
	Ltime9    int `orm:"column(ltime9)" xorm:"ltime9" json:"ltime9"` //201~
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatGate) TableName() string {
	return "stat_gate"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGate) TableEngine() string {
	return "INNODB"
}

// StatGate2 关卡分布统计数据日志
type StatGate2 struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	//
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册玩家关卡分布统计数据
	LogData string `orm:"column(log_data);null;type(text)" xorm:"Text null json 'log_data'" json:"log_data"` // 登录玩家关卡分布统计数据
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//RegData GateMap `xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册数据
	//LogData GateMap `xorm:"Text null json 'log_data'" json:"log_data"` // 登录数据
}

// TableName 表名
func (u *StatGate2) TableName() string {
	return "stat_gate2"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGate2) TableEngine() string {
	return "INNODB"
}

// StatPet 宠物解锁分布统计数据日志
type StatPet struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	//
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册玩家关卡分布统计数据
	LogData string `orm:"column(log_data);null;type(text)" xorm:"Text null json 'log_data'" json:"log_data"` // 登录玩家关卡分布统计数据
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//RegData GateMap `xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册数据
	//LogData GateMap `xorm:"Text null json 'log_data'" json:"log_data"` // 登录数据
}

// TableName 表名
func (u *StatPet) TableName() string {
	return "stat_pet"
}

// TableEngine 设置引擎为 INNODB
func (u *StatPet) TableEngine() string {
	return "INNODB"
}

// StatUseSkill 技能使用统计数据日志
type StatUseSkill struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	SkillID   int `orm:"column(skill_id)" xorm:"'skill_id'" json:"skill_id"` // 技能id
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1    int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"` //1 次数
	Rnum1     int `orm:"column(rnum1)" xorm:"rnum1" json:"rnum1"`    //1 人数
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1    int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"` //1 次数
	Lnum1     int `orm:"column(lnum1)" xorm:"lnum1" json:"lnum1"`    //1 人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatUseSkill) TableName() string {
	return "stat_use_skill"
}

// TableEngine 设置引擎为 INNODB
func (u *StatUseSkill) TableEngine() string {
	return "INNODB"
}

// StatGateGrow 前进关卡(关卡增长)分布统计数据日志
type StatGateGrow struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	//
	Regman int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman int `orm:"column(logman)" xorm:"logman" json:"logman"`
	//
	RegData string `orm:"column(reg_data);null;type(text)" xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册玩家关卡分布统计数据
	LogData string `orm:"column(log_data);null;type(text)" xorm:"Text null json 'log_data'" json:"log_data"` // 登录玩家关卡分布统计数据
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	//RegData GateGrow `xorm:"Text null json 'reg_data'" json:"reg_data"` // 注册数据
	//LogData GateGrow `xorm:"Text null json 'log_data'" json:"log_data"` // 登录数据
}

// TableName 表名
func (u *StatGateGrow) TableName() string {
	return "stat_gate_grow"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateGrow) TableEngine() string {
	return "INNODB"
}

//   - 七巧板
//  0  0
//  1  1~2
//  2  3~6
//  3  7~10
//  4  11~15
//  5  16~25
//  6  26~45
//  7  46~70
//  8  71~
// GateGrow 前进关卡分布
type GateGrow struct {
	Type0 int `json:"type0,omitempty"` // type0
	Type1 int `json:"type1,omitempty"` // type1
	Type2 int `json:"type2,omitempty"` // type2
	Type3 int `json:"type3,omitempty"` // type3
	Type4 int `json:"type4,omitempty"` // type4
	Type5 int `json:"type5,omitempty"` // type5
	Type6 int `json:"type6,omitempty"` // type6
	Type7 int `json:"type7,omitempty"` // type7
	Type8 int `json:"type8,omitempty"` // type8
}

// StatGateStop 关卡停留分布统计数据日志
type StatGateStop struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1    int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"`
	Rtime2    int `orm:"column(rtime2)" xorm:"rtime2" json:"rtime2"`
	Rtime3    int `orm:"column(rtime3)" xorm:"rtime3" json:"rtime3"`
	Rtime4    int `orm:"column(rtime4)" xorm:"rtime4" json:"rtime4"`
	Rtime5    int `orm:"column(rtime5)" xorm:"rtime5" json:"rtime5"`
	Rtime6    int `orm:"column(rtime6)" xorm:"rtime6" json:"rtime6"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1    int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"`
	Ltime2    int `orm:"column(ltime2)" xorm:"ltime2" json:"ltime2"`
	Ltime3    int `orm:"column(ltime3)" xorm:"ltime3" json:"ltime3"`
	Ltime4    int `orm:"column(ltime4)" xorm:"ltime4" json:"ltime4"`
	Ltime5    int `orm:"column(ltime5)" xorm:"ltime5" json:"ltime5"`
	Ltime6    int `orm:"column(ltime6)" xorm:"ltime6" json:"ltime6"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatGateStop) TableName() string {
	return "stat_gate_stop"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateStop) TableEngine() string {
	return "INNODB"
}

// StatGateStop2 关卡停留分布统计数据日志
type StatGateStop2 struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	Today     int `orm:"column(today);index" xorm:"notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Regman    int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Logman    int `orm:"column(logman)" xorm:"logman" json:"logman"`
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
func (u *StatGateStop2) TableName() string {
	return "stat_gate_stop2"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateStop2) TableEngine() string {
	return "INNODB"
}

// StatOnline 在线实时统计数据日志
type StatOnline struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"project_id" json:"project_id"`
	Timestamp int `orm:"column(timestamp);" xorm:"INT notnull 'timestamp'" json:"timestamp"`     // 时间截
	Regist    int `orm:"column(regist);default(0)" xorm:"INT default(0) 'regist'" json:"regist"` // 注册人数
	Dau       int `orm:"column(dau);default(0)" xorm:"INT default(0) 'dau'" json:"dau"`          // dau
	Login     int `orm:"column(login);default(0)" xorm:"INT default(0) 'login'" json:"login"`    // 登录人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	//
	Type int `orm:"column(type);default(0)" xorm:"INT default(0) 'type'" json:"type"` // 0:在线实时人数, 1:实时在玩人数
}

// TableName 表名
func (u *StatOnline) TableName() string {
	return "stat_online"
}

// TableEngine 设置引擎为 INNODB
func (u *StatOnline) TableEngine() string {
	return "INNODB"
}

// StatShare 场景操作类型统计数据日志
type StatShare struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull unique(a) 'project_id'" json:"project_id"`                       // 项目id
	Today     int `orm:"column(today)" xorm:"INT notnull unique(a) 'today'" json:"today"`                                      // 时间yyyymmdd
	SourceID  int `orm:"column(source_id);default(0);size(4)" xorm:"INT(4) default(0) unique(a) 'source_id'" json:"source_id"` // 来源
	//
	Regist int `orm:"column(regist);default(0)" xorm:"INT default(0) 'regist'" json:"regist"` // 注册人数
	Dau    int `orm:"column(dau);default(0)" xorm:"INT default(0) 'dau'" json:"dau"`          // dau
	Login  int `orm:"column(login);default(0)" xorm:"INT default(0) 'login'" json:"login"`    // 登录人数
	Number int `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"` // 分享人数
	Times  int `orm:"column(times);default(0)" xorm:"INT default(0) 'times'" json:"times"`    // 分享次数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatShare) TableName() string {
	return "stat_share"
}

// TableEngine 设置引擎为 INNODB
func (u *StatShare) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *StatShare) TableUnique() [][]string {
	return [][]string{
		[]string{"ProjectId", "Today", "SourceID"},
	}
}

// TodayData 今日实时数据
type TodayData2 struct {
	SourceId int `json:"source_id"` // 来源
	Regist   int `json:"regist"`    // 注册人数
	Dau      int `json:"dau"`       // dau
	Login    int `json:"login"`     // 登录人数
}

// TodayData 今日实时数据
type TodayData struct {
	Count    int `json:"count"`     // 人数
	SourceId int `json:"source_id"` // 来源
}

//   - 七巧板
// StatGateGold 关卡分布金币数据统计
type StatGateGold struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today);index" xorm:"INT notnull index 'today'" json:"today"`    // 时间yyyymmdd
	//类型0:金币产出、消耗,1:金币存储量/转盘获得金币数
	Type   int `orm:"column(type);size(4);index" xorm:"INT(4) index default(0) 'type'" json:"type"`
	Regman int `orm:"column(regman);default(0)" xorm:"regman" json:"regman"`
	Gate   int `orm:"column(gate);default(0)" xorm:"INT default(0) 'gate'" json:"gate"` //关卡(5-100)
	//玩家获得金币平均数/当前拥有金币数量的平均值
	Num1 int `orm:"column(num1);default(0)" xorm:"INT default(0) 'num1'" json:"num1"`
	//玩家消耗金币平均数/转盘获得金币数量的平均值
	Num2   int `orm:"column(num2);default(0)" xorm:"INT default(0) 'num2'" json:"num2"`
	Logman int `orm:"column(logman);default(0)" xorm:"logman" json:"logman"`
	//玩家获得金币平均数/当前拥有金币数量的平均值
	Num3 int `orm:"column(num3);default(0)" xorm:"INT default(0) 'num3'" json:"num3"`
	//玩家消耗金币平均数/转盘获得金币数量的平均值
	Num4 int `orm:"column(num4);default(0)" xorm:"INT default(0) 'num4'" json:"num4"`
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatGateGold) TableName() string {
	return "stat_gate_gold"
}

// TableEngine 设置引擎为 INNODB
func (u *StatGateGold) TableEngine() string {
	return "INNODB"
}

// StatDraw 启动次数统计数据日志
type StatDraw struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today);index" xorm:"INT notnull index 'today'" json:"today"`    // 时间yyyymmdd
	//
	Regman int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1 int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"` //炫耀次数
	Rtime2 int `orm:"column(rtime2)" xorm:"rtime2" json:"rtime2"` //解锁次数
	Logman int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1 int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"` //炫耀次数
	Ltime2 int `orm:"column(ltime2)" xorm:"ltime2" json:"ltime2"` //解锁次数
	//
	Rnum1 int `orm:"column(rnum1)" xorm:"rnum1" json:"rnum1"` //炫耀人数
	Rnum2 int `orm:"column(rnum2)" xorm:"rnum2" json:"rnum2"` //解锁人数
	Lnum1 int `orm:"column(lnum1)" xorm:"lnum1" json:"lnum1"` //炫耀人数
	Lnum2 int `orm:"column(lnum2)" xorm:"lnum2" json:"lnum2"` //解锁人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatDraw) TableName() string {
	return "stat_draw"
}

// TableEngine 设置引擎为 INNODB
func (u *StatDraw) TableEngine() string {
	return "INNODB"
}

// StatSign sign次数统计数据日志
type StatSign struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today);index" xorm:"INT notnull index 'today'" json:"today"`    // 时间yyyymmdd
	//
	Regman int `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1 int `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"` //次数
	Logman int `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1 int `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"` //次数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatSign) TableName() string {
	return "stat_sign"
}

// TableEngine 设置引擎为 INNODB
func (u *StatSign) TableEngine() string {
	return "INNODB"
}

// StatWall 流量墙数据统计日志
type StatWall struct {
	Id        int `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`
	ProjectId int `orm:"column(project_id)" xorm:"INT notnull 'project_id'" json:"project_id"` // 项目id
	Today     int `orm:"column(today);index" xorm:"INT notnull index 'today'" json:"today"`    // 时间yyyymmdd
	//
	AppID  string `orm:"column(appid);size(50)" xorm:"appid varchar(50)" json:"appid"`
	Type   uint32 `orm:"column(type);size(4)" xorm:"type int(4)" json:"type"`
	Regman int    `orm:"column(regman)" xorm:"regman" json:"regman"`
	Rtime1 int    `orm:"column(rtime1)" xorm:"rtime1" json:"rtime1"` //次数
	Rnum1  int    `orm:"column(rnum1)" xorm:"rnum1" json:"rnum1"`    //人数
	Logman int    `orm:"column(logman)" xorm:"logman" json:"logman"`
	Ltime1 int    `orm:"column(ltime1)" xorm:"ltime1" json:"ltime1"` //次数
	Lnum1  int    `orm:"column(lnum1)" xorm:"lnum1" json:"lnum1"`    //人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *StatWall) TableName() string {
	return "stat_wall"
}

// TableEngine 设置引擎为 INNODB
func (u *StatWall) TableEngine() string {
	return "INNODB"
}

// StatFast 极速模式数据统计日志
type StatFast struct {
	Id        int `xorm:"pk autoincr 'id'" json:"id"`
	Today     int `xorm:"INT notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `xorm:"project_id" json:"project_id"`
	//
	Data string `orm:"column(data);null;type(text)" xorm:"Text null json 'data'" json:"data"` //
	// time
	CreatedAt time.Time `xorm:"created" json:"created_at"` // 创建时间
	//
	//Data FastMap `xorm:"Text null json 'data'" json:"data"` // 注册数据
}

// TableName 表名
func (u *StatFast) TableName() string {
	return "stat_fast"
}

// TableEngine 设置引擎为 INNODB
func (u *StatFast) TableEngine() string {
	return "INNODB"
}

// StatAstrology 占星运势数据统计日志
type StatAstrology struct {
	Id        int `xorm:"pk autoincr 'id'" json:"id"`
	Today     int `xorm:"INT notnull index 'today'" json:"today"` // 时间yyyymmdd
	ProjectId int `xorm:"project_id" json:"project_id"`
	//
	Data string `orm:"column(data);null;type(text)" xorm:"Text null json 'data'" json:"data"` //
	// time
	CreatedAt time.Time `xorm:"created" json:"created_at"` // 创建时间
	//
	//Data FastMap `xorm:"Text null json 'data'" json:"data"` // 注册数据
}

// TableName 表名
func (u *StatAstrology) TableName() string {
	return "stat_astrology"
}

// TableEngine 设置引擎为 INNODB
func (u *StatAstrology) TableEngine() string {
	return "INNODB"
}
