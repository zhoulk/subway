package db

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/jinzhu/gorm"

	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	// DBDriver 数据库地址
	// DBDriver = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/wxgame?charset=utf8&&parseTime=true"
	// DBDriverTest 测试数据库地址
	DBDriverTest = "root:A845240287a@tcp(rm-wz9sw694mi8020vigo.mysql.rds.aliyuncs.com:3306)/subway_test?charset=utf8&&parseTime=true"
)

// ConnectDB 连接数据库
func (m *Module) ConnectDB() {
	driver := DBDriverTest
	db, err := gorm.Open("mysql", driver)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// defer db.Close()
	m.db = db
}

// PersistentData 数据库固化
func (m *Module) PersistentData() {
	beego.Debug("persistent start ==================================== ")

	beego.Debug("persistent end ==================================== ")
}

// CreateTables 创建表
func (m *Module) CreateTables() {
	if !m.db.HasTable(&User{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	if !m.db.HasTable(&UserBaseInfo{}) {
		if err := m.db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&UserBaseInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

// LoadFromDB 加载数据
func (m *Module) LoadFromDB() {
	beego.Debug("LoadFromDB start ==================================== ")

	beego.Debug("LoadFromDB end ==================================== ")
}
