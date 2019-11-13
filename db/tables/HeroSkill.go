package tables

import (
	"subway/db/context"
	"github.com/jinzhu/gorm"
)

// 英雄技能表
type HeroSkill struct{
	Uid     		string `gorm:"size:64;unique;not null"`
	HeroUid  		string
	SkillId 		string 
	Level   		int32

	gorm.Model
}

func init()  {
	if !context.DB().HasTable(&HeroSkill{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroSkill{}).Error; err != nil {
			panic(err)
		}
	}
}