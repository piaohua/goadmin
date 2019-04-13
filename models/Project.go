package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"goadmin/libs"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetProjectList ...
func GetProjectList() []Project {
	project := new(Project)
	var projectList, _ = project.GetProjectList()
	return projectList
}

// GetProjectListBy ...
func GetProjectListBy(role_id int32) []Project {
	//用户项目选择权限设置
	r := new(UserRole)
	role_list, _ := r.GetRoleListBy([]int32{role_id})
	if role, ok := role_list[role_id]; ok {
		if len(role.ProjectList) != 0 {
			return role.ProjectList
		}
	}
	project := new(Project)
	var projectList, _ = project.GetProjectList()
	//没有设置选择全部
	projectList, _ = project.GetProjectList()
	return projectList
}

// GetProjectPath 获取项目静态文件路径
func GetProjectPath(projectID int) (path, api, name string, err error) {
	project := new(Project)
	project.ProjectID = projectID
	err = project.GetProjectByID()
	if err != nil {
		beego.Error("GetProjectPath err: ", err)
		return
	}
	api = project.API + "assets/static/"
	path = project.Path + "assets/static/"
	name = project.Name
	//test
	//path = "C:\\Users\\009\\go\\src\\17mmm\\config\\data\\"
	return
}

// GetProjectAPI 获取项目api
func GetProjectAPI(projectID int) (api string, err error) {
	project := new(Project)
	project.ProjectID = projectID
	err = project.GetProjectByID()
	if err != nil {
		beego.Error("err: ", err)
		return
	}
	api = project.API
	return
}

// GetProjectList 项目列表
func (u *Project) GetProjectList() (list []Project, err error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	num, err := query.Filter("status", 1).All(&list)
	beego.Info("GetProjectList query all num: ", num)
	if err != nil {
		beego.Error("GetProjectList query all err: ", err)
		return
	}
	return
}

// GetProjectByID 获取项目
func (u *Project) GetProjectByID() (err error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	err = query.Filter("project_id", u.ProjectID).One(u)
	if err != nil {
		beego.Error("GetProjectByID query one err: ", err)
		return
	}
	return
}

// GetProjectAll 项目列表
func (u *Project) GetProjectAll() (list []Project, err error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	num, err := query.All(&list)
	beego.Info("GetProjectAll query all num: ", num)
	if err != nil {
		beego.Error("GetProjectAll query all err: ", err)
		return
	}
	return
}

// Save 增加或修改
func (u *Project) Save() (bool, error) {
	o := newOrm()
	query := o.QueryTable(u.TableName())
	exist := query.Filter("id", u.Id).Exist()
	if exist {
		query = o.QueryTable(u.TableName())
		query = query.Filter("id", u.Id)
		query = query.Filter("project_id", u.ProjectID)
		if _, err := query.Update(orm.Params{
			"name":   u.Name,
			"api":    u.API,
			"path":   u.Path,
			"status": u.Status,
		}); err != nil {
			return false, err
		}
		//if _, err := o.Update(u); err != nil {
		//	return false, err
		//}
		return true, nil
	}
	query = o.QueryTable(u.TableName())
	i, err := query.PrepareInsert()
	if err != nil {
		return false, err
	}
	if _, err := i.Insert(u); err != nil {
		return false, err
	}
	return true, nil
}

// GetShareSwitchList2 分享开关列表
func GetShareSwitchList2(projectID int) (list []SwitchStatus, err error) {
	var u = new(GameJson)
	err = u.GetSwitch(projectID)
	if err != nil {
		beego.Error("err: ", err)
		return
	}
	var conf SwitchConfig
	err = json.Unmarshal([]byte(u.Content), &conf)
	if err != nil {
		beego.Error("switch load err ", err)
		return
	}
	if conf.VersionList == nil {
		err = errors.New("version empty")
		return
	}
	for version, status := range conf.VersionList {
		s := SwitchStatus{
			ProjectID: projectID,
			Name:      u.Name,
			Version:   version,
			Status:    status > 0,
		}
		list = append(list, s)
	}
	return
}

// ShareSwitchEdit2 分享开关状态编辑
func ShareSwitchEdit2(projectID, status int, version string) (err error) {
	var u = new(GameJson)
	err = u.GetSwitch(projectID)
	if err != nil {
		beego.Error("query all err: ", err)
		return
	}
	var conf SwitchConfig
	err = json.Unmarshal([]byte(u.Content), &conf)
	if err != nil {
		beego.Error("switch load err ", err)
		return
	}
	conf.VersionList[version] = status
	data, err := json.Marshal(&conf)
	if err != nil {
		beego.Error("switch write err ", err)
		return
	}
	u.Content = string(data)
	err = u.UpdateSwitch(projectID)
	if err != nil {
		beego.Error("switch write err ", err)
		return
	}
	return
}

// GetShareSwitchList 分享开关列表
func GetShareSwitchList() (list []SwitchStatus) {
	projectList := GetProjectList()
	for _, v := range projectList {
		filePath := v.Path + "assets/static/switch.json"
		ok, err := libs.PathExists(filePath)
		if !ok || err != nil {
			continue
		}
		conf := new(SwitchConfig)
		err = libs.Load(filePath, conf)
		if err != nil {
			beego.Error("switch load err ", err)
			continue
		}
		if conf.VersionList == nil {
			continue
		}
		for version, status := range conf.VersionList {
			s := SwitchStatus{
				ProjectID: v.ProjectID,
				Name:      v.Name,
				Version:   version,
				Status:    status > 0,
			}
			list = append(list, s)
		}
	}
	return
}

// ShareSwitchEdit 分享开关状态编辑
func ShareSwitchEdit(projectID, status int, name, version string) (err error) {
	project := new(Project)
	project.ProjectID = projectID
	err = project.GetProjectByID()
	if err != nil {
		return
	}
	filePath := project.Path + "assets/static/switch.json"
	ok, err := libs.PathExists(filePath)
	if !ok {
		err = errors.New("file path error")
		return
	}
	if err != nil {
		return
	}
	conf := new(SwitchConfig)
	err = libs.Load(filePath, conf)
	if err != nil {
		beego.Error("switch load err ", err)
		return
	}
	if conf.VersionList == nil {
		err = errors.New("version error")
		return
	}
	conf.VersionList[version] = status
	err = libs.WriteFile(filePath, conf)
	if err != nil {
		beego.Error("switch write err ", err)
		return
	}
	return
}

// GetStaticFileList 获取静态配置文件列表
func GetStaticFileList(dirPath string) (list []FileInfo) {
	list = make([]FileInfo, 0)
	//遍历目录，读出文件名、大小
	filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) != ".json" {
			return nil
		}
		str := strings.Split(path, "assets/static/")
		if len(str) != 2 {
			return nil
		}
		var i FileInfo
		//i.Name = info.Name()
		i.Name = str[1]
		i.Size = info.Size()
		i.Time = info.ModTime()
		list = append(list, i)
		return nil
	})
	return
}

// GetGameGlobalList 项目变量配置列表
func (u *GameGlobal) GetGameGlobalList(projectID int) (list []GameGlobal, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s;", tableName)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("GetGameGlobalList query all num: ", num)
	if err != nil {
		beego.Error("GetGameGlobalList query all err: ", err)
		return
	}
	return
}

// Save 增加或修改
func (u *GameGlobal) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameGlobal
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `info` = '%s', `num` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.Info, u.Num, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`name`, `info`, `num`, `updated_at`, `created_at`) values ('%s', '%s', %d, '%s', '%s');",
		tableName, u.Name, u.Info, u.Num, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// GetSourceList 项目来源配置列表
func GetSourceList(projectID int) (ls []GameSource) {
	p := new(GameSource)
	list, _, _ := p.GetBaseSourceList(projectID, 1, 1000, true)
	ls = []GameSource{
		{
			SourceID: -1,
			Name:     "全部",
		},
		{
			SourceID: 0,
			Name:     "非渠道",
		},
	}
	ls = append(ls, list...)
	return
}

// GetBaseSourceList 项目来源配置列表
func (u *GameSource) GetBaseSourceList(projectID int, page, limit int,
	isPost bool) (list []GameSource, pag Pagination, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by `updated_at` desc limit %d, %d;", tableName, ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("GetBaseSourceList query all num: ", num)
	if err != nil {
		beego.Error("GetBaseSourceList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	var count int64
	sql = fmt.Sprintf("select count(*) as count from %s;", tableName)
	err = o.Raw(sql).QueryRow(&count)
	if err != nil {
		beego.Error("GetBaseSourceList query all err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// Save 增加或修改
func (u *GameSource) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameSource
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `info` = '%s', `source_id` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.Info, u.SourceID, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`name`, `info`, `source_id`, `updated_at`, `created_at`) values ('%s', '%s', %d, '%s', '%s');",
		tableName, u.Name, u.Info, u.SourceID, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// GetGameUserList 项目变量配置列表
func (u *GameUser) GetGameUserList(projectID int) (list []GameUser, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s;", tableName)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("GetGameUserList query all num: ", num)
	if err != nil {
		beego.Error("GetGameUserList query all err: ", err)
		return
	}
	return
}

// Save 增加或修改
func (u *GameUser) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameUser
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	var prize = 0
	if u.Prize {
		prize = 1
	}
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `prize` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, prize, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`userid`, `prize`, `updated_at`, `created_at`) values (%d, %d, '%s', '%s');",
		tableName, u.UserID, prize, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// Save 删除
func (u *GameUser) Remove(projectID int) error {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	//num, err := o.QueryTable(tableName).Filter("id", u.ID).Delete()
	//beego.Info("BaseChest Num: ", num, ", err: ", err)
	//
	sql := fmt.Sprintf("delete from %s where id = %d;", tableName, u.ID)
	res, err := o.Raw(sql).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

// GetGameTaskList 项目任务配置列表
func (u *GameTask) GetGameTaskList(projectID int, page, limit int,
	isPost bool) (list []GameTask, pag Pagination, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by id desc limit %d, %d;", tableName, ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("GetGameTaskList query all num: ", num)
	if err != nil {
		beego.Error("GetGameTaskList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	var count int64
	sql = fmt.Sprintf("select count(*) as count from %s;", tableName)
	err = o.Raw(sql).QueryRow(&count)
	if err != nil {
		beego.Error("GetGameTaskList query all err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// Save 增加或修改
func (u *GameTask) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameTask
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `taskid` = %d, `content` = '%s', `info` = '%s', `before` = %d, `after` = %d, `prize` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.TaskID, u.Content, u.Info, u.Before, u.After, u.Prize, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`taskid`, `content`, `info`, `before`, `after`, `prize`, `updated_at`, `created_at`) values (%d, '%s', '%s', %d, %d, %d, '%s', '%s');",
		tableName, u.TaskID, u.Content, u.Info, u.Before, u.After, u.Prize, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// GetOpenidBy 获取项目玩家openid
func GetOpenidBy(projectID, userid int) (openid string,
	sourceID int, err error) {
	o, prefix := getOrmBy(projectID)
	var u IcBaseData
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select `openid` from %s where `id` = %d;", tableName, userid)
	err = o.Raw(sql).QueryRow(&openid)
	sql = fmt.Sprintf("select `source_id` from %s where `id` = %d;", tableName, userid)
	err = o.Raw(sql).QueryRow(&sourceID)
	return
}

// GetGameAchiementList 项目成就配置列表
func (u *GameAchiement) GetGameAchiementList(projectID, page, limit int,
	isPost bool) (list []GameAchiement, pag Pagination, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by id desc limit %d, %d;", tableName, ((page - 1) * limit), limit)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("GetGameAchiementList query all num: ", num)
	if err != nil {
		beego.Error("GetGameAchiementList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	var count int64
	sql = fmt.Sprintf("select count(*) as count from %s;", tableName)
	err = o.Raw(sql).QueryRow(&count)
	if err != nil {
		beego.Error("GetGameAchiementList query all err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// Save 增加或修改
func (u *GameAchiement) Save(projectID int) (bool, error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameAchiement
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `achid` = %d, `type` = %d, `describe` = '%s', `condition` = '%s', `before` = %d, `after` = %d, `jewel` = %d, `integral` = %d, `updated_at` = '%s' where `id` = %d;",
			tableName, u.AchID, u.Type, u.Describe, u.Condition, u.Before, u.After, u.Jewel, u.Integral, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`achid`, `type`, `describe`, `condition`, `before`, `after`, `jewel`, `integral`, `updated_at`, `created_at`) values (%d, %d, '%s', '%s', %d, %d, %d, %d, '%s', '%s');",
		tableName, u.AchID, u.Type, u.Describe, u.Condition, u.Before, u.After, u.Jewel, u.Integral, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// GetGameJsonList 项目变量配置列表
func (u *GameJson) GetGameJsonList(projectID int) (list []GameJson, err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s order by `updated_at` desc;", tableName)
	num, err := o.Raw(sql).QueryRows(&list)
	beego.Info("GetGameJsonList query all num: ", num)
	if err != nil {
		beego.Error("GetGameJsonList query all err: ", err)
		return
	}
	return
}

// Save 增加或修改
func (u *GameJson) Save(projectID int) (bool, error) {
	if !strings.HasSuffix(u.Name, ".json") || len(u.Name) > 50 {
		return false, errors.New("filename error")
	}
	u.Content = strings.Replace(u.Content, " ", "", -1)    //去掉空格
	u.Content = strings.Replace(u.Content, "\r\n", "", -1) //去掉换行符
	u.Content = strings.Replace(u.Content, "\r", "", -1)   //去掉换行符
	u.Content = strings.Replace(u.Content, "\n", "", -1)   //去掉换行符
	u.Content = strings.Replace(u.Content, "\t", "", -1)   //去掉制表符
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	var u2 GameJson
	err := o.Raw(sql).QueryRow(&u2)
	now := libs.Time2LocalStr(time.Now())
	if err == nil && u2.ID != 0 {
		sql = fmt.Sprintf("update %s set `type` = %d, `content` = '%s', `updated_at` = '%s' where `id` = %d;",
			tableName, u.Type, u.Content, now, u.ID)
		res, err := o.Raw(sql).Exec()
		if err != nil {
			return false, err
		}
		if _, err = res.RowsAffected(); err != nil {
			return false, err
		}
		//备份u2
		u2.ID = 0
		u2.Type = 2 //备份
		u2.Name = libs.CurrTimeStr() + u2.Name
		ok, err := u2.Save(projectID)
		beego.Info("gameJson backup: ", u2.Name, ok, err)
		return true, nil
	}
	sql = fmt.Sprintf("insert into %s (`type`, `name`, `content`, `updated_at`, `created_at`) values (%d, '%s', '%s', '%s', '%s');",
		tableName, u.Type, u.Name, u.Content, now, now)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return false, err
	}
	if _, err = res.RowsAffected(); err != nil {
		return false, err
	}
	return true, nil
}

// Save 删除
func (u *GameJson) Remove(projectID int) error {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	//
	sql := fmt.Sprintf("delete from %s where id = %d;", tableName, u.ID)
	res, err := o.Raw(sql).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		beego.Info("mysql row affected nums: ", num)
	}
	return err
}

// Restore 恢复
func (u *GameJson) Restore(projectID int) error {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	//查询
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	err := o.Raw(sql).QueryRow(u)
	if err != nil {
		return err
	}
	if len(u.Name) < 14 {
		return errors.New("restore fail")
	}
	//恢复
	now := libs.Time2LocalStr(time.Now())
	sql = fmt.Sprintf("update %s set `content` = '%s', `updated_at` = '%s' where `type` = 0 and `name` = '%s';",
		tableName, u.Content, now, string(u.Name[14:]))
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return err
	}
	if _, err = res.RowsAffected(); err != nil {
		return err
	}
	return nil
}

// Sync 同步
func (u *GameJson) Sync(projectID int) error {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	//查询
	sql := fmt.Sprintf("select * from %s where `id` = %d;", tableName, u.ID)
	err := o.Raw(sql).QueryRow(u)
	if err != nil {
		return err
	}
	var name string
	if strings.HasSuffix(u.Name, "Test.json") {
		name = strings.Replace(u.Name, "Test.json", ".json", -1) //Test -> prod
	} else {
		name = strings.Replace(u.Name, ".json", "Test.json", -1) //prod -> Test
	}
	//同步测试文件到正式文件
	now := libs.Time2LocalStr(time.Now())
	sql = fmt.Sprintf("update %s set `content` = '%s', `updated_at` = '%s' where `type` = 0 and `name` = '%s';",
		tableName, u.Content, now, name)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return err
	}
	var affected int64
	if affected, err = res.RowsAffected(); err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("update failed")
	}
	return nil
}

// GetSwitch 获取分享开关配置
func (u *GameJson) GetSwitch(projectID int) (err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	sql := fmt.Sprintf("select `name`, `content` from %s where `name` = 'switch.json' and `type` = 0;", tableName)
	err = o.Raw(sql).QueryRow(u)
	if err != nil {
		beego.Error("query err: ", err)
	}
	return
}

// UpdateSwitch 更新分享开关配置
func (u *GameJson) UpdateSwitch(projectID int) (err error) {
	o, prefix := getOrmBy(projectID)
	tableName := u.TableName()
	tableName = prefix + tableName
	now := libs.Time2LocalStr(time.Now())
	sql := fmt.Sprintf("update %s set `content` = '%s', `updated_at` = '%s' where `type` = 0 and `name` = '%s';",
		tableName, u.Content, now, u.Name)
	res, err := o.Raw(sql).Exec()
	if err != nil {
		return err
	}
	if _, err = res.RowsAffected(); err != nil {
		beego.Error("update err: ", err)
	}
	return
}
