package tables

import (
	"subway/db/context"
	"github.com/jinzhu/gorm"
)

// 英雄装备表
type HeroEquip struct{
	Uid     		string `gorm:"size:64;unique;not null"`
	HeroUid  		string
	EquipId 		string 
	Status   		int8  // 0 未装备  1 已装备  2 已使用

	gorm.Model
}

func init()  {
	if !context.DB().HasTable(&HeroEquip{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroEquip{}).Error; err != nil {
			panic(err)
		}
	}
}