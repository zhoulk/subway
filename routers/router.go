// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/beego/bee/plugins/cors"
	"subway/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/subway",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/zone",
			beego.NSInclude(
				&controllers.ZoneController{},
			),
		),
		beego.NSNamespace("/hero",
			beego.NSInclude(
				&controllers.HeroController{},
			),
		),
		beego.NSNamespace("/tech",
			beego.NSInclude(
				&controllers.TechController{},
			),
		),
		beego.NSNamespace("/gk",
			beego.NSInclude(
				&controllers.GuanKaController{},
			),
		),
		beego.NSNamespace("/bag",
			beego.NSInclude(
				&controllers.BagController{},
			),
		),
		beego.NSNamespace("/battle",
			beego.NSInclude(
				&controllers.BattleController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:true,
		AllowOrigins:[]string{"https://192.168.0.102", "http://localhost:7456"},
		AllowMethods:[]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:[]string{"token", "key", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:[]string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials:true,
		}))
}
