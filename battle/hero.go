package battle

import (
	"subway/models"
)

type Hero struct {
	*models.Hero
	Group  int        // 1  2
	Runing BattleInfo // 运行过程中 产生的变量

	MaxHP int32
	MaxMP int32
}

type BattleInfo struct {
	AD          int32
	AP          int32
	APDef       int32
	MP          int32
	HP          int32
	Strength    int32
	Agility     int32
	Intelligent int32

	Dizzy   int32
	Silence int32
}

func (h *Hero) IncreaseHP(from *Hero, hp int32) int32 {
	if h.Props.HP+hp > h.MaxHP {
		hp = h.MaxHP - h.Props.HP
		h.Props.HP = h.MaxHP
	} else {
		h.Props.HP += hp
	}
	return hp
}

func (h *Hero) DecreaseHP(from *Hero, hp int32) int32 {
	if h.Props.HP-hp < 0 {
		hp = h.Props.HP
		h.Props.HP = 0

		from.IncreaseMP(h, 30)
	} else {
		h.Props.HP -= hp
	}

	h.IncreaseMP(from, hp*100/h.MaxHP)

	return hp
}

func (h *Hero) IncreaseMP(from *Hero, mp int32) int32 {
	if h.Props.MP+mp > h.MaxMP {
		mp = h.MaxMP - h.Props.MP
		h.Props.MP = h.MaxMP
	} else {
		h.Props.MP += mp
	}
	return mp
}

func (h *Hero) DecreaseMP(from *Hero, mp int32) int32 {
	if h.Props.MP-mp < 0 {
		mp = h.Props.MP
		h.Props.MP = 0
	} else {
		h.Props.MP -= mp
	}
	return mp
}

func (h *Hero) SetDizzy(dizzy int32) {
	if dizzy < 0 {
		dizzy = 0
	}
	h.Runing.Dizzy = dizzy
}

func (h *Hero) SetSilence(silence int32) {
	if silence < 0 {
		silence = 0
	}
	h.Runing.Silence = silence
}
