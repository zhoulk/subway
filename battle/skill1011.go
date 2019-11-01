package battle

//  "箭雨",  "对敌人连续射出多只箭矢，造成成吨的伤害。技能等级越高，射出的箭矢数量越多。
//   level * 500
///  mp 达到 100 施放
func Skill1011Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1001  execute")

	if h.Props.MP >= h.MaxMP {
		h.DecreaseMP(h, h.MaxMP)

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
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: damageEff}}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
				})
		}
	}
}
