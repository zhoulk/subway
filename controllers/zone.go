package controllers

import (
	"subway/models"
	"github.com/astaxie/beego"
)

// Operations about Zones
type ZoneController struct {
	beego.Controller
}

// @Title AllZone
// @Description get all Zones
// @Success 200 {object} models.Zone
// @router / [get]
func (z *ZoneController) AllZone() {
	zones := models.GetAllZones()
	z.Data["json"] = zones
	z.ServeJSON()
}
