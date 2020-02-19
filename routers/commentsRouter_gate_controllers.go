package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["subway/gate/controllers:UserController"] = append(beego.GlobalControllerRouter["subway/gate/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/gate/controllers:UserController"] = append(beego.GlobalControllerRouter["subway/gate/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/gate/controllers:UserController"] = append(beego.GlobalControllerRouter["subway/gate/controllers:UserController"],
        beego.ControllerComments{
            Method: "PreLogin",
            Router: `/preLogin`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/gate/controllers:UserController"] = append(beego.GlobalControllerRouter["subway/gate/controllers:UserController"],
        beego.ControllerComments{
            Method: "UserInfo",
            Router: `/userInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/gate/controllers:ZoneController"] = append(beego.GlobalControllerRouter["subway/gate/controllers:ZoneController"],
        beego.ControllerComments{
            Method: "AllZone",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/gate/controllers:ZoneController"] = append(beego.GlobalControllerRouter["subway/gate/controllers:ZoneController"],
        beego.ControllerComments{
            Method: "SelfZone",
            Router: `/self`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
