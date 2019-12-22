package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 装备定义表
type EquipDefine struct {
	EquipId     string `gorm:"size:64;unique;not null"`
	Name        string `gorm:"size:64"`
	Level       int32  // 级别
	Strength    int32  // 力量
	HP          int32  // 生命值
	Agility     int32  // 敏捷
	MP          int32  // 魔法强度
	Intelligent int32
	AD          int32 // 物理攻击
	ADCrit      int32 // 物理暴击
	ADDef       int32 // 物理护甲
	HPGain      int32 // 生命恢复
	MPGain      int32 // 能量恢复
	HPReGain    int32 // 战斗后补充生命
	MPReGain    int32 // 战斗后补充能量
	Desc        string
	From        string
	Power       int32
	Cost        int32

	MixCnt int32
	Mix    string
}

func init() {
	beego.Debug("EquipDefine init")
	if beego.AppConfig.DefaultBool("updateConfigData", true) {
		createEquipDefineTable()
		initEquipData()
	}
}

func createEquipDefineTable() {
	if !context.DB().HasTable(&EquipDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&EquipDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initEquipData() {

	data, err := ioutil.ReadFile("./static/data/equip.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var equips []EquipDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &equips)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&EquipDefine{})

	tx := context.DB().Begin()
	for _, h := range equips {
		tx.Create(h)
	}
	tx.Commit()
}

func LoadEquipDefine() []*EquipDefine {
	var equips []*EquipDefine
	context.DB().Where("1=1").Find(&equips)

	return equips
}
