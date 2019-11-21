package models

import (
	"math"
	"subway/db/tables"

	"github.com/astaxie/beego"
)

var (
	GuanKaList []*GuanKa
)

func init() {
	GuanKaList = make([]*GuanKa, 0)

	defines := tables.LoadGuanKaData()
	for _, def := range defines {

		heros := make([]*GuanKaHero, 0)
		for _, gkHero := range def.Heros {
			heros = append(heros, &GuanKaHero{HeroId: gkHero.HeroId, Level: gkHero.Level, Floor: gkHero.Floor, Star: gkHero.Star, SkillLevels: gkHero.SkillLevels})
		}
		GuanKaList = append(GuanKaList, &GuanKa{
			Info: GuanKaInfo{
				GuanKaId: def.GuanKaId,
				Name:     def.Name,
			},
			Heros: heros,
		})
	}
}

type GuanKa struct {
	Uid   string
	Info  GuanKaInfo
	Heros []*GuanKaHero
}

type GuanKaInfo struct {
	GuanKaId int
	Name     string
}

type GuanKaHero struct {
	*Hero

	HeroId      string
	Level       int32
	Floor       int16 // 阶别
	Star        int16 // 星星
	SkillLevels []int32
}

func GetGuanKa(gkId int) *GuanKa {
	beego.Debug("GetGuanKa  ", gkId)
	index := int(math.Ceil(float64(gkId) / float64(10)))
	if index-1 >= 0 && index-1 < len(GuanKaList) {
		gk := GuanKaList[index-1]
		if gk.Heros != nil {
			for _, h := range gk.Heros {
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
		gk.Info.GuanKaId = gkId
		return gk
	}
	return nil
}

func GetNearGuanKa(uid string) []*GuanKa {
	res := make([]*GuanKa, 0)
	u, _ := GetUser(uid)
	if u != nil {
		// preGuanKaId := u.Profile.GuanKaId - 1
		// NextGuanKaId := u.Profile.GuanKaId + 1
		//res = append(res, GetGuanKa(preGuanKaId))
		res = append(res, GetGuanKa(u.Profile.GuanKaId))
		//res = append(res, GetGuanKa(NextGuanKaId))
	}
	return res
}
