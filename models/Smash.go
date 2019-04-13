package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取smash玩家
func (b *SmBaseData) GetDataList(userid int, startTime, endTime, gold string, diamond,
	page, limit int, isPost bool) (tempData []SmBaseData, pag Pagination, err error) {
	o := newOrmBy(SMDIV)
	tempData = []SmBaseData{}
	query := o.QueryTable((prefixBy(SMDIV) + b.TableName()))
	if userid != 0 {
		query = query.Filter("id", userid)
	}
	if gold != "" {
		query = query.Filter("gold", gold)
	}
	if diamond != 0 {
		query = query.Filter("diamond", diamond)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetDataList query all num: ", num)
	if err != nil {
		beego.Error("GetDataList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetDataList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdateBase 更新smash玩家数据
func (b *SmBaseData) UpdateBase() (err error) {
	o := newOrmBy(SMDIV)
	query := o.QueryTable((prefixBy(SMDIV) + b.TableName()))
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":       b.Gold,
		"diamond":    b.Diamond,
		"guide_step": b.GuideStep,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
