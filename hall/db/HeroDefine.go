package tables

import (
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 英雄定义表
type HeroDefine struct {
	HeroId  int32  `gorm:"unique;not null"`
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
	beego.Debug("HeroDefine init")
	createHeroDefineTable()
}

func createHeroDefineTable() {
	if !context.DB().HasTable(&HeroDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func LoadHeroDefine(heroId int32) *HeroDefine {
	var heroDefine HeroDefine
	context.DB().Where("hero_id=?", heroId).First(&heroDefine)

	return &heroDefine
}
