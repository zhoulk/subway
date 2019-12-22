package tables

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"subway/db/context"

	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

// 副本定义
type CopyDefine struct {
	CopyId int
	Name   string
}

// 副本小关卡定义
type CopyItemDefine struct {
	ChapterId  int
	CopyId     int
	Name       string
	Equips     []int
	EquipParts []int
	HeroParts  []int
	Heros      []*GuanKaHeroDefine
}

func init() {
	// createCopyDefineTable()
	// initCopyDefineData()
}

func createCopyDefineTable() {
	if !context.DB().HasTable(&CopyDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CopyDefine{}).Error; err != nil {
			panic(err)
		}
	}

	if !context.DB().HasTable(&CopyItemDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&CopyItemDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initCopyDefineData() {

	data, err := ioutil.ReadFile("./static/data/copy.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var copys []CopyDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &copys)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&CopyDefine{})
	context.DB().Unscoped().Delete(&CopyItemDefine{})

	tx := context.DB().Begin()
	for _, cp := range copys {
		tx.Create(cp)
		initCopyItemDefineData(tx, cp.CopyId)
	}
	tx.Commit()
}

func initCopyItemDefineData(tx *gorm.DB, copyId int) {
	data, err := ioutil.ReadFile("./static/data/copy/copy" + strconv.Itoa(copyId) + ".json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var copyItems []CopyItemDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &copyItems)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	for _, cp := range copyItems {
		tx.Create(cp)
	}
}

func LoadCopyData() []CopyDefine {

	data, err := ioutil.ReadFile("./static/data/copy.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	var copys []CopyDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &copys)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	return copys
}

func LoadCopyItems(copyId int) []CopyItemDefine {
	data, err := ioutil.ReadFile("./static/data/copy/copy" + strconv.Itoa(copyId) + ".json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	var copyItems []CopyItemDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &copyItems)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	return copyItems
}
