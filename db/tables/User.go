package tables

import (
	"subway/db/context"
	"github.com/jinzhu/gorm"
	"time"
)

// 用户登录信息
type User struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Account    string `gorm:"size:128"`
	Password   string `gorm:"size:64"`
	OpenId     string `gorm:"size:64"`
	ZoneId     int
	LoginTime  time.Time
	LogoutTime time.Time

	gorm.Model
}

func init()  {
	if !context.DB().HasTable(&User{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}