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
	case "1011":
		Skill1011Execute(h, s, c)
		break
	case "1012":
		Skill1012Execute(h, s, c)
		break
	case "1013":
		Skill1013Execute(h, s, c)
		break
	case "1014":
		Skill1014Execute(h, s, c)
		break
	case "1021":
		Skill1021Execute(h, s, c)
		break
	case "1022":
		Skill1022Execute(h, s, c)
		break
	case "1023":
		Skill1023Execute(h, s, c)
		break
	case "1024":
		Skill1024Execute(h, s, c)
		break
	}
}
