package battle

import "math/rand"

func init() {
	RegisterSkillExecute("1023", Skill1023Execute)
}

//   "水刀", Level: 1, Desc: "对自身小范围内的敌人造成魔法伤害。"
//  造成  lv * 20 伤害  随机两个目标
//  4 秒
func Skill1023Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute")

	if context.MilliSeconds%4000 == 0 {

		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			toHeros := make([]ReportHero, 0)

			one := rand.Intn(len(targets))
			target := targets[one]

			damageEff := s.Info.Level * 30
			damageEff = target.DecreaseHP(h, damageEff)

			toHeros = append(toHeros, ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: damageEff}})

			two := (one + 1) % len(targets)
			if one != two {
				target2 := targets[two]

				damageEff = s.Info.Level * 30
				damageEff = target2.DecreaseHP(h, damageEff)

				toHeros = append(toHeros, ReportHero{HeroId: target2.Uid, HP: target2.Props.HP, Deffect: &BattleInfo{HP: damageEff}})

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
