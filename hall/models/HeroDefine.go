package models

import (
	tables "subway/hall/db"
)

var (
	HeroDefineList map[int32]*HeroDefineInfo
)

func init() {
	HeroDefineList = make(map[int32]*HeroDefineInfo)
}

type HeroDefineInfo struct {
	HeroId int32
	Name   string
	Type   int8
}

func GetHeroDefine(heroId int32) *HeroDefineInfo {
	if hero, ok := HeroDefineList[heroId]; ok {
		return hero
	}

	t_b := tables.LoadHeroDefine(heroId)
	if t_b != nil {
		hero := CreateHeroDefineInfoFromTableHeroDefineItemInfo(t_b)
		HeroDefineList[hero.HeroId] = hero
		return hero
	}

	return nil
}

func CreateHeroDefineInfoFromTableHeroDefineItemInfo(a *tables.HeroDefine) *HeroDefineInfo {
	res := &HeroDefineInfo{
		HeroId: a.HeroId,
		Name:   a.Name,
		Type:   a.Type,
	}

	return res
}
