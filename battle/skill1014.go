package battle

import (
	"subway/models"
)

func init() {
	RegisterSkillExecute("1014", Skill1014Execute)
}

//    "射手天赋", "增加全体队友攻击力。"},
//  1级 增加  4点 攻击力
//  被动
func Skill1014Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute")

	if context.MilliSeconds == -1 {
		targets := context.GetSelfHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			eff := s.Info.Level * 4

			if eff > 0 {
				toHeros := make([]ReportHero, 0)
				for _, target := range targets {
					if target.Info.AtkType == models.HeroAtkTypeAD {
						target.Runing.AD += eff
						toHeros = append(toHeros, ReportHero{HeroId: target.Uid, HP: target.Props.HP, Effect: &BattleInfo{AD: eff}})
					}
					if target.Info.AtkType == models.HeroAtkTypeAP {
						target.Runing.AP += eff
						toHeros = append(toHeros, ReportHero{HeroId: target.Uid, HP: target.Props.HP, Effect: &BattleInfo{AP: eff}})
					}
				}

				// 记录
				context.AddBattleItem(
					&BattleItem{
						MilliSeconds: context.MilliSeconds,
						FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
						ToHeros:      toHeros,
						Skill:        ReportSkill{SkillId: s.Info.SkillId},
					})
			}
		}
	}
}
