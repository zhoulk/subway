package models

import (
	tables "subway/hall/db"
	"time"
)

var (
	HouseList map[string]*HouseInfo
)

func init() {
	HouseList = make(map[string]*HouseInfo)
}

type HouseInfo struct {
	RoleId            string
	GoldTimes         int32
	TotalGoldTimes    int32
	DiamondTimes      int32
	TotalDiamondTimes int32
	LastGoldTime      time.Time
	LastDiamondTime   time.Time
}

func GetHouseInfo(roleId string) *HouseInfo {
	if r, ok := HouseList[roleId]; ok {
		return r
	}

	t_r := tables.LoadHouseInfo(roleId)
	if t_r != nil {
		r := CreateHouseInfoFromTableHouseInfo(t_r)
		HouseList[roleId] = r
		return HouseList[roleId]
	}

	return nil
}

func AddHouseInfo(houseInfo *HouseInfo) *HouseInfo {
	if _, ok := HouseList[houseInfo.RoleId]; ok {
		HouseList[houseInfo.RoleId] = houseInfo
	} else {
		HouseList[houseInfo.RoleId] = houseInfo
	}
	return HouseList[houseInfo.RoleId]
}

func GoldRandom(roleId string) *ProductInfo {
	houseInfo := GetHouseInfo(roleId)
	if houseInfo == nil && houseInfo.GoldTimes > 0 {
		return nil
	}
	// 核减
	houseInfo.GoldTimes--
	houseInfo.LastGoldTime = time.Now()
	// 随机出碎片 物品
	product := CreateAProduct(ProductTypeHeroPart)
	// 存储到背包
	AddProduct(roleId, product)
	return product
}

func DiamondRandom(roleId string) *ProductInfo {
	houseInfo := GetHouseInfo(roleId)
	if houseInfo == nil && houseInfo.DiamondTimes > 0 {
		return nil
	}
	houseInfo.DiamondTimes--
	houseInfo.LastDiamondTime = time.Now()
	// 随机出英雄 装备
	hero := RandAHero(roleId)
	product := &ProductInfo{
		ProductId: hero.Uid,
		ItemId:    hero.HeroId,
		Type:      ProductTypeHero,
		Name:      hero.Name,
		Count:     1,
	}
	// 存储到用户
	AddRoleHero(roleId, hero)
	return product
}

func PersistentHouseInfo() {
	houseInfos := make([]*tables.HouseInfo, 0)
	for _, a := range HouseList {
		houseInfos = append(houseInfos, CreateTableHouseInfoFromHouseInfo(a))
	}
	tables.PersistentHouseInfo(houseInfos)
}

func CreateHouseInfoFromTableHouseInfo(a *tables.HouseInfo) *HouseInfo {
	return &HouseInfo{
		RoleId:            a.RoleId,
		GoldTimes:         a.GoldTimes,
		TotalGoldTimes:    a.TotalGoldTimes,
		DiamondTimes:      a.DiamondTimes,
		TotalDiamondTimes: a.TotalDiamondTimes,
		LastGoldTime:      a.LastGoldTime,
		LastDiamondTime:   a.LastDiamondTime,
	}
}

func CreateTableHouseInfoFromHouseInfo(a *HouseInfo) *tables.HouseInfo {
	return &tables.HouseInfo{
		RoleId:            a.RoleId,
		GoldTimes:         a.GoldTimes,
		TotalGoldTimes:    a.TotalGoldTimes,
		DiamondTimes:      a.DiamondTimes,
		TotalDiamondTimes: a.TotalDiamondTimes,
		LastGoldTime:      a.LastGoldTime,
		LastDiamondTime:   a.LastDiamondTime,
	}
}
