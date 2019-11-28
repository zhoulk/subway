package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户扩展信息
type UserExtendInfo struct {
	Uid      string `gorm:"size:64;unique;not null"`
	GuanKaId int
	Tech     int

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&UserExtendInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserExtendInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentUserExtendInfo(userExtendInfos []*UserExtendInfo) {
	tx := context.DB().Begin()

	for _, u_e_i := range userExtendInfos {
		var oldUserExtendInfo UserExtendInfo
		tx.Where("uid = ?", u_e_i.Uid).First(&oldUserExtendInfo)
		if u_e_i.Uid != oldUserExtendInfo.Uid {
			tx.Create(&u_e_i)
		} else {
			tx.Model(&u_e_i).Where("uid = ? ", u_e_i.Uid).Updates(u_e_i)
		}
	}

	tx.Commit()
}

func LoadUserExtendInfo(userUid string) *UserExtendInfo {
	var userExtendInfo UserExtendInfo
	context.DB().Where("uid = ?", userUid).First(&userExtendInfo)

	return &userExtendInfo
}
