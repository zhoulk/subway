package controllers

import (
	"subway/hall/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type HeroController struct {
	beego.Controller
}

// @Title heroList
// @Description heroList
// @Param	roleId		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /heroList [post]
func (u *HeroController) HeroList() {
	roleId := u.GetString("roleId")
	roelHero := models.GetRoleHero(roleId)
	u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: ConvertRoleHeroToResponse(roelHero)}
	u.ServeJSON()
}

// @Title heroDetail
// @Description heroDetail
// @Param	roleId		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /heroDetail [post]
func (u *HeroController) HeroDetail() {
	heroUid := u.GetString("heroUid")
	heroInfo := models.GetHero(heroUid)
	if heroInfo == nil {
		u.Data["json"] = models.Response{Code: 201, Msg: "fail", Data: nil}
	} else {
		equipList := models.GetHeroEquipList(heroUid)
		skillList := models.GetHeroSkillList(heroUid)
		res := ConvertHeroInfoToResponse(heroInfo)
		res.EquipArr = make([]*ResponseEquipInfo, 0)
		for _, equip := range equipList {
			res.EquipArr = append(res.EquipArr, ConvertHeroEquipToResponse(equip))
		}
		res.SkillArr = make([]*ResponseSkillInfo, 0)
		for _, skill := range skillList {
			res.SkillArr = append(res.SkillArr, ConvertHeroSkillToResponse(skill))
		}
		u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: res}
	}
	u.ServeJSON()
}

func ConvertRoleHeroToResponse(roleHero *models.RoleHero) *ResponseRoleHero {
	res := &ResponseRoleHero{
		Items: make([]*ResponseHeroInfo, 0),
	}
	for _, item := range roleHero.Heros {
		res.Items = append(res.Items, ConvertRoleHeroItemToResponse(item))
	}
	return res
}

func ConvertRoleHeroItemToResponse(heroInfo *models.HeroInfo) *ResponseHeroInfo {
	return &ResponseHeroInfo{
		Uid:    heroInfo.Uid,
		HeroId: heroInfo.HeroId,
		Type:   heroInfo.Type,
		Floor:  heroInfo.Floor,
		Level:  heroInfo.Level,
		Star:   heroInfo.Star,
	}
}

func ConvertHeroInfoToResponse(heroInfo *models.HeroInfo) *ResponseHeroDetailInfo {
	return &ResponseHeroDetailInfo{
		Uid:    heroInfo.Uid,
		HeroId: heroInfo.HeroId,
		Type:   heroInfo.Type,
		Floor:  heroInfo.Floor,
		Level:  heroInfo.Level,
		Star:   heroInfo.Star,
	}
}

func ConvertHeroEquipToResponse(equip *models.EquipInfo) *ResponseEquipInfo {
	return &ResponseEquipInfo{
		Uid:     equip.Uid,
		EquipId: equip.EquipId,
		Name:    equip.Name,
	}
}

func ConvertHeroSkillToResponse(skill *models.SkillInfo) *ResponseSkillInfo {
	return &ResponseSkillInfo{
		Uid:     skill.Uid,
		SkillId: skill.SkillId,
		Name:    skill.Name,
	}
}

type ResponseRoleHero struct {
	Items []*ResponseHeroInfo
}

type ResponseHeroInfo struct {
	Uid    string
	HeroId int32
	Type   int8
	Floor  int32
	Level  int32
	Star   int32
}

type ResponseHeroDetailInfo struct {
	Uid      string
	HeroId   int32
	Type     int8
	Floor    int32
	Level    int32
	Star     int32
	EquipArr []*ResponseEquipInfo
	SkillArr []*ResponseSkillInfo
}

type ResponseEquipInfo struct {
	Uid     string
	EquipId int32
	Name    string
}

type ResponseSkillInfo struct {
	Uid     string
	SkillId int32
	Name    string
}
