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
	Expends    map[string]*ProductInfo
	Equips     map[string]*ProductInfo
	HeroParts  map[string]*ProductInfo
	EquipParts map[string]*ProductInfo
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
		Expends:    make(map[string]*ProductInfo),
		Equips:     make(map[string]*ProductInfo),
		HeroParts:  make(map[string]*ProductInfo),
		EquipParts: make(map[string]*ProductInfo),
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
		bagInfo.HeroParts[productInfo.ProductId] = productInfo
	}
	if productInfo.Type == ProductTypeExpend {
		bagInfo.Expends[productInfo.ProductId] = productInfo
	}
	if productInfo.Type == ProductTypeEquip {
		bagInfo.Equips[productInfo.ProductId] = productInfo
	}
	if productInfo.Type == ProductTypeEquipPart {
		bagInfo.EquipParts[productInfo.ProductId] = productInfo
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
		Expends:    make(map[string]*ProductInfo),
		Equips:     make(map[string]*ProductInfo),
		HeroParts:  make(map[string]*ProductInfo),
		EquipParts: make(map[string]*ProductInfo),
	}
	for _, t_item := range a.Items {
		if t_item.ItemType == ProductTypeHeroPart {
			bagInfo.HeroParts[t_item.Uid] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
		if t_item.ItemType == ProductTypeExpend {
			bagInfo.Expends[t_item.Uid] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
		if t_item.ItemType == ProductTypeEquip {
			bagInfo.Equips[t_item.Uid] = CreateProductInfoFromTableBagItemInfo(t_item)
		}
		if t_item.ItemType == ProductTypeEquipPart {
			bagInfo.EquipParts[t_item.Uid] = CreateProductInfoFromTableBagItemInfo(t_item)
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
