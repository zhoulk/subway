package models

var (
	ZoneList []*Zone
)

func init() {
	ZoneList = make([]*Zone, 0)
	z1 := &Zone{1, "北京-1号线", 0}
	ZoneList = append(ZoneList, z1)
}

type Zone struct {
	Id       int
	ZoneName string
	Status int
}

func GetAllZones() []*Zone {
	return ZoneList
}
