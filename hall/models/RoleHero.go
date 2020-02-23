package models

import (
	tables "subway/hall/db"
	"subway/tool"
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

	t_b := tables.LoadRoleHeroInfo(roleId)
	if t_b != nil {
		b := CreateRoleHeroFromTableRoleHero(t_b)
		RoleHeroList[roleId] = b
		return RoleHeroList[roleId]
	}

	roleHero := &RoleHero{
		RoleId: roleId,
		Heros:  make(map[string]*HeroInfo),
	}
	RoleHeroList[roleId] = roleHero

	return roleHero
}

func AddHero(roleId string, hero *HeroInfo) {
	if len(hero.Uid) == 0 {
		hero.Uid = tool.UniqueId()
	}

	roleHero := GetRoleHero(roleId)
	roleHero.Heros[hero.Uid] = hero
}

func PersistentRoleHeroInfo() {
	roleHeroInfos := make([]*tables.RoleHeroItemInfo, 0)
	for _, a := range RoleHeroList {
		roleHero := CreateTableRoleHeroInfoFromRoleHeroInfo(a)
		roleHeroInfos = append(roleHeroInfos, roleHero.Items...)
	}
	tables.PersistentRoleHeroItems(roleHeroInfos)
}

func CreateRoleHeroFromTableRoleHero(a *tables.RoleHeroInfo) *RoleHero {
	roleHero := &RoleHero{
		RoleId: a.RoleId,
		Heros:  make(map[string]*HeroInfo),
	}
	for _, t_item := range a.Items {
		roleHero.Heros[t_item.Uid] = CreateRoleHeroItemInfoFromTableRoleHeroItemInfo(t_item)
	}
	return roleHero
}

func CreateRoleHeroItemInfoFromTableRoleHeroItemInfo(a *tables.RoleHeroItemInfo) *HeroInfo {
	return &HeroInfo{
		Uid:    a.Uid,
		HeroId: a.HeroId,
		Name:   "",
	}
}

func CreateTableRoleHeroInfoFromRoleHeroInfo(a *RoleHero) *tables.RoleHeroInfo {
	t_roleHero := &tables.RoleHeroInfo{
		RoleId: a.RoleId,
		Items:  make([]*tables.RoleHeroItemInfo, 0),
	}

	for _, hero := range a.Heros {
		t_roleHero.Items = append(t_roleHero.Items, CreateTableRoleHeroItemInfoFromRoleHeroItemInfo(a.RoleId, hero))
	}

	return t_roleHero
}

func CreateTableRoleHeroItemInfoFromRoleHeroItemInfo(roleId string, a *HeroInfo) *tables.RoleHeroItemInfo {
	return &tables.RoleHeroItemInfo{
		Uid:    a.Uid,
		RoleId: roleId,
		HeroId: a.HeroId,
	}
}
