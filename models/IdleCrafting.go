package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// GetDataList 获取IdleCrafting玩家
func (ic *IcBaseData) GetDataList(userid int, startTime, endTime string, level, gate, jewel,
	page, limit int, isPost bool) (tempData []IcBaseData, pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBaseData{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("id", userid)
	}
	if jewel != 0 {
		query = query.Filter("jewel", jewel)
	}
	if level != 0 {
		query = query.Filter("level", level)
	}
	if gate != 0 {
		query = query.Filter("gate", gate)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetDataList query all num: ", num)
	if err != nil {
		beego.Error("GetDataList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetDataList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdateBase 更新IdleCrafting玩家数据
func (ic *IcBaseData) UpdateBase() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("id", ic.ID)
	num, err := query.Update(orm.Params{
		"gold":       ic.Gold,
		"jewel":      ic.Jewel,
		"key_num":    ic.KeyNum,
		"gate":       ic.Gate,
		"level":      ic.Level,
		"weapon_lev": ic.WeaponLev,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// UpdateDelete 删除IdleCrafting玩家数据
func (ic *IcBaseData) UpdateDelete() (err error) {
	o := newOrmBy(ICDIV)
	//清除数据
	num, err := o.QueryTable((prefixBy(ICDIV) + ic.TableName())).Filter("id", ic.ID).Update(orm.Params{
		"level": 1, "gate": 1, "gold": "0", "jewel": 0, "key_num": 0,
		"top_gold": "0", "top_jewel": 0, "top_key_num": 0, "score": 0,
		"brick_num": 0, "lottery_times": 0, "weapon_times": 0, "beat_boss": 0,
		"leave_time": 0, "guide_step": "", "signin": 0, "sign_time": "", "auto_atk": 0,
		"weapon": "", "gift": 0, "weapon_lev": 1,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	//
	var pt IcBaseBuyTimes
	num, err = o.QueryTable((prefixBy(ICDIV) + pt.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("BuyTimes Num: ", num, ", err: ", err)
	//
	var bc IcBaseChest
	num, err = o.QueryTable((prefixBy(ICDIV) + bc.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("BaseChest Num: ", num, ", err: ", err)
	//任务
	var bt IcBaseTask
	num, err = o.QueryTable((prefixBy(ICDIV) + bt.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("basetask Num: ", num, ", err: ", err)
	//成就
	var ba IcBaseAchiement
	num, err = o.QueryTable((prefixBy(ICDIV) + ba.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("baseachiement Num: ", num, ", err: ", err)
	//
	var be IcBaseEffect
	num, err = o.QueryTable((prefixBy(ICDIV) + be.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("IcBaseEffect Num: ", num, ", err: ", err)
	//
	var bm IcBaseMaterials
	num, err = o.QueryTable((prefixBy(ICDIV) + bm.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("IcBaseMaterials Num: ", num, ", err: ", err)
	//
	var bp IcBasePet
	num, err = o.QueryTable((prefixBy(ICDIV) + bp.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("IcBasePet Num: ", num, ", err: ", err)
	//
	var bk IcBaseSkill
	num, err = o.QueryTable((prefixBy(ICDIV) + bk.TableName())).Filter("userid", ic.ID).Delete()
	beego.Info("IcBaseSkill Num: ", num, ", err: ", err)
	return
}

// GetPetList 获取IdleCrafting玩家宠物
func (ic *IcBasePet) GetPetList(userid int, startTime, endTime, petid string,
	page, limit int, isPost bool) (tempData []IcBasePet, pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBasePet{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if petid != "" {
		query = query.Filter("pet_id", petid)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetPetList query all num: ", num)
	if err != nil {
		beego.Error("GetPetList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetPetList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdatePet 更新IdleCrafting玩家宠物
func (ic *IcBasePet) UpdatePet() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("userid", ic.UserID)
	query = query.Filter("pet_id", ic.PetID)
	num, err := query.Update(orm.Params{
		"level": ic.Level,
		"star":  ic.Star,
		"extra": ic.Extra,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetEffectList 获取IdleCrafting玩家额外效果
func (ic *IcBaseEffect) GetEffectList(userid int, startTime, endTime, effectid string,
	page, limit int, isPost bool) (tempData []IcBaseEffect, pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBaseEffect{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if effectid != "" {
		query = query.Filter("effect_id", effectid)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetEffectList query all num: ", num)
	if err != nil {
		beego.Error("GetEffectList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetEffectList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdateEffect 更新IdleCrafting玩家效果
func (ic *IcBaseEffect) UpdateEffect() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("userid", ic.UserID)
	query = query.Filter("effect_id", ic.EffectID)
	num, err := query.Update(orm.Params{
		"number": ic.Number,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetMaterialList 获取IdleCrafting玩家额外效果
func (ic *IcBaseMaterials) GetMaterialList(userid int, startTime, endTime, materialid string,
	page, limit int, isPost bool) (tempData []IcBaseMaterials, pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBaseMaterials{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if materialid != "" {
		query = query.Filter("materials_id", materialid)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetMaterialList query all num: ", num)
	if err != nil {
		beego.Error("GetMaterialList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetMaterialList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdateMaterial 更新IdleCrafting玩家效果
func (ic *IcBaseMaterials) UpdateMaterial() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("userid", ic.UserID)
	query = query.Filter("materials_id", ic.MaterialsID)
	num, err := query.Update(orm.Params{
		"number": ic.Number,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetSkillList 获取IdleCrafting玩家额外技能
func (ic *IcBaseSkill) GetSkillList(userid int, startTime, endTime, skillid string,
	page, limit int, isPost bool) (tempData []IcBaseSkill, pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBaseSkill{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if skillid != "" {
		query = query.Filter("skill_id", skillid)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetSkillList query all num: ", num)
	if err != nil {
		beego.Error("GetSkillList query all err: ", err)
		return
	}
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetSkillList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// UpdateSkill 更新IdleCrafting玩家技能
func (ic *IcBaseSkill) UpdateSkill() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("userid", ic.UserID)
	query = query.Filter("skill_id", ic.SkillID)
	num, err := query.Update(orm.Params{
		"number": ic.Number,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetAchList 获取IdleCrafting玩家额外成就
func (ic *IcBaseAchiement) GetAchList(userid int, startTime, endTime string,
	achid uint32, page, limit int, isPost bool) (tempData []IcBaseAchiement, ls []GameAchiement,
	pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBaseAchiement{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if achid != 0 {
		query = query.Filter("achid", achid)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetAchList query all num: ", num)
	if err != nil {
		beego.Error("GetAchList query all err: ", err)
		return
	}
	// 配置信息
	ls = ic.GetAchList2(tempData)
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetAchList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetAchList2 配置信息
func (ic *IcBaseAchiement) GetAchList2(list []IcBaseAchiement) (ls []GameAchiement) {
	var u GameAchiement
	o := newOrmBy(ICDIV)
	ids := make([]int, 0)
	ts := make([]int, 0)
	for _, v := range list {
		ids = append(ids, int(v.AchID))
		ts = append(ts, int(v.Type))
	}
	query := o.QueryTable((prefixBy(ICDIV) + u.TableName()))
	query = query.Filter("achid__in", ids)
	query = query.Filter("type__in", ts)
	query = query.GroupBy("achid", "type")
	num, err := query.All(&ls, "achid", "type", "describe", "condition")
	beego.Info("GetAchList2 query all num: ", num)
	if err != nil {
		beego.Error("GetAchList2 query all err: ", err)
		return
	}
	return
}

// UpdateAch 更新IdleCrafting玩家成就
func (ic *IcBaseAchiement) UpdateAch() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("userid", ic.UserID)
	query = query.Filter("achid", ic.AchID)
	num, err := query.Update(orm.Params{
		"status": ic.Status,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}

// GetTaskList 获取IdleCrafting玩家额外任务
func (ic *IcBaseTask) GetTaskList(userid int, startTime, endTime string, taskid uint32,
	page, limit int, isPost bool) (tempData []IcBaseTask, ls []GameTask, pag Pagination, err error) {
	o := newOrmBy(ICDIV)
	tempData = []IcBaseTask{}
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	if userid != 0 {
		query = query.Filter("userid", userid)
	}
	if taskid != 0 {
		query = query.Filter("taskid", taskid)
	}
	query = queryCreatedAt(query, startTime, endTime)
	query = queryPage(query, page, limit, "-id")
	num, err := query.All(&tempData)
	beego.Info("GetTaskList query all num: ", num)
	if err != nil {
		beego.Error("GetTaskList query all err: ", err)
		return
	}
	// 配置信息
	ls = ic.GetTaskList2(tempData)
	if isPost { //优化
		return
	}
	count, err := query.Count()
	if err != nil {
		beego.Error("GetTaskList count err: ", err)
		return
	}
	pag = setPag(page, count)
	return
}

// GetTaskList2 配置信息
func (ic *IcBaseTask) GetTaskList2(list []IcBaseTask) (ls []GameTask) {
	var u GameTask
	o := newOrmBy(ICDIV)
	ids := make([]int, 0)
	for _, v := range list {
		ids = append(ids, int(v.TaskID))
	}
	query := o.QueryTable((prefixBy(ICDIV) + u.TableName()))
	query = query.Filter("taskid__in", ids)
	query = query.GroupBy("taskid")
	num, err := query.All(&ls, "taskid", "content", "info")
	beego.Info("GetTaskList2 query all num: ", num)
	if err != nil {
		beego.Error("GetTaskList2 query all err: ", err)
		return
	}
	return
}

// UpdateTask 更新IdleCrafting玩家任务
func (ic *IcBaseTask) UpdateTask() (err error) {
	o := newOrmBy(ICDIV)
	query := o.QueryTable((prefixBy(ICDIV) + ic.TableName()))
	query = query.Filter("userid", ic.UserID)
	query = query.Filter("taskid", ic.TaskID)
	num, err := query.Update(orm.Params{
		"status": ic.Status,
	})
	beego.Info("Affected Num: ", num, ", err: ", err)
	return
}
