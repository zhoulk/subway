package models

import (
	tables "subway/hall/db"
	"subway/tool"
)

var (
	BagList map[string]*BagInfo
)

func init() {
	BagList = make(map[string]*BagInfo)
}

type BagInfo struct {
	RoleId     string
	Expends    map[int32]*ProductInfo
	Equips     map[int32]*ProductInfo
	HeroParts  map[int32]*ProductInfo
	EquipParts map[int32]*ProductInfo
}

func GetBag(roleId string) *BagInfo {
	if b, ok := BagList[roleId]; ok {
		return b
	}

	t_b := tables.LoadBagInfo(roleId)
	if t_b != nil {
		b := CreateBagInfoFromTableBagInfo(t_b)
		BagList[roleId] = b
		return BagList[roleId]
	}

	bagInfo := &BagInfo{
		RoleId:     roleId,
		Expends:    make(map[int32]*ProductInfo),
		Equips:     make(map[int32]*ProductInfo),
		HeroParts:  make(map[int32]*ProductInfo),
		EquipParts: make(map[int32]*ProductInfo),
	}
	BagList[roleId] = bagInfo

	return bagInfo
}

func AddProduct(roleId string, productInfo *ProductInfo) {
	if len(productInfo.ProductId) == 0 {
		productInfo.ProductId = tool.UniqueId()
	}

	bagInfo := GetBag(roleId)
	if productInfo.Type == ProductTypeHeroPart {
		if p, ok := bagInfo.HeroParts[productInfo.ItemId]; ok {
			p.Count += productInfo.Count
		} else {
			bagInfo.HeroParts[productInfo.ItemId] = productInfo
		}
	}
	if productInfo.Type == ProductTypeExpend {
		if p, ok := bagInfo.Expends[productInfo.ItemId]; ok {
			p.Count += productInfo.Count
		} else {
			bagInfo.Expends[productInfo.ItemId] = productInfo
		}
	}
	if productInfo.Type == ProductTypeEquip {
		if p, ok := bagInfo.Equips[productInfo.ItemId]; ok {
			p.Count += productInfo.Count
		} else {
			bagInfo.Equips[productInfo.ItemId] = productInfo
		}
	}
	if productInfo.Type == ProductTypeEquipPart {
		if p, ok := bagInfo.EquipParts[productInfo.ItemId]; ok {
			p.Count += productInfo.Count
		} else {
			bagInfo.EquipParts[productInfo.ItemId] = productInfo
		}
	}
}

func PersistentBagInfo() {
	bagItems := make([]*tables.BagItemInfo, 0)
	for _, a := range BagList {
		bagInfo := CreateTableBagInfoFromBagInfo(a)
		bagItems = append(bagItems, bagInfo.Items...)
	}
	tables.PersistentBagItems(bagItems)
}

func CreateBagInfoFromTableBagInfo(a *tables.BagInfo) *BagInfo {
	bagInfo := &BagInfo{
		RoleId:     a.RoleId,
		Expends:    make(map[int32]*ProductInfo),
		Equips:     make(map[int32]*ProductInfo),
		HeroParts:  make(map[int32]*ProductInfo),
		EquipParts: make(map[int32]*ProductInfo),
	}
	for _, t_item := range a.Items {
		if t_item.ItemType == ProductTypeHeroPart {
			bagInfo.HeroParts[t_item.ItemId] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
		if t_item.ItemType == ProductTypeExpend {
			bagInfo.Expends[t_item.ItemId] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
		if t_item.ItemType == ProductTypeEquip {
			bagInfo.Equips[t_item.ItemId] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
		if t_item.ItemType == ProductTypeEquipPart {
			bagInfo.EquipParts[t_item.ItemId] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
	}
	return bagInfo
}

func CreateTableBagInfoFromBagInfo(a *BagInfo) *tables.BagInfo {
	bagInfo := &tables.BagInfo{
		RoleId: a.RoleId,
		Items:  make([]*tables.BagItemInfo, 0),
	}
	for _, t_item := range a.Expends {
		bagInfo.Items = append(bagInfo.Items, CreateTableBagItemInfoFromProductInfo(a.RoleId, t_item))
	}
	for _, t_item := range a.HeroParts {
		bagInfo.Items = append(bagInfo.Items, CreateTableBagItemInfoFromProductInfo(a.RoleId, t_item))
	}
	for _, t_item := range a.Equips {
		bagInfo.Items = append(bagInfo.Items, CreateTableBagItemInfoFromProductInfo(a.RoleId, t_item))
	}
	for _, t_item := range a.EquipParts {
		bagInfo.Items = append(bagInfo.Items, CreateTableBagItemInfoFromProductInfo(a.RoleId, t_item))
	}
	return bagInfo
}

func CreateTableBagItemInfoFromProductInfo(roleId string, a *ProductInfo) *tables.BagItemInfo {
	return &tables.BagItemInfo{
		Uid:      a.ProductId,
		RoleId:   roleId,
		ItemId:   a.ItemId,
		ItemType: a.Type,
		Count:    a.Count,
	}
}

func CreateProductInfoFromTableBagItemInfo(a *tables.BagItemInfo) *ProductInfo {
	return &ProductInfo{
		ProductId: a.Uid,
		ItemId:    a.ItemId,
		Type:      a.ItemType,
		Count:     a.Count,
	}
}
