package battle

//  "冰箭", "射出一只冰箭，对目标造成物理伤害。"
//   level * 25
//  4 秒 CD
func Skill1012Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1001  execute")

	if context.MilliSeconds%4000 == 0 {

		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			target := targets[0]

			damageEff := s.Info.Level * 25
			damageEff = target.DecreaseHP(h, damageEff)

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: damageEff}}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
				})
		}
	}
}
