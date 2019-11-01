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

	s1011 := &Skill{
		Info:   SkillInfo{SkillId: "1011", Name: "箭雨", Level: 1, Desc: "对敌人连续射出多只箭矢，造成成吨的伤害。技能等级越高，射出的箭矢数量越多。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 1800, StepGold: 500, StepGold2: 100}}
	s1012 := &Skill{
		Info:   SkillInfo{SkillId: "1012", Name: "冰箭", Level: 1, Desc: "射出一只冰箭，对目标造成物理伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 600, StepGold: 200, StepGold2: 2}}
	s1013 := &Skill{
		Info:   SkillInfo{SkillId: "1013", Name: "沉默", Level: 1, Desc: "使数个敌人陷入沉默，无法释放任何带魔法伤害的技能。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 700, StepGold: 200, StepGold2: 2}}
	s1014 := &Skill{
		Info:   SkillInfo{SkillId: "1014", Name: "射手天赋", Level: 1, Desc: "增加全体队友攻击力。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 800, StepGold: 200, StepGold2: 2}}
	SkillDefineList[s1011.Info.SkillId] = s1011
	SkillDefineList[s1012.Info.SkillId] = s1012
	SkillDefineList[s1013.Info.SkillId] = s1013
	SkillDefineList[s1014.Info.SkillId] = s1014

	s1021 := &Skill{
		Info:   SkillInfo{SkillId: "1021", Name: "幽灵船", Level: 1, Desc: "召唤幽灵船冲向对方，造成大范围的眩晕和魔法伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 1800, StepGold: 500, StepGold2: 100}}
	s1022 := &Skill{
		Info:   SkillInfo{SkillId: "1022", Name: "洪流", Level: 1, Desc: "用水流击飞一个随机敌人，造成魔法伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 600, StepGold: 200, StepGold2: 2}}
	s1023 := &Skill{
		Info:   SkillInfo{SkillId: "1023", Name: "水刀", Level: 1, Desc: "对自身小范围内的敌人造成魔法伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 700, StepGold: 200, StepGold2: 2}}
	s1024 := &Skill{
		Info:   SkillInfo{SkillId: "1024", Name: "力量强化", Level: 1, Desc: "力量强化船长专注地磨练自己的身体，增加力量。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 800, StepGold: 200, StepGold2: 2}}
	SkillDefineList[s1021.Info.SkillId] = s1021
	SkillDefineList[s1022.Info.SkillId] = s1022
	SkillDefineList[s1023.Info.SkillId] = s1023
	SkillDefineList[s1024.Info.SkillId] = s1024

	s1031 := &Skill{
		Info:   SkillInfo{SkillId: "1021", Name: "神灭斩", Level: 1, Desc: "对目标射出一道闪电，造成巨量的魔法伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 1800, StepGold: 500, StepGold2: 100}}
	s1032 := &Skill{
		Info:   SkillInfo{SkillId: "1022", Name: "龙破斩", Level: 1, Desc: "用火焰斩击，造成大范围魔法伤害。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 600, StepGold: 200, StepGold2: 2}}
	s1033 := &Skill{
		Info:   SkillInfo{SkillId: "1023", Name: "光击阵", Level: 1, Desc: "从随机敌人脚下召唤火圈，造成小范围魔法伤害和眩晕。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 700, StepGold: 200, StepGold2: 2}}
	s1034 := &Skill{
		Info:   SkillInfo{SkillId: "1024", Name: "焰魂", Level: 1, Desc: "增加物理攻击的暴击。"},
		Secret: SkillSecretInfo{OriginLevelUpGold: 800, StepGold: 200, StepGold2: 2}}
	SkillDefineList[s1031.Info.SkillId] = s1031
	SkillDefineList[s1032.Info.SkillId] = s1032
	SkillDefineList[s1033.Info.SkillId] = s1033
	SkillDefineList[s1034.Info.SkillId] = s1034

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

func GetSkillDefine(skillId string) *Skill {
	if s, ok := SkillDefineList[skillId]; ok {
		res := new(Skill)
		tool.Clone(s, res)
		res.Uid = tool.UniqueId()
		return res
	}
	return nil
}

func GetSkills(skillIds []string) []*Skill {
	skills := make([]*Skill, 0)
	for _, skillId := range skillIds {
		if s := GetSkillDefine(skillId); s != nil {
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

func (s *Skill) SetSkillLevel(lv int32) {
	s.Info.Level = lv
}
