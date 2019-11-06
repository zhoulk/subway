package db

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 大区信息
type Zone struct {
	ZoneId     int    `unique;not null"`
	Name       string `gorm:"size:64"`
	LoginTime  time.Time
	LogoutTime time.Time

	gorm.Model
}

// 用户登录信息
type User struct {
	Uid        string `gorm:"size:64;unique;not null"`
	Account    string `gorm:"size:128"`
	Password   string `gorm:"size:64"`
	OpenId     string `gorm:"size:64"`
	ZoneId     int
	LoginTime  time.Time
	LogoutTime time.Time

	gorm.Model
}

// 用户基本信息
type UserBaseInfo struct {
	Uid     string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	HeadUrl string `gorm:"size:128"`
	Gold    int64
	Diamond int64

	gorm.Model
}

// 英雄定义表
// 技能定义表
// 装备定义表
// 英雄技能关系表
// 英雄装备关系表

// 用户英雄表
