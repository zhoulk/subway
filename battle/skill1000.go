package battle

import (
	"subway/models"
)

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

			// 属性加成
			if h.Info.Type == models.HeroTypeStrength {
				hpEff += h.Runing.Strength * 2
			}
			if h.Info.Type == models.HeroTypeAgility {
				hpEff += h.Runing.Agility * 2
			}
			if h.Info.Type == models.HeroTypeIntelligent {
				hpEff += h.Runing.Intelligent * 2
			}
			if h.Props.AD > 0 {
				hpEff += h.Runing.AD
			}
			if h.Props.AP > 0 {
				hpEff += h.Runing.AP
			}

			hpEff = target.DecreaseHP(h, hpEff)

			// 增加能量
			h.IncreaseMP(h, 5)

			// 记录
			context.AddBattleItem(
				&BattleItem{
					MilliSeconds: context.MilliSeconds,
					FromHero:     ReportHero{HeroId: h.Uid, HP: h.Props.HP},
					ToHeros:      []ReportHero{ReportHero{HeroId: target.Uid, HP: target.Props.HP, Deffect: &BattleInfo{HP: hpEff}}},
					Skill:        ReportSkill{SkillId: "1000"},
				})
		}
	}
}
