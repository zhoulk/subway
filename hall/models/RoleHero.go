package models

import (
	tables "subway/hall/db"
)

var (
	RoleHeroList map[string]*RoleHero
)

func init() {
	RoleHeroList = make(map[string]*RoleHero)
}

type RoleHero struct {
	RoleId string
	Heros  map[string]*HeroInfo
}

func GetRoleHero(roleId string) *RoleHero {
	if b, ok := RoleHeroList[roleId]; ok {
		return b
	}

	roleHero := &RoleHero{
		RoleId: roleId,
		Heros:  make(map[string]*HeroInfo),
	}

	t_b := tables.LoadRoleHeroInfo(roleId)
	if t_b != nil {
		for _, t_item := range t_b.Items {
			heroInfo := GetHero(t_item.Uid)
			if heroInfo != nil{
				roleHero.Heros[t_item.Uid] = heroInfo
			}
		}
	}

	RoleHeroList[roleId] = roleHero

	return roleHero
}

func AddRoleHero(roleId string, hero *HeroInfo) {
	roleHero := GetRoleHero(roleId)
	roleHero.Heros[hero.Uid] = hero
}
