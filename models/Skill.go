package models

import (
	"subway/db/tables"
	"subway/tool"
)

var (
	SkillDefineList map[string]*Skill
)

func init() {
	SkillDefineList = make(map[string]*Skill)

	defines := tables.LoadSkillDefine()
	for _, def := range defines {
		SkillDefineList[def.SkillId] = CreateSkillFromSkillDefine(def)
	}
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

func GetSkillDefines(skillIds []string) []*Skill {
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

func CreateSkillFromSkillDefine(t_s_d *tables.SkillDefine) *Skill {
	return &Skill{
		Info: SkillInfo{
			SkillId: t_s_d.SkillId,
			Name:    t_s_d.Name,
			Level:   t_s_d.Level,
			Type:    t_s_d.Type,
			Desc:    t_s_d.Desc,
		},
	}
}

func CreateSkillFromHeroSkill(t_h_s *tables.HeroSkill) *Skill {
	if s, ok := SkillDefineList[t_h_s.SkillId]; ok {
		res := new(Skill)
		tool.Clone(s, res)
		res.Uid = t_h_s.Uid
		res.Info.Level = t_h_s.Level
		return res
	}
	return nil
}

func CreateHeroSkillFromSkill(heroUid string, u_h_s *Skill) *tables.HeroSkill {
	return &tables.HeroSkill{
		Uid:     u_h_s.Uid,
		HeroUid: heroUid,
		SkillId: u_h_s.Info.SkillId,
		Level:   u_h_s.Info.Level,
	}
}
