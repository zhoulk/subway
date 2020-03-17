package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户背包
type RoleHeroPartInfo struct {
	RoleId string
	Items  []*RoleHeroPartItemInfo
}

type RoleHeroPartItemInfo struct {
	Uid    string `gorm:"size:64;unique;not null"`
	RoleId string
	HeroId int32
	Num    int32

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&RoleHeroPartItemInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&RoleHeroPartItemInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentRoleHeroPartItems(roleHeroItems []*RoleHeroPartItemInfo) {
	tx := context.DB().Begin()

	for _, u_b := range roleHeroItems {
		var oldRoleHeroItem RoleHeroPartItemInfo
		tx.Where("uid = ?", u_b.Uid).First(&oldRoleHeroItem)
		if u_b.Uid != oldRoleHeroItem.Uid {
			tx.Create(&u_b)
		} else {
			tx.Model(&u_b).Where("uid = ? ", u_b.Uid).Updates(u_b)
		}
	}

	tx.Commit()
}

func LoadRoleHeroPartInfo(roelId string) *RoleHeroPartInfo {
	var roleHeroItems []*RoleHeroPartItemInfo
	context.DB().Where("role_id = ?", roelId).Find(&roleHeroItems)
	if len(roleHeroItems) > 0 {
		return &RoleHeroPartInfo{
			RoleId: roelId,
			Items:  roleHeroItems,
		}
	}
	return nil
}

func LoadRoleHeroPartItemInfo(heroUid string) *RoleHeroPartItemInfo {
	var roleHeroItem RoleHeroPartItemInfo
	context.DB().Where("uid = ?", heroUid).First(&roleHeroItem)
	return &roleHeroItem
}
