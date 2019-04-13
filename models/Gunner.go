package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取gunner玩家
func (b *GnBaseData) GetDataList(userid int, startTime, endTime string, gold,
	page, limit int, isPost bool) (tempData []GnBaseData, pag Pagination, err error) {
	o := newOrmBy(GNDIV)
	tempData = []GnBaseData{}
	query := o.QueryTable((prefixBy(GNDIV) + b.TableName()))
	if userid != 0 {
		query = query.Filter("id", userid)
	}
	if gold != 0 {
		query = query.Filter("gold", gold)
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

// UpdateBase 更新gunner玩家数据
func (b *GnBaseData) UpdateBase() (err error) {
	o := newOrmBy(GNDIV)
	query := o.QueryTable((prefixBy(GNDIV) + b.TableName()))
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":      b.Gold,
		"top_score": b.TopScore,
		"atk_prop":  b.AtkProp,
		"boom_prop": b.BoomProp,
		"out_time":  b.OutTime,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
