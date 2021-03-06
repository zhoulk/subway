package controllers

import (
	"subway/gate/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title PreLogin
// @Description Logs user into the system
// @Param	openId		query 	string	true		"The username for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /preLogin [post]
func (u *UserController) PreLogin() {
	openId := u.GetString("openId")
	if account := models.GetAccount(openId); account != nil {
		u.Data["json"] = models.Response{Code: 200, Msg: "login success", Data: ResponsePreLogin{
			AccountId: account.AccountId,
		}}
	} else {
		if account := models.AddAccount(openId); account != nil {
			u.Data["json"] = models.Response{Code: 200, Msg: "login success", Data: ResponsePreLogin{
				AccountId: account.AccountId,
			}}
		} else {
			u.Data["json"] = models.Response{Code: 201, Msg: "login fail", Data: nil}
		}
	}
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	zoneId		query 	int	true		"The username for login"
// @Param	openId		query 	string	true		"The username for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	zoneId, _ := u.GetInt("zoneId")
	openId := u.GetString("openId")
	if role := models.Login(zoneId, openId); role != nil {
		u.Data["json"] = models.Response{Code: 200, Msg: "login success", Data: ResponseLogin{
			RoleId: role.RoleId,
			Ip:     "127.0.0.1",
			Port:   10086,
		}}
	} else {
		if role := models.AddRole(zoneId, openId); role != nil {
			u.Data["json"] = models.Response{Code: 200, Msg: "login success", Data: ResponseLogin{
				RoleId: role.RoleId,
				Ip:     "127.0.0.1",
				Port:   10086,
			}}
		} else {
			u.Data["json"] = models.Response{Code: 201, Msg: "login fail", Data: nil}
		}
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}

// @Title userInfo
// @Description userInfo
// @Param	uid		query 	string	true		"The username for login"
// @Success 200 {string}
// @router /userInfo [post]
func (u *UserController) UserInfo() {
	// uid := u.GetString("uid")
	// user, _ := models.GetUser(uid)
	// u.Data["json"] = models.Response{Code: 200, Msg: "login success", Data: user.Profile}
	// u.ServeJSON()
}

type ResponsePreLogin struct {
	AccountId string
}

type ResponseLogin struct {
	RoleId string
	Ip     string
	Port   int
}
