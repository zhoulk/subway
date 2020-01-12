package controllers

import (
	"subway/models"

	"github.com/astaxie/beego"
)

type MapController struct {
	beego.Controller
}

// @Title RandomAPath
// @Description random a path
// @Param	uid		query 	string	true		"The username for login"
// @Param	from		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /randomAPath [post]
func (m *MapController) RandomAPath() {
	uid := m.GetString("uid")
	originNode, _ := m.GetInt("from")

	step, res := models.RandomAPath(uid, int32(originNode))

	m.Data["json"] = models.Response{Code: 200, Msg: "", Data: RandomAPathData{
		Step: step,
		Path: res,
	}}

	m.ServeJSON()
}

type RandomAPathData struct {
	Step int32
	Path []models.MapItem
}
