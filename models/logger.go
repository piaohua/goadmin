package models

import (
	"fmt"
	"goadmin/libs"
	"wcgames/common/config"

	"github.com/astaxie/beego"
)

// GetStatList ...
func (u *StatSendClick) GetStatList(timestamp, projectID, version, cover,
	page, limit int) (tempData []StatSendClick, pag Pagination, err error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	if timestamp != 0 {
		query = query.Filter("timestamp", timestamp)
	}
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	if version != 0 {
		query = query.Filter("version", version)
	}
	if cover != 0 {
		query = query.Filter("cover", cover)
	}
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetStatList query all num: ", num)
	if err != nil {
		beego.Error("GetStatList query all err: ", err)
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetStatList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetChatList 获取最新10条记录
func GetChatList() []SystemChat {
	chat := new(SystemChat)
	var list, _ = chat.GetChatList(1, 10)
	return list
}

// GetChatLast 获取id之后的记录
func GetChatLast(id int) (list []SystemChat) {
	if id == 0 {
		return
	}
	chat := new(SystemChat)
	list, _ = chat.GetChatLast(id)
	return
}

// GetChatList ...
func (u *SystemChat) GetChatList(page, limit int) (tempData []SystemChat, err error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	query = queryPage(query, page, limit)
	num, err := query.All(&tempData)
	beego.Info("GetChatList query all num: ", num)
	if err != nil {
		beego.Error("GetChatList query all err: ", err)
		return
	}
	return
}

// GetChatLast ...
func (u *SystemChat) GetChatLast(id int) (tempData []SystemChat, err error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	if id != 0 {
		query = query.Filter("id__gt", id)
	}
	_, err = query.All(&tempData)
	//beego.Info("GetChatLast query all num: ", num)
	if err != nil {
		//beego.Error("GetChatLast query all err: ", err)
		return
	}
	return
}

// Save ...
func (u *SystemChat) Save() (bool, error) {
	o := newOrm()
	_, err := o.Insert(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

// log

// GetLogLoginList 获取登录记录日志
func GetLogLoginList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogLogin,
	pag Pagination, err error) {
	var u LogLogin
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogRegisterList 获取注册记录日志
func GetLogRegisterList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogRegister,
	pag Pagination, err error) {
	var u LogRegister
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogSendClickList 获取分享图点击记录日志
func GetLogSendClickList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogSendClick,
	pag Pagination, err error) {
	var u LogSendClick
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogProgressList 获取点击转化记录日志
func GetLogProgressList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogProgress,
	pag Pagination, err error) {
	var u LogProgress
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogSceneActList 获取场景漏斗记录日志
func GetLogSceneActList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogSceneAct,
	pag Pagination, err error) {
	var u LogSceneAct
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogPlaytimeList 获取在线时长记录日志
func GetLogPlaytimeList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogPlaytime,
	pag Pagination, err error) {
	var u LogPlaytime
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogTipList 获取关卡提示记录日志
func GetLogTipList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogTip,
	pag Pagination, err error) {
	var u LogTip
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogWallList 获取流量墙记录日志
func GetLogWallList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogWall,
	pag Pagination, err error) {
	var u LogWall
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogGateGrowList 获取前进关卡记录日志
func GetLogGateGrowList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogGateGrow,
	pag Pagination, err error) {
	var u LogGateGrow
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogGoldList 获取金币记录日志
func GetLogGoldList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogGold,
	pag Pagination, err error) {
	var u LogGold
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// GetLogGateList 获取关卡记录日志
func GetLogGateList(datetime string, projectID, userid,
	page, limit int, isPost bool) (ls []LogGate,
	pag Pagination, err error) {
	var u LogGate
	pag, err = getLogList(u.TableName(), datetime, projectID,
		userid, page, limit, isPost, &ls)
	return
}

// getLogList 获取记录日志
func getLogList(tableName, datetime string, projectID,
	userid, page, limit int, isPost bool,
	ls interface{}) (pag Pagination, err error) {
	o := newOrmBy(RPDIV)
	//table name
	prefix := prefixByID(projectID)
	today := libs.Datetime2Day(datetime)
	tableName = config.GetSubTableByDay(today,
		projectID, prefix, tableName)
	//select sql
	var sql, where, order string
	sql = fmt.Sprintf("SELECT * FROM `%s`", tableName)
	where = logWhere(datetime, userid)
	order = logPage(page, limit)
	_, err = o.Raw(sql + where + order).QueryRows(ls)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost {
		return
	}
	var count int64
	sql = fmt.Sprintf("SELECT COUNT(*) FROM `%s` ", tableName)
	err = o.Raw(sql + where).QueryRow(&count)
	if err != nil {
		beego.Error("count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

func logWhere(datetime string, userid int) (where string) {
	if userid != 0 {
		where += fmt.Sprintf(" WHERE `userid` = %d", userid)
	}
	if datetime == "" {
		return
	}
	start, end := libs.Datetime2Format(datetime)
	if start == "" || end == "" {
		return
	}
	if where == "" {
		where = " WHERE "
	} else {
		where += " AND "
	}
	where += fmt.Sprintf("`created_at` >= '%s' AND `created_at` < '%s'", start, end)
	return
}

func logPage(page, limit int) (sql string) {
	sql = fmt.Sprintf(" ORDER BY `id` DESC LIMIT %d, %d", ((page - 1) * limit), limit)
	return
}
