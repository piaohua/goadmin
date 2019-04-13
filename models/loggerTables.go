package models

import (
	"time"
)

// StatSendClick 发送点击统计数据日志
type StatSendClick struct {
	ID          int `orm:"column(id);pk;auto" json:"id"`
	Timestamp   int `orm:"column(timestamp);" json:"timestamp"`                 // 当天凌晨时间截
	Cover       int `orm:"column(cover);default(0)" json:"cover"`               // 图片id
	Scene       int `orm:"column(scene);default(0)" json:"scene"`               // 场景id
	Version     int `orm:"column(version);default(0)" json:"version"`           // 客户端版本
	ProjectId   int `orm:"column(project_id);default(0)" json:"project_id"`     // 项目id
	SendNumber  int `orm:"column(send_number);default(0)" json:"send_number"`   // 发送人数
	SendTimes   int `orm:"column(send_times);default(0)" json:"send_times"`     // 发送次数
	ClickNumber int `orm:"column(click_number);default(0)" json:"click_number"` // 点击人数
	ClickTimes  int `orm:"column(click_times);default(0)" json:"click_times"`   // 点击次数
	Regist      int `orm:"column(regist);default(0)" json:"regist"`             // 点击注册人数
	Login       int `orm:"column(login);default(0)" json:"login"`               // 点击登录人数
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"` // 创建时间
}

// TableName 表名
func (u *StatSendClick) TableName() string {
	return "stat_send_click"
}

// TableEngine 设置引擎为 INNODB
func (u *StatSendClick) TableEngine() string {
	return "INNODB"
}

// TableUnique 多字段唯一键
func (u *StatSendClick) TableUnique() [][]string {
	return [][]string{
		[]string{"Timestamp", "Cover", "Scene", "Version", "ProjectId"},
	}
}

// SystemChat 消息请求日志
type SystemChat struct {
	ID       int    `orm:"column(id);pk;auto" json:"id"`
	UserId   int32  `orm:"column(user_id)" json:"user_id"`
	UserName string `orm:"column(username)" json:"username"`
	Message  string `orm:"column(message);null" json:"message"`            // 提交数据
	Content  string `orm:"column(content);type(text);null" json:"content"` // 消息结果内容
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *SystemChat) TableName() string {
	return "system_chat"
}

// TableEngine 设置引擎为 INNODB
func (u *SystemChat) TableEngine() string {
	return "INNODB"
}

// log

// LogLogin 用户登录日志
type LogLogin struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index 'userid'" json:"userid"` // userid
	// register
	Scene   int    `orm:"column(scene);default(0)" xorm:"INT(11) default(0) 'scene'" json:"scene"`           // 登录场景id
	Cover   int    `orm:"column(cover);default(0)" xorm:"INT(11) default(0) 'cover'" json:"cover"`           // 登录分享图片id
	Version int    `orm:"column(version);default(0)" xorm:"INT(11) default(0) 'version'" json:"version"`     // 登录客户端版本
	LoginIP string `orm:"column(login_ip);size(15);null" xorm:"varchar(15) null 'login_ip'" json:"login_ip"` // 登录ip
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogLogin) TableName() string {
	return "log_login"
}

// TableEngine 设置引擎为 INNODB
func (u *LogLogin) TableEngine() string {
	return "INNODB"
}

// LogRegister 用户注册日志
type LogRegister struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index 'userid'" json:"userid"` // userid
	// register
	Scene    int    `orm:"column(scene);default(0)" xorm:"INT(11) default(0) 'scene'" json:"scene"`              // 场景id
	Cover    int    `orm:"column(cover);default(0)" xorm:"INT(11) default(0) 'cover'" json:"cover"`              // 分享图片id
	Version  int    `orm:"column(version);default(0)" xorm:"INT(11) default(0) 'version'" json:"version"`        // 客户端版本
	RegistIP string `orm:"column(regist_ip);size(15);null" xorm:"varchar(15) null 'regist_ip'" json:"regist_ip"` // 注册ip
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogRegister) TableName() string {
	return "log_register"
}

// TableEngine 设置引擎为 INNODB
func (u *LogRegister) TableEngine() string {
	return "INNODB"
}

// LogSendClick 发送和点击数据日志
type LogSendClick struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(send_click) 'userid'" json:"userid"` // userid
	//
	Timestamp int `orm:"column(timestamp);default(0)" xorm:"INT(11) notnull unique(send_click) 'timestamp'" json:"timestamp"` // 当天凌晨时间截
	Cover     int `orm:"column(cover);default(0)" xorm:"INT(11) default(0) unique(send_click) 'cover'" json:"cover"`          // 图片id
	Scene     int `orm:"column(scene);default(0)" xorm:"INT(11) default(0) unique(send_click) 'scene'" json:"scene"`          // 场景id
	Version   int `orm:"column(version);default(0)" xorm:"INT(11) default(0) unique(send_click) 'version'" json:"version"`    // 客户端版本
	Type      int `orm:"column(type);default(0)" xorm:"TINYINT default(0) unique(send_click) 'type'" json:"type"`             // 类型1:发送,2:点击
	Number    int `orm:"column(number);default(0)" xorm:"INT(11) default(0) 'number'" json:"number"`                          // 发送/点击次数，人数group统计得出
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *LogSendClick) TableName() string {
	return "log_send_click"
}

// TableEngine 设置引擎为 INNODB
func (u *LogSendClick) TableEngine() string {
	return "INNODB"
}

// LogProgress 点击进入转化数据日志
type LogProgress struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index(rds_idx_0) 'userid'" json:"userid"` // userid
	//
	Today int `orm:"column(today)" xorm:"INT notnull index 'today'" json:"today"`         // 当天时间yyyymmdd
	Type  int `orm:"column(type)" xorm:"INT notnull index(rds_idx_0) 'type'" json:"type"` // type,记录类型
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogProgress) TableName() string {
	return "log_progress"
}

// TableEngine 设置引擎为 INNODB
func (u *LogProgress) TableEngine() string {
	return "INNODB"
}

// LogSceneAct 各场景分享
type LogSceneAct struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index 'userid'" json:"userid"` // userid
	//
	Today int `orm:"column(today)" xorm:"INT notnull index 'today'" json:"today"`         // 当天时间yyyymmdd
	Scene int `orm:"column(scene);default(0)" xorm:"INT default(0) 'scene'" json:"scene"` // 场景id
	Type  int `orm:"column(type);default(0)" xorm:"INT default(0) 'type'" json:"type"`    // 操作类型
	//
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT(4) default(0) 'source_id'" json:"source_id"` // source id
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogSceneAct) TableName() string {
	return "log_scene_act"
}

// TableEngine 设置引擎为 INNODB
func (u *LogSceneAct) TableEngine() string {
	return "INNODB"
}

// LogPlaytime 在线时长
type LogPlaytime struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(userid_playtime) 'userid'" json:"userid"` // userid
	//
	Today    int    `orm:"column(today)" xorm:"INT notnull unique(userid_playtime) 'today'" json:"today"` // 当天时间yyyymmdd
	Playtime uint32 `orm:"column(playtime)" xorm:"INT(4) notnull 'playtime'" json:"playtime"`             // playtime (秒)
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *LogPlaytime) TableName() string {
	return "log_playtime"
}

// TableEngine 设置引擎为 INNODB
func (u *LogPlaytime) TableEngine() string {
	return "INNODB"
}

// LogTip 关卡提示次数记录
type LogTip struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index 'userid'" json:"userid"` // userid
	//
	Today int    `orm:"column(today)" xorm:"INT notnull index 'today'" json:"today"` // 当天时间yyyymmdd
	Gate  uint32 `orm:"column(gate)" xorm:"INT notnull 'gate'" json:"gate"`          // 关卡id
	Type  uint32 `orm:"column(type)" xorm:"INT notnull 'type'" json:"type"`          // 操作类型 1:金币提示,2:免费提示
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogTip) TableName() string {
	return "log_tip"
}

// TableEngine 设置引擎为 INNODB
func (u *LogTip) TableEngine() string {
	return "INNODB"
}

// LogWall 流量墙次数记录
type LogWall struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(wall) 'userid'" json:"userid"` // userid
	//
	Today  int    `orm:"column(today)" xorm:"INT notnull unique(wall) 'today'" json:"today"`                  // 当天时间yyyymmdd
	AppID  string `orm:"column(appid);size(32)" xorm:"varchar(32) notnull unique(wall) 'appid'" json:"appid"` // appid
	Type   uint32 `orm:"column(type);default(0)" xorm:"TINYINT default(0) unique(wall) 'type'" json:"type"`   // 类型0:流量墙点击,1:领取奖励,2:抽屉点击,3:浮窗点击,4:猜你喜欢
	Number uint32 `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"`              // 变更数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *LogWall) TableName() string {
	return "log_wall"
}

// TableEngine 设置引擎为 INNODB
func (u *LogWall) TableEngine() string {
	return "INNODB"
}

// LogGateGrow 用户前进关卡日志
type LogGateGrow struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull unique(userid_gate) 'userid'" json:"userid"` // userid
	//
	Today  int    `orm:"column(today)" xorm:"INT notnull unique(userid_gate) 'today'" json:"today"` // 当天时间yyyymmdd
	Number uint32 `orm:"column(number)" xorm:"INT default(0) 'number'" json:"number"`               // 变更数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *LogGateGrow) TableName() string {
	return "log_gate_grow"
}

// TableEngine 设置引擎为 INNODB
func (u *LogGateGrow) TableEngine() string {
	return "INNODB"
}

// LogGold 用户金币日志
type LogGold struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index 'userid'" json:"userid"` // userid
	// register
	Type   int32  `orm:"column(type);default(0)" xorm:"INT default(0) 'type'" json:"type"`           // 变更类型
	Number int32  `orm:"column(number);default(0)" xorm:"INT default(0) 'number'" json:"number"`     // 变更数量
	Rest   uint32 `orm:"column(rest);default(0)" xorm:"INT default(0) 'rest'" json:"rest"`           // 剩余数量
	Gate   uint32 `orm:"column(gate);default(0)" xorm:"INT(11) index default(0) 'gate'" json:"gate"` // 关卡id
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogGold) TableName() string {
	return "log_gold"
}

// TableEngine 设置引擎为 INNODB
func (u *LogGold) TableEngine() string {
	return "INNODB"
}

// LogGate 用户关卡日志
type LogGate struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid)" xorm:"INT notnull index 'userid'" json:"userid"` // userid
	// register
	Gate    uint32 `orm:"column(gate);default(0)" xorm:"INT(11) index default(0) 'gate'" json:"gate"`       //关卡id
	Step    uint32 `orm:"column(step);default(0)" xorm:"INT(4) default(0) 'step'" json:"step"`              //步数
	UseTime uint32 `orm:"column(use_time);default(0)" xorm:"INT(11) default(0) 'use_time'" json:"use_time"` //关卡用时(毫秒ms)
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
}

// TableName 表名
func (u *LogGate) TableName() string {
	return "log_gate"
}

// TableEngine 设置引擎为 INNODB
func (u *LogGate) TableEngine() string {
	return "INNODB"
}
