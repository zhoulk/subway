package tables

import (
	"subway/db/context"
)

// 英雄技能定义表
type HeroSkillDefine struct {
	HeroId  int32
	SkillId int32
}

func init() {
	createHeroSkillDefineTable()
}

func createHeroSkillDefineTable() {
	if !context.DB().HasTable(&HeroSkillDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroSkillDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func LoadHeroSkillDefine(heroId int32) []*HeroSkillDefine {
	var heroSkills []*HeroSkillDefine
	context.DB().Where("hero_id=?", heroId).Find(&heroSkills)

	return heroSkills
}
