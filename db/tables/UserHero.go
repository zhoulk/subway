package tables

import (
	"subway/db/context"
	"github.com/jinzhu/gorm"
)

// 用户英雄表
type UserHero struct{
	Uid     		string `gorm:"size:64;unique;not null"`
	UserId  		string
	HeroId      	string `gorm:"size:64;unique;not null"`
	Level       	int32
	Floor       	int16 // 阶别
	Star        	int16 // 星星
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

	gorm.Model
}

func init()  {
	if !context.DB().HasTable(&UserHero{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserHero{}).Error; err != nil {
			panic(err)
		}
	}
}