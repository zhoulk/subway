package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 英雄装备定义表
type HeroEquipDefine struct {
	HeroId  string
	Floor   int16 // 阶别
	EquipId string
}

func init() {
	beego.Debug("HeroEquipDefine init")
	if beego.AppConfig.DefaultBool("updateConfigData", true) {
		createHeroEquipDefineTable()
		initHeroEquipData()
	}
}

func createHeroEquipDefineTable() {
	if !context.DB().HasTable(&HeroEquipDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroEquipDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initHeroEquipData() {

	data, err := ioutil.ReadFile("./static/data/heroEquip.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var heroEquips []HeroEquipDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &heroEquips)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&HeroEquipDefine{})

	tx := context.DB().Begin()
	for _, h := range heroEquips {
		tx.Create(h)
	}
	tx.Commit()
}

func LoadHeroEquipDefine() []*HeroEquipDefine {
	var heroEquips []*HeroEquipDefine
	context.DB().Where("1=1").Find(&heroEquips)

	return heroEquips
}
