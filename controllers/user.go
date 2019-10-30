package controllers

import (
	"subway/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description Logs user into the system
// @Param	openId		query 	string	true		"The username for login"
// @Param	userName		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	openId := u.GetString("openId")
	userName := u.GetString("userName")
	if user:= models.Login(openId, userName); user != nil {
		u.Data["json"] = models.Response{Code:200, Msg:"login success", Data:user.Info}
	} else {
		if user :=models.AddUser(openId, userName); user != nil {
			u.Data["json"] = models.Response{Code:200, Msg:"login success", Data:user.Info}
		}else {
			u.Data["json"] = models.Response{Code:201, Msg:"login fail", Data:nil}
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

