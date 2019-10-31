package battle

//  敌法加速旋转他的刀刃，增加他的敏捷。
//  1 级增加 5点
//  被动
func Skill1003Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1001  execute")
	if context.MilliSeconds == -1 {
		eff := s.Info.Level * 5

		if eff > 0 {
			h.Runing.Agility += eff
			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: h.Uid, HP: h.Props.HP}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
					Effect:       &BattleInfo{Agility: eff},
				})
		}
	}
}
