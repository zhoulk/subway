package controllers

import (
	"subway/models"

	"github.com/astaxie/beego"
)

// Operations about Heros
type HeroController struct {
	beego.Controller
}

// @Title AllHeros
// @Description get all Heros
// @Success 200 {object} models.Hero
// @router /all [post]
func (h *HeroController) AllHeros() {
	heros := models.GetAllHeros()
	h.Data["json"] = heros
	h.ServeJSON()
}

// @Title SelfHeros
// @Description get all Heros
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /self [post]
func (h *HeroController) SelfHeros() {
	uid := h.GetString("uid")
	heros := models.GetSelfHeros(uid)
	h.Data["json"] = models.Response{Code: 200, Msg: "", Data: heros}
	h.ServeJSON()
}

// @Title LevelUpHero
// @Description level up hero
// @Param	uid		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /levelUp [post]
func (h *HeroController) LevelUpHero() {
	uid := h.GetString("uid")
	heroUid := h.GetString("heroUid")
	if models.HeroLevelUp(uid, heroUid) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}

// @Title Wear
// @Description Wear equip
// @Param	uid		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Param	equipId		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /Wear [post]
func (h *HeroController) Wear() {
	uid := h.GetString("uid")
	heroUid := h.GetString("heroUid")
	equipId := h.GetString("equipId")
	if models.Wear(uid, heroUid, equipId) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}

// @Title FloorUpHero
// @Description floor up hero
// @Param	uid		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /floorUpHero [post]
func (h *HeroController) FloorUpHero() {
	uid := h.GetString("uid")
	heroUid := h.GetString("heroUid")
	if models.HeroFloorUp(uid, heroUid) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}

// @Title LevelUpSkill
// @Description level up skill
// @Param	uid		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Param	skillUid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /levelUpSkill [post]
func (h *HeroController) LevelUpSkill() {
	uid := h.GetString("uid")
	heroUid := h.GetString("heroUid")
	skillUid := h.GetString("skillUid")
	if models.SkillLevelUp(uid, heroUid, skillUid) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}

// @Title SelectHero
// @Description select hero
// @Param	uid		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /selectHero [post]
func (h *HeroController) SelectHero() {
	uid := h.GetString("uid")
	heroUid := h.GetString("heroUid")
	if models.SelectHero(uid, heroUid) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}

// @Title UnSelectHero
// @Description level up skill
// @Param	uid		query 	string	true		"The username for login"
// @Param	heroUid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /unSelectHero [post]
func (h *HeroController) UnSelectHero() {
	uid := h.GetString("uid")
	heroUid := h.GetString("heroUid")
	if models.UnSelectHero(uid, heroUid) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}

// @Title ExchangeHero
// @Description level up skill
// @Param	uid		query 	string	true		"The username for login"
// @Param	fromHeroUid		query 	string	true		"The username for login"
// @Param	toHeroUid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /exchangeHero [post]
func (h *HeroController) ExchangeHero() {
	uid := h.GetString("uid")
	fromHeroUid := h.GetString("fromHeroUid")
	toHeroUid := h.GetString("toHeroUid")
	if models.ExchangeHero(uid, fromHeroUid, toHeroUid) {
		h.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		h.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	h.ServeJSON()
}
