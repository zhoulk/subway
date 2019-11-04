package battle

func init() {
	RegisterSkillExecute("1024", Skill1024Execute)
}

//   "力量强化"  "力量强化船长专注地磨练自己的身体，增加力量。"
//  1级 增加  2点 力量
//  被动
func Skill1024Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute")

	if context.MilliSeconds == -1 {
		eff := s.Info.Level * 2

		if eff > 0 {
			h.Runing.Strength += eff
			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: h.Uid, HP: h.Props.HP, Effect: &BattleInfo{Strength: eff}}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
				})
		}
	}
}
