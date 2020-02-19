package models

var (
	ZoneList []*Zone
)

func init() {
	ZoneList = make([]*Zone, 0)
	z1 := &Zone{Id: 1, ZoneName: "(1区)  舰船统帅"}
	ZoneList = append(ZoneList, z1)
}

type Zone struct {
	Id       int
	ZoneName string
	Status   int
	Name     string
	Level    int
}

func GetAllZones() []*Zone {
	return ZoneList
}

func GetSelfZones() []*Zone {
	_ZoneList := make([]*Zone, 0)
	z1 := &Zone{Id: 1, ZoneName: "(1区)  舰船统帅", Name: "小小", Level: 10}
	_ZoneList = append(_ZoneList, z1)
	return _ZoneList
}
