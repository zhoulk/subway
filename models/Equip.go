package models

import (
	"encoding/json"
	"strconv"
	"subway/db/tables"
	"subway/tool"

	"github.com/astaxie/beego"
)

var (
	EquipDefineList map[string]*Equip
)

func init() {
	EquipDefineList = make(map[string]*Equip)

	defines := tables.LoadEquipDefine()
	for _, def := range defines {
		e := &Equip{
			Info: EquipInfo{
				EquipId:     def.EquipId,
				Name:        def.Name,
				Level:       def.Level,
				Strength:    def.Strength,
				HP:          def.HP,
				Agility:     def.Agility,
				MP:          def.MP,
				Intelligent: def.Intelligent,
				AD:          def.AD,
				ADCrit:      def.ADCrit,
				ADDef:       def.ADDef,
				HPGain:      def.HPGain,
				MPGain:      def.MPGain,
				HPReGain:    def.HPReGain,
				MPReGain:    def.MPReGain,
				Desc:        def.Desc,
				From:        def.From,
				Power:       def.Power,
				Cost:        def.Cost,
			},
		}
		EquipDefineList[def.EquipId] = e
	}

	for _, def := range defines {

		if len(def.Mix) > 0 {
			var mixArr []int
			//读取的数据为json格式，需要进行解码
			err := json.Unmarshal([]byte(def.Mix), &mixArr)
			if err == nil {
				e := EquipDefineList[def.EquipId]
				e.Mix = make([]*Equip, 0)
				for _, equipId := range mixArr {
					e.Mix = append(e.Mix, EquipDefineList[strconv.Itoa(equipId)])
				}
			}
		}

	}
}

const (
	EquipStatusWearNormal   int8 = 1
	EquipStatusWearAcquire  int8 = 2
	EquipStatusWearComplete int8 = 3
)

type Equip struct {
	Uid    string
	Info   EquipInfo
	Mix    []*Equip
	Parts  int
	Status int8 // EquipStatusWear
}

type EquipInfo struct {
	EquipId     string
	Name        string
	Level       int32 // 级别
	Strength    int32 // 力量
	HP          int32 // 生命值
	Agility     int32 // 敏捷
	MP          int32 // 魔法强度
	Intelligent int32
	AD          int32 // 物理攻击
	ADCrit      int32 // 物理暴击
	ADDef       int32 // 物理护甲
	HPGain      int32 // 生命恢复
	MPGain      int32 // 能量恢复
	HPReGain    int32 // 战斗后补充生命
	MPReGain    int32 // 战斗后补充能量
	Desc        string
	From        string
	Power       int32
	Cost        int32
	MixCnt      int32
}

func GetEquipDefines(equipIds []string) []*Equip {

	beego.Debug("GetEquipDefines ", equipIds)

	equips := make([]*Equip, 0)
	for _, equipId := range equipIds {
		if e, ok := EquipDefineList[equipId]; ok {
			res := new(Equip)
			tool.Clone(e, res)
			res.Uid = tool.UniqueId()
			res.Status = EquipStatusWearNormal
			equips = append(equips, res)
		}
	}
	return equips
}

func CreateEquipFromHeroEquip(t_h_e *tables.HeroEquip) *Equip {
	if e, ok := EquipDefineList[t_h_e.EquipId]; ok {
		res := new(Equip)
		tool.Clone(e, res)
		res.Uid = t_h_e.Uid
		res.Status = t_h_e.Status
		return res
	}
	return nil
}

func CreateHeroEquipFromEquip(heroUid string, u_h_e *Equip) *tables.HeroEquip {
	return &tables.HeroEquip{
		Uid:     u_h_e.Uid,
		HeroUid: heroUid,
		EquipId: u_h_e.Info.EquipId,
		Status:  u_h_e.Status,
	}
}
