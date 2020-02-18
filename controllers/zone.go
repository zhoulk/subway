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
	resZones := make([]ResponseAllZone, 0)
	for _, zone := range zones {
		resZones = append(resZones, ResponseAllZone{
			Id:       zone.Id,
			ZoneName: zone.ZoneName,
		})
	}
	z.Data["json"] = models.Response{Code: 200, Data: resZones}
	z.ServeJSON()
}

// @Title SelfZone
// @Description get self Zones
// @Success 200 {object} models.Zone
// @router /self [get]
func (z *ZoneController) SelfZone() {
	zones := models.GetSelfZones()
	resZones := make([]ResponseUserZone, 0)
	for _, zone := range zones {
		resZones = append(resZones, ResponseUserZone{
			Id:       zone.Id,
			ZoneName: zone.ZoneName,
			Name:     zone.Name,
			Level:    zone.Level,
		})
	}
	z.Data["json"] = models.Response{Code: 200, Data: zones}
	z.ServeJSON()
}

type ResponseAllZone struct {
	Id       int
	ZoneName string
}

type ResponseUserZone struct {
	Id       int
	ZoneName string
	Name     string
	Level    int
}
