package models

import (
	"time"
)

// GameFuDai 福袋数据配置表
type GameFuDai struct {
	ID  int    `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`         // id
	Bag uint32 `orm:"column(bag);unique;size(4)" xorm:"INT(4) notnull unique 'bag'" json:"bag"` // 福袋ID
	Num uint32 `orm:"column(num);default(0);size(4)" xorm:"INT(4) default(0) 'num'" json:"num"` // 福券数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at,omitempty"`     // 更新时间
}

// TableName 表名
func (u *GameFuDai) TableName() string {
	return "game_fu_dai"
}

// TableEngine 设置引擎为 INNODB
func (u *GameFuDai) TableEngine() string {
	return "INNODB"
}

// GameFuGate 福袋出现数据配置表
type GameFuGate struct {
	ID  int    `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`         // id
	Bag uint32 `orm:"column(bag);unique;size(4)" xorm:"INT(4) notnull unique 'bag'" json:"bag"` // 福袋ID
	Num uint32 `orm:"column(num);default(0);size(4)" xorm:"INT(4) default(0) 'num'" json:"num"` // 通关数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at,omitempty"`     // 更新时间
}

// TableName 表名
func (u *GameFuGate) TableName() string {
	return "game_fu_gate"
}

// TableEngine 设置引擎为 INNODB
func (u *GameFuGate) TableEngine() string {
	return "INNODB"
}

// GameFuDay 福袋每天获得数量配置表
type GameFuDay struct {
	ID  int    `orm:"column(id);pk;auto" xorm:"INT notnull pk autoincr 'id'" json:"id"`         // id
	Day uint32 `orm:"column(day);unique;size(4)" xorm:"INT(4) notnull unique 'day'" json:"day"` // day 1-15
	Num uint32 `orm:"column(num);default(0);size(4)" xorm:"INT(4) default(0) 'num'" json:"num"` // 福袋数量
	// time
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" xorm:"created" json:"created_at,omitempty"` // 创建时间
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" xorm:"updated" json:"updated_at,omitempty"`     // 更新时间
}

// TableName 表名
func (u *GameFuDay) TableName() string {
	return "game_fu_day"
}

// TableEngine 设置引擎为 INNODB
func (u *GameFuDay) TableEngine() string {
	return "INNODB"
}
