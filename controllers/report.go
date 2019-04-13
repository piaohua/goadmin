package controllers

import (
	"goadmin/libs"
	"goadmin/models"

	"github.com/astaxie/beego"
)

// ReportController ...
type ReportController struct {
	baseAdminController
}

// Active ...
func (c *ReportController) Active() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	if !c.isPost() {
		sourceID = -1 //默认显示全部
	}
	report := new(models.StatActive)
	lists, pagination, _ := report.GetActiveList(dateat, sourceID, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["ActiveData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Active", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	sourceList := models.GetSourceList(projectID)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["ActiveData"] = &lists
	c.Data["Pagination"] = &pagination
	c.Data["Sources"] = &sourceList
	c.Data["SourceID"] = sourceID
	c.display("report/active.html")
}

// Online ...
func (c *ReportController) Online() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatPlaytime)
	lists, pagination, _ := report.GetOnlineList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Online", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["OnlineData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/online.html")
}

// Start ...
func (c *ReportController) Start() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatStart2)
	lists, pagination, _ := report.GetStartList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Start", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["StartData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/start.html")
}

// Draw ...
func (c *ReportController) Draw() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatDraw)
	lists, pagination, _ := report.GetDrawList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Draw", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["DrawData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/draw.html")
}

// Sign ...
func (c *ReportController) Sign() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatSign)
	lists, pagination, _ := report.GetSignList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Sign", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["SignData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("stat/sign.html")
}

// Wall ...
func (c *ReportController) Wall() {
	page, _ := c.GetInt("page", 1)
	limit, _ := c.GetInt("limit", c.pageSize)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	appType, _ := c.GetInt("apptype", 0)
	appid := c.GetString("appid", "")
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	beego.Info("report index appType: ", appType)
	beego.Info("report index appid: ", appid)
	report := new(models.StatWall)
	if limit <= 0 {
		limit = c.pageSize
	}
	lists, pagination, _ := report.GetWallList(dateat, appid, projectID, appType, page, limit, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["WallData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Wall", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["WallData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/wall.html")
}

// Fast ...
func (c *ReportController) Fast() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatFast)
	lists, pagination, _ := report.GetFastList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Fast", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["FastData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/fast.html")
}

// Astrology ...
func (c *ReportController) Astrology() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatAstrology)
	lists, pagination, _ := report.GetAstrologyList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Astrology", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	c.Data["Dateat"] = dateat
	c.Data["AstrologyData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/astrology.html")
}

// Progress ...
func (c *ReportController) Progress() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatProgress)
	var pageSize = c.pageSize
	if projectID == 11 {
		pageSize = 1
	}
	lists, pagination, _ := report.GetProgressList(dateat, projectID, page, pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Progress", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["ProgressData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/progress.html")
}

// Scene ...
func (c *ReportController) Scene() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	if !c.isPost() {
		//sourceID = -1 //默认显示全部
	}
	report := new(models.StatScene)
	lists1, pagination, _ := report.GetSceneList(dateat, sourceID, 1, projectID, page, c.pageSize, c.isPost())
	lists2, _, _ := report.GetSceneList(dateat, sourceID, 2, projectID, page, c.pageSize, c.isPost())
	lists3, _, _ := report.GetSceneList(dateat, sourceID, 3, projectID, page, c.pageSize, c.isPost())
	lists4, _, _ := report.GetSceneList(dateat, sourceID, 4, projectID, page, c.pageSize, c.isPost())
	var lists5, lists6, lists7, lists8 []models.StatScene
	switch projectID {
	case 11: //idleCrafting
		lists5, _, _ = report.GetSceneList(dateat, sourceID, 5, projectID, page, c.pageSize, c.isPost())
		lists6, _, _ = report.GetSceneList(dateat, sourceID, 6, projectID, page, c.pageSize, c.isPost())
	case 21, 22: //electro, card2048
		lists5, _, _ = report.GetSceneList(dateat, sourceID, 5, projectID, page, c.pageSize, c.isPost())
		lists6, _, _ = report.GetSceneList(dateat, sourceID, 6, projectID, page, c.pageSize, c.isPost())
		lists7, _, _ = report.GetSceneList(dateat, sourceID, 7, projectID, page, c.pageSize, c.isPost())
		lists8, _, _ = report.GetSceneList(dateat, sourceID, 8, projectID, page, c.pageSize, c.isPost())
	}
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		//m := make(map[string][]models.StatScene)
		m["SceneData1"] = lists1
		m["SceneData2"] = lists2
		m["SceneData3"] = lists3
		m["SceneData4"] = lists4
		m["SceneData5"] = lists5
		m["SceneData6"] = lists6
		m["SceneData7"] = lists7
		m["SceneData8"] = lists8
		c.RenderJson(200, "处理成功", m)
		return
	}
	beego.Info("report lists1: ", len(lists1), len(lists2), len(lists3))
	pagination.BaseUrl = beego.URLFor("ReportController.Scene", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	sourceList := models.GetSourceList(projectID)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["SceneData1"] = &lists1
	c.Data["SceneData2"] = &lists2
	c.Data["SceneData3"] = &lists3
	c.Data["SceneData4"] = &lists4
	c.Data["SceneData5"] = &lists5
	c.Data["SceneData6"] = &lists6
	c.Data["SceneData7"] = &lists7
	c.Data["SceneData8"] = &lists8
	c.Data["Sources"] = &sourceList
	c.Data["SourceID"] = sourceID
	c.Data["Pagination"] = &pagination
	c.display("report/scene.html")
}

// Gate 关卡分布
func (c *ReportController) Gate() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatGate2)
	lists, pagination, _ := report.GetGateList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Gate", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["GateList"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/gate.html")
}

// Pet 宠物解锁
func (c *ReportController) Pet() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatPet)
	lists, pagination, _ := report.GetPetList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Pet", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["PetData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/pet.html")
}

// Useskill 技能使用
func (c *ReportController) Useskill() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatUseSkill)
	lists, pagination, _ := report.GetUseSkillList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.Useskill", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["UseSkillData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/useskill.html")
}

// GateGrow 关卡增长
func (c *ReportController) GateGrow() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatGateGrow)
	lists, pagination, _ := report.GetGateGrowList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.GateGrow", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["GateGrowData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/gategrow.html")
}

// GateStop 关卡停留
func (c *ReportController) GateStop() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatGateStop2)
	lists, pagination, _ := report.GetGateStopList(dateat, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.GateStop", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	//projectList := models.GetProjectList()
	c.Data["Dateat"] = dateat
	//c.Data["Projects"] = &projectList
	//c.Data["ProjectID"] = projectID
	c.Data["GateStopData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/gatestop.html")
}

// GateData 关卡数据统计分布(七巧板)
func (c *ReportController) GateData() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	gate, _ := c.GetInt("gate", 0)
	Type, _ := c.GetInt("type", 0)
	if Type <= 0 {
		Type = 1 // 默认1
	}
	if gate <= 0 {
		gate = 1 // 默认1
	}
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatGateData)
	lists, pagination, _ := report.GetGateDataList(dateat, projectID, Type, gate, page, c.pageSize, c.isPost())
	lists2, _, _ := report.GetGateDataList(dateat, projectID, (Type * 10), gate, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	beego.Info("report lists2: ", lists2)
	if c.isPost() {
		m := make(map[string][]models.StatGateData)
		m["RegGateData"] = lists
		m["LogGateData"] = lists2
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.GateData", "project_id", projectID, "type", Type)
	beego.Info("report pagination: ", pagination)
	var TypeList = []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
	}{
		{Type: 1, Name: "使用步数分布"},
		{Type: 2, Name: "通关时间分布"},
		{Type: 3, Name: "金币提示次数分布"},
		{Type: 4, Name: "免费提示次数分布"},
	}
	c.Data["Dateat"] = dateat
	c.Data["RegGateData"] = &lists
	c.Data["LogGateData"] = &lists2
	c.Data["Gate"] = gate
	c.Data["Type"] = Type
	c.Data["TypeList"] = &TypeList
	c.Data["Pagination"] = &pagination
	c.display("report/gatedata.html")
}

// Realtime ...
func (c *ReportController) Realtime() {
	page, _ := c.GetInt("page", 1)
	projectID, _ := c.GetInt("project_id", 0)
	dateat := c.GetString("dateat", "")
	timeat := c.GetString("timeat", "")
	timestamp, _ := c.GetInt("timestamp", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	beego.Info("report index timeat: ", timeat)
	beego.Info("report index timestamp: ", timestamp)
	// 整天
	t, err := libs.Str2LocalTime(dateat)
	if err != nil {
		//beego.Error("dateat err: ", err)
		t = libs.TodayTime()
	}
	timestamp = int(t.Unix())
	//
	report := new(models.StatOnline)
	//Type == 0 //累计数据
	lists, _ := report.GetRealtimeList(0, timestamp, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	c.Data["Dateat"] = dateat
	c.Data["Timeat"] = timeat
	c.Data["OnlineData"] = &lists
	c.display("report/realtime.html")
}

// Realplay ...
func (c *ReportController) Realplay() {
	page, _ := c.GetInt("page", 1)
	projectID, _ := c.GetInt("project_id", 0)
	dateat := c.GetString("dateat", "")
	timeat := c.GetString("timeat", "")
	timestamp, _ := c.GetInt("timestamp", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	beego.Info("report index timeat: ", timeat)
	beego.Info("report index timestamp: ", timestamp)
	// 整天
	t, err := libs.Str2LocalTime(dateat)
	if err != nil {
		//beego.Error("dateat err: ", err)
		t = libs.TodayTime()
	}
	timestamp = int(t.Unix())
	//
	report := new(models.StatOnline)
	lists1, lists2, lists3, _ := report.GetRealplayList(timestamp, projectID, page, c.pageSize, c.isPost())
	//beego.Info("report lists: ", lists1)
	if c.isPost() {
		m := make(map[string]interface{})
		m["TodayPlay"] = lists1
		m["YesterdayPlay"] = lists2
		m["LastweekPlay"] = lists3
		c.RenderJson(200, "处理成功", m)
		return
	}
	c.Data["Dateat"] = dateat
	c.Data["Timeat"] = timeat
	c.Data["TodayPlay"] = &lists1
	c.Data["YesterdayPlay"] = &lists2
	c.Data["LastweekPlay"] = &lists3
	c.display("report/realplay.html")
}

// Realtrend ...
func (c *ReportController) Realtrend() {
	page, _ := c.GetInt("page", 1)
	projectID, _ := c.GetInt("project_id", 0)
	dateat := c.GetString("dateat", "")
	timeat := c.GetString("timeat", "")
	timestamp, _ := c.GetInt("timestamp", 0)
	granularity, _ := c.GetInt("granularity", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	beego.Info("report index timeat: ", timeat)
	beego.Info("report index timestamp: ", timestamp)
	beego.Info("report index granularity: ", granularity)
	if granularity <= 0 {
		granularity = 1
	}
	// 整天
	t, err := libs.Str2LocalTime(dateat)
	if err != nil {
		//beego.Error("dateat err: ", err)
		t = libs.TodayTime()
	}
	timestamp = int(t.Unix())
	//
	report := new(models.StatOnline)
	lists, _ := report.GetRealtimeList(granularity, timestamp, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	c.Data["Granularity"] = granularity
	c.Data["Dateat"] = dateat
	c.Data["Timeat"] = timeat
	c.Data["Timeat"] = timeat
	c.Data["OnlineData"] = &lists
	c.display("report/realtrend.html")
}

// Realtoday ...
func (c *ReportController) Realtoday() {
	page, _ := c.GetInt("page", 1)
	projectID, _ := c.GetInt("project_id", 0)
	beego.Info("report index projectID: ", projectID)
	rs, ls, _ := models.GetRealtodayList(projectID, page, c.pageSize, c.isPost())
	beego.Info("report rs: ", rs)
	beego.Info("report ls: ", ls)
	sourceList := models.GetSourceList(projectID)
	if c.isPost() {
		m := make(map[string]interface{})
		m["RegistData"] = rs
		m["LoginData"] = ls
		c.RenderJson(200, "处理成功", m)
		return
	}
	//rs = []models.TodayData{
	//	{
	//		SourceId: 0,
	//		Count:    10,
	//	},
	//	{
	//		SourceId: 1,
	//		Count:    11,
	//	},
	//}
	//ls = []models.TodayData{
	//	{
	//		SourceId: 0,
	//		Count:    23,
	//	},
	//	{
	//		SourceId: 1,
	//		Count:    21,
	//	},
	//}
	m := make(map[int]models.TodayData2)
	for _, v := range rs {
		if val, ok := m[v.SourceId]; ok {
			val.Regist = v.Count
			m[v.SourceId] = val
		} else {
			m[v.SourceId] = models.TodayData2{
				SourceId: v.SourceId,
				Regist:   v.Count,
			}
		}
	}
	for _, v := range ls {
		if val, ok := m[v.SourceId]; ok {
			val.Dau = v.Count
			m[v.SourceId] = val
		} else {
			m[v.SourceId] = models.TodayData2{
				SourceId: v.SourceId,
				Dau:      v.Count,
			}
		}
	}
	c.Data["RegistData"] = &rs
	c.Data["LoginData"] = &ls
	c.Data["TodayData"] = m
	c.Data["Sources"] = &sourceList
	c.display("report/realtoday.html")
}

// GateLoss 关卡流失数据分布(七巧板)
func (c *ReportController) GateLoss() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	gate, _ := c.GetInt("gate", 0)
	Type1, _ := c.GetInt("type1", 0) //分布数据类型
	Type2, _ := c.GetInt("type2", 0) //注册/登录
	if Type1 <= 0 {
		Type1 = 1 // 默认1
	}
	if Type2 <= 0 {
		Type2 = 1 // 默认1
	}
	if gate <= 0 {
		gate = 1 // 默认1
	}
	if dateat == "" {
		dateat = libs.TodayStr(-1)
	}
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	Type := Type1
	if Type2 == 2 { //登录类型
		Type = Type * 10
	}
	report := new(models.StatGateData)
	lists1, lists2, gate2, gate1, pagination, _ := report.GetGateLossList(dateat,
		projectID, Type, gate, page, c.pageSize, c.isPost())
	beego.Info("report lists1: ", lists1) //当前天
	beego.Info("report lists2: ", lists2) //前一天
	beego.Info("report gate1: ", gate1)   //第一关
	if c.isPost() {
		m := make(map[string]interface{})
		m["GateData1"] = lists1
		m["GateData2"] = lists2
		m["Gate1Data"] = gate1
		m["Gate2Data"] = gate2
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.GateLoss", "project_id", projectID, "type1", Type1, "type2", Type2, "dateat", dateat)
	beego.Info("report pagination: ", pagination)
	var TypeList1 = []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
	}{
		{Type: 1, Name: "使用步数分布"},
		{Type: 2, Name: "通关时间分布"},
		{Type: 3, Name: "金币提示次数分布"},
		{Type: 4, Name: "免费提示次数分布"},
	}
	var TypeList2 = []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
	}{
		{Type: 1, Name: "注册人数"},
		{Type: 2, Name: "登录人数"},
	}
	c.Data["Dateat"] = dateat
	c.Data["GateData1"] = &lists1
	c.Data["GateData2"] = &lists2
	c.Data["Gate"] = gate
	c.Data["Gate1Data"] = &gate1
	c.Data["Gate2Data"] = &gate2
	c.Data["Type1"] = Type1
	c.Data["Type2"] = Type2
	c.Data["TypeList1"] = &TypeList1
	c.Data["TypeList2"] = &TypeList2
	c.Data["Pagination"] = &pagination
	c.display("report/gateloss.html")
}

// Share 分享数据
func (c *ReportController) Share() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	sourceID, _ := c.GetInt("source_id", 0)
	if !c.isPost() {
		sourceID = -1 //默认显示全部
	}
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatShare)
	lists, pagination, _ := report.GetShareList(dateat, sourceID, projectID, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["Pagination"] = pagination
		m["ShareList"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	sourceList := models.GetSourceList(projectID)
	pagination.BaseUrl = beego.URLFor("ReportController.Share", "project_id", projectID)
	beego.Info("report pagination: ", pagination)
	c.Data["Sources"] = &sourceList
	c.Data["SourceID"] = sourceID
	c.Data["Dateat"] = dateat
	c.Data["ShareList"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("report/share.html")
}

// GateGold 关卡分布金币数据统计(七巧板)
func (c *ReportController) GateGold() {
	page, _ := c.GetInt("page", 1)
	dateat := c.GetString("dateat", "")
	projectID, _ := c.GetInt("project_id", 0)
	gate, _ := c.GetInt("gate", 0)
	Type, _ := c.GetInt("type", 0)
	beego.Info("report index projectID: ", projectID)
	beego.Info("report index dateat: ", dateat)
	report := new(models.StatGateGold)
	lists, pagination, _ := report.GetGateGoldList(dateat, projectID, Type, gate, page, c.pageSize, c.isPost())
	beego.Info("report lists: ", lists)
	if c.isPost() {
		m := make(map[string][]models.StatGateGold)
		m["GateGoldData"] = lists
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("ReportController.GateGold", "project_id", projectID, "type", Type)
	beego.Info("report pagination: ", pagination)
	var TypeList = []struct {
		Type int    `json:"type"`
		Name string `json:"name"`
	}{
		{Type: 0, Name: "金币产出消耗"},
		{Type: 1, Name: "金币存储转盘"},
	}
	c.Data["Dateat"] = dateat
	c.Data["GateGoldData"] = &lists
	c.Data["Gate"] = gate
	c.Data["Type"] = Type
	c.Data["TypeList"] = &TypeList
	c.Data["Pagination"] = &pagination
	c.display("report/gategold.html")
}
