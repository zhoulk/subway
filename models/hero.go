package models

import (
	"subway/db/tables"
	"subway/tool"

	"github.com/astaxie/beego"
)

var (
	HeroDefineList map[string]*Hero

	HeroFloorDefine map[string]map[int16][]string
	HeroSkillDefine map[string][]string
)

func init() {
	HeroDefineList = make(map[string]*Hero)

	defines := tables.LoadHeroDefine()
	for _, def := range defines {
		HeroDefineList[def.HeroId] = CreateHeroFromHeroDefine(def)
	}

	HeroFloorDefine = make(map[string]map[int16][]string)
	heroEquipDefines := tables.LoadHeroEquipDefine()
	for _, def := range heroEquipDefines {
		if floor, ok := HeroFloorDefine[def.HeroId]; ok {
			if _, ok := floor[def.Floor]; ok {
				floor[def.Floor] = append(floor[def.Floor], def.EquipId)
			}
		} else {
			heroFloor := make(map[int16][]string)
			heroFloor[def.Floor] = []string{def.EquipId}
			HeroFloorDefine[def.HeroId] = heroFloor
		}
	}

	HeroSkillDefine = make(map[string][]string)
	heroSkillDefines := tables.LoadHeroSkillDefine()
	for _, def := range heroSkillDefines {
		if _, ok := HeroSkillDefine[def.HeroId]; ok {
			HeroSkillDefine[def.HeroId] = append(HeroSkillDefine[def.HeroId], def.SkillId)
		} else {
			HeroSkillDefine[def.HeroId] = []string{def.SkillId}
		}
	}
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
		if u.Heros == nil {
			u.Heros = make([]*Hero, 0)
			t_u_hs := tables.LoadUserHeros(u.Info.Uid)
			for _, t_u_h := range t_u_hs {

				hero := CreateHeroFromUserHero(t_u_h)

				t_h_es := tables.LoadHeroEquips(t_u_h.Uid)
				for _, t_h_e := range t_h_es {
					hero.Equips = append(hero.Equips, CreateEquipFromHeroEquip(t_h_e))
				}

				t_h_ss := tables.LoadHeroSkills(t_u_h.Uid)
				for _, t_h_s := range t_h_ss {
					hero.Skills = append(hero.Skills, CreateSkillFromHeroSkill(t_h_s))
				}

				u.Heros = append(u.Heros, hero)
			}
		}
		return u.Heros
	}
	return nil
}

func AddHero(uid string, heroId string) *Hero {
	u, _ := GetUser(uid)
	// beego.Debug(u)

	if u != nil {
		if u.Heros == nil {
			u.Heros = make([]*Hero, 0)
		}

		h := GetHeroDefine(heroId)
		// beego.Debug(h)

		beego.Debug(HeroFloorDefine)

		if h != nil {
			h.Info.LevelUpGold = h.Secret.OriginLevelUpGold
			h.Equips = GetEquipDefines(HeroFloorDefine[h.Info.HeroId][h.Info.Floor])
			h.Skills = GetSkillDefines(HeroSkillDefine[h.Info.HeroId])
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
			target.Equips = GetEquipDefines(HeroFloorDefine[target.Info.HeroId][target.Info.Floor])
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

func CreateHeroFromHeroDefine(def *tables.HeroDefine) *Hero {
	return &Hero{
		Info: HeroInfo{
			HeroId: def.HeroId,
			Type:   def.Type,
			Name:   def.Name,
			Level:  def.Level,
			Floor:  def.Floor,
			Star:   def.Star,
			Desc:   def.Desc,
		},
		Props: HeroProperties{
			HP:              def.HP,
			MP:              def.MP,
			AD:              def.AD,
			AP:              def.AP,
			ADDef:           def.ADDef,
			APDef:           def.APDef,
			SPD:             def.SPD,
			Agility:         def.Agility,
			Intelligent:     def.Intelligent,
			Strength:        def.Strength,
			ADCrit:          def.ADCrit,
			StrengthGrow:    def.StrengthGrow,
			AgilityGrow:     def.AgilityGrow,
			IntelligentGrow: def.IntelligentGrow,
		},
	}
}

func CreateHeroFromUserHero(t_u_h *tables.UserHero) *Hero {
	return &Hero{
		Uid: t_u_h.Uid,
		Info: HeroInfo{
			HeroId: t_u_h.HeroId,
			Level:  t_u_h.Level,
			Floor:  t_u_h.Floor,
			Star:   t_u_h.Star,
		},
		Props: HeroProperties{
			HP:              t_u_h.HP,
			MP:              t_u_h.MP,
			AD:              t_u_h.AD,
			AP:              t_u_h.AP,
			ADDef:           t_u_h.ADDef,
			APDef:           t_u_h.APDef,
			SPD:             t_u_h.SPD,
			Agility:         t_u_h.Agility,
			Intelligent:     t_u_h.Intelligent,
			Strength:        t_u_h.Strength,
			ADCrit:          t_u_h.ADCrit,
			StrengthGrow:    t_u_h.StrengthGrow,
			AgilityGrow:     t_u_h.AgilityGrow,
			IntelligentGrow: t_u_h.IntelligentGrow,
		},
		Equips: make([]*Equip, 0),
		Skills: make([]*Skill, 0),
	}
}

func CreateUserHeroFromHero(uid string, u_h *Hero) *tables.UserHero {
	return &tables.UserHero{
		Uid:             u_h.Uid,
		UserId:          uid,
		HeroId:          u_h.Info.HeroId,
		Level:           u_h.Info.Level,
		Floor:           u_h.Info.Floor,
		Star:            u_h.Info.Star,
		HP:              u_h.Props.HP,
		MP:              u_h.Props.MP,
		AD:              u_h.Props.AD,
		AP:              u_h.Props.AP,
		ADDef:           u_h.Props.ADDef,
		APDef:           u_h.Props.APDef,
		SPD:             u_h.Props.SPD,
		Agility:         u_h.Props.Agility,
		Intelligent:     u_h.Props.Intelligent,
		Strength:        u_h.Props.Strength,
		ADCrit:          u_h.Props.ADCrit,
		StrengthGrow:    u_h.Props.StrengthGrow,
		AgilityGrow:     u_h.Props.AgilityGrow,
		IntelligentGrow: u_h.Props.IntelligentGrow,
	}
}
