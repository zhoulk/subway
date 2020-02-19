package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户登录信息
type Role struct {
	RoleId    string `gorm:"size:64;unique;not null"`
	AccountId string `gorm:"size:64"`
	ZoneId    int

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&Role{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Role{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentRole(users []*Role) {
	tx := context.DB().Begin()

	for _, t_u := range users {
		var oldRole Role
		tx.Where("role_id = ?", t_u.RoleId).First(&oldRole)
		if t_u.RoleId != oldRole.RoleId {
			tx.Create(&t_u)
		} else {
			tx.Model(&t_u).Where("role_id = ?", t_u.RoleId).Updates(t_u)
		}
	}

	tx.Commit()
}

func LoadRole(roleId string) *Role {
	var role Role
	context.DB().Where("role_id = ?", roleId).First(&role)
	if role.RoleId == roleId {
		return &role
	}
	return nil
}

func LoadRoleByAccount(zoneId int, accountId string) *Role {
	var role Role
	context.DB().Where("zone_id = ? AND account_id = ?", zoneId, accountId).First(&role)
	if role.ZoneId == zoneId && role.AccountId == accountId {
		return &role
	}
	return nil
}
