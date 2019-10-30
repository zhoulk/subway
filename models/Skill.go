package models

import (
	"subway/tool"
)

var (
	SkillDefineList map[string]*Skill
)

func init() {
	SkillDefineList = make(map[string]*Skill)
	s := &Skill{
		Info:   SkillInfo{SkillId: "1000", Name: "闪烁", Level: 0, Desc: ""},
		Secret: SkillSecretInfo{OriginLevelUpGold: 800, StepGold: 200, StepGold2: 2}}
	SkillDefineList[s.Info.SkillId] = s
}

type Skill struct {
	Uid    string
	Info   SkillInfo
	Secret SkillSecretInfo
}

type SkillInfo struct {
	SkillId string
	Name    string
	Level   int32
	Desc    string
}

type SkillSecretInfo struct {
	OriginLevelUpGold int32
	StepGold          int32
	StepGold2         int32
}

func GetSkills(skillIds []string) []*Skill {
	skills := make([]*Skill, 0)
	for _, skillId := range skillIds {
		if s, ok := SkillDefineList[skillId]; ok {
			s.Uid = tool.UniqueId()
			skills = append(skills, s)
		}
	}
	return skills
}

func SkillLevelUp(uid string, heroUid string, skillUid string) bool {
	var targetHero *Hero
	var targetSkill *Skill

	u, _ := GetUser(uid)
	if u != nil {
		heros := u.Heros
		if heros != nil {
			for _, h := range heros {
				if h.Uid == heroUid {
					targetHero = h
					break
				}
			}
		}
	}

	if targetHero != nil {
		skills := targetHero.Skills
		if skills != nil {
			for _, s := range skills {
				if s.Uid == skillUid {
					targetSkill = s
					break
				}
			}
		}
	}

	if targetSkill != nil {
		endLevel := targetSkill.Info.Level + 1
		// 不超过英雄等级 * 2
		if targetHero.Info.Level*2 >= endLevel {
			// 计算升级需要的金币
			needGold := int64(targetSkill.Secret.OriginLevelUpGold + targetSkill.Secret.StepGold*endLevel + endLevel*endLevel)
			if u.Profile.Gold >= needGold {
				targetSkill.Info.Level = endLevel
				u.Profile.Gold -= needGold
				return true
			}
		}
	}

	return false
}
