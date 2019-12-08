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
	Equips    map[int]*BagItem
	HeroParts map[int]*BagItem

	Items []*BagItem
}

const (
	BagItemEquip    int8 = 1
	BagItemHeroPart int8 = 2
	BagItemOther    int8 = 3
)

type BagItem struct {
	Uid string

	Type    int8
	GoodsId int
	Count   int
	Name    string
	Cost    int32
	Desc    string

	EquipInfo EquipInfo
	HeroInfo  HeroInfo
	HeroProps HeroProperties
}

func GetBag(uid string) *Bag {
	if b, ok := BagList[uid]; ok {
		return b
	}

	b := new(Bag)
	b.Equips = make(map[int]*BagItem)
	b.HeroParts = make(map[int]*BagItem)

	t_u_bs := tables.LoadUserBags(uid)
	for _, t_u_b := range t_u_bs {
		item := CreateBagItemFromUserBag(t_u_b)
		if item.Type == BagItemEquip {
			completeBagItemEquip(item)
			b.Equips[item.GoodsId] = item
		} else if item.Type == BagItemHeroPart {
			completeBagItemHeroPart(item)
			b.HeroParts[item.GoodsId] = item
		}
		b.Items = append(b.Items, item)
	}

	BagList[uid] = b
	return b
}

func BagContainEquip(uid string, equipId string) bool {
	b := GetBag(uid)
	if b != nil {
		for _, e := range b.Equips {
			if strconv.Itoa(e.GoodsId) == equipId && e.Count > 0 {
				return true
			}
		}
	}
	return false
}

func GetBagItemOfHeroPart(uid string, heroId string) *BagItem {
	b := GetBag(uid)
	num, err := strconv.Atoi(heroId)
	if err == nil {
		if item, ok := b.HeroParts[num]; ok {
			return item
		}
	}
	return nil
}

// 获得一个物品
func GainABagItem(uid string, item *BagItem) {
	b := GetBag(uid)

	item.Uid = tool.UniqueId()
	if item.Type == BagItemEquip {
		if _, ok := b.Equips[item.GoodsId]; ok {
			b.Equips[item.GoodsId].Count += item.Count
		} else {
			completeBagItemEquip(item)
			b.Equips[item.GoodsId] = item
			b.Items = append(b.Items, item)
		}
	} else if item.Type == BagItemHeroPart {
		if _, ok := b.HeroParts[item.GoodsId]; ok {
			b.HeroParts[item.GoodsId].Count += item.Count
		} else {
			completeBagItemHeroPart(item)
			b.HeroParts[item.GoodsId] = item
			b.Items = append(b.Items, item)
		}
	}
}

// 消耗物品
func UseAEquip(uid string, equipId string) bool {
	b := GetBag(uid)
	if b != nil {
		for _, e := range b.Equips {
			if strconv.Itoa(e.GoodsId) == equipId && e.Count > 0 {
				e.Count--
				return true
			}
		}
	}
	return false
}

// 消耗英雄碎片
func UseHeroParts(uid string, heroId string, count int) bool {
	b := GetBag(uid)
	if b != nil {
		for _, e := range b.HeroParts {
			if strconv.Itoa(e.GoodsId) == heroId && e.Count >= count {
				e.Count -= count
				return true
			}
		}
	}
	return false
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

func CreateBagItemFromUserBag(t_u_b *tables.UserBag) *BagItem {
	return &BagItem{
		Uid:     t_u_b.Uid,
		Type:    t_u_b.ItemType,
		GoodsId: t_u_b.ItemId,
		Count:   t_u_b.Count,
	}
}

func completeBagItemEquip(item *BagItem) {
	if e, ok := EquipDefineList[strconv.Itoa(item.GoodsId)]; ok {
		item.Name = e.Info.Name
		item.Desc = e.Info.Desc
		item.Cost = e.Info.Cost

		item.EquipInfo = e.Info
	}
}

func completeBagItemHeroPart(item *BagItem) {
	if h, ok := HeroDefineList[strconv.Itoa(item.GoodsId)]; ok {
		item.Name = h.Info.Name + "(碎片)"
		item.Desc = h.Info.Desc
		item.Cost = 1

		item.HeroInfo = h.Info
		item.HeroProps = h.Props
	}
}
