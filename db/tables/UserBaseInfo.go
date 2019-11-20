package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户基本信息
type UserBaseInfo struct {
	Uid     string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	HeadUrl string `gorm:"size:128"`
	Gold    int64
	Diamond int64

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&UserBaseInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserBaseInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentUserBaseInfo(userBaseInfos []*UserBaseInfo) {
	tx := context.DB().Begin()

	for _, u_b_i := range userBaseInfos {
		var oldUserBaseInfo UserBaseInfo
		tx.Where("uid = ?", u_b_i.Uid).First(&oldUserBaseInfo)
		if u_b_i.Uid != oldUserBaseInfo.Uid {
			tx.Create(&u_b_i)
		} else {
			tx.Model(&u_b_i).Where("uid = ? ", u_b_i.Uid).Updates(u_b_i)
		}
	}

	tx.Commit()
}

func LoadUserBaseInfo(userUid string) *UserBaseInfo {
	var userBaseInfo UserBaseInfo
	context.DB().Where("uid = ?", userUid).First(&userBaseInfo)

	return &userBaseInfo
}
