package models

import (
	"subway/tool"
)

var (
	SkillDefineList map[string]*Skill
)

func init() {
	SkillDefineList = make(map[string]*Skill)
	s1001 := &Skill{
		Info:   SkillInfo{SkillId: "1001", Name: "能量虚空", Level: 1, Desc: "敌法闪烁到智力最高的敌人身后，以目标为中心施放能量虚空，造成巨大的魔法伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 1800, StepGold: 500, StepGold2: 100}}
	s1002 := &Skill{
		Info:   SkillInfo{SkillId: "1002", Name: "能量燃烧", Level: 1, Desc: "敌法闪烁到智力最高的敌人身边攻击，造成小范围物理伤害并额外损毁敌人的能量。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 600, StepGold: 200, StepGold2: 2}}
	s1003 := &Skill{
		Info:   SkillInfo{SkillId: "1003", Name: "战刃技巧", Level: 1, Desc: "敌法加速旋转他的刀刃，增加他的敏捷。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 700, StepGold: 200, StepGold2: 2}}
	s1004 := &Skill{
		Info:   SkillInfo{SkillId: "1004", Name: "魔法盾", Level: 1, Desc: "敌法利用对魔法能量的理解减少受到的伤害，增加魔抗。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 800, StepGold: 200, StepGold2: 2}}
	SkillDefineList[s1001.Info.SkillId] = s1001
	SkillDefineList[s1002.Info.SkillId] = s1002
	SkillDefineList[s1003.Info.SkillId] = s1003
	SkillDefineList[s1004.Info.SkillId] = s1004

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
	Type    int8 // 1 主动  2  被动
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
