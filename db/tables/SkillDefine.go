package tables

import (
	"io/ioutil"
	"github.com/astaxie/beego"
	"encoding/json"
	"subway/db/context"
)

// 技能定义表
type SkillDefine struct {
	SkillId string `gorm:"size:64;unique;not null"`
	Name    string `gorm:"size:64"`
	Level   int32
	Type    int8 // 1 主动  2  被动
	Desc    string
}

func init()  {
	createSkillDefineTable()
	initSkillData()
}

func createSkillDefineTable()  {
	if !context.DB().HasTable(&SkillDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&SkillDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initSkillData(){

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
	for _,h := range skills{
		tx.Create(h)
	}
	tx.Commit()
}