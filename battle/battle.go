package battle

import (
	"subway/models"
	"subway/tool"

	"github.com/astaxie/beego"
)

type BattleResult struct {
	Result int // 1 胜利  2 失败   3  平局
}

func BattleGuanKa(uid string, gkId int) *BattleResult {
	return Battle(models.GetSelfHeros(uid), models.GetSelfHeros(uid))
}

func Battle(heros1 []*models.Hero, heros2 []*models.Hero) *BattleResult {
	result := BattleResult{}

	if heros1 == nil || heros2 == nil {
		return &result
	}

	beego.Informational(heros1, heros2)

	selfHeros := make([]*models.Hero, 0)
	otherHeros := make([]*models.Hero, 1)

	for _, h := range heros1 {
		hh := new(models.Hero)
		tool.Clone(h, hh)
		selfHeros = append(selfHeros, hh)
	}

	tool.Clone(heros2, otherHeros)

	for _, h := range otherHeros {
		beego.Informational(h.Props.HP)

	}

	// tool.Clone(heros1, selfHeros)

	beego.Informational(selfHeros, otherHeros)

	for isLive(selfHeros) && isLive(otherHeros) {
		// 按速度排序
		// 执行技能
		for _, h := range selfHeros {
			h.Props.HP -= 10
		}
		for _, h := range otherHeros {
			h.Props.HP -= 10
		}
	}

	return &result
}

func isLive(heros []*models.Hero) bool {
	if heros == nil {
		return false
	}
	for _, hero := range heros {
		if hero.Props.HP > 0 {
			return true
		}
	}
	return false
}
