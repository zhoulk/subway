package battle

import (
	"subway/models"
)

var (
	SkillExecuteMap map[string]func(h *Hero, s *Skill, context *BattleContext)
)

func RegisterSkillExecute(skillId string, exe func(h *Hero, s *Skill, context *BattleContext)) {
	if SkillExecuteMap == nil {
		SkillExecuteMap = make(map[string]func(h *Hero, s *Skill, context *BattleContext))
	}
	SkillExecuteMap[skillId] = exe
}

type Skill struct {
	*models.Skill
}

func ExecuteSkill(h *Hero, s *Skill, c *BattleContext) {

	if s == nil {
		// Skill1000Execute(h, s, c)
		if e, ok := SkillExecuteMap["1000"]; ok {
			e(h, s, c)
		}
		return
	}

	if e, ok := SkillExecuteMap[s.Info.SkillId]; ok {
		e(h, s, c)
	}

	// switch s.Info.SkillId {
	// case "1001":
	// 	Skill1001Execute(h, s, c)
	// 	break
	// case "1002":
	// 	Skill1002Execute(h, s, c)
	// 	break
	// case "1003":
	// 	Skill1003Execute(h, s, c)
	// 	break
	// case "1004":
	// 	Skill1004Execute(h, s, c)
	// 	break
	// case "1011":
	// 	Skill1011Execute(h, s, c)
	// 	break
	// case "1012":
	// 	Skill1012Execute(h, s, c)
	// 	break
	// case "1013":
	// 	Skill1013Execute(h, s, c)
	// 	break
	// case "1014":
	// 	Skill1014Execute(h, s, c)
	// 	break
	// case "1021":
	// 	Skill1021Execute(h, s, c)
	// 	break
	// case "1022":
	// 	Skill1022Execute(h, s, c)
	// 	break
	// case "1023":
	// 	Skill1023Execute(h, s, c)
	// 	break
	// case "1024":
	// 	Skill1024Execute(h, s, c)
	// 	break
	// }
}
