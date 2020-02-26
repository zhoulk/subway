package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户背包
type HeroEquipInfo struct {
	HeroUid string
	Items   []*HeroEquipItemInfo
}

type HeroEquipItemInfo struct {
	Uid     string `gorm:"size:64;unique;not null"`
	HeroUid string
	EquipId int32

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&HeroEquipItemInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroEquipItemInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentHeroEquipItems(heroEquipItems []*HeroEquipItemInfo) {
	tx := context.DB().Begin()

	for _, u_b := range heroEquipItems {
		var oldHeroEquipItem HeroEquipItemInfo
		tx.Where("uid = ?", u_b.Uid).First(&oldHeroEquipItem)
		if u_b.Uid != oldHeroEquipItem.Uid {
			tx.Create(&u_b)
		} else {
			tx.Model(&u_b).Where("uid = ? ", u_b.Uid).Updates(u_b)
		}
	}

	tx.Commit()
}

func LoadHeroEquipInfo(heroUid string) *HeroEquipInfo {
	var heroEquipItems []*HeroEquipItemInfo
	context.DB().Where("hero_uid = ?", heroUid).Find(&heroEquipItems)
	if len(heroEquipItems) > 0 {
		return &HeroEquipInfo{
			HeroUid: heroUid,
			Items:   heroEquipItems,
		}
	}
	return nil
}

func LoadHeroEquipItemInfo(equipUid string) *HeroEquipItemInfo {
	var heroEquipItem HeroEquipItemInfo
	context.DB().Where("uid = ?", equipUid).First(&heroEquipItem)
	return &heroEquipItem
}
