package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取qise玩家
func (b *QsBaseData) GetDataList(sourceID, userid int, startTime, endTime string, gold, gate,
	page, limit int, isPost bool) (tempData []QsBaseData, pag Pagination, err error) {
	o := newOrmBy(QSDIV)
	tempData = []QsBaseData{}
	query := o.QueryTable((prefixBy(QSDIV) + b.TableName()))
	if sourceID >= 0 {
		query = query.Filter("source_id", sourceID)
	}
	if userid != 0 {
		query = query.Filter("id", userid)
	}
	if gold != 0 {
		query = query.Filter("gold", gold)
	}
	if gate != 0 {
		query = query.Filter("gate", gate)
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
		//return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetDataList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdateBase 更新qise玩家数据
func (b *QsBaseData) UpdateBase() (err error) {
	o := newOrmBy(QSDIV)
	query := o.QueryTable((prefixBy(QSDIV) + b.TableName()))
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":       b.Gold,
		"gate":       b.Gate,
		"guide_step": b.GuideStep,
		"version":    b.Version,
		"add":        b.Add,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetDataList 获取qise玩家
func (b *Qs99BaseData) GetDataList(sourceID, userid int, startTime, endTime string, gold, gate,
	page, limit int, isPost bool) (tempData []Qs99BaseData, pag Pagination, err error) {
	o := newOrmBy(QS99DIV)
	tempData = []Qs99BaseData{}
	query := o.QueryTable((prefixBy(QS99DIV) + b.TableName()))
	query = query.Filter("source_id", sourceID)
	if userid != 0 {
		query = query.Filter("id", userid)
	}
	if gold != 0 {
		query = query.Filter("gold", gold)
	}
	if gate != 0 {
		query = query.Filter("gate", gate)
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

// UpdateBase 更新qise玩家数据
func (b *Qs99BaseData) UpdateBase() (err error) {
	o := newOrmBy(QS99DIV)
	query := o.QueryTable((prefixBy(QS99DIV) + b.TableName()))
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":       b.Gold,
		"gate":       b.Gate,
		"guide_step": b.GuideStep,
		"version":    b.Version,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
