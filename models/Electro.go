package models

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取玩家
func (b *EmBaseData) GetDataList(userid int, startTime, endTime string, gold,
	page, limit int, isPost bool) (tempData []EmBaseData, pag Pagination, err error) {
	o := newOrmBy(EMDIV)
	tempData = []EmBaseData{}
	query := o.QueryTable(prefixBy(EMDIV) + b.TableName())
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
func (b *EmBaseData) UpdateBase() (err error) {
	o := newOrmBy(EMDIV)
	query := o.QueryTable(prefixBy(EMDIV) + b.TableName())
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"gold":        b.Gold,
		"gate":        b.Gate,
		"guide_step":  b.GuideStep,
		"add":         b.Add,
		"unlock_gate": b.UnlockGate,
		"signin":      b.Signin,
		"sign_time":   b.SignTime,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// Remove 删除玩家数据
func (b *EmBaseData) Remove() (err error) {
	o := newOrmBy(EMDIV)
	if _, err = o.Delete(b); err != nil {
		beego.Error("error: ", err)
		return
	}
	return
}

// GetData 获取玩家
func (b *EmBaseFu) GetData() (err error) {
	o := newOrmBy(EMDIV)
	query := o.QueryTable(prefixBy(EMDIV) + b.TableName())
	query = query.Filter("userid", b.UserID)
	err = query.One(b)
	beego.Info("query b: ", b)
	if err != nil {
		beego.Error("query all err: ", err)
	}
	return
}

// UpdateBase 更新玩家数据
func (b *EmBaseFu) UpdateBase() (err error) {
	o := newOrmBy(EMDIV)
	query := o.QueryTable(prefixBy(EMDIV) + b.TableName())
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"today":   b.Today,
		"days":    b.Days,
		"level":   b.Level,
		"bag":     b.Bag,
		"bag_num": b.BagNum,
		"number":  b.Number,
		"ticket":  b.Ticket,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetDataList 获取玩家
func (b *EmBaseFu) GetDataList(userid int, startTime, endTime string, ticket,
	page, limit int, isPost bool) (tempData []EmBaseFu, pag Pagination, err error) {
	o := newOrmBy(EMDIV)
	tempData = []EmBaseFu{}
	query := o.QueryTable(prefixBy(EMDIV) + b.TableName())
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if ticket > 0 {
		query = query.Filter("ticket__gte", ticket)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		//return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

//update em_base_rank_num set num = num + 1 where num > 0 and score in (select score from em_base_rank where userid = 1)
//update em_base_rank_num as a inner join (select score from em_base_rank where userid = 1) as c on a.score = c.score set a.num = a.num - 1 where a.num > 0;
// RankRemove 删除玩家rank数据
func (b *EmBaseData) RankRemove() (err error) {
	o := newOrmBy(EMDIV)
	tableName := "em_base_rank"
	tableName2 := "em_base_rank_num"
	sql := fmt.Sprintf("update %s as a inner join (select score from %s where userid = %d) as c on a.score = c.score set a.num = a.num - 1 where a.num > 0", tableName2, tableName, b.ID)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return err
	}
	if _, err = res.RowsAffected(); err != nil {
		return err
	}
	sql = fmt.Sprintf("delete from %s where userid = %d", tableName, b.ID)
	res, err = o.Raw(sql).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}
