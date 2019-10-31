package battle

import (
	"subway/models"
)

type Skill struct {
	*models.Skill
}

func ExecuteSkill(h *Hero, s *Skill, c *BattleContext) {

	if s == nil {
		Skill1000Execute(h, s, c)
		return
	}

	switch s.Info.SkillId {
	case "1001":
		Skill1001Execute(h, s, c)
		break
	case "1002":
		Skill1002Execute(h, s, c)
		break
	case "1003":
		Skill1003Execute(h, s, c)
		break
	case "1004":
		Skill1004Execute(h, s, c)
		break
	}
}
