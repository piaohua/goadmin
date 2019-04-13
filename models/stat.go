package models

import (
	"fmt"

	"goadmin/libs"

	"github.com/astaxie/beego"
)

// GetGateRetimesList ...
func (u *StatGateRetimes) GetGateRetimesList(dateat string, projectID, Type, page,
	limit int, isPost bool) (tempData []StatGateRetimes, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	query = query.Filter("type", Type)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetGateRetimesList query all num: ", num)
	if err != nil {
		beego.Error("GetGateRetimesList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateRetimesList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetGateScoreList ...
func (u *StatGateScore) GetGateScoreList(dateat string, projectID, Type, page,
	limit int, isPost bool) (tempData []StatGateScore, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	query = query.Filter("type", Type)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetGateScoreList query all num: ", num)
	if err != nil {
		beego.Error("GetGateScoreList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateScoreList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetGateLossList 新手关卡流失
func (u *StatGateLoss) GetGateLossList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGateLoss, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit, "gate")
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.GroupBy("today").Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

// GetPetLevelList 宠物主角升级
func (u *StatPetLevel) GetPetLevelList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatPetLevel, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

// GetGateWeaponList 合成的武器
func (u *StatGateWeapon) GetGateWeaponList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGateWeapon, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit, "gate")
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.GroupBy("today").Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

//select A.*, B.data from stat_gate_gold_dis as A, stat_gate_data as B where A.project_id = B.project_id and A.today = B.today and A.gate = B.gate and A.project_id = 21 and A.today = 20190218;

//select A.*, B.data, C.data as logdata from stat_gate_gold_dis as A, stat_gate_data as B, stat_gate_data as C where
//A.project_id = B.project_id and A.today = B.today and A.gate = B.gate and B.type = 2 and
//A.project_id = C.project_id and A.today = C.today and A.gate = C.gate and C.type = 20 and
//A.project_id = 21 and A.today = 20190214;

// GetGateGoldDisList 金币产出消耗
func (u *StatGateGoldDis) GetGateGoldDisList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGateGoldDis, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit, "gate")
	//num, err := query.All(&tempData)
	//
	tableName1 := u.TableName()
	tableName2 := new(StatGateData).TableName()
	sql := fmt.Sprintf("select A.*, B.data, C.data as logdata from %s as A, %s as B, %s as C where ",
		tableName1, tableName2, tableName2)
	sql += "A.project_id = B.project_id and A.today = B.today and A.gate = B.gate and B.type = 2 and "
	sql += "A.project_id = C.project_id and A.today = C.today and A.gate = C.gate and C.type = 20 and "
	sql += fmt.Sprintf("A.project_id = %d and A.today = %d ", projectID, libs.Datetime2Day(dateat))
	sql += fmt.Sprintf("order by A.gate asc limit %d, %d;", ((page - 1) * limit), limit)
	//sql := fmt.Sprintf("select A.*, B.data from %s as A, %s as B where A.project_id = B.project_id and A.today = b.today and A.gate = B.gate and A.project_id = %d and A.today = %d order by `gate` asc limit %d, %d;",
	//	tableName1, tableName2, projectID, libs.Datetime2Day(dateat), ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.GroupBy("today").Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

// GetGateGoldDepList 金币存量
func (u *StatGateGoldDep) GetGateGoldDepList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGateGoldDep, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit, "gate")
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.GroupBy("today").Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

// GetWeaponLevList 武器等级停留
func (u *StatWeaponLev) GetWeaponLevList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatWeaponLev, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

// GetPropList 道具
func (u *StatProp) GetPropList(Type uint32, dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatProp, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	query = query.Filter("type", Type)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}
