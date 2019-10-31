package battle

//  普通攻击
//  依赖 英雄攻速
func Skill1000Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1000  execute")

	if context.MilliSeconds%h.Props.SPD == 0 {
		targets := context.GetOtherHeros(h.Group)
		if targets != nil && len(targets) > 0 {
			target := targets[0]

			hpEff := h.Props.AD
			if h.Props.AP > h.Props.AD {
				hpEff = h.Props.AP
			}

			hpEff = target.DecreaseHP(h, hpEff)

			// 增加能量
			h.IncreaseMP(h, 5)

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP}},
					Skill:        ReportSkill{SkillId: "1000"},
					Deffect:      &BattleInfo{HP: hpEff},
				})
		}
	}
}
