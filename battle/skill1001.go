package battle

//  敌法闪烁到智力最高的敌人身后，以目标为中心施放能量虚空，造成巨大的魔法伤害。
//  造成  lv *  500 点物理伤害
//  mp 达到 100 施放
func Skill1001Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1001  execute")
	if context.MilliSeconds%5000 == 0 {
		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			target := targets[0]
			for _, h := range targets {
				if h.Props.Intelligent > target.Props.Intelligent {
					target = h
				}
			}

			damageEff := s.Info.Level * 500
			damageEff = target.DecreaseHP(h, damageEff)

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
					Deffect:      &BattleInfo{HP: damageEff},
				})
		}
	}
}
