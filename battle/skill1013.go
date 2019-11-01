package battle

import "math/rand"

//  "沉默",  "使数个敌人陷入沉默，无法释放任何带魔法伤害的技能。"
//   3个
// 沉默 3秒
//  6 秒 CD
func Skill1013Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1001  execute")

	if context.MilliSeconds%6000 == 0 {
		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			toHeros := make([]ReportHero, 0)

			one := rand.Intn(len(targets))
			target := targets[one]

			silence := int32(3000)
			target.SetSilence(silence)

			toHeros = append(toHeros, ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{Silence: silence}})

			for i := 1; i < 3; i++ {
				two := (one + i) % len(targets)

				if one != two {
					target2 := targets[two]

					target2.SetSilence(silence)

					toHeros = append(toHeros, ReportHero{HeroId: target2.Uid, HP: target2.Props.HP, Deffect: &BattleInfo{Silence: silence}})
				}
			}

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      toHeros,
					Skill:        ReportSkill{SkillId: s.Info.SkillId},
				})
		}
	}
}
