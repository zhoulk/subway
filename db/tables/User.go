package tables

import (
	"subway/db/context"
	"time"

	"github.com/jinzhu/gorm"
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

func init() {
	if !context.DB().HasTable(&User{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentUser(users []*User) {
	tx := context.DB().Begin()

	for _, t_u := range users {
		var oldUser User
		tx.Where("uid = ?", t_u.Uid).First(&oldUser)
		if t_u.Uid != oldUser.Uid {
			tx.Create(&t_u)
		} else {
			tx.Model(&t_u).Where("uid = ?", t_u.Uid).Updates(t_u)
		}
	}

	tx.Commit()
}

func LoadUser(zoneId int, openId string) *User {
	var user User
	context.DB().Where("zone_id = ? AND open_id = ?", zoneId, openId).First(&user)
	if user.OpenId == openId {
		return &user
	}
	return nil
}

func LoadUserByUid(uid string) *User {
	var user User
	context.DB().Where("uid = ?", uid).First(&user)
	if user.Uid == uid {
		return &user
	}
	return nil
}
