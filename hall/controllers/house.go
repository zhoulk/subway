package controllers

import (
	"subway/hall/models"
	"time"

	"github.com/astaxie/beego"
)

// Operations about Users
type HouseController struct {
	beego.Controller
}

// @Title houseInfo
// @Description houseInfo
// @Param	roleId		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /houseInfo [post]
func (u *HouseController) HouseInfo() {
	roleId := u.GetString("roleId")
	houseInfo := models.GetHouseInfo(roleId)
	if houseInfo == nil {
		houseInfo = models.AddHouseInfo(&models.HouseInfo{
			RoleId:            roleId,
			GoldTimes:         5,
			TotalGoldTimes:    5,
			DiamondTimes:      1,
			TotalDiamondTimes: 1,
		})
	}
	lastGoldStr := houseInfo.LastGoldTime.Format("2006-01-02")
	nowGoldStr := time.Now().Format("2006-01-02")
	goldSeconds := int32(0)
	if lastGoldStr == nowGoldStr {
		nextTime := houseInfo.LastGoldTime.Add(5 * time.Minute)
		goldSeconds = int32(nextTime.Unix() - time.Now().Unix())
		if goldSeconds < 0 {
			goldSeconds = 0
		}
	}

	lastDiamondStr := houseInfo.LastDiamondTime.Format("2006-01-02")
	nowDiamondStr := time.Now().Format("2006-01-02")
	diamondSeconds := int32(0)
	if lastDiamondStr == nowDiamondStr {
		nextTime, _ := time.ParseInLocation("2006-01-02", nowDiamondStr, time.Local)
		nextTime = nextTime.Add(24 * time.Hour)
		beego.Debug(nextTime, time.Now(), nextTime.Sub(time.Now()))
		diamondSeconds = int32(nextTime.Unix() - time.Now().Unix())
	}
	u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: ResponseHouseInfo{
		GoldTimes:          houseInfo.GoldTimes,
		TotalGoldTimes:     houseInfo.TotalGoldTimes,
		DiamondTimes:       houseInfo.DiamondTimes,
		TotalDiamondTimes:  houseInfo.TotalDiamondTimes,
		NextGoldSeconds:    goldSeconds,
		NextDiamondSeconds: diamondSeconds,
	}}
	u.ServeJSON()
}

// @Title goldRandom
// @Description goldRandom
// @Param	roleId		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /goldRandom [post]
func (u *HouseController) GoldRandom() {
	roleId := u.GetString("roleId")

	productInfo := models.GoldRandom(roleId)

	if productInfo == nil {
		u.Data["json"] = models.Response{Code: 201, Msg: "fail", Data: nil}
	} else {
		u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: ResponseProductInfo{
			ProductId: productInfo.ProductId,
			ItemId:    productInfo.ItemId,
			Type:      productInfo.Type,
			Name:      productInfo.Name,
		}}
	}

	u.ServeJSON()
}

// @Title diamondRandom
// @Description diamondRandom
// @Param	roleId		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /diamondRandom [post]
func (u *HouseController) DiamondRandom() {
	roleId := u.GetString("roleId")

	productInfo := models.DiamondRandom(roleId)

	if productInfo == nil {
		u.Data["json"] = models.Response{Code: 201, Msg: "fail", Data: nil}
	} else {
		u.Data["json"] = models.Response{Code: 200, Msg: "success", Data: ResponseProductInfo{
			ProductId: productInfo.ProductId,
			ItemId:    productInfo.ItemId,
			Type:      productInfo.Type,
			Name:      productInfo.Name,
			Count:     productInfo.Count,
		}}
	}

	u.ServeJSON()
}

type ResponseHouseInfo struct {
	GoldTimes          int32
	TotalGoldTimes     int32
	DiamondTimes       int32
	TotalDiamondTimes  int32
	NextGoldSeconds    int32
	NextDiamondSeconds int32
}

type ResponseProductInfo struct {
	ProductId string
	ItemId    int32
	Type      int8
	Name      string
	Count     int32
}
