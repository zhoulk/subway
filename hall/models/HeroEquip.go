package models

import tables "subway/hall/db"

var (
	HeroEquipList map[string]*HeroEquip
)

func init() {
	HeroEquipList = make(map[string]*HeroEquip)
}

type HeroEquip struct {
	HeroUid string
	Equips  map[string]*EquipInfo
}

func GetHeroEquipList(heroUid string) map[string]*EquipInfo {
	if heroEquip, ok := HeroEquipList[heroUid]; ok {
		return heroEquip.Equips
	}

	t_b := tables.LoadHeroEquipInfo(heroUid)
	if t_b != nil {
		b := CreateHeroEquipFromTableHeroEquip(heroUid, t_b)
		HeroEquipList[heroUid] = b
		return HeroEquipList[heroUid].Equips
	}

	heroInfo := GetHero(heroUid)
	if heroInfo == nil {
		return nil
	}
	heroFloorEquip := GetHeroFloorEquipDefine(heroInfo.HeroId, heroInfo.Floor)
	if heroFloorEquip == nil {
		return nil
	}
	heroEquip := &HeroEquip{
		HeroUid: heroUid,
		Equips:  make(map[string]*EquipInfo),
	}
	for _, v := range heroFloorEquip.Equips {
		equip := CreateAEquip(v.EquipId)
		heroEquip.Equips[equip.Uid] = equip
	}
	HeroEquipList[heroUid] = heroEquip

	return heroEquip.Equips
}

func CreateHeroEquipFromTableHeroEquip(heroUid string, t_heroEquip *tables.HeroEquipInfo) *HeroEquip {
	heroEquip := &HeroEquip{
		HeroUid: heroUid,
		Equips:  make(map[string]*EquipInfo),
	}
	for _, t_item := range t_heroEquip.Items {
		heroEquip.Equips[t_item.Uid] = CreateHeroEquipItemInfoFromTableHeroequipItemInfo(t_item)
	}
	return heroEquip
}

func CreateHeroEquipItemInfoFromTableHeroequipItemInfo(a *tables.HeroEquipItemInfo) *EquipInfo {
	return &EquipInfo{
		Uid:     a.HeroUid,
		EquipId: a.EquipId,
		Name:    "",
	}
}
