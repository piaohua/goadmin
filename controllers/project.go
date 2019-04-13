package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"goadmin/libs"
	"goadmin/models"
	"wcgames/common/middleware"

	"github.com/astaxie/beego"
)

// ProjectController ...
type ProjectController struct {
	baseAdminController
}

// Index ...
func (c *ProjectController) Index() {
	p := new(models.Project)
	lists, _ := p.GetProjectAll()
	//beego.Info("project lists ", lists)
	c.Data["ProjectData"] = &lists
	c.display("project/index.html")
}

// Save ...
func (c *ProjectController) Save() {
	var project_id, err = c.GetInt("project_id")
	var name = c.GetString("name")
	var path = c.GetString("path")
	var api = c.GetString("api")
	var status, _ = c.GetInt8("status", int8(0))
	var id, _ = c.GetInt("id", 0)
	if project_id == 0 || err != nil || name == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	project := &models.Project{
		Id:        id,
		Name:      name,
		Path:      path,
		API:       api,
		ProjectID: project_id,
		Status:    status,
	}
	if ok, err := project.Save(); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// SwitchList 开关状态列表
func (c *ProjectController) SwitchList() {
	var project_id, _ = c.GetInt("project_id")
	//lists := models.GetShareSwitchList()
	lists, _ := models.GetShareSwitchList2(project_id)
	beego.Info("switch lists ", lists)
	c.Data["SwitchData"] = &lists
	c.display("project/switch.html")
}

// SwitchEdit 开关状态编辑
func (c *ProjectController) SwitchEdit() {
	var project_id, err = c.GetInt("project_id")
	var name = c.GetString("name")
	var version = c.GetString("version")
	var status, _ = c.GetInt("status", 0)
	beego.Info("switch edit ", name, ", project_id ",
		project_id, ", version ", version, ", status ", status)
	if project_id == 0 || err != nil || name == "" || version == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	//if err = models.ShareSwitchEdit(project_id, status, name, version); err != nil {
	if err = models.ShareSwitchEdit2(project_id, status, version); err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// 文件列表
func (c *ProjectController) Files() {
	var project_id, _ = c.GetInt("project_id")
	var projectPath, api, projectName, err = models.GetProjectPath(project_id)
	if err != nil {
		beego.Error(err.Error())
	} else {
		list := models.GetStaticFileList(projectPath)
		for k, v := range list {
			v.Href = api + v.Name
			v.Project = projectName
			list[k] = v
		}
		c.Data["FilesData"] = &list
	}
	beego.Info("files projectPath ", projectPath, ", project_id ", project_id)
	c.Data["ProjectID"] = project_id
	c.display("project/files.html")
}

// 编辑文件
func (c *ProjectController) FileEdit() {
	var fileName = c.GetString("name")
	var content = c.GetString("content")
	var project_id, _ = c.GetInt("project_id")
	beego.Info("edit file ", fileName, ", project_id ", project_id)
	var projectPath, _, _, err = models.GetProjectPath(project_id)
	if err != nil {
		beego.Error(err.Error())
		c.RenderJson(300, err.Error(), nil)
		return
	}
	//TODO 备份文件
	err = ioutil.WriteFile((projectPath + fileName), []byte(content), 0666)
	if err != nil {
		beego.Error(err.Error())
		c.RenderJson(300, err.Error(), nil)
		return
	}
	c.RenderJson(200, "处理成功", nil)
}

// 删除文件
func (c *ProjectController) FileDelete() {
	fileName := c.GetString("fileName")
	var project_id, _ = c.GetInt("project_id")
	beego.Info("delete file ", fileName, ", project_id ", project_id)
	var projectPath, _, _, err = models.GetProjectPath(project_id)
	if err != nil {
		beego.Error(err.Error())
		c.RenderJson(300, err.Error(), nil)
		return
	}
	err = os.Remove(projectPath + fileName) //删除文件
	if err != nil {
		beego.Error(err.Error())
		c.RenderJson(300, err.Error(), nil)
		return
	}
	beego.Info("delete file ", fileName)
	c.RenderJson(200, "处理成功", nil)
}

//// 下载文件
//func (c *ProjectController) FileDownload() {
//	fileName := c.GetString("fileName")
//	beego.Info("download file ", fileName)
//	c.Ctx.Output.Download(fileName)
//}

// 上传文件
func (c *ProjectController) FileUpload() {
	var project_id, _ = c.GetInt("project_id")
	c.Data["ProjectID"] = project_id
	var projectPath, _, _, err = models.GetProjectPath(project_id)
	if err != nil {
		beego.Error(err.Error())
		//c.RenderJson(300, err.Error(), nil)
		c.redirect(beego.URLFor("ProjectController.Files", "project_id", project_id))
		return
	}
	beego.Info("file upload ", projectPath, ", project_id ", project_id)
	file, h, err := c.GetFile("uploadFile") //获取上传的文件
	if err != nil {
		beego.Error(err.Error())
		//c.RenderJson(300, err.Error(), nil)
		c.redirect(beego.URLFor("ProjectController.Files", "project_id", project_id))
		return
	}
	if file == nil || h == nil {
		beego.Error("file error")
		//c.RenderJson(300, "file error", nil)
		c.redirect(beego.URLFor("ProjectController.Files", "project_id", project_id))
		return
	}
	file.Close() //关闭上传的文件
	//file.Size()
	var fileName = projectPath + h.Filename
	beego.Info("file name: ", fileName)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		beego.Error(err.Error())
		//c.RenderJson(300, err.Error(), nil)
		c.redirect(beego.URLFor("ProjectController.Files", "project_id", project_id))
		return
	}
	if f != nil {
		io.Copy(f, file)
		f.Close()
	}
	//c.RenderJson(200, "处理成功", nil)
	c.redirect(beego.URLFor("ProjectController.Files", "project_id", project_id))
}

// Global ...
func (c *ProjectController) Global() {
	projectID, _ := c.GetInt("project_id", 0)
	var lists []models.GameGlobal
	if projectID != 0 {
		p := new(models.GameGlobal)
		lists, _ = p.GetGameGlobalList(projectID)
		beego.Info("game global lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	//projectList := models.GetProjectList()
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["GlobalData"] = &lists
	c.display("project/global.html")
}

// GlobalEdit ...
func (c *ProjectController) GlobalEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var name = c.GetString("name")
	var info = c.GetString("info")
	var num, _ = c.GetInt32("num", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || name == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameGlobal{
		ID:   id,
		Name: name,
		Info: info,
		Num:  num,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// Gameuser ...
func (c *ProjectController) Gameuser() {
	projectID, _ := c.GetInt("project_id", 0)
	var lists []models.GameUser
	if projectID != 0 {
		p := new(models.GameUser)
		lists, _ = p.GetGameUserList(projectID)
		beego.Info("game user lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	c.Data["GameUserData"] = &lists
	c.display("project/gameuser.html")
}

// GameuserEdit ...
func (c *ProjectController) GameuserEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var prize, err = c.GetInt32("prize", 0)
	var id, _ = c.GetInt("id", 0)
	var userid, _ = c.GetInt("userid", 0)
	if projectID == 0 || err != nil {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameUser{
		ID:     id,
		UserID: userid,
	}
	if prize == 1 {
		g.Prize = true
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// 删除文件
func (c *ProjectController) GameuserDelete() {
	var project_id, _ = c.GetInt("project_id")
	var id, _ = c.GetInt("id")
	beego.Info("delete gameuser ", id, ", project_id ", project_id)
	g := &models.GameUser{
		ID: id,
	}
	if err := g.Remove(project_id); err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// GameJson ...
func (c *ProjectController) GameJson() {
	projectID, _ := c.GetInt("project_id", 0)
	var lists []models.GameJson
	var api string
	if projectID != 0 {
		p := new(models.GameJson)
		lists, _ = p.GetGameJsonList(projectID)
		//beego.Info("gamejson lists ", lists)
		_, api, _, _ = models.GetProjectPath(projectID)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	c.Data["ProjectAPI"] = api
	c.Data["GameJsonData"] = &lists
	c.display("project/gamejson.html")
}

// GameJsonEdit ...
func (c *ProjectController) GameJsonEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var name = c.GetString("name")
	var content = c.GetString("content")
	var id, _ = c.GetInt("id", 0)
	var Type, err = c.GetUint32("type", 0)
	if projectID == 0 || name == "" || err != nil {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameJson{
		ID:      id,
		Name:    name,
		Content: content,
		Type:    Type,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// 删除文件
func (c *ProjectController) GameJsonDelete() {
	var project_id, _ = c.GetInt("project_id")
	var id, _ = c.GetInt("id")
	beego.Info("delete GameJson ", id, ", project_id ", project_id)
	g := &models.GameJson{
		ID: id,
	}
	if err := g.Remove(project_id); err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// 恢复文件
func (c *ProjectController) GameJsonRestore() {
	var project_id, _ = c.GetInt("project_id")
	var id, _ = c.GetInt("id")
	beego.Info("restore GameJson ", id, ", project_id ", project_id)
	g := &models.GameJson{
		ID: id,
	}
	if err := g.Restore(project_id); err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// 同步文件
func (c *ProjectController) GameJsonSync() {
	var project_id, _ = c.GetInt("project_id")
	var id, _ = c.GetInt("id")
	beego.Info("sync GameJson ", id, ", project_id ", project_id)
	g := &models.GameJson{
		ID: id,
	}
	if err := g.Sync(project_id); err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// Task ...
func (c *ProjectController) Task() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.GameTask
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.GameTask)
		lists, pagination, _ = p.GetGameTaskList(projectID, page, c.pageSize, c.isPost())
		beego.Info("game task lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ProjectController.Task", "project_id", projectID)
	c.Data["Pagination"] = &pagination
	c.Data["TaskData"] = &lists
	c.display("project/task.html")
}

// TaskEdit ...
func (c *ProjectController) TaskEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var content = c.GetString("content")
	var info = c.GetString("info")
	var before, _ = c.GetUint32("before", 0)
	var after, _ = c.GetUint32("after", 0)
	var prize, _ = c.GetUint32("prize", 0)
	var taskid, _ = c.GetUint32("taskid", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || taskid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameTask{
		ID:      id,
		TaskID:  taskid,
		Content: content,
		Info:    info,
		Before:  before,
		After:   after,
		Prize:   prize,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// Token ...
func (c *ProjectController) Token() {
	projectID, _ := c.GetInt("project_id", 0)
	userid, _ := c.GetInt("userid", 0)
	dateat := c.GetString("dateat", "")
	openid := c.GetString("openid", "")
	token := c.GetString("token", "")
	sourceID, _ := c.GetInt("source_id", 0)
	beego.Info("token project_id ", projectID)
	beego.Info("token source_id ", sourceID)
	beego.Info("token userid ", userid)
	beego.Info("token dateat ", dateat)
	beego.Info("token openid ", openid)
	beego.Info("token token ", token)
	var tokenData string
	var err error
	// 设置jwt secret
	secret := beego.AppConfig.String("jwt.secret")
	//beego.Info("token secret ", secret)
	middleware.SetSignKey(secret)
	if projectID != 0 && userid != 0 && dateat != "" {
		t, _ := libs.Str2LocalTime(dateat)
		second := t.Unix() - time.Now().Unix()
		if second > 0 {
			middleware.SetExpiresAt(second)
		}
		if openid == "" {
			openid, sourceID, err = models.GetOpenidBy(projectID, userid)
		}
		if openid != "" {
			tokenData, err = middleware.GenToken(openid, userid, projectID, sourceID)
		}
	} else if token != "" {
		openid, userid, projectID, sourceID, err = middleware.ParseToken(token)
	}
	if err != nil {
		tokenData = err.Error()
	}
	if userid != 0 {
		c.Data["UserID"] = userid
	}
	//projectList := models.GetProjectList()
	c.Data["OpenID"] = openid
	c.Data["Dateat"] = dateat
	c.Data["SourceID"] = sourceID
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["TokenData"] = tokenData
	c.display("project/token.html")
}

// Achiement ...
func (c *ProjectController) Achiement() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.GameAchiement
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.GameAchiement)
		lists, pagination, _ = p.GetGameAchiementList(projectID, page, c.pageSize, c.isPost())
		beego.Info("game achiement lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	//projectList := models.GetProjectList()
	pagination.BaseUrl = beego.URLFor("ProjectController.Achiement", "project_id", projectID)
	beego.Info("achiement pagination: ", pagination)
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["AchiementData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("project/achiement.html")
}

// AchiementEdit ...
func (c *ProjectController) AchiementEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var describe = c.GetString("describe")
	var condition = c.GetString("condition")
	var before, _ = c.GetUint32("before", 0)
	var after, _ = c.GetUint32("after", 0)
	var jewel, _ = c.GetUint32("jewel", 0)
	var integral, _ = c.GetUint32("integral", 0)
	var achid, _ = c.GetUint32("achid", 0)
	var Type, _ = c.GetInt32("type", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || achid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameAchiement{
		ID:        id,
		AchID:     achid,
		Type:      Type,
		Describe:  describe,
		Condition: condition,
		Before:    before,
		After:     after,
		Jewel:     jewel,
		Integral:  integral,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ApiDoc apidoc列表展示
func (c *ProjectController) ApiDoc() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.ApiDoc
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.ApiDoc)
		lists, pagination, _ = p.List(projectID, page, c.pageSize, c.isPost())
		//beego.Info("api doc lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ProjectController.ApiDoc", "project_id", projectID)
	beego.Info("apidoc pagination: ", pagination)
	c.Data["ApiDocData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("apidoc/index.html")
}

// ApiDocView apidoc展示或编辑
func (c *ProjectController) ApiDocView() {
	//projectID, _ := c.GetInt("project_id", 0)
	var id, _ = c.GetInt("id", 0)
	g := &models.ApiDoc{
		ID: id,
	}
	err := g.Get()
	if err != nil {
		//c.RenderJson(300, err.Error(), nil)
	}
	c.Data["ApiDocData"] = g
	c.display("apidoc/edit.html")
}

// ApiDocEdit apidoc修改或添加
func (c *ProjectController) ApiDocEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var url = c.GetString("url")
	var name = c.GetString("title")
	var detail = c.GetString("detail")
	var sort, _ = c.GetInt("sort", 0)
	var method, _ = c.GetInt("method", 0)
	var status, _ = c.GetInt("status", 0)
	var Type, _ = c.GetInt("type", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || name == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.ApiDoc{
		Sort:   sort,
		URL:    url,
		Method: method,
		Name:   name,
		Detail: detail,
		Status: status,
		Type:   Type,
	}
	if id == 0 { //add
		g.CreateID = int(c.userId)
		g.ProjectId = int(projectID)
		if err := g.Add(); err != nil {
			c.RenderJson(300, err.Error(), nil)
		}
	} else { //update
		g.ID = id
		g.UpdateID = int(c.userId)
		if err := g.Update(); err != nil {
			c.RenderJson(300, err.Error(), nil)
		}
	}
	c.RenderJson(200, "处理成功", nil)
}

// ApiDocAudit apidoc 审核
func (c *ProjectController) ApiDocAudit() {
	projectID, _ := c.GetInt("project_id", 0)
	var status, _ = c.GetInt("status", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || id == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.ApiDoc{
		ID:      id,
		Status:  status,
		AuditID: int(c.userId),
	}
	if err := g.Audit(); err != nil {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// Source ...
func (c *ProjectController) Source() {
	projectID, _ := c.GetInt("project_id", 0)
	page, _ := c.GetInt("page", 1)
	var lists []models.GameSource
	var pagination models.Pagination
	if projectID != 0 {
		p := new(models.GameSource)
		lists, pagination, _ = p.GetBaseSourceList(projectID, page, c.pageSize, c.isPost())
		beego.Info("base source lists ", lists)
	}
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ProjectController.Source", "project_id", projectID)
	c.Data["Pagination"] = &pagination
	c.Data["SourceData"] = &lists
	c.display("project/source.html")
}

// SourceEdit ...
func (c *ProjectController) SourceEdit() {
	projectID, _ := c.GetInt("project_id", 0)
	var name = c.GetString("name")
	var info = c.GetString("info")
	var source_id, _ = c.GetInt("source_id", 0)
	var id, _ = c.GetInt("id", 0)
	if projectID == 0 || name == "" {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	g := &models.GameSource{
		ID:       id,
		Name:     name,
		Info:     info,
		SourceID: source_id,
	}
	if ok, err := g.Save(projectID); !ok {
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ApiReq ...
func (c *ProjectController) ApiReq() {
	projectID, _ := c.GetInt("project_id", 0)
	report, _ := c.GetInt("report", 0)
	userid, _ := c.GetInt("userid", 0)
	path := c.GetString("path", "")
	body := c.GetString("body", "")
	//if c.isPost() && path == "" || userid <= 0 {
	//	c.RenderJson(300, "缺少必要参数", nil)
	//}
	beego.Info("api reqest userid ", userid)
	var resultData, token, openid, api string
	var sourceID int
	var err error
	if c.isPost() {
		if openid, sourceID, err = models.GetOpenidBy(projectID, userid); err != nil {
			resultData = err.Error()
		}
		if resultData == "" && openid == "" {
			resultData = "缺少必要参数"
		}
		if resultData == "" {
			// 设置jwt secret
			secret := beego.AppConfig.String("jwt.secret")
			//beego.Info("token secret ", secret)
			middleware.SetSignKey(secret)
			middleware.SetExpiresAt(60)
			if token, err = middleware.GenToken(openid, userid, projectID, sourceID); err != nil {
				resultData = err.Error()
			}
		}
		if resultData == "" {
			var projectID2 = projectID
			if report != 0 {
				projectID2 = report
			}
			if api == "" {
				if api, err = models.GetProjectAPI(projectID2); err != nil {
					resultData = err.Error()
				}
			}
			if report != 0 {
				switch projectID {
				case 11, 20: //idlecrafting, triangle
					api = strings.Replace(api, "4", "6", -1)
				}
			}
		}
		var b []byte
		if resultData == "" {
			body2 := strings.Replace(body, "\"", "\\\"", -1)
			// default post
			cmd := fmt.Sprintf("curl -H \"Content-Type:application/json\" -H \"Authorization:%s\" -X POST -d \"%s\" %s", token, body2, api+path)
			beego.Info("cmd ", cmd)
			b, err = libs.ExecCmd(cmd)
			beego.Info("b ", string(b), err)
			if err != nil {
				resultData = err.Error()
			} else {
				resultData = string(b)
			}
		}
	}
	if userid != 0 {
		c.Data["UserID"] = userid
	}
	if report != 0 {
		c.Data["Report"] = report
	}
	c.Data["PATH"] = path
	c.Data["BODY"] = body
	c.Data["ResultData"] = resultData
	c.display("project/apitest.html")
}
