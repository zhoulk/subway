package battle

import (
	"math/rand"
)

func init() {
	RegisterSkillExecute("1031", Skill1031Execute)
}

//   神灭斩",  "对目标射出一道闪电，造成巨量的魔法伤害。
//  造成  lv *  1000 点物理伤害
//  mp 达到 100 施放
func Skill1031Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute ", h.Props.MP, h.MaxMP)

	if h.Props.MP >= h.MaxMP {
		h.DecreaseMP(h, h.MaxMP)

		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			target := targets[rand.Intn(len(targets))]

			damageEff := s.Info.Level * 1000
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
