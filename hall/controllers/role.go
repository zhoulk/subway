package controllers

import (
	"subway/hall/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type RoleController struct {
	beego.Controller
}

// @Title roleInfo
// @Description roleInfo
// @Param	roleId		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /roleInfo [post]
func (u *RoleController) RoleInfo() {
	roleId := u.GetString("roleId")
	roleInfo := models.GetRoleInfo(roleId)
	if roleInfo == nil {
		roleInfo = models.AddRoleInfo(&models.RoleInfo{
			Name:    "",
			HeadUrl: "",
			Gold:    0,
			Diamond: 0,
			Power:   0,
			Level:   1,
			Exp:     0,
			Point:   60,
		})
	}
	u.Data["json"] = models.Response{Code: 200, Msg: "login success", Data: ResponseRoleInfo{
		Name:       roleInfo.Name,
		HeadUrl:    roleInfo.HeadUrl,
		Gold:       roleInfo.Gold,
		Diamond:    roleInfo.Diamond,
		Power:      roleInfo.Power,
		Level:      roleInfo.Level,
		LevelUpExp: 100,
		Exp:        roleInfo.Exp,
		Point:      roleInfo.Point,
	}}
	u.ServeJSON()
}

// @Title updateRoleInfo
// @Description updateRoleInfo
// @Param	roleId		query 	string true		"The username for login"
// @Param	name		query 	string true		"The username for login"
// @Success 200 {string}
// @router /updateRoleInfo [post]
func (u *RoleController) UpdateRoleInfo() {
	roleId := u.GetString("roleId")
	name := u.GetString("name")
	roleInfo := models.GetRoleInfo(roleId)
	if roleInfo != nil {
		roleInfo.RoleId = roleId
		roleInfo.Name = name
		roleInfo = models.AddRoleInfo(roleInfo)

		u.Data["json"] = models.Response{Code: 200, Msg: "update success", Data: ResponseRoleInfo{
			Name:       roleInfo.Name,
			HeadUrl:    roleInfo.HeadUrl,
			Gold:       roleInfo.Gold,
			Diamond:    roleInfo.Diamond,
			Power:      roleInfo.Power,
			Level:      roleInfo.Level,
			LevelUpExp: 100,
			Exp:        roleInfo.Exp,
			Point:      roleInfo.Point,
		}}
	} else {
		u.Data["json"] = models.Response{Code: 201, Msg: "update fail", Data: nil}
	}

	u.ServeJSON()
}

type ResponseRoleInfo struct {
	Name       string
	HeadUrl    string
	Gold       int64
	Diamond    int64
	Power      int64
	Level      int32
	Exp        int64
	LevelUpExp int64
	Point      int32
}
