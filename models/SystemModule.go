package models

import (
	"goadmin/libs"

	"github.com/astaxie/beego/orm"
)

// GetModuleList 获取模块配置列表
func (module *SystemModule) GetModuleList() (modules []SystemModule, err error) {
	return module.GetModuleListAll()
}

// GetModuleGroup 获取所有父模块
func (module *SystemModule) GetModuleGroup() (tempData []SystemModule, err error) {
	om := orm.NewOrm()
	query := om.QueryTable(module)
	query = query.Filter("parent_id", 0)
	_, err = query.All(&tempData)
	return
}

// GetModuleListAll 获取所有模块
func (module *SystemModule) GetModuleListAll() (modules []SystemModule, err error) {
	om := orm.NewOrm()
	tempData := []SystemModule{}
	_, err = om.QueryTable(module).OrderBy("parent_id", "sort").All(&tempData)
	if err != nil {
		return
	}
	m := make(map[int32]int)
	for _, item := range tempData {
		if item.ParentId == 0 {
			modules = append(modules, item)
			m[item.ModuleId] = len(modules) - 1
		} else {
			mod := modules[m[item.ParentId]]
			mod.Children = append(mod.Children, item)
			modules[m[item.ParentId]] = mod
		}
	}
	return
}

func (module *SystemModule) Save() (bool, error) {
	old_module := SystemModule{ModuleId: module.ModuleId}
	if module.Type == 1 {
		module.MethodName = ""
	}
	o := orm.NewOrm()
	if module.ModuleId != 0 && o.Read(&old_module) == nil {
		if _, err := o.Update(module); err != nil {
			return false, err
		}
	} else {
		_, err := o.Insert(module)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

func (module *SystemModule) GetAccessModuleList(role_id int32) (map[string]bool, error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	access := map[string]bool{}
	tempData := []SystemModule{}
	table_module := module.TableName()
	table_access := new(UserPermission).TableName()
	qb.Select(table_module + ".*").From(table_module).
		LeftJoin(table_access).On(table_module + ".module_id = " + table_access + ".module_id").
		Where("( " + table_access + ".role_id = ? )")

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	_, err := o.Raw(sql, role_id).QueryRows(&tempData)
	if err != nil {
		return access, err
	}
	for _, item := range tempData {
		item.ModuleName = libs.UcFirst(item.ModuleName)
		item.MethodName = libs.UcFirst(item.MethodName)
		access[(item.ModuleName + "." + item.MethodName)] = true
	}
	return access, nil
}
