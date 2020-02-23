package controllers

import (
	"subway/hall/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type BagController struct {
	beego.Controller
}

// @Title bagInfo
// @Description bagInfo
// @Param	roleId		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /bagInfo [post]
func (u *BagController) BagInfo() {
	roleId := u.GetString("roleId")
	bagInfo := models.GetBag(roleId)
	u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: ConvertBagInfoToResponse(bagInfo)}
	u.ServeJSON()
}

func ConvertBagInfoToResponse(a *models.BagInfo) *ResponseBagInfo {
	res := &ResponseBagInfo{
		Items: make([]*ResponseProductInfo, 0),
	}
	for _, t_item := range a.Expends {
		res.Items = append(res.Items, CreateResponseProductItemInfoFromProductInfo(t_item))
	}
	for _, t_item := range a.HeroParts {
		res.Items = append(res.Items, CreateResponseProductItemInfoFromProductInfo(t_item))
	}
	for _, t_item := range a.Equips {
		res.Items = append(res.Items, CreateResponseProductItemInfoFromProductInfo(t_item))
	}
	for _, t_item := range a.EquipParts {
		res.Items = append(res.Items, CreateResponseProductItemInfoFromProductInfo(t_item))
	}
	return res
}

type ResponseBagInfo struct {
	Items []*ResponseProductInfo
}

func CreateResponseProductItemInfoFromProductInfo(a *models.ProductInfo) *ResponseProductInfo {
	return &ResponseProductInfo{
		ProductId: a.ProductId,
		ItemId:    a.ItemId,
		Type:      a.Type,
		Count:     a.Count,
	}
}
