package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户副本
type UserBag struct {
	Uid      string `gorm:"size:64;unique;not null"`
	UserId   string
	ItemId   int
	Count    int
	ItemType int8

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&UserBag{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserBag{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentUserBags(userBags []*UserBag) {
	tx := context.DB().Begin()

	for _, u_b := range userBags {
		var oldUserBag UserBag
		tx.Where("uid = ?", u_b.Uid).First(&oldUserBag)
		if u_b.Uid != oldUserBag.Uid {
			tx.Create(&u_b)
		} else {
			tx.Model(&u_b).Where("uid = ? ", u_b.Uid).Updates(u_b)
		}
	}

	tx.Commit()
}

func LoadUserBags(userUid string) []*UserBag {
	var userBags []*UserBag
	context.DB().Where("user_id = ?", userUid).Find(&userBags)

	return userBags
}
