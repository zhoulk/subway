package tables

import (
	"subway/db/context"

	"github.com/jinzhu/gorm"
)

// 用户背包
type BagInfo struct {
	RoleId string
	Items  []*BagItemInfo
}

type BagItemInfo struct {
	Uid      string `gorm:"size:64;unique;not null"`
	RoleId   string
	ItemId   int32
	Count    int32
	ItemType int8
	Extend   string

	gorm.Model
}

func init() {
	if !context.DB().HasTable(&BagItemInfo{}) {
		if err := context.DB().Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&BagItemInfo{}).Error; err != nil {
			panic(err)
		}
	}
}

func PersistentBagItems(bagItems []*BagItemInfo) {
	tx := context.DB().Begin()

	for _, u_b := range bagItems {
		var oldBagItem BagItemInfo
		tx.Where("uid = ?", u_b.Uid).First(&oldBagItem)
		if u_b.Uid != oldBagItem.Uid {
			tx.Create(&u_b)
		} else {
			data := make(map[string]interface{})
			data["Count"] = u_b.Count
			data["Extend"] = u_b.Extend
			tx.Model(&u_b).Where("uid = ? ", u_b.Uid).Updates(data)
		}
	}

	tx.Commit()
}

func LoadBagInfo(roelId string) *BagInfo {
	var bagItems []*BagItemInfo
	context.DB().Where("role_id = ?", roelId).Find(&bagItems)
	if len(bagItems) > 0 {
		return &BagInfo{
			RoleId: roelId,
			Items:  bagItems,
		}
	}
	return nil
}
