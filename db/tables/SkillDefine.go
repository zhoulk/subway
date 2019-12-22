package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 技能定义表
type SkillDefine struct {
	SkillId string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	Level   int32
	Type    int8 // 1 主动  2  被动
	Desc    string
}

func init() {
	beego.Debug("SkillDefine init")
	if beego.AppConfig.DefaultBool("updateConfigData", true) {
		createSkillDefineTable()
		initSkillData()
	}
}

func createSkillDefineTable() {
	if !context.DB().HasTable(&SkillDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&SkillDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initSkillData() {

	data, err := ioutil.ReadFile("./static/data/skill.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var skills []SkillDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &skills)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&SkillDefine{})

	tx := context.DB().Begin()
	for _, h := range skills {
		tx.Create(h)
	}
	tx.Commit()
}

func LoadSkillDefine() []*SkillDefine {
	var skills []*SkillDefine
	context.DB().Where("1=1").Find(&skills)

	return skills
}
