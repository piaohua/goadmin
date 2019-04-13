package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取玩家
func (b *CaBaseData) GetDataList(userid int, startTime, endTime string, gold,
	page, limit int, isPost bool) (tempData []CaBaseData, pag Pagination, err error) {
	o := newOrmBy(CADIV)
	tempData = []CaBaseData{}
	query := o.QueryTable(prefixBy(CADIV) + b.TableName())
	if userid != 0 {
		query = query.Filter("id", userid)
	}
	if gold != 0 {
		query = query.Filter("gold", gold)
	}
	if b.Score != 0 {
		query = query.Filter("score", b.Score)
	}
	if b.Bomb != 0 {
		query = query.Filter("bomb", b.Bomb)
	}
	if b.Throw != 0 {
		query = query.Filter("throw", b.Throw)
	}
	if b.Refresh != 0 {
		query = query.Filter("refresh", b.Refresh)
	}
	if b.Gate != 0 {
		query = query.Filter("gate", b.Gate)
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
func (b *CaBaseData) UpdateBase() (err error) {
	o := newOrmBy(CADIV)
	query := o.QueryTable(prefixBy(CADIV) + b.TableName())
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":        b.Gold,
		"score":       b.Score,
		"bomb":        b.Bomb,
		"throw":       b.Throw,
		"refresh":     b.Refresh,
		"gate":        b.Gate,
		"guide_step":  b.GuideStep,
		"add":         b.Add,
		"unlock_gate": b.UnlockGate,
		"prize":       b.Prize,
		"playday":     b.Playday,
		"playtime":    b.Playtime,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// Remove 删除玩家数据
func (b *CaBaseData) Remove() (err error) {
	o := newOrmBy(CADIV)
	//if _, err = o.Delete(b); err != nil {
	//	beego.Error("error: ", err)
	//	return
	//}
	//清除数据
	num, err := o.QueryTable((prefixBy(CADIV) + b.TableName())).
		Filter("id", b.ID).Update(orm.Params{
		"gold": 0, "gate": 1, "score": 0, "bomb": 0, "throw": 0,
		"add": 0, "guide_step": 0, "unlock_gate": 0, "prize": 0,
		"playday": 0, "playtime": 0, "refresh": 0,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
