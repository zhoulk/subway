package tables

import (
	"subway/db/context"
	"github.com/jinzhu/gorm"
)
// 大区信息
type Zone struct {
	ZoneId     int    `gorm:"unique;not null"`
	Name       string `gorm:"size:64"`
	Status int

	gorm.Model
}

func init()  {
	if !context.DB().HasTable(&Zone{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Zone{}).Error; err != nil {
			panic(err)
		}
	}

	// context.DB().Unscoped().Delete(&Zone{})
	// for _,d := range data.ZoneData {
	// 	var s Zone
	// 	json.Unmarshal([]byte(d), &s)
	// 	context.DB().Create(&s)
	// }
}