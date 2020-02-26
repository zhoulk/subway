package controllers

import (
	"subway/hall/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type EquipController struct {
	beego.Controller
}

// @Title equipDetail
// @Description equipDetail
// @Param	roleId		query 	string	true		"The username for login"
// @Param	equipUid		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /equipDetail [post]
func (u *EquipController) EquipDetail() {
	equipUid := u.GetString("equipUid")
	equipInfo := models.GetEquip(equipUid)
	if equipInfo == nil {
		u.Data["json"] = models.Response{Code: 201, Msg: "fail", Data: nil}
	} else {
		res := ConvertEquipInfoToResponse(equipInfo)
		equipList := models.GetEquipMixList(equipUid)
		if equipList != nil {
			res.MixArr = make([]*ResponseEquipInfo, 0)
			for _, equip := range equipList {
				res.MixArr = append(res.MixArr, ConvertEquipItemInfoToResponse(equip))
			}
		}
		u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: res}
	}
	u.ServeJSON()
}

func ConvertEquipInfoToResponse(equipInfo *models.EquipInfo) *ResponseEquipDetailInfo {
	return &ResponseEquipDetailInfo{
		Uid:     equipInfo.Uid,
		EquipId: equipInfo.EquipId,
		Name:    "",
	}
}

func ConvertEquipItemInfoToResponse(equipInfo *models.EquipInfo) *ResponseEquipInfo {
	return &ResponseEquipInfo{
		Uid:     equipInfo.Uid,
		EquipId: equipInfo.EquipId,
		Name:    "",
	}
}

type ResponseEquipDetailInfo struct {
	Uid     string
	EquipId int32
	Name    string
	MixArr  []*ResponseEquipInfo

	Strength    int32 // 力量
	HP          int32 // 生命值
	Agility     int32 // 敏捷
	MP          int32 // 魔法强度
	Intelligent int32
	AD          int32 // 物理攻击
	ADCrit      int32 // 物理暴击
	ADDef       int32 // 物理护甲
	HPGain      int32 // 生命恢复
	MPGain      int32 // 能量恢复
	HPReGain    int32 // 战斗后补充生命
	MPReGain    int32 // 战斗后补充能量
}
