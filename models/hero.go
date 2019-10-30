package models

import (
	"subway/tool"
)

var (
	HeroDefineList map[string]*Hero

	HeroFloorDefine map[string]map[int16][]string
	HeroSkillDefine map[string][]string
)

func init() {
	HeroDefineList = make(map[string]*Hero)
	h1 := &Hero{
		Info:   HeroInfo{HeroId: "1000", Name: "敌法师", Level: 1, Desc: "第一个英雄"},
		Props:  HeroProperties{HP: 100, AD: 20},
		Secret: HeroSecretInfo{OriginLevelUpGold: 1000, StepGold: 100, StepGold2: 2}}
	HeroDefineList[h1.Info.HeroId] = h1

	HeroFloorDefine = make(map[string]map[int16][]string)

	Hero1000Floor := make(map[int16][]string)
	Hero1000Floor[1] = []string{"1000", "1000", "1000", "1000", "1002", "1001"}
	Hero1000Floor[2] = []string{"1001", "1001", "1001", "1001", "1002", "1001"}
	Hero1000Floor[3] = []string{"1002", "1002", "1002", "1001", "1002", "1001"}
	HeroFloorDefine["1000"] = Hero1000Floor

	HeroSkillDefine = make(map[string][]string)
	HeroSkillDefine["1000"] = []string{"1000"}
}

type Hero struct {
	Uid    string
	Info   HeroInfo
	Props  HeroProperties
	Secret HeroSecretInfo
	Equips []*Equip
	Skills []*Skill
	Status int8 // 0 正常  1 上阵
}

type HeroInfo struct {
	HeroId      string
	Name        string
	Level       int32
	LevelUpGold int32
	Floor       int16
	Desc        string
}

type HeroProperties struct {
	HP    int
	AD    int
	AP    int
	ADDef int
	APDef int
	SPD   int
}

type HeroSecretInfo struct {
	OriginLevelUpGold int32
	StepGold          int32
	StepGold2         int32
}

func GetAllHeros() []*Hero {
	heros := make([]*Hero, 0)
	for _, h := range HeroDefineList {
		heros = append(heros, h)
	}
	return heros
}

func GetSelfHeros(uid string) []*Hero {
	u, _ := GetUser(uid)
	if u != nil {
		return u.Heros
	}
	return nil
}

func AddHero(uid string, heroId string) bool {
	u, _ := GetUser(uid)
	if u != nil {
		if u.Heros == nil {
			u.Heros = make([]*Hero, 0)
		}
		if h, ok := HeroDefineList[heroId]; ok {
			h.Uid = tool.UniqueId()
			h.Info.Floor = 1
			h.Info.LevelUpGold = h.Secret.OriginLevelUpGold
			h.Equips = GetEquips(HeroFloorDefine[h.Info.HeroId][h.Info.Floor])
			h.Skills = GetSkills(HeroSkillDefine[h.Info.HeroId])
			u.Heros = append(u.Heros, h)
			return true
		}
	}
	return false
}

func HeroLevelUp(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)
	u, _ := GetUser(uid)

	if target != nil {
		endLevel := target.Info.Level + 1
		// 计算升级需要的金币
		needGold := int64(target.Secret.OriginLevelUpGold + target.Secret.StepGold*endLevel + endLevel*endLevel)
		if u.Profile.Gold >= needGold {
			target.Info.Level = endLevel
			u.Profile.Gold -= needGold
			return true
		}
	}

	return false
}

func HeroFloorUp(uid string, heroUid string) bool {
	var target *Hero = GetHero(uid, heroUid)

	if target != nil {
		endFloor := target.Info.Floor + 1
		// 装备是否穿戴完毕
		var isAll = true
		for _, e := range target.Equips {
			if e.Status == 0 {
				isAll = false
			}
		}
		if isAll {
			target.Info.Floor = endFloor
			target.Equips = GetEquips(HeroFloorDefine[target.Info.HeroId][target.Info.Floor])
			return true
		}
	}

	return false
}

func Wear(uid string, heroUid string, equipId string) bool {
	var target *Hero = GetHero(uid, heroUid)

	if target != nil {
		// 装备是否穿戴完毕
		var targetEquip *Equip
		for _, e := range target.Equips {
			if e.Status == 0 && e.Info.EquipId == equipId {
				targetEquip = e
				break
			}
		}
		// 是否拥有
		if targetEquip != nil {
			if BagContainEquip(uid, equipId) {
				targetEquip.Status = 1
				return true
			}
		}
	}
	return false
}

func GetSelectedHeros(uid string) []*Hero {
	res := make([]*Hero, 0)
	u, _ := GetUser(uid)
	if u != nil {
		heros := u.Heros
		if heros != nil {
			for _, h := range heros {
				if h.Status == 1 {
					res = append(res, h)
				}
			}
		}
	}
	return res
}

func SelectedHerosCount(uid string) int {
	res := 0
	u, _ := GetUser(uid)
	if u != nil {
		heros := u.Heros
		if heros != nil {
			for _, h := range heros {
				if h.Status == 1 {
					res++
				}
			}
		}
	}
	return res
}

func GetHero(uid string, heroUid string) *Hero {
	var target *Hero

	u, _ := GetUser(uid)
	if u != nil {
		heros := u.Heros
		if heros != nil {
			for _, h := range heros {
				if h.Uid == heroUid {
					target = h
				}
			}
		}
	}
	return target
}

func SelectHero(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)

	if target != nil {
		if SelectedHerosCount(uid) < 5 {
			target.Status = 1
			return true
		}
	}

	return false
}

func UnSelectHero(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)

	if target != nil {
		target.Status = 0
		return true
	}

	return false
}

func ExchangeHero(uid string, fromHeroUid string, toHeroUid string) bool {
	if UnSelectHero(uid, fromHeroUid) {
		return SelectHero(uid, toHeroUid)
	}
	return false
}
