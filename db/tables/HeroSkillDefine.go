package tables

import (
	"encoding/json"
	"io/ioutil"
	"subway/db/context"

	"github.com/astaxie/beego"
)

// 英雄技能定义表
type HeroSkillDefine struct {
	HeroId  string
	SkillId string
}

func init() {
	beego.Debug("HeroSkillDefine init")
	if beego.AppConfig.DefaultBool("updateConfigData", true) {
		createHeroSkillDefineTable()
		initHeroSkillData()
	}
}

func createHeroSkillDefineTable() {
	if !context.DB().HasTable(&HeroSkillDefine{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&HeroSkillDefine{}).Error; err != nil {
			panic(err)
		}
	}
}

func initHeroSkillData() {

	data, err := ioutil.ReadFile("./static/data/heroSkill.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	var heroSkills []HeroSkillDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &heroSkills)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return
	}

	context.DB().Unscoped().Delete(&HeroSkillDefine{})

	tx := context.DB().Begin()
	for _, h := range heroSkills {
		tx.Create(h)
	}
	tx.Commit()
}

func LoadHeroSkillDefine() []*HeroSkillDefine {
	var heroSkills []*HeroSkillDefine
	context.DB().Where("1=1").Find(&heroSkills)

	return heroSkills
}
