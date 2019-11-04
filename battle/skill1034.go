package battle

func init() {
	RegisterSkillExecute("1034", Skill1034Execute)
}

//    "焰魂",   "增加物理攻击的暴击。"
//  1级 增加  1点暴击
//  被动
func Skill1034Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute")

	if context.MilliSeconds == -1 {

		eff := s.Info.Level * 1

		if eff > 0 {
			h.Runing.ADCrit += eff

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: h.Uid, HP: h.Props.HP, Effect: &BattleInfo{ADCrit: eff}}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
				})
		}
	}
}
