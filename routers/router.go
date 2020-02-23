// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"subway/controllers"

	gate "subway/gate/controllers"
	hall "subway/hall/controllers"

	"github.com/beego/bee/plugins/cors"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/subway",
		beego.NSNamespace("/gate",
			beego.NSNamespace("/user",
				beego.NSInclude(
					&gate.UserController{},
				),
			), beego.NSNamespace("/zone",
				beego.NSInclude(
					&gate.ZoneController{},
				),
			),
		),
		beego.NSNamespace("/hall",
			beego.NSNamespace("/role",
				beego.NSInclude(
					&hall.RoleController{},
				),
			),
			beego.NSNamespace("/house",
				beego.NSInclude(
					&hall.HouseController{},
				),
			),
			beego.NSNamespace("/bag",
				beego.NSInclude(
					&hall.BagController{},
				),
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
		beego.NSNamespace("/battle",
			beego.NSInclude(
				&controllers.BattleController{},
			),
		),
		beego.NSNamespace("/map",
			beego.NSInclude(
				&controllers.MapController{},
			),
		),
	)
	beego.AddNamespace(ns)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowOrigins:     []string{"https://192.168.0.102", "http://localhost:7456"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"token", "key", "Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}
