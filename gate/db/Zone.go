package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 大区信息
type Zone struct {
	ZoneId int    `gorm:"unique;not null"`
	Name   string `gorm:"size:64"`
	Status int
}

func init() {
	beego.Debug("Zone init")
	createZoneTable()
	initZoneData()
}

func createZoneTable() {
	if !context.DB().HasTable(&Zone{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Zone{}).Error; err != nil {
			panic(err)
		}
	}
}

func initZoneData() {

	data, err := ioutil.ReadFile("./static/data/zone.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var zones []Zone

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &zones)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&Zone{})

	tx := context.DB().Begin()
	for _, h := range zones {
		tx.Create(h)
	}
	tx.Commit()
}
