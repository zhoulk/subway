package main

import (
	_ "subway/db"
	_ "subway/gate/timer"
	_ "subway/hall/timer"
	_ "subway/routers"
	_ "subway/timer"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
