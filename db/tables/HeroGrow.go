package tables

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego"
)

// 英雄成长
type HeroGrowDefine struct {
	HeroId          int
	Star            int
	StrengthGrow    int
	AgilityGrow     int
	IntelligentGrow int
}

func init() {
}

func LoadHeroGrowData() []HeroGrowDefine {

	data, err := ioutil.ReadFile("./static/data/heroGrow.json")
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	var grows []HeroGrowDefine

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, &grows)
	if err != nil {
		beego.Error("initData failed", err.Error())
		return nil
	}

	return grows
}
