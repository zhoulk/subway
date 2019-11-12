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

func init()  {
	if !context.DB().HasTable(&UserBaseInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserBaseInfo{}).Error; err != nil {
			panic(err)
		}
	}
}