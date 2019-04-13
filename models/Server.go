package models

import (
	"encoding/json"
	"fmt"
	"goadmin/libs"
	"io/ioutil"
	"time"

	"wcgames/common/model"

	"github.com/astaxie/beego"
	yaml "gopkg.in/yaml.v2"
)

// List 获取
func (b *Server) List(name, startTime, endTime string,
	status, projectID, page, limit int,
	isPost bool) (tempData []Server, pag Pagination, err error) {
	o := newOrm()
	tempData = []Server{}
	query := o.QueryTable(b.TableName())
	if name != "" {
		query = query.Filter("name", name)
	}
	if status >= 0 {
		query = query.Filter("status", status)
	}
	if projectID != 0 {
		query = query.Filter("project_id", projectID)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-host", "-updated_at")
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

// Upsert 更新
func (b *Server) Upsert() (err error) {
	o := newOrm()
	if b.ID == 0 {
		if _, err = o.Insert(b); err != nil {
			beego.Error("error: ", err)
			return
		}
	}
	b.UpdatedAt = time.Now().Local()
	if _, err = o.Update(b, "name", "host", "port",
		"path", "project_id", "jwt_secret", "token",
		"appid", "app_secret", "group", "prefixs",
		"default_db", "default_prefix",
		"logger_db", "logger_prefix", "updated_at"); err != nil {
		beego.Error("error: ", err)
		return
	}
	return
}

// UpdateStatus update status
func (b *Server) UpdateStatus(actStatus int) (err error) {
	switch actStatus {
	case 1:
		b.Status = 1
	case 2:
		b.Status = 0
	case 3:
		b.Status = 2
	case 4:
		b.Status = 3
	default:
		return
	}
	o := newOrm()
	b.UpdatedAt = time.Now().Local()
	if _, err = o.Update(b, "status", "updated_at"); err != nil {
		beego.Error("error: ", err)
		return
	}
	return
}

// Remove 删除
func (u *Server) Remove() (err error) {
	o := newOrm()
	if _, err = o.Delete(u); err != nil {
		beego.Error("error: ", err)
		return
	}
	return
}

// Get 查询
func (u *Server) Get() (err error) {
	o := newOrm()
	if err = o.Read(u); err != nil {
		beego.Error("error: ", err)
		return
	}
	return
}

// Act 执行操作指令
func (u *Server) Act(actStatus int, chatLog *SystemChat) (err error) {
	var b []byte
	var cmd string
	switch actStatus {
	case 1: //start
		cmd = fmt.Sprintf("sh ./conf/act.sh start %s %s %s %s", u.Host, u.Path, u.Name, u.Port)
	case 2: //stop
		cmd = fmt.Sprintf("sh ./conf/act.sh stop %s %s %s", u.Host, u.Path, u.Name)
	case 3: //open
		cmd = fmt.Sprintf("sh ./conf/act.sh open %s %s %s %s", u.Host, u.Path, u.Name, u.Port)
	case 4: //close
		cmd = fmt.Sprintf("sh ./conf/act.sh close %s %s %s %s", u.Host, u.Path, u.Name, u.Port)
	case 5: //create
		tplYaml("/mnt/bin/config/", u)
		tplPm2Json("/mnt/bin/pm2/", u)
		tplPm2Sh("/mnt/bin/pm2/", u)
		cmd = fmt.Sprintf("sh ./conf/act.sh create %s %s %s %s", u.Host, u.Path, u.Name, u.Port)
	case 6: //update
		cmd = fmt.Sprintf("sh ./conf/act.sh update %s %s %s", u.Host, u.Path, u.Name)
	}
	beego.Info("cmd ", cmd)
	b, err = libs.ExecCmd(cmd)
	beego.Info("b ", string(b), err)
	chatLog.Message = cmd
	if err != nil {
		chatLog.Content = err.Error()
		return err
	}
	chatLog.Content = string(b)
	return
}

// tplYaml template yaml
func tplYaml(path string, s *Server) {
	c := model.Config{}
	c.Name = s.Name
	c.Version = "0.1.1"
	c.GinMode = "release"
	c.JWT = true
	c.JWTSecret = s.JwtSecret
	c.JWTExp = 43200
	c.WXAppid = s.Appid
	c.WXAppSecret = s.AppSecret
	c.ProjectID = s.ProjectID
	c.Env = "prod"
	c.SwitchFile = "./assets/static/switch.json"
	c.DatxFile = "./config/17monipdb.datx"
	c.HeaderToken = s.Token
	c.Envs = map[string]model.Environments{
		"prod": model.Environments{
			URL:   "http://prod.xxx.com",
			Name:  "Prod",
			Addr:  ":" + s.Port,
			Cors:  true,
			Group: s.Group,
		},
	}
	c.Xorm = make(map[string]model.XormEngine)
	if s.DefaultDB != "" {
		loggerFile := fmt.Sprintf("./logs/sql_%s_default.log", s.Name)
		c.Xorm["default"] = model.XormEngine{
			DriverName:     "mysql",
			DataSourceName: s.DefaultDB,
			ShowSQL:        false,
			LoggerLevel:    1,
			MaxIdleConns:   5,
			MaxOpenConns:   30,
			Location:       "Asia/Shanghai",
			LoggerFile:     loggerFile,
			Prefix:         s.DefaultPrefix,
		}
	}
	if s.LoggerDB != "" {
		loggerFile := fmt.Sprintf("./logs/sql_%s_logger.log", s.Name)
		c.Xorm["logger"] = model.XormEngine{
			DriverName:     "mysql",
			DataSourceName: s.LoggerDB,
			ShowSQL:        false,
			LoggerLevel:    1,
			MaxIdleConns:   5,
			MaxOpenConns:   30,
			Location:       "Asia/Shanghai",
			LoggerFile:     loggerFile,
			Prefix:         s.LoggerPrefix,
		}
	}
	c.Prefixs = make(map[int]string)
	if s.Prefixs != "" {
		if err := json.Unmarshal([]byte(s.Prefixs), &c.Prefixs); err != nil {
			beego.Error("error: ", err)
			return
		}
	}
	//
	d, err := yaml.Marshal(&c)
	if err != nil {
		beego.Error("error: ", err)
		return
	}
	fileName := path + "config." + c.Name + ".yaml"
	err = ioutil.WriteFile(fileName, []byte(d), 0666)
	if err != nil {
		beego.Error("error: ", err)
	}
}

// tplPm2Json template yaml
func tplPm2Json(path string, s *Server) {
	var pm2json = `{
    "apps":
    {
        "name": "%s",
        "cwd": "%s/",
        "script": "%s/pm2/%s.sh",
        "exec_interpreter": "bash",
        "min_uptime": "60s",
        "max_restarts": 3,
        "exec_mode" : "fork_mode",
        "error_file" : "%s/logs/%s-err.log",
        "out_file": "%s/logs/%s-out.log",
        "pid_file": "%s/pids/%s.pid",
        "watch": false
    }
}`
	//pm2 json
	pmJson := fmt.Sprintf(pm2json, s.Name, s.Path,
		s.Path, s.Name, s.Path, s.Name,
		s.Path, s.Name, s.Path, s.Name)
	fileName := path + s.Name + ".json"
	err := ioutil.WriteFile(fileName, []byte(pmJson), 0666)
	if err != nil {
		beego.Error("error: ", err)
	}
}

// tplPm2Sh template shell
func tplPm2Sh(path string, s *Server) {
	var pm2sh = `#!/bin/bash
set -e
ulimit -HSn 65535
cd %s/
chmod +x ./%s-bin
./%s-bin -config=config/config.%s.yaml -log_dir=logs -stderrthreshold=ERROR`
	//
	pmSh := fmt.Sprintf(pm2sh, s.Path, s.Name, s.Name, s.Name)
	fileName := path + s.Name + ".sh"
	err := ioutil.WriteFile(fileName, []byte(pmSh), 0666)
	if err != nil {
		beego.Error("error: ", err)
	}
}
