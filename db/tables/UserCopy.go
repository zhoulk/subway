package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户副本
type UserCopy struct {
	Uid    string `gorm:"size:64;unique;not null"`
	UserId string
	CopyId int
	Star   int
	Status int8

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&UserCopy{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserCopy{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentUserCopy(userCopys []*UserCopy) {
	tx := context.DB().Begin()

	for _, u_c := range userCopys {
		var oldUserCopy UserCopy
		tx.Where("uid = ?", u_c.Uid).First(&oldUserCopy)
		if u_c.Uid != oldUserCopy.Uid {
			tx.Create(&u_c)
		} else {
			tx.Model(&u_c).Where("uid = ? ", u_c.Uid).Updates(u_c)
		}
	}

	tx.Commit()
}

func LoadUserCopys(userUid string) []*UserCopy {
	var userCopys []*UserCopy
	context.DB().Where("user_id = ?", userUid).Find(&userCopys)

	return userCopys
}
