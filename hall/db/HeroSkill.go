package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

//
type HeroSkillInfo struct {
	HeroUid string
	Items   []*HeroSkillItemInfo
}

type HeroSkillItemInfo struct {
	Uid     string `gorm:"size:64;unique;not null"`
	HeroUid string
	SkillId int32

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&HeroSkillItemInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroSkillItemInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentHeroSkillItems(heroSkillItems []*HeroSkillItemInfo) {
	tx := context.DB().Begin()

	for _, u_b := range heroSkillItems {
		var oldHeroSkillItem HeroSkillItemInfo
		tx.Where("uid = ?", u_b.Uid).First(&oldHeroSkillItem)
		if u_b.Uid != oldHeroSkillItem.Uid {
			tx.Create(&u_b)
		} else {
			tx.Model(&u_b).Where("uid = ? ", u_b.Uid).Updates(u_b)
		}
	}

	tx.Commit()
}

func LoadHeroSkillInfo(heroUid string) *HeroSkillInfo {
	var heroSkillItems []*HeroSkillItemInfo
	context.DB().Where("hero_uid = ?", heroUid).Find(&heroSkillItems)
	if len(heroSkillItems) > 0 {
		return &HeroSkillInfo{
			HeroUid: heroUid,
			Items:   heroSkillItems,
		}
	}
	return nil
}
