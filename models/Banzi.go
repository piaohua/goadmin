package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取banzi玩家
func (b *BzBaseData) GetDataList(userid int, startTime, endTime string, gold,
	page, limit int, isPost bool) (tempData []BzBaseData, pag Pagination, err error) {
	o := newOrmBy(BZDIV)
	tempData = []BzBaseData{}
	query := o.QueryTable((prefixBy(BZDIV) + b.TableName()))
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

// UpdateBase 更新banzi玩家数据
func (b *BzBaseData) UpdateBase() (err error) {
	o := newOrmBy(BZDIV)
	query := o.QueryTable((prefixBy(BZDIV) + b.TableName()))
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":         b.Gold,
		"skin_lock":    b.SkinLock,
		"pet_count":    b.PetCount,
		"top_gate":     b.TopGate,
		"perfect_pass": b.PerfectPass,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
