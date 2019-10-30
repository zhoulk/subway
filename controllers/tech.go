package controllers

import (
	"subway/models"

	"github.com/astaxie/beego"
)

type TechController struct {
	beego.Controller
}

// @Title GainFirstHero
// @Description get all Heros
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /gainFirstHero [post]
func (t *TechController) GainFirstHero() {
	uid := t.GetString("uid")
	if models.AddHero(uid, "1000") {
		t.Data["json"] = models.Response{Code: 200, Msg: "", Data: nil}
	} else {
		t.Data["json"] = models.Response{Code: 201, Msg: "", Data: nil}
	}
	t.ServeJSON()
}
