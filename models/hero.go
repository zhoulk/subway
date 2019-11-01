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
		Info:   HeroInfo{HeroId: "1000", Type: 2, Name: "敌法师", Level: 1, Floor: 1, Star: 2, Desc: "前排刺客，可以出现在敌人背后，对法师有致命的威胁。"},
		Props:  HeroProperties{HP: 100, AD: 10, SPD: 1000, StrengthGrow: 170, AgilityGrow: 420, IntelligentGrow: 270},
		Secret: HeroSecretInfo{OriginLevelUpGold: 1000, StepGold: 100, StepGold2: 2}}
	HeroDefineList[h1.Info.HeroId] = h1

	h2 := &Hero{
		Info:   HeroInfo{HeroId: "1001", Type: 2, Name: "小黑", Level: 1, Floor: 1, Star: 1, Desc: "后排物理输出，拥有强大的单体和群体伤害技能，但注意不要被人近身。"},
		Props:  HeroProperties{HP: 200, AD: 10, SPD: 1000, StrengthGrow: 200, AgilityGrow: 240, IntelligentGrow: 190},
		Secret: HeroSecretInfo{OriginLevelUpGold: 1000, StepGold: 100, StepGold2: 2}}
	HeroDefineList[h2.Info.HeroId] = h2

	h3 := &Hero{
		Info:   HeroInfo{HeroId: "1002", Type: 1, Name: "船长", Level: 1, Floor: 1, Star: 1, Desc: "前排坦克，能肉能输出能控场的全能英雄，无可争议的团队领袖。"},
		Props:  HeroProperties{HP: 300, AD: 10, SPD: 1000, StrengthGrow: 330, AgilityGrow: 130, IntelligentGrow: 220},
		Secret: HeroSecretInfo{OriginLevelUpGold: 1000, StepGold: 100, StepGold2: 2}}
	HeroDefineList[h3.Info.HeroId] = h3

	h4 := &Hero{
		Info:   HeroInfo{HeroId: "1003", Type: 1, Name: "火女", Level: 1, Floor: 1, Star: 1, Desc: "中排爆发型法师，娇弱的身体中蕴藏着恐怖的法力，技能很强很暴力。"},
		Props:  HeroProperties{HP: 150, AD: 10, SPD: 1000, StrengthGrow: 150, AgilityGrow: 320, IntelligentGrow: 150},
		Secret: HeroSecretInfo{OriginLevelUpGold: 1000, StepGold: 100, StepGold2: 2}}
	HeroDefineList[h4.Info.HeroId] = h4

	HeroFloorDefine = make(map[string]map[int16][]string)

	Hero1000Floor := make(map[int16][]string)
	Hero1000Floor[1] = []string{"1000", "1000", "1001", "1002", "1003", "1004"}
	HeroFloorDefine["1000"] = Hero1000Floor
	Hero2000Floor := make(map[int16][]string)
	Hero2000Floor[1] = []string{"1000", "1000", "1001", "1002", "1003", "1004"}
	HeroFloorDefine["1001"] = Hero2000Floor

	HeroSkillDefine = make(map[string][]string)
	HeroSkillDefine["1000"] = []string{"1001", "1002", "1003", "1004"}
	HeroSkillDefine["1001"] = []string{"1011", "1012", "1013", "1014"}
	HeroSkillDefine["1002"] = []string{"1021", "1022", "1023", "1024"}
	HeroSkillDefine["1003"] = []string{"1031", "1032", "1033", "1034"}
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

const (
	HeroTypeStrength    int8 = 1
	HeroTypeAgility     int8 = 2
	HeroTypeIntelligent int8 = 3
)

type HeroInfo struct {
	HeroId      string
	Type        int8
	Name        string
	Level       int32
	LevelUpGold int32
	Floor       int16 // 阶别
	Star        int16 // 星星
	Desc        string
}

type HeroProperties struct {
	HP              int32
	MP              int32
	AD              int32
	AP              int32
	ADDef           int32
	APDef           int32
	SPD             int32 // 毫秒数
	Agility         int32
	Intelligent     int32
	Strength        int32
	ADCrit          int32 // 物理暴击
	StrengthGrow    int32
	AgilityGrow     int32
	IntelligentGrow int32
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

func GetHeroDefine(heroId string) *Hero {
	if h, ok := HeroDefineList[heroId]; ok {
		res := new(Hero)
		tool.Clone(h, res)
		res.Uid = tool.UniqueId()
		return res
	}
	return nil
}

func GetSelfHeros(uid string) []*Hero {
	u, _ := GetUser(uid)
	if u != nil {
		return u.Heros
	}
	return nil
}

func AddHero(uid string, heroId string) *Hero {
	u, _ := GetUser(uid)
	if u != nil {
		if u.Heros == nil {
			u.Heros = make([]*Hero, 0)
		}

		h := GetHeroDefine(heroId)
		if h != nil {
			h.Info.Floor = 1
			h.Info.LevelUpGold = h.Secret.OriginLevelUpGold
			h.Equips = GetEquips(HeroFloorDefine[h.Info.HeroId][h.Info.Floor])
			h.Skills = GetSkills(HeroSkillDefine[h.Info.HeroId])
			u.Heros = append(u.Heros, h)
			return h
		}
	}
	return nil
}

func HeroLevelUp(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)
	u, _ := GetUser(uid)

	if target != nil {
		endLevel := target.Info.Level + 1
		// 计算升级需要的金币
		needGold := int64(target.Secret.OriginLevelUpGold + target.Secret.StepGold*endLevel + endLevel*endLevel)
		if u.Profile.Gold >= needGold {
			target.SetHeroLevel(endLevel)
			u.Profile.Gold -= needGold
			return true
		}
	}

	return false
}

func HeroFloorUp(uid string, heroUid string) bool {
	var target *Hero = GetHero(uid, heroUid)

	if target != nil {
		// 装备是否穿戴完毕
		var isAll = true
		for _, e := range target.Equips {
			if e.Status == 0 {
				isAll = false
			}
		}
		if isAll {
			target.SetFloorLevel(target.Info.Floor + 1)
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

func (h *Hero) SetHeroLevel(level int32) {
	h.Info.Level = level
}

func (h *Hero) SetFloorLevel(floor int16) {
	h.Info.Floor = floor
}

func (h *Hero) SetStar(star int16) {
	h.Info.Star = star
}
