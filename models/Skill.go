package models

import (
	"subway/db/tables"
	"subway/tool"

	"github.com/astaxie/beego"
)

var (
	SkillDefineList map[string]*Skill
)

func init() {
	beego.Debug("Skill init")

	SkillDefineList = make(map[string]*Skill)

	defines := tables.LoadSkillDefine()
	for _, def := range defines {
		s := CreateSkillFromSkillDefine(def)
		SkillDefineList[def.SkillId] = s
		s.SetSkillLevel(1)
	}
}

type Skill struct {
	Uid    string
	Info   SkillInfo
	Secret SkillSecretInfo
}

type SkillInfo struct {
	SkillId     string
	Name        string
	Level       int32
	Type        int8 // 1 主动  2  被动
	LevelUpGold int32
	Desc        string
}

type SkillSecretInfo struct {
	OriginLevelUpGold int32
	StepGold          int32
	StepGold2         int32
}

func (s *Skill) SetSkillLevel(lv int32) {
	s.Info.Level = lv

	// 增加效果
	//RefreshHero(h)

	// 修改升级需要的金币
	nextLevel := s.Info.Level + 1
	s.Info.LevelUpGold = int32(s.Secret.OriginLevelUpGold + s.Secret.StepGold*nextLevel + nextLevel*nextLevel)
}

func GetSkillDefine(skillId string) *Skill {
	if s, ok := SkillDefineList[skillId]; ok {
		res := new(Skill)
		err := tool.Clone(s, res)
		if err == nil {
			res.Uid = tool.UniqueId()
		} else {
			beego.Error("GetSkillDefine  error ", err)
		}
		return res
	}
	return nil
}

func GetSkillDefines(skillIds []string) []*Skill {
	beego.Debug("GetSkillDefines ", skillIds)
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
			needGold := int64(targetSkill.Info.LevelUpGold)
			if u.Profile.Gold >= needGold {
				targetSkill.SetSkillLevel(endLevel)
				u.Profile.Gold -= needGold
				return true
			}
		}
	}

	return false
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
		res.SetSkillLevel(res.Info.Level)
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
