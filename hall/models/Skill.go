package models

import "subway/tool"

type SkillInfo struct {
	Uid     string
	SkillId int32
	Name    string
}

func CreateASkill(skillId int32) *SkillInfo {
	return &SkillInfo{
		Uid:     tool.UniqueId(),
		SkillId: skillId,
		Name:    "",
	}
}
