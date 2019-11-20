package tables

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego"
)

// 副本定义
type GuanKaDefine struct {
	GuanKaId int
	Name     string

	Heros []*GuanKaHeroDefine
}

type GuanKaHeroDefine struct {
	HeroId      string
	Level       int32
	Floor       int16 // 阶别
	Star        int16 // 星星
	SkillLevels []int32
}

func init() {
}

func LoadGuanKaData() []GuanKaDefine {

	data, err := ioutil.ReadFile("./static/data/gk/gk1.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	var gks []GuanKaDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &gks)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	return gks
}
