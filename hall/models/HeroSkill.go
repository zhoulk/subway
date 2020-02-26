package models

import tables "subway/hall/db"

var (
	HeroSkillList map[string]*HeroSkill
)

func init() {
	HeroSkillList = make(map[string]*HeroSkill)
}

type HeroSkill struct {
	HeroUid string
	Skills  map[string]*SkillInfo
}

func GetHeroSkillList(heroUid string) map[string]*SkillInfo {
	if heroSkill, ok := HeroSkillList[heroUid]; ok {
		return heroSkill.Skills
	}

	t_b := tables.LoadHeroSkillInfo(heroUid)
	if t_b != nil {
		b := CreateHeroSkillFromTableHeroSkill(heroUid, t_b)
		HeroSkillList[heroUid] = b
		return HeroSkillList[heroUid].Skills
	}

	heroInfo := GetHero(heroUid)
	if heroInfo == nil {
		return nil
	}
	heroSkillDefine := GetHeroSkillDefine(heroInfo.HeroId)
	if heroSkillDefine == nil {
		return nil
	}
	heroSkill := &HeroSkill{
		HeroUid: heroUid,
		Skills:  make(map[string]*SkillInfo),
	}
	for _, v := range heroSkillDefine.Skills {
		skill := CreateASkill(v.SkillId)
		heroSkill.Skills[skill.Uid] = skill
	}
	HeroSkillList[heroUid] = heroSkill

	return heroSkill.Skills
}

func CreateHeroSkillFromTableHeroSkill(heroUid string, t_heroSkill *tables.HeroSkillInfo) *HeroSkill {
	heroSkill := &HeroSkill{
		HeroUid: heroUid,
		Skills:  make(map[string]*SkillInfo),
	}
	for _, t_item := range t_heroSkill.Items {
		heroSkill.Skills[t_item.Uid] = CreateHeroSkillItemInfoFromTableHeroSkillItemInfo(t_item)
	}
	return heroSkill
}

func CreateHeroSkillItemInfoFromTableHeroSkillItemInfo(a *tables.HeroSkillItemInfo) *SkillInfo {
	return &SkillInfo{
		Uid:     a.HeroUid,
		SkillId: a.SkillId,
		Name:    "",
	}
}
