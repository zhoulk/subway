package models

import (
	"subway/tool"

	tables "subway/hall/db"
)

const (
	HeroTypeStrength    int8 = 1
	HeroTypeAgility     int8 = 2
	HeroTypeIntelligent int8 = 3
)

var (
	HeroList map[string]*HeroInfo
)

func init() {
	HeroList = make(map[string]*HeroInfo)
}

type HeroInfo struct {
	Uid    string
	RoleId string
	HeroId int32
	Type   int8
	Name   string
	Floor  int32
	Star   int32
	Level  int32
}

func RandAHero(roleId string) *HeroInfo {
	return CreateAHero(roleId, 1002)
}

func CreateAHero(roleId string, heroId int32) *HeroInfo {
	hero := &HeroInfo{
		RoleId: roleId,
		Uid:    tool.UniqueId(),
		HeroId: heroId,
		Name:   "",
		Floor:  1,
		Star:   1,
		Level:  1,
		Type:   HeroTypeIntelligent,
	}
	AddHero(hero)
	return hero
}

func AddHero(hero *HeroInfo) {
	HeroList[hero.Uid] = hero
}

func GetHero(uid string) *HeroInfo {
	if hero, ok := HeroList[uid]; ok {
		return hero
	}

	t_b := tables.LoadRoleHeroItemInfo(uid)
	if t_b != nil {
		return CreateRoleHeroItemInfoFromTableRoleHeroItemInfo(t_b)
	}

	return nil
}

func PersistentHeroInfo() {
	roleHeroInfos := make([]*tables.RoleHeroItemInfo, 0)
	for _, a := range HeroList {
		roleHeroInfos = append(roleHeroInfos, CreateTableRoleHeroItemInfoFromRoleHeroItemInfo(a))
	}
	tables.PersistentRoleHeroItems(roleHeroInfos)
}

func CreateRoleHeroItemInfoFromTableRoleHeroItemInfo(a *tables.RoleHeroItemInfo) *HeroInfo {
	return &HeroInfo{
		Uid:    a.Uid,
		RoleId: a.RoleId,
		HeroId: a.HeroId,
		Name:   "",
		Floor:  a.Floor,
		Level:  a.Level,
		Star:   a.Star,
	}
}

func CreateTableRoleHeroItemInfoFromRoleHeroItemInfo(a *HeroInfo) *tables.RoleHeroItemInfo {
	return &tables.RoleHeroItemInfo{
		Uid:    a.Uid,
		RoleId: a.RoleId,
		HeroId: a.HeroId,
		Floor:  a.Floor,
		Level:  a.Level,
		Star:   a.Star,
	}
}
