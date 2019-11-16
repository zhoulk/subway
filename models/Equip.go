package models

import (
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
		EquipDefineList[def.EquipId] = &Equip{
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
			},
		}
	}

	// e1 := &Equip{
	// 	Info: EquipInfo{EquipId: "1000", Name: "树枝", Level: 1, Strength: 1, Agility: 1, Intelligent: 1, Desc: "带上它以确保一局GG", Power: 3}}
	// e2 := &Equip{
	// 	Info: EquipInfo{EquipId: "1001", Name: "敏捷丝袜", Level: 1, Agility: 3, Desc: "蜘蛛侠Cosplay套装组件", Power: 3}}
	// e3 := &Equip{
	// 	Info: EquipInfo{EquipId: "1002", Name: "小圆盾", Level: 1, ADDef: 2, Desc: "曾经是某人的马桶盖子", Power: 4}}
	// e4 := &Equip{
	// 	Info: EquipInfo{EquipId: "1003", Name: "补刀斧", Level: 1, AD: 6, Desc: "无", Power: 4}}
	// e5 := &Equip{
	// 	Info: EquipInfo{EquipId: "1004", Name: "贵族头环", Level: 1, Strength: 2, Agility: 2, Intelligent: 2, Desc: "明明是屌丝头环啊亲！", Power: 5}}

	// EquipDefineList[e1.Info.EquipId] = e1
	// EquipDefineList[e2.Info.EquipId] = e2
	// EquipDefineList[e3.Info.EquipId] = e3
	// EquipDefineList[e4.Info.EquipId] = e4
	// EquipDefineList[e5.Info.EquipId] = e5
}

const (
	EquipStatusWearOff int8 = 0
	EquipStatusWearOn  int8 = 1
)

type Equip struct {
	Uid    string
	Info   EquipInfo
	Mix    []*Equip
	Status int8 // 0 未穿戴  1 穿上
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
}

func GetEquipDefines(equipIds []string) []*Equip {

	beego.Debug("GetEquipDefines ", equipIds)

	equips := make([]*Equip, 0)
	for _, equipId := range equipIds {
		if e, ok := EquipDefineList[equipId]; ok {
			res := new(Equip)
			tool.Clone(e, res)
			res.Uid = tool.UniqueId()
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
