package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 英雄装备表
type HeroEquip struct {
	Uid     string `gorm:"size:64;unique;not null"`
	HeroUid string
	EquipId string
	Status  int8 // 0 未装备  1 已装备  2 已使用

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&HeroEquip{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroEquip{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentHeroEquip(heroEquips []*HeroEquip) {
	tx := context.DB().Begin()

	for _, h_e := range heroEquips {
		var oldHeroEquip HeroEquip
		tx.Where("uid = ?", h_e.Uid).First(&oldHeroEquip)
		if h_e.Uid != oldHeroEquip.Uid {
			tx.Create(&h_e)
		} else {
			tx.Model(&h_e).Where("uid = ? ", h_e.Uid).Updates(h_e)
		}
	}

	tx.Commit()
}

func LoadHeroEquips(heroUid string) []*HeroEquip {
	var heroEquips []*HeroEquip
	context.DB().Where("hero_uid = ?", heroUid).Find(&heroEquips)

	return heroEquips
}
