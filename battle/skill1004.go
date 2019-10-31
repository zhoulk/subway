package battle

//  敌法利用对魔法能量的理解减少受到的伤害，增加魔抗。
//  1级 增加  1点
//  被动
func Skill1004Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1001  execute")

	if context.MilliSeconds == -1 {
		eff := s.Info.Level * 1

		if eff > 0 {
			h.Runing.APDef += eff
			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: h.Uid, HP: h.Props.HP}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
					Effect:       &BattleInfo{APDef: eff},
				})
		}
	}
}
