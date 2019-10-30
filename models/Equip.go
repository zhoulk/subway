package models

var (
	EquipDefineList map[string]*Equip
)

func init() {
	EquipDefineList = make(map[string]*Equip)
	e1 := &Equip{
		Info: EquipInfo{EquipId: "1000", Strength: 1, Name: "树枝"}}
	e2 := &Equip{
		Info: EquipInfo{EquipId: "1001", Strength: 2, Name: "魔棒"}}
	e3 := &Equip{
		Info: EquipInfo{EquipId: "1002", Strength: 3, Name: "大魔棒"},
		Mix:  []*Equip{e1, e1, e1, e2}}

	EquipDefineList[e1.Info.EquipId] = e1
	EquipDefineList[e2.Info.EquipId] = e2
	EquipDefineList[e3.Info.EquipId] = e3
}

type Equip struct {
	Uid    string
	Info   EquipInfo
	Mix    []*Equip
	Status int8 // 0 未穿戴  1 穿上
}

type EquipInfo struct {
	EquipId  string
	Name     string
	Strength int32
}

func GetEquips(equipIds []string) []*Equip {
	equips := make([]*Equip, 0)
	for _, equipId := range equipIds {
		if e, ok := EquipDefineList[equipId]; ok {
			equips = append(equips, e)
		}
	}
	return equips
}
