package models

import (
	tables "subway/hall/db"
)

var (
	RoleHeroPartList map[string]*RoleHeroPart
)

func init() {
	RoleHeroPartList = make(map[string]*RoleHeroPart)
}

type RoleHeroPart struct {
	RoleId    string
	HeroParts map[string]*HeroPartInfo
}

func GetRoleHeroPart(roleId string) *RoleHeroPart {
	if b, ok := RoleHeroPartList[roleId]; ok {
		return b
	}

	roleHero := &RoleHeroPart{
		RoleId:    roleId,
		HeroParts: make(map[string]*HeroPartInfo),
	}

	t_b := tables.LoadRoleHeroPartInfo(roleId)
	if t_b != nil {
		for _, t_item := range t_b.Items {
			heroInfo := GetHeroPart(t_item.Uid)
			if heroInfo != nil {
				roleHero.HeroParts[t_item.Uid] = heroInfo
			}
		}
	}

	RoleHeroPartList[roleId] = roleHero

	return roleHero
}

func AddRoleHeroPart(roleId string, hero *HeroPartInfo) {
	roleHero := GetRoleHeroPart(roleId)
	roleHero.HeroParts[hero.Uid] = hero
}
