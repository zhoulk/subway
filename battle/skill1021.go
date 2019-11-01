package battle

import (
	"math/rand"
)

//   "幽灵船", "召唤幽灵船冲向对方，造成大范围的眩晕和魔法伤害。"
//  造成  lv *  100 点物理伤害  三个目标
//  眩晕  3秒
//  mp 达到 100 施放
func Skill1021Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute ", h.Props.MP, h.MaxMP)

	if h.Props.MP >= h.MaxMP {
		h.DecreaseMP(h, h.MaxMP)

		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			toHeros := make([]ReportHero, 0)

			one := rand.Intn(len(targets))
			target := targets[one]

			damageEff := s.Info.Level * 100
			damageEff = target.DecreaseHP(h, damageEff)

			dizzy := int32(3000)
			target.SetDizzy(dizzy)

			toHeros = append(toHeros, ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: damageEff, Dizzy: dizzy}})

			for i := 1; i < 3; i++ {
				two := (one + i) % len(targets)

				if one != two {
					target2 := targets[two]

					damageEff = s.Info.Level * 30
					damageEff = target2.DecreaseHP(h, damageEff)

					target2.SetDizzy(dizzy)

					toHeros = append(toHeros, ReportHero{HeroId: target2.Uid, HP: target2.Props.HP, Deffect: &BattleInfo{HP: damageEff, Dizzy: dizzy}})
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
