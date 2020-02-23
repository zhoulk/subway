package tables

import (
	"subway/db/context"
	"time"

	"github.com/jinzhu/gorm"
)

// 用户基本信息
type HouseInfo struct {
	RoleId            string `gorm:"size:64;unique;not null"`
	GoldTimes         int32
	TotalGoldTimes    int32
	DiamondTimes      int32
	TotalDiamondTimes int32
	LastGoldTime      time.Time
	LastDiamondTime   time.Time

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&HouseInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HouseInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentHouseInfo(houseInfos []*HouseInfo) {
	tx := context.DB().Begin()

	for _, h_i := range houseInfos {
		var oldHouseInfo HouseInfo
		tx.Where("role_id = ?", h_i.RoleId).First(&oldHouseInfo)
		if h_i.RoleId != oldHouseInfo.RoleId {
			tx.Create(&h_i)
		} else {
			tx.Model(&h_i).Where("role_id = ? ", h_i.RoleId).Updates(h_i)
		}
	}

	tx.Commit()
}

func LoadHouseInfo(roleId string) *HouseInfo {
	var houseInfo HouseInfo
	context.DB().Where("role_id = ?", roleId).First(&houseInfo)
	if houseInfo.RoleId == roleId {
		return &houseInfo
	}
	return nil
}
