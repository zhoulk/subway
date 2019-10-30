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
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /all [post]
func (h *HeroController) AllHeros() {
	uid := h.GetString("uid")
	beego.Debug(uid)
	heros := models.GetAllHeros(uid)
	h.Data["json"] = heros
	h.ServeJSON()
}
