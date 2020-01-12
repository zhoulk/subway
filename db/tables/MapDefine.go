package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

//
type MapDefine struct {
	Id   int32
	Name string
	Next int32
}

func init() {
	beego.Debug("MapDefine init")
	if beego.AppConfig.DefaultBool("updateConfigData", true) {
		createMapDefineTable()
		initMapData()
	}
}

func createMapDefineTable() {
	if !context.DB().HasTable(&MapDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&MapDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initMapData() {

	data, err := ioutil.ReadFile("./static/data/map/beijing.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var maps []MapDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &maps)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&MapDefine{})

	tx := context.DB().Begin()
	for _, m := range maps {
		tx.Create(m)
	}
	tx.Commit()
}

func LoadMapDefine() []*MapDefine {
	var maps []*MapDefine
	context.DB().Where("1=1").Find(&maps)

	return maps
}
