package models

import (
	tables "subway/hall/db"
)

var (
	HeroSkillDefineList map[int32]*HeroSkillDefine
)

func init() {
	HeroSkillDefineList = make(map[int32]*HeroSkillDefine)
}

type HeroSkillDefine struct {
	HeroId int32
	Skills []*SkillInfo
}

func GetHeroSkillDefine(heroId int32) *HeroSkillDefine {
	if heroSkillDefine, ok := HeroSkillDefineList[heroId]; ok {
		return heroSkillDefine
	}

	t_bs := tables.LoadHeroSkillDefine(heroId)
	if t_bs != nil {
		HeroSkillDefineList[heroId] = &HeroSkillDefine{
			HeroId: heroId,
			Skills: make([]*SkillInfo, 0),
		}
		for _, t_b := range t_bs {
			HeroSkillDefineList[heroId].Skills = append(HeroSkillDefineList[heroId].Skills, CreateSkillInfoFromTableHeroSkillDefine(t_b))
		}
		return HeroSkillDefineList[heroId]
	}

	return nil
}

func CreateSkillInfoFromTableHeroSkillDefine(a *tables.HeroSkillDefine) *SkillInfo {
	res := &SkillInfo{
		SkillId: a.SkillId,
	}
	return res
}
