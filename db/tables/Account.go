package tables

import (
	"subway/db/context"
	"time"

	"github.com/jinzhu/gorm"
)

// 用户登录信息
type Account struct {
	AccountId  string `gorm:"size:64;unique;not null"`
	Account    string `gorm:"size:128"`
	Password   string `gorm:"size:64"`
	OpenId     string `gorm:"size:64"`
	LoginTime  time.Time
	LogoutTime time.Time

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&Account{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Account{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentAccount(users []*Account) {
	tx := context.DB().Begin()

	for _, t_u := range users {
		var oldAccount Account
		tx.Where("account_id = ?", t_u.AccountId).First(&oldAccount)
		if t_u.AccountId != oldAccount.AccountId {
			tx.Create(&t_u)
		} else {
			tx.Model(&t_u).Where("account_id = ?", t_u.AccountId).Updates(t_u)
		}
	}

	tx.Commit()
}

func LoadAccount(openId string) *Account {
	var account Account
	context.DB().Where("open_id = ?", openId).First(&account)
	if account.OpenId == openId {
		return &account
	}
	return nil
}

func LoadAccountByUid(accountId string) *Account {
	var account Account
	context.DB().Where("account_id = ?", accountId).First(&account)
	if account.AccountId == accountId {
		return &account
	}
	return nil
}
