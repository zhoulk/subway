package controllers

import (
	"subway/models"

	"github.com/astaxie/beego"
)

type BagController struct {
	beego.Controller
}

// @Title GetSelfBag
// @Description get near GuanKa
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /getSelfBag [post]
func (b *BagController) GetSelfBag() {
	uid := b.GetString("uid")
	bag := models.GetBag(uid)
	b.Data["json"] = models.Response{Code: 200, Msg: "", Data: bag.Items}
	b.ServeJSON()
}
