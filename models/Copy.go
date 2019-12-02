package models

import (
	"subway/db/tables"
	"subway/tool"

	"github.com/astaxie/beego"
)

var (
	CopyList     []*Copy
	CopyItemList []*CopyItem
	CopyItemDic  map[int]*CopyItem
	CopyItems    map[int][]*CopyItem
)

func init() {
	CopyList = make([]*Copy, 0)
	CopyItemList = make([]*CopyItem, 0)
	CopyItemDic = make(map[int]*CopyItem)
	CopyItems = make(map[int][]*CopyItem)

	defines := tables.LoadCopyData()
	for _, def := range defines {
		itemDefines := tables.LoadCopyItems(def.CopyId)

		items := make([]*CopyItem, 0)
		for _, itemDef := range itemDefines {
			heros := make([]*GuanKaHero, 0)
			for _, gkHero := range itemDef.Heros {
				heros = append(heros, &GuanKaHero{HeroId: gkHero.HeroId, Level: gkHero.Level, Floor: gkHero.Floor, Star: gkHero.Star, SkillLevels: gkHero.SkillLevels})
			}

			goods := make([]*CopyGoodItem, 0)
			for _, cpEquip := range itemDef.Equips {
				goods = append(goods, &CopyGoodItem{
					Type:    CopyGoodItemTypeEquip,
					GoodId:  cpEquip,
					Count:   1,
					Percent: 100,
				})
			}

			cpItem := &CopyItem{
				CopyItemId: itemDef.CopyId,
				Name:       itemDef.Name,
				TotalStar:  3,
				Heros:      heros,
				Goods:      goods,
				Status:     CopyStatusLock,
			}
			items = append(items, cpItem)

			CopyItemList = append(CopyItemList, cpItem)
			CopyItemDic[cpItem.CopyItemId] = cpItem
		}
		CopyItems[def.CopyId] = items

		cp := &Copy{
			Info: CopyInfo{
				CopyId: def.CopyId,
				Name:   def.Name,
			},
			TotalStar: len(itemDefines) * 3,
			Status:    CopyStatusLock,
		}
		CopyList = append(CopyList, cp)
	}
}

const (
	CopyStatusNone      int8 = 0
	CopyStatusNormal    int8 = 1
	CopyStatusLock      int8 = 2
	CopyStatusCompleted int8 = 3
)

type Copy struct {
	Uid       string
	Info      CopyInfo
	Star      int
	TotalStar int
	Status    int8 // 0 未知   1 已解锁   2 未解锁  3 已通关
}

type CopyInfo struct {
	CopyId int
	Name   string
}

type CopyItem struct {
	Uid        string
	CopyItemId int
	Name       string
	Star       int
	TotalStar  int
	Status     int8 // 0 未知   1 已解锁   2 未解锁  3 已通关

	Heros []*GuanKaHero
	Goods []*CopyGoodItem
}

const (
	CopyGoodItemTypeEquip int8 = 1 // 装备
)

type CopyGoodItem struct {
	Type    int8
	GoodId  int
	Count   int
	Percent int8
}

func GetAllCopy() []*Copy {
	return CopyList
}

func GetCopyItems(copyId int) []*CopyItem {
	return CopyItems[copyId]
}

func GetSelfCopy(uid string) ([]*Copy, map[int]*Copy) {
	u, _ := GetUser(uid)
	if u != nil {
		if u.Copys == nil {
			u.Copys = make([]*Copy, 0)
			u.CopyDic = make(map[int]*Copy)
			t_u_cs := tables.LoadUserCopys(u.Info.Uid)
			if len(t_u_cs) > 0 {
				for _, t_u_c := range t_u_cs {
					cp := CreateCopyFromUserCopy(t_u_c)
					u.AddCopy(cp)
				}
			} else {
				cp := new(Copy)
				tool.Clone(CopyList[0], cp)
				cp.Uid = tool.UniqueId()
				cp.Status = CopyStatusNormal
				u.AddCopy(cp)
			}
		}
		return u.Copys, u.CopyDic
	}
	return nil, nil
}

func GetSelfCopyItem(uid string) ([]*CopyItem, map[int]*CopyItem) {
	u, _ := GetUser(uid)
	if u != nil {
		if u.CopyItems == nil {
			u.CopyItems = make([]*CopyItem, 0)
			u.CopyItemDic = make(map[int]*CopyItem)
			t_u_c_is := tables.LoadUserCopyItems(u.Info.Uid)
			if len(t_u_c_is) > 0 {
				for _, t_u_c_i := range t_u_c_is {
					cp := CreateCopyItemFromUserCopyItem(t_u_c_i)
					u.AddCopyItem(cp)
				}
			} else {
				cp := new(CopyItem)
				tool.Clone(CopyItemList[0], cp)
				cp.Uid = tool.UniqueId()
				cp.Status = CopyStatusNormal
				u.AddCopyItem(cp)
			}
		}
		return u.CopyItems, u.CopyItemDic
	}
	return nil, nil
}

func GetCopyItemDefine(cpId int) *CopyItem {
	beego.Debug("GetCopyItemDefine  ", cpId)
	if cp, ok := CopyItemDic[cpId]; ok {
		if cp.Heros != nil {
			for _, h := range cp.Heros {
				if h.Hero == nil {
					h.Hero = GetHeroDefine(h.HeroId)
					h.Hero.SetHeroLevel(h.Level)
					h.Hero.SetFloorLevel(h.Floor)
					h.Hero.SetStar(h.Star)
					h.Hero.Skills = GetSkillDefines(HeroSkillDefine[h.HeroId])
					if h.Hero.Skills != nil {
						for i, s := range h.Hero.Skills {
							if i < len(h.SkillLevels) {
								s.SetSkillLevel(h.SkillLevels[i])
							}
						}
					}
				}
			}
		}
		return cp
	}
	return nil
}

// 过关
func CompleteCopy(uid string, cpId int) {
	beego.Debug("CompleteCopy  ", cpId)

	u, _ := GetUser(uid)
	if u != nil {
		if cpItem, ok := u.CopyItemDic[cpId]; ok {
			cpItem.Status = CopyStatusCompleted
			cpItem.Star = 3
		}

		// 刷新章节 star  状态
		for _, cp := range u.Copys {
			if cp.Status == CopyStatusNormal {
				star := 0
				status := CopyStatusCompleted
				for _, cpItemDef := range GetCopyItems(cp.Info.CopyId) {
					if cpItem, ok := u.CopyItemDic[cpItemDef.CopyItemId]; ok {
						if cpItem.Status == CopyStatusNormal {
							status = CopyStatusNormal
						}
						star += cpItem.Star
					} else {
						status = CopyStatusNormal
					}
				}
				cp.Status = status
				cp.Star = star
			}
		}

		nextCpId := getNextCopyItemId(cpId)
		if nextCpId != cpId {
			if cpItem, ok := CopyItemDic[nextCpId]; ok {
				cp := new(CopyItem)
				tool.Clone(cpItem, cp)
				cp.Uid = tool.UniqueId()
				cp.Status = CopyStatusNormal
				u.AddCopyItem(cp)
			}
		}
	}
}

// 获取下一关副本
func getNextCopyItemId(cpId int) int {
	for index, cp := range CopyItemList {
		if cp.CopyItemId == cpId {
			if index+1 < len(CopyItemList) {
				return CopyItemList[index+1].CopyItemId
			}
		}
	}
	return cpId
}

func CreateCopyFromUserCopy(t_u_c *tables.UserCopy) *Copy {
	return &Copy{
		Uid: t_u_c.Uid,
		Info: CopyInfo{
			CopyId: t_u_c.CopyId,
		},
		Star:   t_u_c.Star,
		Status: t_u_c.Status,
	}
}

func CreateUserCopyFromCopy(uid string, cp *Copy) *tables.UserCopy {
	return &tables.UserCopy{
		Uid:    cp.Uid,
		UserId: uid,
		CopyId: cp.Info.CopyId,
		Star:   cp.Star,
		Status: cp.Status,
	}
}

func CreateCopyItemFromUserCopyItem(t_u_c_i *tables.UserCopyItem) *CopyItem {
	return &CopyItem{
		Uid:        t_u_c_i.Uid,
		CopyItemId: t_u_c_i.CopyItemId,
		Star:       t_u_c_i.Star,
		Status:     t_u_c_i.Status,
	}
}

func CreateUserCopyItemFromCopyItem(uid string, cpItem *CopyItem) *tables.UserCopyItem {
	return &tables.UserCopyItem{
		Uid:        cpItem.Uid,
		UserId:     uid,
		CopyItemId: cpItem.CopyItemId,
		Star:       cpItem.Star,
		Status:     cpItem.Status,
	}
}
