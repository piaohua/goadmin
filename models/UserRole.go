package models

import (
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Menu struct {
	Url    string `json:"url"`
	Icon   string `json:"icon"`
	Title  string `json:"title"`
	Module string `json:"module_name"`
	Action string `json:"method_name"`
}

type ParentMenu struct {
	Url      string `json:"url"`
	Icon     string `json:"icon"`
	Title    string `json:"title"`
	Module   string `json:"module_name"`
	Action   string `json:"method_name"`
	Children []Menu `json:"children,omitempty"`
}

func (ur *UserRole) GetMenus(role_id int32) (menus []ParentMenu, err error) {
	namespace := beego.AppConfig.String("namespace")
	if role_id == 0 {
		return menus, errors.New("Role ID Empty")
	}
	qb, _ := orm.NewQueryBuilder("mysql")
	var (
		ups []*SystemModule
	)
	table_module := new(SystemModule).TableName()
	table_access := new(UserPermission).TableName()
	qb.Select(table_module + ".*").From(table_module).
		LeftJoin(table_access).On(table_module + ".module_id = " + table_access + ".module_id").
		Where("( " + table_access + ".role_id = ? OR parent_id=0 ) AND " + table_module + ".show=1").
		OrderBy("type,sort")

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	_, err = o.Raw(sql, role_id).QueryRows(&ups)
	m := make(map[int32]int)
	if err == nil {
		for _, up := range ups {
			if up.ParentId == 0 {
				val := ParentMenu{
					Url:      namespace + "/" + up.ModuleName + "/" + up.MethodName,
					Icon:     up.Icon,
					Title:    up.Title,
					Module:   up.ModuleName,
					Action:   up.MethodName,
					Children: []Menu{},
				}
				menus = append(menus, val)
				m[up.ModuleId] = len(menus) - 1
			} else {
				var temp_data = ParentMenu{}
				if v, ok := m[up.ParentId]; ok {
					temp_data = menus[v]
				} else {
					continue
				}
				temp_data.Children = append(temp_data.Children, Menu{
					Url:    namespace + "/" + up.ModuleName + "/" + up.MethodName,
					Icon:   up.Icon,
					Title:  up.Title,
					Module: up.ModuleName,
					Action: up.MethodName,
				})
				menus[m[up.ParentId]] = temp_data
			}
		}
	}
	return menus, err
}

func (role *UserRole) Save() (bool, error) {
	old_role := UserRole{RoleId: role.RoleId}
	o := orm.NewOrm()
	if role.RoleId != 0 && o.Read(&old_role) == nil {
		if _, err := o.Update(role); err != nil {
			return false, err
		}
	} else {
		_, err := o.Insert(role)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

// GetRoleListBy 获取用户所属角色
func (role *UserRole) GetRoleListBy(ids []int32) (map[int32]UserRole, error) {
	om := orm.NewOrm()
	roles := map[int32]UserRole{}
	tempData := []UserRole{}
	_, err := om.QueryTable(role).Filter("role_id__in", ids).All(&tempData)
	if err != nil {
		return roles, err
	}
	for _, item := range tempData {
		item.Count, _ = om.QueryTable(new(SystemUser)).
			Filter("role_id", item.RoleId).Count()
		//转换项目选择
		var ids []int
		if err := json.Unmarshal([]byte(item.ProjectIds), &ids); err == nil {
			item.Projects = ids
		} else {
			beego.Error("projectIds error ", item)
		}
		roles[item.RoleId] = item
	}
	roles = RoleProjectList(roles)
	return roles, err
}

// GetRoleList 获取用户列表
func (role *UserRole) GetRoleList(page, limit int) (roles map[int32]UserRole,
	pag Pagination, err error) {
	om := orm.NewOrm()
	roles = map[int32]UserRole{}
	tempData := []UserRole{}
	query := om.QueryTable(role)
	query = queryPage(query, page, limit, "-role_id")
	_, err = query.All(&tempData)
	if err != nil {
		return
	}
	for _, item := range tempData {
		item.Count, _ = om.QueryTable(new(SystemUser)).
			Filter("role_id", item.RoleId).Count()
		//转换项目选择
		var ids []int
		if err := json.Unmarshal([]byte(item.ProjectIds), &ids); err == nil {
			item.Projects = ids
		} else {
			beego.Error("projectIds error ", item)
		}
		roles[item.RoleId] = item
	}
	roles = RoleProjectList(roles)
	var count int64
	count, err = query.Count()
	if err != nil {
		return
	}
	pag = setPag(page, count)
	return
}

// RoleProjectList 角色项目列表
func RoleProjectList(role_lists map[int32]UserRole) map[int32]UserRole {
	//角色项目列表,为空默认全部
	m := make(map[int]Project)
	projectList := GetProjectList()
	for k, v := range projectList {
		m[v.ProjectID] = projectList[k]
	}
	//beego.Info("projectList ", projectList)
	//beego.Info("role_lists ", role_lists)
	for k, v := range role_lists {
		if len(v.Projects) == 0 {
			v.ProjectList = projectList
			role_lists[k] = v
			continue
		}
		for _, id := range v.Projects {
			if val, ok := m[id]; ok {
				v.ProjectList = append(v.ProjectList, val)
				role_lists[k] = v
			}
		}
	}
	//beego.Info("role_lists ", role_lists)
	return role_lists
}

func (userp *UserPermission) GetAccessList(role_id int32) map[int32]int {
	om := orm.NewOrm()
	access := map[int32]int{}
	tempData := []UserPermission{}
	_, err := om.QueryTable(userp).Filter("RoleId", role_id).All(&tempData)
	if err == nil {
		for _, item := range tempData {
			access[item.ModuleId] = 1
		}
	}
	return access
}

func (role *UserRole) GetRoleName(role_id int32) string {
	var name = ""
	err := orm.NewOrm().QueryTable(role).Filter("RoleId", role_id).One(role)
	if err == nil {
		name = role.RoleName
	}
	return name
}

func (userp *UserPermission) SetPermission(role_id int32, permissios []UserPermission) bool {
	if role_id == 0 || permissios == nil {
		return false
	}
	o := orm.NewOrm()
	if _, err := o.QueryTable(userp).Filter("RoleId", role_id).Delete(); err != nil {
		return false
	}
	if _, err := o.InsertMulti(100, permissios); err != nil {
		return false
	}
	return true
}
