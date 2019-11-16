package models

var (
	GuanKaList []*GuanKa
)

func init() {
	GuanKaList = make([]*GuanKa, 0)

	g1 := &GuanKa{Info: GuanKaInfo{GuanKaId: 1, Name: "1号线-苹果园"},
		Heros: []*GuanKaHero{&GuanKaHero{HeroId: "1001", Level: 1, Floor: 0, Star: 0, SkillLevels: []int32{1, 1, 1, 1}}}}
	g2 := &GuanKa{Info: GuanKaInfo{GuanKaId: 2, Name: "1号线-苹果园1"},
		Heros: []*GuanKaHero{&GuanKaHero{HeroId: "1003", Level: 1, Floor: 0, Star: 0, SkillLevels: []int32{1, 1, 1, 1}}}}
	g3 := &GuanKa{Info: GuanKaInfo{GuanKaId: 3, Name: "1号线-苹果园2"}}
	g4 := &GuanKa{Info: GuanKaInfo{GuanKaId: 4, Name: "1号线-苹果园3"}}
	g5 := &GuanKa{Info: GuanKaInfo{GuanKaId: 5, Name: "1号线-苹果园4"}}
	GuanKaList = append(GuanKaList, g1)
	GuanKaList = append(GuanKaList, g2)
	GuanKaList = append(GuanKaList, g3)
	GuanKaList = append(GuanKaList, g4)
	GuanKaList = append(GuanKaList, g5)
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

func GetGuanKa(guanKaId int) *GuanKa {
	if guanKaId-1 >= 0 && guanKaId-1 < len(GuanKaList) {
		gk := GuanKaList[guanKaId-1]
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
		return gk
	}
	return nil
}

func GetNearGuanKa(uid string) []*GuanKa {
	res := make([]*GuanKa, 0)
	u, _ := GetUser(uid)
	if u != nil {
		preGuanKaId := u.Profile.GuanKaId - 1
		NextGuanKaId := u.Profile.GuanKaId + 1
		res = append(res, GetGuanKa(preGuanKaId))
		res = append(res, GetGuanKa(u.Profile.GuanKaId))
		res = append(res, GetGuanKa(NextGuanKaId))
	}
	return res
}
