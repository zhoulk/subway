package models

import (
	"encoding/json"
	tables "subway/hall/db"
)

var (
	HeroEquipDefineList map[int32]*HeroEquipDefine
)

func init() {
	HeroEquipDefineList = make(map[int32]*HeroEquipDefine)
}

type HeroEquipDefine struct {
	HeroId int32
	Floors map[int32]*HeroFloorEquipDefine
}

type HeroFloorEquipDefine struct {
	Floor  int32
	Equips []*EquipInfo
}

func GetHeroFloorEquipDefine(heroId int32, floor int32) *HeroFloorEquipDefine {
	if heroEquipDefine, ok := HeroEquipDefineList[heroId]; ok {
		return heroEquipDefine.Floors[floor]
	}

	t_bs := tables.LoadHeroEquipDefine(heroId)
	if t_bs != nil {
		HeroEquipDefineList[heroId] = &HeroEquipDefine{
			HeroId: heroId,
			Floors: make(map[int32]*HeroFloorEquipDefine, 0),
		}
		for _, t_b := range t_bs {
			HeroEquipDefineList[heroId].Floors[t_b.Floor] = CreateHeroFloorEquipDefineFromTableHeroEquipDefine(t_b)
		}
		return HeroEquipDefineList[heroId].Floors[floor]
	}

	return nil
}

func CreateHeroFloorEquipDefineFromTableHeroEquipDefine(a *tables.HeroEquipDefine) *HeroFloorEquipDefine {
	res := &HeroFloorEquipDefine{
		Floor:  a.Floor,
		Equips: make([]*EquipInfo, 0),
	}

	if len(a.EquipId) > 0 {
		var equipArr []int32
		//读取的数据为json格式，需要进行解码
		err := json.Unmarshal([]byte(a.EquipId), &equipArr)
		if err == nil {
			for _, equipId := range equipArr {
				res.Equips = append(res.Equips, &EquipInfo{
					EquipId: equipId,
				})
			}
		}
	}

	return res
}
