package models

import (
	"subway/db/tables"
	"subway/tool"
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
		h := CreateHeroFromHeroDefine(def)
		h.Status = HeroStatusPart
		HeroDefineList[def.HeroId] = h
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
	Status int8 // 1 正常  2 上阵
}

const (
	HeroStatusNormal   int8 = 1
	HeroStatusSelected int8 = 2
	HeroStatusPart     int8 = 3 //碎片模式  未获得
)

const (
	HeroTypeStrength    int8 = 1
	HeroTypeAgility     int8 = 2
	HeroTypeIntelligent int8 = 3
)

const (
	HeroAtkTypeAD int8 = 1
	HeroAtkTypeAP int8 = 2
)

type HeroInfo struct {
	HeroId      string
	AtkType     int8
	Type        int8
	Name        string
	Level       int32
	LevelUpGold int32
	Floor       int16 // 阶别
	Star        int16 // 星星
	Desc        string
	Parts       int32 // 碎片数
	StarUp      int32 // 升星碎片数
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

func (h *Hero) SetHeroLevel(level int32) {
	h.Info.Level = level

	// 增加属性
	RefreshHero(h)

	// 修改升级需要的金币
	nextLevel := h.Info.Level + 1
	h.Info.LevelUpGold = int32(h.Secret.OriginLevelUpGold + h.Secret.StepGold*nextLevel + nextLevel*nextLevel)
}

func (h *Hero) SetFloorLevel(floor int16) {
	h.Info.Floor = floor
}

func (h *Hero) SetStar(star int16) {
	h.Info.Star = star

	h.Info.StarUp = int32(20 + (h.Info.Star-1)*10)
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
		res.Info.LevelUpGold = h.Secret.OriginLevelUpGold
		return res
	}
	return nil
}

// 获取已获得英雄
func GetSelfHeros(uid string) ([]*Hero, map[string]*Hero) {
	u, _ := GetUser(uid)
	if u != nil {
		if u.Heros == nil {
			u.Heros = make([]*Hero, 0)
			u.HeroDic = make(map[string]*Hero)
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
				u.HeroDic[hero.Info.HeroId] = hero
			}
		} else {

		}
		return u.Heros, u.HeroDic
	}
	return nil, nil
}

// 获取未获得英雄
func GetUnCollectHeros(uid string) []*Hero {
	_, selfHeros := GetSelfHeros(uid)
	heros := make([]*Hero, 0)
	if selfHeros != nil {
		for _, h := range HeroDefineList {
			if _, ok := selfHeros[h.Info.HeroId]; !ok {

				// 计算碎片
				bagItem := GetBagItemOfHeroPart(uid, h.Info.HeroId)
				if bagItem != nil {
					h.Info.Parts = int32(bagItem.Count)
				}
				h.SetStar(h.Info.Star)
				heros = append(heros, h)
			}
		}
	}

	return heros
}

func AddHero(uid string, heroId string) *Hero {
	u, _ := GetUser(uid)
	// beego.Debug(u)

	if u != nil {
		if u.Heros == nil {
			u.Heros = make([]*Hero, 0)
			u.HeroDic = make(map[string]*Hero)
		}

		if target, ok := u.HeroDic[heroId]; ok {
			target.Status = HeroStatusNormal
			return target
		} else {
			h := GetHeroDefine(heroId)
			h.Status = HeroStatusNormal
			if h != nil {
				h.Equips = GetEquipDefines(HeroFloorDefine[h.Info.HeroId][h.Info.Floor])
				h.Skills = GetSkillDefines(HeroSkillDefine[h.Info.HeroId])
				u.Heros = append(u.Heros, h)
				u.HeroDic[h.Info.HeroId] = h
				return h
			}
		}
	}
	return nil
}

// 英雄升级
func HeroLevelUp(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)
	u, _ := GetUser(uid)

	if target != nil {
		endLevel := target.Info.Level + 1
		// 计算升级需要的金币
		needGold := int64(target.Info.LevelUpGold) //int64(target.Secret.OriginLevelUpGold + target.Secret.StepGold*endLevel + endLevel*endLevel)
		if u.Profile.Gold >= needGold {
			target.SetHeroLevel(endLevel)
			u.Profile.Gold -= needGold
			return true
		}
	}

	return false
}

// 英雄升星
func HeroStarUp(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)
	if target != nil {
		bagItem := GetBagItemOfHeroPart(uid, target.Info.HeroId)
		if bagItem != nil {
			if int32(bagItem.Count) >= target.Info.StarUp {
				if UseHeroParts(uid, target.Info.HeroId, int(target.Info.StarUp)) {
					target.SetStar(target.Info.Star + 1)
					return true
				}
			}
		}
	}

	return false
}

// 英雄进阶
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

// 英雄合成
func HeroCompose(uid string, heroId string) bool {
	bagItem := GetBagItemOfHeroPart(uid, heroId)
	if bagItem != nil {
		if def, ok := HeroDefineList[heroId]; ok {
			def.SetStar(def.Info.Star)
			if int32(bagItem.Count) >= def.Info.StarUp {
				if UseHeroParts(uid, heroId, int(def.Info.StarUp)) {
					AddHero(uid, heroId)
					return true
				}
			}
		}
	}
	return false
}

func Wear(uid string, heroUid string, equipUid string) bool {
	var target *Hero = GetHero(uid, heroUid)

	if target != nil {
		// 装备是否穿戴完毕
		var targetEquip *Equip
		for _, e := range target.Equips {
			if e.Status != EquipStatusWearComplete && e.Uid == equipUid {
				targetEquip = e
				break
			}
		}
		// 是否拥有
		if targetEquip != nil {
			if UseAEquip(uid, targetEquip.Info.EquipId) {
				targetEquip.Status = EquipStatusWearComplete
				return true
			}
		}
	}
	return false
}

func GetSelectedHeros(uid string) []*Hero {
	res := make([]*Hero, 0)
	heros, _ := GetSelfHeros(uid)
	if heros != nil {
		for _, h := range heros {
			if h.Status == HeroStatusSelected {
				res = append(res, h)
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
				if h.Status == HeroStatusSelected {
					res++
				}
			}
		}
	}
	return res
}

// 获取英雄详情
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

	// 计算装备状态
	for _, e := range target.Equips {
		if e.Status != EquipStatusWearComplete {
			if BagContainEquip(uid, e.Info.EquipId) {
				e.Status = EquipStatusWearAcquire
			}
		}
	}
	// 计算碎片
	bagItem := GetBagItemOfHeroPart(uid, target.Info.HeroId)
	if bagItem != nil {
		target.Info.Parts = int32(bagItem.Count)
	}

	return target
}

func SelectHero(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)

	if target != nil {
		if SelectedHerosCount(uid) < 5 {
			target.Status = HeroStatusSelected
			return true
		}
	}

	return false
}

func UnSelectHero(uid string, heroUid string) bool {
	target := GetHero(uid, heroUid)

	if target != nil {
		target.Status = HeroStatusNormal
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

func CreateHeroFromHeroDefine(def *tables.HeroDefine) *Hero {
	h := &Hero{
		Info: HeroInfo{
			HeroId:  def.HeroId,
			Type:    def.Type,
			AtkType: def.AtkType,
			Name:    def.Name,
			Floor:   def.Floor,
			Desc:    def.Desc,
		},
	}
	h.Props = HeroProperties{
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
	}

	h.SetStar(def.Star)

	return h
}

func CreateHeroFromUserHero(t_u_h *tables.UserHero) *Hero {
	if h, ok := HeroDefineList[t_u_h.HeroId]; ok {
		res := new(Hero)
		tool.Clone(h, res)
		res.Uid = t_u_h.Uid
		res.Info.HeroId = t_u_h.HeroId
		res.Info.Floor = t_u_h.Floor
		res.SetHeroLevel(t_u_h.Level)
		res.SetStar(t_u_h.Star)
		res.Props = HeroProperties{
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
		}
		res.Status = t_u_h.Status
		res.Equips = make([]*Equip, 0)
		res.Skills = make([]*Skill, 0)
		return res
	}
	return nil
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

		Status: u_h.Status,
	}
}

// 	力量：一点力量增加18点生命值上限和0.14护甲，额外增加力量英雄1点物理攻击。
// 　　智力：一点智力增加2.4魔法强度和0.1魔法抗性，额外增加智力英雄1点物理攻击。
// 　　敏捷：一点敏捷增加0.4攻击强度和0.07护甲以及0.4物理暴击，额外增加敏捷英雄1点物理攻击。
// 　　生命最大值：一点力量增加18点生命最大值。
// 　　物理攻击力：一点主属性增加一点物理攻击力，一点敏捷增加0.4物理攻击力，对方无护甲情况下一点物理攻击力等于一点平砍伤害，增加物理技能伤害，伤害数值由各个技能公式决定。
// 　　魔法强度：一点智力增加2.4魔法强度，增加法术技能伤害，伤害数值由各个技能公式决定。
// 　　物理护甲：一点力量增加0.14护甲，一点敏捷增加0.07护甲，护甲会降低受到物理伤害的暴击率，护甲减伤为百分比减伤，目前公式暂缺。
// 　　魔法抗性：一点智力增加0.1魔法抗性，魔法抗性会降低受到法术伤害的暴击率，魔法抗性减伤为百分比减伤，目前公式暂缺。
// 　　物理暴击：一点敏捷增加0.4物理暴击，暴击伤害为200%,目前暴击等级与暴击率转换公式暂缺。
// 　　魔法暴击：魔法暴击与三围无关，由装备提供，暴击伤害为200%,目前暴击等级与暴击率转换公式暂缺。
// 　　生命回复：过场时回复的生命值，与三围无关，由装备提供。
// 　　能量回复：过场时回复的能量值，与三围无关，由装备提供。
func RefreshHero(h *Hero) {

	def := HeroDefineList[h.Info.HeroId]

	h.Props.Strength = h.Props.StrengthGrow * h.Info.Level / 100
	h.Props.Agility = h.Props.AgilityGrow * h.Info.Level / 100
	h.Props.Intelligent = h.Props.IntelligentGrow * h.Info.Level / 100

	h.Props.HP = def.Props.HP + h.Props.Strength*18
	h.Props.ADDef = def.Props.ADDef + (h.Props.Strength*14+h.Props.Agility*7)/100
	h.Props.AP = def.Props.AP + h.Props.Intelligent*24/10
	h.Props.APDef = def.Props.APDef + h.Props.Intelligent*1/10
	h.Props.ADCrit = def.Props.ADCrit + h.Props.Agility*4/10

	if def.Info.Type == HeroTypeStrength {
		h.Props.AD = def.Props.AD + (h.Props.Strength*10+h.Props.Agility*4)/10
	}
	if def.Info.Type == HeroTypeIntelligent {
		h.Props.AD = def.Props.AD + (h.Props.Intelligent*10+h.Props.Agility*4)/10
	}
	if def.Info.Type == HeroTypeAgility {
		h.Props.AD = def.Props.AD + h.Props.Agility*14/10
	}
}
