package models

import (
	tables "subway/hall/db"
	"subway/tool"
)

var (
	EquipList map[string]*EquipInfo
)

func init() {
	EquipList = make(map[string]*EquipInfo)
}

type EquipInfo struct {
	Uid     string
	EquipId int32
	Name    string
}

func CreateAEquip(equipId int32) *EquipInfo {
	equip := &EquipInfo{
		Uid:     tool.UniqueId(),
		EquipId: equipId,
		Name:    "",
	}
	AddEquip(equip)
	return equip
}

func AddEquip(equip *EquipInfo) {
	EquipList[equip.Uid] = equip
}

func GetEquip(equipUid string) *EquipInfo {
	if equip, ok := EquipList[equipUid]; ok {
		return equip
	}

	t_b := tables.LoadHeroEquipItemInfo(equipUid)
	if t_b != nil {
		equip := CreateHeroEquipItemInfoFromTableHeroEquipItemInfo(t_b)
		EquipList[equip.Uid] = equip
		return equip
	}

	return nil
}

func GetEquipMixList(equipUid string) []*EquipInfo {
	equip := GetEquip(equipUid)
	if equip == nil {
		return nil
	}

	equipDefineInfo := GetEquipDefine(equip.EquipId)
	if equipDefineInfo == nil {
		return nil
	}

	return equipDefineInfo.MixArr
}

func PersistentEquipInfo() {
	heroEquipInfos := make([]*tables.HeroEquipItemInfo, 0)
	for _, a := range EquipList {
		heroEquipInfos = append(heroEquipInfos, CreateTableHeroEquipItemInfoFromHeroEquipItemInfo(a))
	}
	tables.PersistentHeroEquipItems(heroEquipInfos)
}

func CreateHeroEquipItemInfoFromTableHeroEquipItemInfo(a *tables.HeroEquipItemInfo) *EquipInfo {
	return &EquipInfo{
		Uid:     a.Uid,
		EquipId: a.EquipId,
		Name:    "",
	}
}

func CreateTableHeroEquipItemInfoFromHeroEquipItemInfo(a *EquipInfo) *tables.HeroEquipItemInfo {
	return &tables.HeroEquipItemInfo{
		Uid:     a.Uid,
		EquipId: a.EquipId,
	}
}
