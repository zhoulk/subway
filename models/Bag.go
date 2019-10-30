package models

var (
	BagList map[string]*Bag
)

func init() {
	BagList = make(map[string]*Bag)
}

type Bag struct {
	Equips []*Equip
}

func GetBag(uid string) *Bag {
	if b, ok := BagList[uid]; ok {
		return b
	}

	b := new(Bag)
	BagList[uid] = b
	return b
}

func BagContainEquip(uid string, equipId string) bool {
	b := GetBag(uid)
	if b != nil {
		for _, e := range b.Equips {
			if e.Info.EquipId == equipId {
				return true
			}
		}
	}
	return false
}
