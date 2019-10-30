package controllers

import (
	"subway/battle"
	"subway/models"

	"github.com/astaxie/beego"
)

type BattleController struct {
	beego.Controller
}

// @Title BattleGK
// @Description get near GuanKa
// @Param	uid		query 	string	true		"The username for login"
// @Param	guankaId		query 	int	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /battleGK [post]
func (b *BattleController) BattleGK() {
	uid := b.GetString("uid")
	guankaId, _ := b.GetInt("guankaId")
	res := battle.BattleGuanKa(uid, guankaId)
	b.Data["json"] = models.Response{Code: 200, Msg: "", Data: res}
	b.ServeJSON()
}
