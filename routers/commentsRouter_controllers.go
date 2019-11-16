package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["subway/controllers:BagController"] = append(beego.GlobalControllerRouter["subway/controllers:BagController"],
        beego.ControllerComments{
            Method: "GetSelfBag",
            Router: `/getSelfBag`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:BattleController"] = append(beego.GlobalControllerRouter["subway/controllers:BattleController"],
        beego.ControllerComments{
            Method: "BattleGK",
            Router: `/battleGK`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:GuanKaController"] = append(beego.GlobalControllerRouter["subway/controllers:GuanKaController"],
        beego.ControllerComments{
            Method: "GetAllCopy",
            Router: `/getAllCopy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:GuanKaController"] = append(beego.GlobalControllerRouter["subway/controllers:GuanKaController"],
        beego.ControllerComments{
            Method: "GetNearGuanKa",
            Router: `/getNearGuanKa`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:GuanKaController"] = append(beego.GlobalControllerRouter["subway/controllers:GuanKaController"],
        beego.ControllerComments{
            Method: "GetSelfCopy",
            Router: `/getSelfCopy`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "Wear",
            Router: `/Wear`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "AllHeros",
            Router: `/all`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "ExchangeHero",
            Router: `/exchangeHero`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "FloorUpHero",
            Router: `/floorUpHero`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "HeroDetail",
            Router: `/heroDetail`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "LevelUpHero",
            Router: `/levelUp`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "LevelUpSkill",
            Router: `/levelUpSkill`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "SelectHero",
            Router: `/selectHero`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "SelfHeros",
            Router: `/self`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:HeroController"] = append(beego.GlobalControllerRouter["subway/controllers:HeroController"],
        beego.ControllerComments{
            Method: "UnSelectHero",
            Router: `/unSelectHero`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:TechController"] = append(beego.GlobalControllerRouter["subway/controllers:TechController"],
        beego.ControllerComments{
            Method: "GainFirstHero",
            Router: `/gainFirstHero`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:UserController"] = append(beego.GlobalControllerRouter["subway/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:UserController"] = append(beego.GlobalControllerRouter["subway/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["subway/controllers:ZoneController"] = append(beego.GlobalControllerRouter["subway/controllers:ZoneController"],
        beego.ControllerComments{
            Method: "AllZone",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
