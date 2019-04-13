package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// List 获取apidoc
func (b *ApiDoc) List(projectID int, page,
	limit int, isPost bool) (tempData []ApiDoc, pag Pagination, err error) {
	o := newOrm()
	tempData = []ApiDoc{}
	query := o.QueryTable(b.TableName())
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = query.Filter("doc_id", 0) //不是修改记录
	query = queryPage(query, page, limit, "sort")
	num, err := query.All(&tempData)
	beego.Info("List query all num: ", num)
	if err != nil {
		beego.Error("List query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("List count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// TplList 获取apidoc模板
func (b *ApiDoc) TplList() (tempData []ApiDoc, err error) {
	o := newOrm()
	tempData = []ApiDoc{}
	query := o.QueryTable(b.TableName())
	//query = query.Filter("project_id", 0) //TODO 暂时不区分项目
	query = query.Filter("type", 1) //基础模板
	query = query.Filter("type", 2) //接口模板
	num, err := query.All(&tempData)
	beego.Info("TplList query all num: ", num)
	if err != nil {
		beego.Error("TplList query all err: ", err)
		return
	}
	return
}

// Add 添加api
func (b *ApiDoc) Add() (err error) {
	o := newOrm()
	num, err := o.Insert(b)
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// Update 修改api
func (b *ApiDoc) Update() (err error) {
	o := newOrm()
	query := o.QueryTable(b.TableName())
	query = query.Filter("id", b.ID)
	//获取修改前内容
	var b2 = &ApiDoc{ID: b.ID}
	err = b2.Get()
	beego.Info("apidoc get err: ", err)
	//update
	num, err := query.Update(orm.Params{
		"type":      b.Type,
		"status":    b.Status,
		"sort":      b.Sort,
		"url":       b.URL,
		"method":    b.Method,
		"name":      b.Name,
		"detail":    b.Detail,
		"update_id": b.UpdateID,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	//添加备份修改记录
	b2.DocID = b.ID
	b2.ID = 0
	num, err = o.Insert(b2)
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// Audit 修改api status
func (b *ApiDoc) Audit() (err error) {
	o := newOrm()
	query := o.QueryTable(b.TableName())
	query = query.Filter("id", b.ID)
	num, err := query.Update(orm.Params{
		"status":     b.Status,
		"audit_id":   b.AuditID,
		"audit_time": time.Now().Local(),
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// Get 获取apidoc
func (b *ApiDoc) Get() (err error) {
	o := newOrm()
	err = o.Read(b)
	return
}

// Get 获取apidoc
func (b *ApiDoc) One(projectID int) (err error) {
	o := newOrm()
	query := o.QueryTable(b.TableName())
	query = query.Filter("id", b.ID)
	query = query.Filter("project_id", projectID)
	query = query.Filter("doc_id", 0) //不是修改记录
	err = query.One(b)
	return
}
