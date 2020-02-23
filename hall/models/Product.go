package models

import "subway/tool"

const (
	ProductTypeHero      int8 = 1
	ProductTypeHeroPart  int8 = 2
	ProductTypeExpend    int8 = 3
	ProductTypeEquip     int8 = 4
	ProductTypeEquipPart int8 = 5
)

type ProductInfo struct {
	ProductId string
	ItemId    int32
	Type      int8
	Name      string
	Count     int32
}

func CreateAProduct(t int8) *ProductInfo {
	return &ProductInfo{
		ProductId: tool.UniqueId(),
		ItemId:    1002,
		Type:      t,
		Name:      "测试",
	}
}
