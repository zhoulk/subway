package db

import (
	"sync"

	"github.com/jinzhu/gorm"
)

// Module ...
type Module struct {
	db *gorm.DB
}

var _instance *Module
var once sync.Once

// GetInstance 单例
func GetInstance() *Module {
	once.Do(func() {
		_instance = &Module{}
	})
	return _instance
}

func init() {
	GetInstance().ConnectDB()
	GetInstance().CreateTables()
}
