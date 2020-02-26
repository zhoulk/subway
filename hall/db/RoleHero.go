package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户背包
type RoleHeroInfo struct {
	RoleId string
	Items  []*RoleHeroItemInfo
}

type RoleHeroItemInfo struct {
	Uid    string `gorm:"size:64;unique;not null"`
	RoleId string
	HeroId int32
	Level  int32
	Floor  int32
	Star   int32

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&RoleHeroItemInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&RoleHeroItemInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentRoleHeroItems(roleHeroItems []*RoleHeroItemInfo) {
	tx := context.DB().Begin()

	for _, u_b := range roleHeroItems {
		var oldRoleHeroItem RoleHeroItemInfo
		tx.Where("uid = ?", u_b.Uid).First(&oldRoleHeroItem)
		if u_b.Uid != oldRoleHeroItem.Uid {
			tx.Create(&u_b)
		} else {
			tx.Model(&u_b).Where("uid = ? ", u_b.Uid).Updates(u_b)
		}
	}

	tx.Commit()
}

func LoadRoleHeroInfo(roelId string) *RoleHeroInfo {
	var roleHeroItems []*RoleHeroItemInfo
	context.DB().Where("role_id = ?", roelId).Find(&roleHeroItems)
	if len(roleHeroItems) > 0 {
		return &RoleHeroInfo{
			RoleId: roelId,
			Items:  roleHeroItems,
		}
	}
	return nil
}

func LoadRoleHeroItemInfo(heroUid string) *RoleHeroItemInfo {
	var roleHeroItem RoleHeroItemInfo
	context.DB().Where("uid = ?", heroUid).First(&roleHeroItem)
	return &roleHeroItem
}
