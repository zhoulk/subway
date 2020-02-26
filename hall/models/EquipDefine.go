package models

import (
	"encoding/json"
	tables "subway/hall/db"
)

var (
	EquipDefineList map[int32]*EquipDefineInfo
)

func init() {
	EquipDefineList = make(map[int32]*EquipDefineInfo)
}

type EquipDefineInfo struct {
	EquipId int32
	Name    string

	MixArr []*EquipInfo
}

func GetEquipDefine(equipId int32) *EquipDefineInfo {
	if equip, ok := EquipDefineList[equipId]; ok {
		return equip
	}

	t_b := tables.LoadEquipDefineInfo(equipId)
	if t_b != nil {
		equip := CreateEquipDefineInfoFromTableEquipDefineItemInfo(t_b)
		EquipDefineList[equip.EquipId] = equip
		return equip
	}

	return nil
}

func CreateEquipDefineInfoFromTableEquipDefineItemInfo(a *tables.EquipDefine) *EquipDefineInfo {
	res := &EquipDefineInfo{
		EquipId: a.EquipId,
		Name:    "",
		MixArr:  make([]*EquipInfo, 0),
	}

	if len(a.Mix) > 0 {
		var equipArr []int32
		//读取的数据为json格式，需要进行解码
		err := json.Unmarshal([]byte(a.Mix), &equipArr)
		if err == nil {
			for _, equipId := range equipArr {
				res.MixArr = append(res.MixArr, &EquipInfo{
					EquipId: equipId,
				})
			}
		}
	}

	return res
}
