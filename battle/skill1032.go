package battle

import (
	"math/rand"
)

func init() {
	RegisterSkillExecute("1032", Skill1032Execute)
}

//   "龙破斩",  "用火焰斩击，造成大范围魔法伤害。"
//  造成  lv *  50 点物理伤害  三个目标
//  CD  4秒
func Skill1032Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute ", h.Props.MP, h.MaxMP)

	if context.MilliSeconds%4000 == 0 {
		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			toHeros := make([]ReportHero, 0)

			one := rand.Intn(len(targets))
			target := targets[one]

			damageEff := s.Info.Level * 50
			damageEff = target.DecreaseHP(h, damageEff)

			toHeros = append(toHeros, ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: damageEff}})

			for i := 1; i < 3; i++ {
				two := (one + i) % len(targets)

				if one != two {
					target2 := targets[two]

					damageEff = s.Info.Level * 30
					damageEff = target2.DecreaseHP(h, damageEff)

					toHeros = append(toHeros, ReportHero{HeroId: target2.Uid, HP: target2.Props.HP, Deffect: &BattleInfo{HP: damageEff}})
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
