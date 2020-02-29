package controllers

import (
	"subway/hall/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type SkillController struct {
	beego.Controller
}

// @Title skillDetail
// @Description skillDetail
// @Param	roleId		query 	string	true		"The username for login"
// @Param	skillUid		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /skillDetail [post]
func (u *SkillController) SkillDetail() {
	skillUid := u.GetString("skillUid")
	skillInfo := models.GetSkill(skillUid)
	if skillInfo == nil {
		u.Data["json"] = models.Response{Code: 201, Msg: "fail", Data: nil}
	} else {
		res := ConvertSkillInfoToResponse(skillInfo)
		u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: res}
	}
	u.ServeJSON()
}

func ConvertSkillInfoToResponse(skillInfo *models.SkillInfo) *ResponseSkillDetailInfo {
	return &ResponseSkillDetailInfo{
		Uid:     skillInfo.Uid,
		SkillId: skillInfo.SkillId,
		Name:    "",
		Level:   skillInfo.Level,
		Gold:    1,
	}
}

type ResponseSkillDetailInfo struct {
	Uid     string
	SkillId int32
	Name    string
	Level   int32
	Gold    int32
}
