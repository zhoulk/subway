package models

import (
	"subway/tool"

	tables "subway/hall/db"
)

var (
	HeroPartList map[string]*HeroPartInfo
)

func init() {
	HeroPartList = make(map[string]*HeroPartInfo)
}

type HeroPartInfo struct {
	Uid    string
	RoleId string
	HeroId int32
	Num    int32
}

func RandAHeroPart(roleId string) *HeroPartInfo {
	return CreateAHeroPart(roleId, 1002)
}

func CreateAHeroPart(roleId string, heroId int32) *HeroPartInfo {
	hero := &HeroPartInfo{
		RoleId: roleId,
		Uid:    tool.UniqueId(),
		HeroId: heroId,
		Num:    1,
	}
	AddHeroPart(hero)
	return hero
}

func AddHeroPart(hero *HeroPartInfo) {
	HeroPartList[hero.Uid] = hero
}

func GetHeroPart(uid string) *HeroPartInfo {
	if hero, ok := HeroPartList[uid]; ok {
		return hero
	}

	t_b := tables.LoadRoleHeroPartItemInfo(uid)
	if t_b != nil {
		return CreateRoleHeroPartItemInfoFromTableRoleHeroPartItemInfo(t_b)
	}

	return nil
}

func PersistentHeroPartInfo() {
	roleHeroInfos := make([]*tables.RoleHeroPartItemInfo, 0)
	for _, a := range HeroPartList {
		roleHeroInfos = append(roleHeroInfos, CreateTableRoleHeroPartItemInfoFromRoleHeroPartItemInfo(a))
	}
	tables.PersistentRoleHeroPartItems(roleHeroInfos)
}

func CreateRoleHeroPartItemInfoFromTableRoleHeroPartItemInfo(a *tables.RoleHeroPartItemInfo) *HeroPartInfo {
	return &HeroPartInfo{
		Uid:    a.Uid,
		RoleId: a.RoleId,
		HeroId: a.HeroId,
		Num:    a.Num,
	}
}

func CreateTableRoleHeroPartItemInfoFromRoleHeroPartItemInfo(a *HeroPartInfo) *tables.RoleHeroPartItemInfo {
	return &tables.RoleHeroPartItemInfo{
		Uid:    a.Uid,
		RoleId: a.RoleId,
		HeroId: a.HeroId,
		Num:    a.Num,
	}
}
