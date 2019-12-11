package tables

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/astaxie/beego"
)

// 副本定义
type CopyDefine struct {
	CopyId int
	Name   string
}

// 副本小关卡定义
type CopyItemDefine struct {
	CopyId     int
	Name       string
	Equips     []int
	EquipParts []int
	HeroParts  []int
	Heros      []*GuanKaHeroDefine
}

func init() {
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
