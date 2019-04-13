package models

import (
	"time"
)

// Project 项目
type Project struct {
	Id        int    `orm:"column(id);pk;auto" xorm:"pk autoincr 'id'" json:"id"`
	ProjectID int    `orm:"column(project_id);unique" xorm:"notnull unique 'project_id'" json:"project_id"` //项目id
	Name      string `orm:"column(name)" xorm:"'name'" json:"name"`                                         //项目名称
	API       string `orm:"column(api)" xorm:"'api'" json:"api"`                                            //项目api接口
	Path      string `orm:"column(path)" xorm:"'path'" json:"path"`                                         //项目路径
	Status    int8   `orm:"column(status)" xorm:"'status'" json:"status"`                                   //项目状态1正常,0禁用
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *Project) TableName() string {
	return "system_project"
}

// TableEngine 设置引擎为 INNODB
func (u *Project) TableEngine() string {
	return "INNODB"
}

// GameGlobal 全局数据配置表
type GameGlobal struct {
	ID   int    `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`                  // id
	Name string `orm:"column(name);unique;size(50)" xorm:"varchar(50) notnull unique 'name'" json:"name"` // name
	Info string `orm:"column(info);size(50)" xorm:"varchar(50) null 'info'" json:"info"`                  // info
	Num  int32  `orm:"column(num);default(0)" xorm:"INT default(0) 'num'" json:"num"`                     // 数值
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *GameGlobal) TableName() string {
	return "game_global"
}

// TableEngine 设置引擎为 INNODB
func (u *GameGlobal) TableEngine() string {
	return "INNODB"
}

// GameSource 用户来源定义数据表
type GameSource struct {
	ID       int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	SourceID int `orm:"column(source_id);default(0)" xorm:"INT default(0) 'source_id'" json:"source_id"` // source id >= 0
	//
	Name string `orma:"column(name);unique;size(30);" xorm:"varchar(30) notnull unique 'name'" json:"name"` // source name
	Info string `orma:"column(info);null" xorm:"null 'info'" json:"info"`                                   // source info
	//
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *GameSource) TableName() string {
	return "game_source"
}

// TableEngine 设置引擎为 INNODB
func (u *GameSource) TableEngine() string {
	return "INNODB"
}

// GameUser 针对玩家配置表
type GameUser struct {
	ID     int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	UserID int `orm:"column(userid);unique" xorm:"INT notnull unique(mp) 'userid'" json:"userid"` // userid
	//
	Prize bool `orm:"column(prize);default(0);bool" xorm:"Bool default(0) 'prize'" json:"prize"` //是否开启公众号(白名单)
	//
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *GameUser) TableName() string {
	return "game_user"
}

// TableEngine 设置引擎为 INNODB
func (u *GameUser) TableEngine() string {
	return "INNODB"
}

// GameJson 静态json文件配置表
type GameJson struct {
	ID int `orm:"column(id);pk;auto" xorm:"INT(11) notnull pk autoincr 'id'" json:"id"`
	//
	Type    uint32 `orm:"column(type);size(4);default(0)" xorm:"INT(4) default(0) 'type'" json:"type"`       // 类型:0正常,1审核,2备份
	Name    string `orm:"column(name);size(50);unique" xorm:"varchar(50) notnull unique 'name'" json:"name"` // file name
	Content string `orm:"column(content);type(text);null" xorm:"Text null json 'content'" json:"content"`    // json content
	//
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at"`     // 更新时间
}

// TableName 表名
func (u *GameJson) TableName() string {
	return "game_json"
}

// TableEngine 设置引擎为 INNODB
func (u *GameJson) TableEngine() string {
	return "INNODB"
}

//

//version == 0 关闭分享
//version == 1 开启地域屏蔽检测和白名单检测
//version == 2 全部开启

// SwitchConfig switch config
type SwitchConfig struct {
	Open         bool            `json:"open"`         //是否开启
	Name         string          `json:"name"`         //项目名称
	Province     map[string]bool `json:"province"`     //省份
	City         map[string]bool `json:"city"`         //城市
	ProvinceDate SwitchDate      `json:"provinceDate"` //存在的province对应的时间状态
	CityDate     SwitchDate      `json:"cityDate"`     //存在的city对应的时间状态
	Date         SwitchDate      `json:"date"`         //其它city对应的时间状态
	WhiteList    map[string]bool `json:"whiteList"`    //openid 白名单
	UseridList   map[string]bool `json:"useridList"`   //userid 白名单
	VersionList  map[string]int  `json:"versionList"`  //version 状态
}

// SwitchDate switch date
type SwitchDate struct {
	Open      bool            `json:"open"`      //是否开启
	Week      map[string]bool `json:"week"`      //星期对应状态
	StartTime string          `json:"startTime"` //开始时间
	EndTime   string          `json:"endTime"`   //结束时间
}

// SwitchStatus switch list
type SwitchStatus struct {
	ProjectID int    `json:"project_id"` //项目ID
	Name      string `json:"name"`       //项目名称
	Version   string `json:"version"`    //版本
	Status    bool   `json:"status"`     //是否开启
}

// 文件
type FileInfo struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Time    time.Time `json:"time"`
	Content string    `json:"content"` //内容
	Href    string    `json:"href"`    //地址
	Project string    `json:"project"` //项目名称
}

// GameTask 任务数据配置表
type GameTask struct {
	ID      int    `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`          // id
	TaskID  uint32 `orm:"column(taskid);unique" xorm:"INT(4) notnull unique 'taskid'" json:"taskid"` // 任务ID
	Content string `orm:"column(content);null" xorm:"null 'content'" json:"content"`                 // 内容
	Info    string `orm:"column(info);null" xorm:"null 'info'" json:"info"`                          // 完成条件
	Before  uint32 `orm:"column(before);default(0)" xorm:"INT(4) default(0) 'before'" json:"before"` // 前置id
	After   uint32 `orm:"column(after);default(0)" xorm:"INT(4) default(0) 'after'" json:"after"`    // 后置id
	Prize   uint32 `orm:"column(prize);default(0)" xorm:"INT(4) default(0) 'prize'" json:"prize"`    // 奖励ID
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at,omitempty"`     // 更新时间
}

// TableName 表名
func (u *GameTask) TableName() string {
	return "game_task"
}

// TableEngine 设置引擎为 INNODB
func (u *GameTask) TableEngine() string {
	return "INNODB"
}

// GameAchiement 成就数据配置表
type GameAchiement struct {
	ID        int    `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`                // id
	AchID     uint32 `orm:"column(achid);unique" xorm:"INT(4) notnull unique 'achid'" json:"achid"`          // 成就ID
	Type      int32  `orm:"column(type);default(0)" xorm:"INT(4) default(0) 'type'" json:"type"`             // 成就类型
	Describe  string `orm:"column(describe);null" xorm:"null 'describe'" json:"describe"`                    // 内容
	Condition string `orm:"column(condition);null" xorm:"null 'condition'" json:"condition"`                 // 完成条件
	Before    uint32 `orm:"column(before);default(0)" xorm:"INT(4) default(0) 'before'" json:"before"`       // 前置id
	After     uint32 `orm:"column(after);default(0)" xorm:"INT(4) default(0) 'after'" json:"after"`          // 后置id
	Jewel     uint32 `orm:"column(jewel);default(0)" xorm:"INT(4) default(0) 'jewel'" json:"jewel"`          // 宝石奖励
	Integral  uint32 `orm:"column(integral);default(0)" xorm:"INT(4) default(0) 'integral'" json:"integral"` // 积分奖励
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at,omitempty"`     // 更新时间
}

// TableName 表名
func (u *GameAchiement) TableName() string {
	return "game_achiement"
}

// TableEngine 设置引擎为 INNODB
func (u *GameAchiement) TableEngine() string {
	return "INNODB"
}
