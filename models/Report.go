package models

import (
	"fmt"

	"goadmin/libs"
	"wcgames/common/config"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetActiveList ...
func (u *StatActive) GetActiveList(dateat string, sourceID, projectID, page,
	limit int, isPost bool) (tempData []StatActive, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	query = query.Filter("source_id", sourceID)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetActiveList query all num: ", num)
	if err != nil {
		beego.Error("GetActiveList query all err: ", err)
		return
	}
	if isPost { //优化
		//return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetActiveList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetOnlineList ...
func (u *StatPlaytime) GetOnlineList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatPlaytime, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetOnlineList query all num: ", num)
	if err != nil {
		beego.Error("GetOnlineList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetOnlineList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetStartList ...
func (u *StatStart2) GetStartList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatStart2, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetStartList query all num: ", num)
	if err != nil {
		beego.Error("GetStartList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetStartList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetGateList ...
func (u *StatGate2) GetGateList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGate2, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetGateList query all num: ", num)
	if err != nil {
		beego.Error("GetGateList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetPetList ...
func (u *StatPet) GetPetList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatPet, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetPetList query all num: ", num)
	if err != nil {
		beego.Error("GetPetList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetPetList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetUseSkillList ...
func (u *StatUseSkill) GetUseSkillList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatUseSkill, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetUseSkillList query all num: ", num)
	if err != nil {
		beego.Error("GetUseSkillList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetUseSkillList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetProgressList ...
func (u *StatProgress) GetProgressList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatProgress, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetProgressList query all num: ", num)
	if err != nil {
		beego.Error("GetProgressList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetProgressList count err: ", err)
		return
	}
	pag = setPag2(page, limit, count)
	return
}

// GetSceneList ...
func (u *StatScene) GetSceneList(dateat string, sourceID, Type, projectID, page,
	limit int, isPost bool) (tempData []StatScene, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if sourceID >= 0 {
		query = query.Filter("source_id", sourceID)
	} else {
		//cond := orm.NewCondition()
		//cond1 := cond.Or("number__gt", 0).Or("times__gt", 0)
		//query = query.SetCond(cond.AndCond(cond1))
		query = query.Filter("source_id", sourceID)
	}
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	if Type != 0 {
		query = query.Filter("type", Type)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetSceneList query all num: ", num)
	if err != nil {
		beego.Error("GetSceneList query all err: ", err)
		return
	}
	if isPost { //优化
		//return
	}
	count, err := query.Count()
	pag = setPag(page, count)
	return
}

// GetGateGrowList ...
func (u *StatGateGrow) GetGateGrowList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGateGrow, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetGateGrowList query all num: ", num)
	if err != nil {
		beego.Error("GetGateGrowList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateGrowList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetGateStopList ...
func (u *StatGateStop2) GetGateStopList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatGateStop2, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetGateStopList query all num: ", num)
	if err != nil {
		beego.Error("GetGateStopList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateStopList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetGateDataList ...
func (u *StatGateData) GetGateDataList(dateat string, projectID, Type, gate,
	page, limit int, isPost bool) (tempData []StatGateData, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	if Type != 0 {
		query = query.Filter("type", Type) //注册类型
		//query = query.Filter("type", (Type * 10)) //登录类型
	}
	if gate != 0 {
		query = query.Filter("gate__gte", gate)
		//query = query.Filter("gate__lt", (gate + 20))
	}
	query = queryPage(query, page, limit, "-today", "gate")
	num, err := query.All(&tempData)
	beego.Info("GetGateDataList query all num: ", num)
	if err != nil {
		beego.Error("GetGateDataList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateDataList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetGateLossList ...
func (u *StatGateData) GetGateLossList(dateat string, projectID, Type, gate,
	page, limit int, isPost bool) (tempData, tempData2 []StatGateData,
	tempData3, tempData4 StatGateData, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	qs := queryDateAt2(query, dateat) //前一天数据
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
		qs = qs.Filter("project_id", projectID)
	}
	if Type != 0 {
		query = query.Filter("type", Type) //注册类型
		//query = query.Filter("type", (Type * 10)) //登录类型
		qs = qs.Filter("type", Type) //注册类型
	}
	qs3 := qs.Filter("gate", 1) //第一关数据
	qs3 = qs3.Limit(1)
	qs4 := query.Filter("gate", 1) //第一关数据
	qs4 = qs4.Limit(1)
	if gate != 0 {
		query = query.Filter("gate__gte", gate)
		//query = query.Filter("gate__lt", (gate + 20))
		qs = qs.Filter("gate__gte", gate)
	}
	//query = queryPage(query, page, limit, "-today", "gate")
	query = query.OrderBy("-today", "gate").Offset((page - 1) * limit).Limit(limit + 1)
	//qs = queryPage(qs, page, limit, "-today", "gate")
	qs = qs.OrderBy("-today", "gate").Offset((page - 1) * limit).Limit(limit + 1)
	num, err := query.All(&tempData)
	beego.Info("GetGateLossList query all num: ", num)
	if err != nil {
		beego.Error("GetGateLossList query all err: ", err)
		return
	}
	num, err = qs.All(&tempData2)
	beego.Info("GetGateLossList query all num: ", num)
	if err != nil {
		beego.Error("GetGateLossList query all err: ", err)
		return
	}
	err = qs3.One(&tempData3)
	if err != nil {
		beego.Error("GetGateLossList query all err: ", err)
		return
	}
	err = qs4.One(&tempData4)
	if err != nil {
		beego.Error("GetGateLossList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateLossList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetRealtimeList ...
func (u *StatOnline) GetRealtimeList(Type, timestamp, projectID, page,
	limit int, isPost bool) (tempData []StatOnline, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = query.Filter("type", Type)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	if timestamp != 0 {
		query = query.Filter("timestamp__gte", timestamp)
		query = query.Filter("timestamp__lte", (timestamp + 86400))
	}
	query = query.OrderBy("-id")
	//query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetRealtimeList query all num: ", num)
	if err != nil {
		beego.Error("GetRealtimeList query all err: ", err)
		return
	}
	return
}

// GetRealplayList ...
func (u *StatOnline) GetRealplayList(timestamp, projectID, page,
	limit int, isPost bool) (tempData, tempData2, tempData3 []StatOnline, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = query.Filter("type", 5) //每5分钟
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = query.OrderBy("-id")
	qs2 := query
	qs3 := query
	if timestamp != 0 {
		query = query.Filter("timestamp__gte", timestamp)
		query = query.Filter("timestamp__lt", (timestamp + 86400))
		qs2 = qs2.Filter("timestamp__gte", (timestamp - (1 * 86400)))
		qs2 = qs2.Filter("timestamp__lt", timestamp)
		qs3 = qs3.Filter("timestamp__gte", (timestamp - (7 * 86400)))
		qs3 = qs3.Filter("timestamp__lt", (timestamp - (6 * 86400)))
	}
	_, err = query.All(&tempData)
	if err != nil {
		//beego.Error("GetRealplayList query all err: ", err)
		//return
	}
	_, err = qs2.All(&tempData2)
	if err != nil {
		//beego.Error("GetRealplayList query all err: ", err)
		//return
	}
	_, err = qs3.All(&tempData3)
	if err != nil {
		//beego.Error("GetRealplayList query all err: ", err)
		//return
	}
	return
}

// GetRealtodayList ...
func GetRealtodayList(projectID, page,
	limit int, isPost bool) (rs, ls []TodayData, err error) {
	o := newOrmBy(RPDIV)
	//todayStr := time.Now().Format("2006-01-02")
	start, end := libs.Today2Str()
	//tableName := prefixBy(QSDIV) + "log_register"
	tableName := config.GetSubSuffix(projectID, prefixByID(projectID)) + "log_register"
	beego.Info("GetRealtodayList ", tableName, projectID)
	//sql := fmt.Sprintf("SELECT COUNT(*) AS `count`, `source_id` FROM `%s` WHERE DATE_FORMAT(`created_at`, '%%Y-%%m-%%d') = '%s' GROUP BY `source_id`", tableName, todayStr)
	sql := fmt.Sprintf("SELECT COUNT(*) AS `count`, `source_id` FROM `%s` WHERE `created_at` >= '%s' AND `created_at` <= '%s' GROUP BY `source_id`", tableName, start, end)
	_, err = o.Raw(sql).QueryRows(&rs)
	if err != nil {
		beego.Error("GetRealtodayList query all err: ", err)
		return
	}
	//
	//tableName = prefixBy(QSDIV) + "log_login"
	tableName = config.GetSubSuffix(projectID, prefixByID(projectID)) + "log_login"
	//sql = fmt.Sprintf("SELECT COUNT(DISTINCT `userid`) AS `count`, `source_id` FROM `%s` WHERE DATE_FORMAT(`created_at`, '%%Y-%%m-%%d') = '%s' GROUP BY `source_id`", tableName, todayStr)
	sql = fmt.Sprintf("SELECT COUNT(DISTINCT `userid`) AS `count`, `source_id` FROM `%s` WHERE `created_at` >= '%s' AND `created_at` <= '%s' GROUP BY `source_id`", tableName, start, end)
	_, err = o.Raw(sql).QueryRows(&ls)
	if err != nil {
		beego.Error("GetRealtodayList query all err: ", err)
		return
	}
	return
}

// GetShareList 分享数据
func (u *StatShare) GetShareList(dateat string, sourceID, projectID, page,
	limit int, isPost bool) (tempData []StatShare, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if sourceID >= 0 {
		query = query.Filter("source_id", sourceID)
	} else {
		cond := orm.NewCondition()
		cond1 := cond.Or("regist__gt", 0).Or("login__gt", 0).
			Or("number__gt", 0).Or("times__gt", 0)
		query = query.SetCond(cond.AndCond(cond1))
	}
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetShareList query all num: ", num)
	if err != nil {
		beego.Error("GetShareList query all err: ", err)
		return
	}
	if isPost { //优化
		//return
	}
	count, err := query.Count()
	pag = setPag(page, count)
	return
}

// GetGateGoldList ...
func (u *StatGateGold) GetGateGoldList(dateat string, projectID, Type, gate,
	page, limit int, isPost bool) (tempData []StatGateGold, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = query.Filter("type", Type) //类型
	if gate != 0 {
		query = query.Filter("gate__gte", gate)
	}
	query = queryPage(query, page, limit, "-today", "gate")
	num, err := query.All(&tempData)
	beego.Info("GetGateGoldList query all num: ", num)
	if err != nil {
		beego.Error("GetGateGoldList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetGateGoldList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetDrawList ...
func (u *StatDraw) GetDrawList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatDraw, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetDrawList query all num: ", num)
	if err != nil {
		beego.Error("GetDrawList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetDrawList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetSignList ...
func (u *StatSign) GetSignList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatSign, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetSignList query all num: ", num)
	if err != nil {
		beego.Error("GetSignList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetSignList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetWallList ...
func (u *StatWall) GetWallList(dateat, appid string, projectID, appType, page,
	limit int, isPost bool) (tempData []StatWall, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	//appType = -1 //默认显示全部
	if appType >= 0 && isPost {
		query = query.Filter("type", appType)
	} else if appType == -2 {
		query = query.Filter("appid", "")
	} else { // appType == -1
		query = query.Exclude("appid", "")
	}
	if appid != "" && appType != -2 {
		query = query.Filter("appid", appid)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetWallList query all num: ", num)
	if err != nil {
		beego.Error("GetWallList query all err: ", err)
		return
	}
	if isPost { //优化
		//return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetWallList count err: ", err)
		return
	}
	//pag = setPag(page, count)
	pag = setPag2(page, limit, count)
	return
}

// GetFastList ...
func (u *StatFast) GetFastList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatFast, pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	query := o.QueryTable(u.TableName())
	query = queryDateAt(query, dateat)
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetFastList query all num: ", num)
	if err != nil {
		beego.Error("GetFastList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetFastList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetAstrologyList ...
func (u *StatAstrology) GetAstrologyList(dateat string, projectID, page,
	limit int, isPost bool) (tempData []StatAstrology, pag Pagination, err error) {
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
	pag = setPag(page, count)
	return
}

// queryDateAt ...
func queryDateAt(qs orm.QuerySeter, dateat string) orm.QuerySeter {
	if dateat == "" {
		return qs
	}
	t, err := libs.Str2LocalTime(dateat)
	if err != nil {
		beego.Error("queryDateAt err: ", err)
		return qs
	}
	//qs = qs.Filter("DATE_FORMAT(dateat, '%Y-%m-%d')", t)
	qs = qs.Filter("today", libs.Time2DayDate(t))
	return qs
}

// queryDateAt2 ...
func queryDateAt2(qs orm.QuerySeter, dateat string) orm.QuerySeter {
	if dateat == "" {
		return qs
	}
	t, err := libs.Str2LocalTime(dateat)
	if err != nil {
		beego.Error("queryDateAt err: ", err)
		return qs
	}
	//qs = qs.Filter("DATE_FORMAT(dateat, '%Y-%m-%d')", t)
	qs = qs.Filter("today", libs.Time2DayDate(t.AddDate(0, 0, -1)))
	return qs
}
