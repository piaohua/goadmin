package controllers

import (
	"goadmin/models"

	"github.com/astaxie/beego"
)

// IdleCraftingController ...
type IdleCraftingController struct {
	baseAdminController
}

// Index ...
func (c *IdleCraftingController) Index() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	gate, _ := c.GetInt("gate", 0)
	level, _ := c.GetInt("level", 0)
	jewel, _ := c.GetInt("jewel", 0)
	beego.Info("ic index startTime: ", startTime)
	beego.Info("ic index endTime: ", endTime)
	beego.Info("ic index userid: ", userid)
	beego.Info("ic index jewel: ", jewel)
	ic := new(models.IcBaseData)
	lists, pagination, _ := ic.GetDataList(userid, startTime, endTime,
		level, gate, jewel, page, c.pageSize, c.isPost())
	//beego.Info("ic lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.Index")
	beego.Info("ic pagination: ", pagination)
	if jewel != 0 {
		c.Data["Jewel"] = jewel
	}
	if gate != 0 {
		c.Data["Gate"] = gate
	}
	if level != 0 {
		c.Data["Level"] = level
	}
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	c.Data["IcData"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/index.html")
}

// BaseEdit ...
func (c *IdleCraftingController) BaseEdit() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	gold := c.GetString("gold")
	jewel, err := c.GetUint32("jewel", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "宝石参数错误", nil)
	}
	key_num, err := c.GetUint32("key_num", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "钥匙参数错误", nil)
	}
	gate, err := c.GetUint32("gate", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	level, err := c.GetUint32("level", 0)
	if err != nil {
		beego.Error("base edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	weapon_lev, _ := c.GetUint32("weapon_lev", 0)
	beego.Info("ic baseEdit gold: ", gold)
	beego.Info("ic baseEdit jewel: ", jewel)
	beego.Info("ic baseEdit key_num: ", key_num)
	beego.Info("ic baseEdit userid: ", userid)
	pet := new(models.IcBaseData)
	pet.ID = userid
	pet.Gold = gold
	pet.Jewel = jewel
	pet.KeyNum = key_num
	pet.Gate = gate
	pet.Level = level
	pet.WeaponLev = weapon_lev
	err = pet.UpdateBase()
	if err != nil {
		beego.Error("ic baseEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// BaseDelete ...
func (c *IdleCraftingController) BaseDelete() {
	userid, _ := c.GetInt("userid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	beego.Info("ic baseEdit userid: ", userid)
	pet := new(models.IcBaseData)
	pet.ID = userid
	err := pet.UpdateDelete()
	if err != nil {
		beego.Error("ic baseDelete err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ViewPets ...
func (c *IdleCraftingController) ViewPets() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	petid := c.GetString("petid", "")
	beego.Info("ic viewPets startTime: ", startTime)
	beego.Info("ic viewPets endTime: ", endTime)
	beego.Info("ic viewPets userid: ", userid)
	beego.Info("ic viewPets petid: ", petid)
	ic := new(models.IcBasePet)
	lists, pagination, _ := ic.GetPetList(userid, startTime, endTime,
		petid, page, c.pageSize, c.isPost())
	beego.Info("ic viewPets lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.ViewPets")
	beego.Info("ic pagination: ", pagination)
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Petid"] = petid
	c.Data["IcPets"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/viewPets.html")
}

// PetEdit ...
func (c *IdleCraftingController) PetEdit() {
	userid, _ := c.GetInt("userid", 0)
	petid, _ := c.GetInt("pet_id", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	level, err := c.GetUint32("level", 0)
	if err != nil {
		beego.Error("pet edit err: ", err)
		c.RenderJson(300, "等级参数错误", nil)
	}
	star, err := c.GetUint32("star", 0)
	if err != nil {
		beego.Error("pet edit err: ", err)
		c.RenderJson(300, "星数参数错误", nil)
	}
	extra, err := c.GetUint32("extra", 0)
	if err != nil {
		beego.Error("pet edit err: ", err)
		c.RenderJson(300, "进度参数错误", nil)
	}
	beego.Info("ic petEdit level: ", level)
	beego.Info("ic petEdit star: ", star)
	beego.Info("ic petEdit extra: ", extra)
	beego.Info("ic petEdit userid: ", userid)
	beego.Info("ic petEdit petid: ", petid)
	pet := new(models.IcBasePet)
	pet.UserID = userid
	pet.PetID = petid
	pet.Level = level
	pet.Star = star
	pet.Extra = extra
	err = pet.UpdatePet()
	if err != nil {
		beego.Error("ic petEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ViewEffects ...
func (c *IdleCraftingController) ViewEffects() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	effectid := c.GetString("effectid", "")
	beego.Info("ic viewEffects startTime: ", startTime)
	beego.Info("ic viewEffects endTime: ", endTime)
	beego.Info("ic viewEffects userid: ", userid)
	beego.Info("ic viewEffects effectid: ", effectid)
	ic := new(models.IcBaseEffect)
	lists, pagination, _ := ic.GetEffectList(userid, startTime, endTime,
		effectid, page, c.pageSize, c.isPost())
	beego.Info("ic viewEffects lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.ViewEffects")
	beego.Info("ic pagination: ", pagination)
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	c.Data["Userid"] = userid
	c.Data["Effectid"] = effectid
	c.Data["IcEffects"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/viewEffects.html")
}

// EffectEdit ...
func (c *IdleCraftingController) EffectEdit() {
	userid, _ := c.GetInt("userid", 0)
	effectid, _ := c.GetInt("effect_id", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	number, err := c.GetUint32("number", 0)
	if err != nil {
		beego.Error("effect edit err: ", err)
		c.RenderJson(300, "等级参数错误", nil)
	}
	beego.Info("ic effectEdit number: ", number)
	beego.Info("ic effectEdit userid: ", userid)
	beego.Info("ic effectEdit effectid: ", effectid)
	effect := new(models.IcBaseEffect)
	effect.UserID = userid
	effect.EffectID = effectid
	effect.Number = number
	err = effect.UpdateEffect()
	if err != nil {
		beego.Error("ic effectEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ViewMaterials ...
func (c *IdleCraftingController) ViewMaterials() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	materialid := c.GetString("materialid", "")
	beego.Info("ic viewMaterials startTime: ", startTime)
	beego.Info("ic viewMaterials endTime: ", endTime)
	beego.Info("ic viewMaterials userid: ", userid)
	beego.Info("ic viewMaterials materialid: ", materialid)
	ic := new(models.IcBaseMaterials)
	lists, pagination, _ := ic.GetMaterialList(userid, startTime, endTime,
		materialid, page, c.pageSize, c.isPost())
	beego.Info("ic viewMaterials lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.ViewMaterials")
	beego.Info("ic pagination: ", pagination)
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Materialid"] = materialid
	c.Data["IcMaterials"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/viewMaterials.html")
}

// MaterialEdit ...
func (c *IdleCraftingController) MaterialEdit() {
	userid, _ := c.GetInt("userid", 0)
	materialid, _ := c.GetInt("material_id", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	number, err := c.GetUint32("number", 0)
	if err != nil {
		beego.Error("material edit err: ", err)
		c.RenderJson(300, "等级参数错误", nil)
	}
	beego.Info("ic materialEdit number: ", number)
	beego.Info("ic materialEdit userid: ", userid)
	beego.Info("ic materialEdit materialid: ", materialid)
	material := new(models.IcBaseMaterials)
	material.UserID = userid
	material.MaterialsID = materialid
	material.Number = number
	err = material.UpdateMaterial()
	if err != nil {
		beego.Error("ic materialEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ViewSkills ...
func (c *IdleCraftingController) ViewSkills() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	skillid := c.GetString("skillid", "")
	beego.Info("ic viewSkills startTime: ", startTime)
	beego.Info("ic viewSkills endTime: ", endTime)
	beego.Info("ic viewSkills userid: ", userid)
	beego.Info("ic viewSkills skillid: ", skillid)
	ic := new(models.IcBaseSkill)
	lists, pagination, _ := ic.GetSkillList(userid, startTime, endTime,
		skillid, page, c.pageSize, c.isPost())
	beego.Info("ic viewSkills lists: ", lists)
	if c.isPost() {
		c.RenderJson(200, "处理成功", lists)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.ViewSkills")
	beego.Info("ic pagination: ", pagination)
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Skillid"] = skillid
	c.Data["IcSkills"] = &lists
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/viewSkills.html")
}

// SkillEdit ...
func (c *IdleCraftingController) SkillEdit() {
	userid, _ := c.GetInt("userid", 0)
	skillid, _ := c.GetInt("skill_id", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	number, err := c.GetUint32("number", 0)
	if err != nil {
		beego.Error("skill edit err: ", err)
		c.RenderJson(300, "等级参数错误", nil)
	}
	beego.Info("ic skillEdit number: ", number)
	beego.Info("ic skillEdit userid: ", userid)
	beego.Info("ic skillEdit skillid: ", skillid)
	skill := new(models.IcBaseSkill)
	skill.UserID = userid
	skill.SkillID = skillid
	skill.Number = number
	err = skill.UpdateSkill()
	if err != nil {
		beego.Error("ic skillEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ViewAchs 成就展示
func (c *IdleCraftingController) ViewAchs() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	achid, _ := c.GetUint32("achid", 0)
	beego.Info("ic viewAchs startTime: ", startTime)
	beego.Info("ic viewAchs endTime: ", endTime)
	beego.Info("ic viewAchs userid: ", userid)
	beego.Info("ic viewAchs achid: ", achid)
	ic := new(models.IcBaseAchiement)
	lists, lists2, pagination, _ := ic.GetAchList(userid, startTime, endTime,
		achid, page, c.pageSize, c.isPost())
	beego.Info("ic viewAchs lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["IcAchs"] = lists
		m["IcAchsInfo"] = lists2
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.ViewAchs")
	beego.Info("ic pagination: ", pagination)
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Achid"] = achid
	c.Data["IcAchs"] = &lists
	c.Data["IcAchsInfo"] = &lists2
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/viewAchiement.html")
}

// AchEdit 成就编辑
func (c *IdleCraftingController) AchEdit() {
	userid, _ := c.GetInt("userid", 0)
	achid, _ := c.GetUint32("achid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	status, err := c.GetUint32("status", 0)
	if err != nil {
		beego.Error("ach edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	beego.Info("ic achEdit status: ", status)
	beego.Info("ic achEdit userid: ", userid)
	beego.Info("ic achEdit achid: ", achid)
	val := new(models.IcBaseAchiement)
	val.UserID = userid
	val.AchID = achid
	val.Status = status
	err = val.UpdateAch()
	if err != nil {
		beego.Error("ic achEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}

// ViewTasks 任务展示
func (c *IdleCraftingController) ViewTasks() {
	page, _ := c.GetInt("page", 1)
	startTime := c.GetString("startTime", "")
	endTime := c.GetString("endTime", "")
	userid, _ := c.GetInt("userid", 0)
	taskid, _ := c.GetUint32("taskid", 0)
	beego.Info("ic viewTasks startTime: ", startTime)
	beego.Info("ic viewTasks endTime: ", endTime)
	beego.Info("ic viewTasks userid: ", userid)
	beego.Info("ic viewTasks taskid: ", taskid)
	ic := new(models.IcBaseTask)
	lists, lists2, pagination, _ := ic.GetTaskList(userid, startTime, endTime,
		taskid, page, c.pageSize, c.isPost())
	beego.Info("ic viewTasks lists: ", lists)
	if c.isPost() {
		m := make(map[string]interface{})
		m["IcTasks"] = lists
		m["IcTasksInfo"] = lists2
		c.RenderJson(200, "处理成功", m)
		return
	}
	pagination.BaseUrl = beego.URLFor("IdleCraftingController.ViewTasks")
	beego.Info("ic pagination: ", pagination)
	c.Data["StartTime"] = startTime
	c.Data["EndTime"] = endTime
	if userid != 0 {
		c.Data["Userid"] = userid
	}
	c.Data["Taskid"] = taskid
	c.Data["IcTasks"] = &lists
	c.Data["IcTasksInfo"] = &lists2
	c.Data["Pagination"] = &pagination
	c.display("idlecrafting/viewTask.html")
}

// TaskEdit 任务编辑
func (c *IdleCraftingController) TaskEdit() {
	userid, _ := c.GetInt("userid", 0)
	taskid, _ := c.GetUint32("taskid", 0)
	if userid == 0 {
		c.RenderJson(300, "缺少必要参数", nil)
	}
	status, err := c.GetUint32("status", 0)
	if err != nil {
		beego.Error("task edit err: ", err)
		c.RenderJson(300, "参数错误", nil)
	}
	beego.Info("ic taskEdit status: ", status)
	beego.Info("ic taskEdit userid: ", userid)
	beego.Info("ic taskEdit taskid: ", taskid)
	val := new(models.IcBaseTask)
	val.UserID = userid
	val.TaskID = taskid
	val.Status = status
	err = val.UpdateTask()
	if err != nil {
		beego.Error("ic taskEdit err: ", err)
		c.RenderJson(300, err.Error(), nil)
	}
	c.RenderJson(200, "处理成功", nil)
}
