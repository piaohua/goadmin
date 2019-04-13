package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取玩家
func (b *TgBaseData) GetDataList(userid int, startTime, endTime string, gold,
	page, limit int, isPost bool) (tempData []TgBaseData, pag Pagination, err error) {
	o := newOrmBy(TGDIV)
	tempData = []TgBaseData{}
	query := o.QueryTable(prefixBy(TGDIV) + b.TableName())
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

// UpdateBase 更新玩家数据
func (b *TgBaseData) UpdateBase() (err error) {
	o := newOrmBy(TGDIV)
	query := o.QueryTable(prefixBy(TGDIV) + b.TableName())
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":        b.Gold,
		"gate":        b.Gate,
		"guide_step":  b.GuideStep,
		"add":         b.Add,
		"unlock_gate": b.UnlockGate,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
