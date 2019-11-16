package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 英雄技能表
type HeroSkill struct {
	Uid     string `gorm:"size:64;unique;not null"`
	HeroUid string
	SkillId string
	Level   int32

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&HeroSkill{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroSkill{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentHeroSkill(heroSkills []*HeroSkill) {
	tx := context.DB().Begin()

	for _, h_s := range heroSkills {
		var oldHeroSkill HeroSkill
		tx.Where("uid = ?", h_s.Uid).First(&oldHeroSkill)
		if h_s.Uid != oldHeroSkill.Uid {
			tx.Create(&h_s)
		} else {
			tx.Model(&h_s).Where("uid = ? ", h_s.Uid).Updates(h_s)
		}
	}

	tx.Commit()
}

func LoadHeroSkills(heroUid string) []*HeroSkill {
	var heroSkills []*HeroSkill
	context.DB().Where("hero_uid = ?", heroUid).Find(&heroSkills)

	return heroSkills
}
