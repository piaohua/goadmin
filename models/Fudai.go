package models

import (
	"fmt"
	"time"

	"goadmin/libs"

	"github.com/astaxie/beego"
)

// GetGameFuDaiList 配置列表
func (u *GameFuDai) GetGameFuDaiList(projectID, page, limit int,
	isPost bool) (list []GameFuDai, pag Pagination, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by id asc limit %d, %d;", tableName, ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	var count int64
	sql = fmt.Sprintf("select count(*) as count from %s;", tableName)
	err = o.Raw(sql).QueryRow(&count)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// Save 增加或修改
func (u *GameFuDai) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameFuDai
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `bag` = %d, `num` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.Bag, u.Num, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`bag`, `num`, `updated_at`, `created_at`) values (%d, %d, '%s', '%s');",
		tableName, u.Bag, u.Num, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// GetGameFuGateList 配置列表
func (u *GameFuGate) GetGameFuGateList(projectID, page, limit int,
	isPost bool) (list []GameFuGate, pag Pagination, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by id asc limit %d, %d;", tableName, ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	var count int64
	sql = fmt.Sprintf("select count(*) as count from %s;", tableName)
	err = o.Raw(sql).QueryRow(&count)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// Save 增加或修改
func (u *GameFuGate) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameFuGate
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `bag` = %d, `num` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.Bag, u.Num, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`bag`, `num`, `updated_at`, `created_at`) values (%d, %d, '%s', '%s');",
		tableName, u.Bag, u.Num, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// GetGameFuDayList 配置列表
func (u *GameFuDay) GetGameFuDayList(projectID, page, limit int,
	isPost bool) (list []GameFuDay, pag Pagination, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by id asc limit %d, %d;", tableName, ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("query all num: ", num)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	var count int64
	sql = fmt.Sprintf("select count(*) as count from %s;", tableName)
	err = o.Raw(sql).QueryRow(&count)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// Save 增加或修改
func (u *GameFuDay) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameFuDay
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `day` = %d, `num` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.Day, u.Num, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`day`, `num`, `updated_at`, `created_at`) values (%d, %d, '%s', '%s');",
		tableName, u.Day, u.Num, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}
