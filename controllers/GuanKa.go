package controllers

import (
	"subway/models"

	"github.com/astaxie/beego"
)

type GuanKaController struct {
	beego.Controller
}

// @Title GetNearGuanKa
// @Description get near GuanKa
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /getNearGuanKa [post]
func (g *GuanKaController) GetNearGuanKa() {
	uid := g.GetString("uid")
	gks := models.GetNearGuanKa(uid)
	res := make([]*models.GuanKaInfo, 0)
	if gks != nil {
		for _, gk := range gks {
			if gk != nil {
				res = append(res, &gk.Info)
			} else {
				res = append(res, nil)
			}
		}
	}
	g.Data["json"] = models.Response{Code: 200, Msg: "", Data: res}
	g.ServeJSON()
}

// @Title GetSelfCopy
// @Description get self Copy
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {object} models.Hero
// @router /getSelfCopy [post]
func (g *GuanKaController) GetSelfCopy() {
	uid := g.GetString("uid")
	allCps := models.GetAllCopy()
	_, cpDic := models.GetSelfCopy(uid)
	for _, cp := range allCps {
		if c, ok := cpDic[cp.Info.CopyId]; ok {
			cp.Status = c.Status
			cp.Star = c.Star
		}
	}
	g.Data["json"] = models.Response{Code: 200, Msg: "", Data: allCps}
	g.ServeJSON()
}

// @Title GetAllCopy
// @Description get near GuanKa
// @Success 200 {object} models.Hero
// @router /getAllCopy [post]
func (g *GuanKaController) GetAllCopy() {
	cps := models.GetAllCopy()
	g.Data["json"] = models.Response{Code: 200, Msg: "", Data: cps}
	g.ServeJSON()
}

// @Title GetCopyItems
// @Description get copy items
// @Param	uid		query 	string	true		"The username for login"
// @Param	copyId		query 	int	true
// @Success 200 {object} models.Hero
// @router /getCopyItems [post]
func (g *GuanKaController) GetCopyItems() {
	uid := g.GetString("uid")
	copyId, _ := g.GetInt("copyId")
	copyItems := models.GetCopyItems(copyId)
	_, cpItemDic := models.GetSelfCopyItem(uid)
	for _, cpItem := range copyItems {
		if c, ok := cpItemDic[cpItem.CopyItemId]; ok {
			cpItem.Status = c.Status
			cpItem.Star = c.Star
		}
	}
	g.Data["json"] = models.Response{Code: 200, Msg: "", Data: copyItems}
	g.ServeJSON()
}
