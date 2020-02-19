package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["subway/hall/controllers:RoleController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:RoleController"],
        beego.ControllerComments{
            Method: "RoleInfo",
            Router: `/roleInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:RoleController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:RoleController"],
        beego.ControllerComments{
            Method: "UpdateRoleInfo",
            Router: `/updateRoleInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
