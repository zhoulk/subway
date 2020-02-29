package models

import (
	tables "subway/hall/db"
	"subway/tool"
)

var (
	SkillList map[string]*SkillInfo
)

func init() {
	SkillList = make(map[string]*SkillInfo)
}

type SkillInfo struct {
	Uid     string
	SkillId int32
	Name    string
	Level   int32
}

func CreateASkill(skillId int32) *SkillInfo {
	skill := &SkillInfo{
		Uid:     tool.UniqueId(),
		SkillId: skillId,
		Name:    "",
	}
	AddSkill(skill)
	return skill
}

func AddSkill(skill *SkillInfo) {
	SkillList[skill.Uid] = skill
}

func GetSkill(SkillUid string) *SkillInfo {
	if Skill, ok := SkillList[SkillUid]; ok {
		return Skill
	}

	t_b := tables.LoadHeroSkillItemInfo(SkillUid)
	if t_b != nil {
		Skill := CreateHeroSkillItemInfoFromTableHeroSkillItemInfo(t_b)
		SkillList[Skill.Uid] = Skill
		return Skill
	}

	return nil
}

func PersistentSkillInfo() {
	heroSkillInfos := make([]*tables.HeroSkillItemInfo, 0)
	for _, a := range SkillList {
		heroSkillInfos = append(heroSkillInfos, CreateTableHeroSkillItemInfoFromHeroSkillItemInfo(a))
	}
	tables.PersistentHeroSkillItems(heroSkillInfos)
}

func CreateTableHeroSkillItemInfoFromHeroSkillItemInfo(a *SkillInfo) *tables.HeroSkillItemInfo {
	return &tables.HeroSkillItemInfo{
		Uid:     a.Uid,
		SkillId: a.SkillId,
	}
}
