package battle

//   "光击阵",  "从随机敌人脚下召唤火圈，造成小范围魔法伤害和眩晕。"},
//  造成  lv * 30 伤害
// 眩晕 2 秒
//  10 秒
func Skill1033Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute")

	if context.MilliSeconds%10000 == 0 {

		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			target := targets[0]

			damageEff := s.Info.Level * 30
			damageEff = target.DecreaseHP(h, damageEff)

			dizzy := int32(3000)
			target.SetDizzy(dizzy)

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: damageEff, Dizzy: dizzy}}},
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
				})
		}
	}
}
