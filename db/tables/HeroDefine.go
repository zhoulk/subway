package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 英雄定义表
type HeroDefine struct {
	HeroId  string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	Type    int8
	AtkType int8
	Level   int32
	Floor   int16 // 阶别
	Star    int16 // 星星

	HP              int32
	MP              int32
	AD              int32
	AP              int32
	ADDef           int32
	APDef           int32
	SPD             int32 // 毫秒数
	Agility         int32
	Intelligent     int32
	Strength        int32
	ADCrit          int32 // 物理暴击
	StrengthGrow    int32
	AgilityGrow     int32
	IntelligentGrow int32

	Desc string
}

func init() {
	createHeroDefineTable()
	initHeroData()
}

func createHeroDefineTable() {
	if !context.DB().HasTable(&HeroDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initHeroData() {

	data, err := ioutil.ReadFile("./static/data/hero.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var heros []HeroDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &heros)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&HeroDefine{})

	tx := context.DB().Begin()
	for _, h := range heros {
		tx.Create(h)
	}
	tx.Commit()
}

func LoadHeroDefine() []*HeroDefine {
	var heros []*HeroDefine
	context.DB().Where("1=1").Find(&heros)

	return heros
}
