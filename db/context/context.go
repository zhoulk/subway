package context

import (
	"fmt"
	"github.com/astaxie/beego"
	"sync"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Context ...
type Context struct {
	db *gorm.DB
}

func init() {
	ConnectDB()
}

var _instance *Context
var once sync.Once

// GetInstance 单例
func GetInstance() *Context {
	once.Do(func() {
		_instance = &Context{}
	})
	return _instance
}

func DB() *gorm.DB {
	return GetInstance().db
}

func ConnectDB(){
	GetInstance().ConnectDB()
}

// ConnectDB 连接数据库
func (m *Context) ConnectDB() {
	driver := beego.AppConfig.String("sqlconn")
	db, err := gorm.Open("mysql", driver)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// defer db.Close()
	m.db = db
}

// PersistentData 数据库固化
func (m *Context) PersistentData() {
	beego.Debug("persistent start ==================================== ")

	beego.Debug("persistent end ==================================== ")
}

// LoadFromDB 加载数据
func (m *Context) LoadFromDB() {
	beego.Debug("LoadFromDB start ==================================== ")

	beego.Debug("LoadFromDB end ==================================== ")
}