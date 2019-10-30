package models

var (
	HeroDefineList map[string]*Hero
)

func init() {
	HeroDefineList = make(map[string]*Hero)
	h1 := &Hero{HeroId:"1000", Name:"敌法师", Detail: HeroProfile{Desc:"第一个英雄"}}
	HeroDefineList[h1.HeroId] = h1
}

type Hero struct{
	Uid string
	HeroId string
	Name string
	Detail HeroProfile
}

type HeroProfile struct{
	Desc string
}

func GetAllHeros(uid string) []*Hero{
	u, _ := GetUser(uid)
	if u != nil{
		return u.Heros
	}
	return nil
}