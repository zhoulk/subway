package battle

import "math/rand"

func init() {
	RegisterSkillExecute("1022", Skill1022Execute)
}

//   "洪流", Level: 1, Desc: "用水流击飞一个随机敌人，造成魔法伤害。"
//  造成  lv * 30 伤害
//  6 秒
func Skill1022Execute(h *Hero, s *Skill, context *BattleContext) {
	// beego.Informational("Skill1024  execute")

	if context.MilliSeconds%6000 == 0 {

		targets := context.GetOtherHeros(h.Group)

		if targets != nil && len(targets) > 0 {
			target := targets[rand.Intn(len(targets))]

			damageEff := s.Info.Level * 30
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
