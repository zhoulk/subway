package models

var (
	CopyList []*Copy
)

func init() {
	CopyList = make([]*Copy, 0)

	items1 := []*CopyItem{&CopyItem{CopyItemId: 1001, Name: "苹果园"}}
	c1 := &Copy{Info: CopyInfo{CopyId: 1, Name: "1号线"}, Items: items1}
	c2 := &Copy{Info: CopyInfo{CopyId: 2, Name: "2号线"},
		Items: []*CopyItem{&CopyItem{
			CopyItemId: 2001,
			Name:       "苹果园",
			Goods: []*CopyGoodItem{
				&CopyGoodItem{
					Type:    1,
					GoodId:  "1000",
					Count:   1,
					Percent: 100,
				},
			},
		}}}
	CopyList = append(CopyList, c1)
	CopyList = append(CopyList, c2)
}

const (
	CopyStatusNone      int8 = 0
	CopyStatusNormal    int8 = 1
	CopyStatusLock      int8 = 2
	CopyStatusCompleted int8 = 3
)

type Copy struct {
	Uid    string
	Info   CopyInfo
	Star   int
	Status int8 // 0 未知   1 已解锁   2 未解锁  3 已通关
	Items  []*CopyItem
}

type CopyInfo struct {
	CopyId int
	Name   string
}

type CopyItem struct {
	CopyItemId int
	Name       string

	Goods []*CopyGoodItem
}

const (
	CopyGoodItemTypeEquip int8 = 1 // 装备
)

type CopyGoodItem struct {
	Type    int8
	GoodId  string
	Count   int8
	Percent int8
}

func GetAllCopy() []*Copy {
	return CopyList
}

func GetSelfCopy(uid string) []*Copy {
	u, _ := GetUser(uid)
	if u != nil {
		return u.Copys
	}
	return nil
}
