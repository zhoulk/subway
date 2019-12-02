package models

import (
	"strconv"
	"subway/db/tables"
	"subway/tool"
)

var (
	BagList map[string]*Bag
)

func init() {
	BagList = make(map[string]*Bag)
}

type Bag struct {
	Equips map[int]*BagItem

	Items []*BagItem
}

const (
	BagItemEquip int8 = 1
)

type BagItem struct {
	Uid string

	Type    int8
	GoodsId int
	Count   int
}

func GetBag(uid string) *Bag {
	if b, ok := BagList[uid]; ok {
		return b
	}

	b := new(Bag)
	b.Equips = make(map[int]*BagItem)
	BagList[uid] = b
	return b
}

func BagContainEquip(uid string, equipId string) bool {
	b := GetBag(uid)
	if b != nil {
		for _, e := range b.Equips {
			if strconv.Itoa(e.GoodsId) == equipId {
				return true
			}
		}
	}
	return false
}

// 获得一个物品
func GainABagItem(uid string, item *BagItem) {
	b := GetBag(uid)

	item.Uid = tool.UniqueId()
	if item.Type == BagItemEquip {
		if _, ok := b.Equips[item.GoodsId]; ok {
			b.Equips[item.GoodsId].Count += item.Count
		} else {
			b.Equips[item.GoodsId] = item
			b.Items = append(b.Items, item)
		}
	}
}

func CreateUserBagFromBagItem(uid string, bagItem *BagItem) *tables.UserBag {
	return &tables.UserBag{
		Uid:      bagItem.Uid,
		UserId:   uid,
		ItemId:   bagItem.GoodsId,
		Count:    bagItem.Count,
		ItemType: bagItem.Type,
	}
}
