package models

var (
	GuanKaList []*GuanKa
)

func init() {
	GuanKaList = make([]*GuanKa, 0)

	g1 := &GuanKa{Info: GuanKaInfo{GuanKaId: 1, Name: "1号线-苹果园"}}
	g2 := &GuanKa{Info: GuanKaInfo{GuanKaId: 2, Name: "1号线-苹果园1"}}
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
	Uid  string
	Info GuanKaInfo
}

type GuanKaInfo struct {
	GuanKaId int
	Name     string
}

func GetGuanKa(guanKaId int) *GuanKa {
	if guanKaId-1 >= 0 && guanKaId-1 < len(GuanKaList) {
		return GuanKaList[guanKaId-1]
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
