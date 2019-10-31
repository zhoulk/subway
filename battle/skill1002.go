package battle

//  敌法闪烁到智力最高的敌人身边攻击，造成小范围物理伤害并额外损毁敌人的能量。
//  造成  lv *  50 点物理伤害
//  消耗  lv  * 1 点魔法
//  5s
func Skill1002Execute(h *Hero, s *Skill, context *BattleContext) {
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

			damageEff := s.Info.Level * 50
			mpEff := s.Info.Level * 1

			damageEff = target.DecreaseHP(h, damageEff)
			mpEff = target.DecreaseMP(h, mpEff)

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
					Deffect:      &BattleInfo{MP: mpEff, HP: damageEff},
				})
		}
	}
}
