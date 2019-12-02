package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户副本
type UserCopyItem struct {
	Uid        string `gorm:"size:64;unique;not null"`
	UserId     string
	CopyItemId int
	Star       int
	Status     int8

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&UserCopyItem{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserCopyItem{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentUserCopyItem(userCopyItems []*UserCopyItem) {
	tx := context.DB().Begin()

	for _, u_c_i := range userCopyItems {
		var oldUserCopyItem UserCopyItem
		tx.Where("uid = ?", u_c_i.Uid).First(&oldUserCopyItem)
		if u_c_i.Uid != oldUserCopyItem.Uid {
			tx.Create(&u_c_i)
		} else {
			tx.Model(&u_c_i).Where("uid = ? ", u_c_i.Uid).Updates(u_c_i)
		}
	}

	tx.Commit()
}

func LoadUserCopyItems(userUid string) []*UserCopyItem {
	var userCopyItems []*UserCopyItem
	context.DB().Where("user_id = ?", userUid).Find(&userCopyItems)

	return userCopyItems
}
