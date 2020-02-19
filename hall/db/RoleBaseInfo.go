package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户基本信息
type RoleBaseInfo struct {
	RoleId  string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	HeadUrl string `gorm:"size:256"`
	Gold    int64
	Diamond int64
	Power   int64
	Level   int32
	Exp     int64
	Point   int32

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&RoleBaseInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&RoleBaseInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentRoleBaseInfo(roleBaseInfos []*RoleBaseInfo) {
	tx := context.DB().Begin()

	for _, r_b_i := range roleBaseInfos {
		var oldRoleBaseInfo RoleBaseInfo
		tx.Where("role_id = ?", r_b_i.RoleId).First(&oldRoleBaseInfo)
		if r_b_i.RoleId != oldRoleBaseInfo.RoleId {
			tx.Create(&r_b_i)
		} else {
			tx.Model(&r_b_i).Where("role_id = ? ", r_b_i.RoleId).Updates(r_b_i)
		}
	}

	tx.Commit()
}

func LoadRoleBaseInfo(roleId string) *RoleBaseInfo {
	var roleBaseInfo RoleBaseInfo
	context.DB().Where("role_id = ?", roleId).First(&roleBaseInfo)

	return &roleBaseInfo
}
