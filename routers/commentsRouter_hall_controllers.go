package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["subway/hall/controllers:BagController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:BagController"],
        beego.ControllerComments{
            Method: "BagInfo",
            Router: `/bagInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:EquipController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:EquipController"],
        beego.ControllerComments{
            Method: "EquipDetail",
            Router: `/equipDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:HeroController"],
        beego.ControllerComments{
            Method: "HeroDetail",
            Router: `/heroDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:HeroController"],
        beego.ControllerComments{
            Method: "HeroList",
            Router: `/heroList`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:HouseController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:HouseController"],
        beego.ControllerComments{
            Method: "DiamondRandom",
            Router: `/diamondRandom`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:HouseController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:HouseController"],
        beego.ControllerComments{
            Method: "GoldRandom",
            Router: `/goldRandom`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/hall/controllers:HouseController"] = append(beego.GlobalControllerRouter["subway/hall/controllers:HouseController"],
        beego.ControllerComments{
            Method: "HouseInfo",
            Router: `/houseInfo`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

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
