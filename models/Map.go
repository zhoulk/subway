package models

import (
	"math/rand"
	"subway/db/tables"
)

var (
	MapDic      map[int32]string
	MapChildDic map[string][]int32
)

func init() {
	MapDic = make(map[int32]string)
	MapChildDic = make(map[string][]int32)

	mapDefins := tables.LoadMapDefine()
	for _, def := range mapDefins {
		MapDic[def.Id] = def.Name
		if childs, ok := MapChildDic[def.Name]; ok {
			childs = append(childs, def.Next)
			MapChildDic[def.Name] = childs
		} else {
			childs := make([]int32, 0)
			childs = append(childs, def.Next)
			MapChildDic[def.Name] = childs
		}
	}
}

type MapItem struct {
	Id   int32
	Name string
	Next []int32
}

func RandomAPath(uid string, from int32) (int32, []MapItem) {
	res := make([]MapItem, 0)

	res = append(res, MapItem{
		Id:   from,
		Name: MapDic[from],
		Next: MapChildDic[MapDic[from]],
	})

	step := rand.Intn(6)

	childs := make([]int32, 0)
	if name, ok := MapDic[from]; ok {
		if cls, ok := MapChildDic[name]; ok {
			childs = append(childs, cls...)
		}
	}

	for i := 0; i <= step; i++ {
		childs2 := make([]int32, 0)
		for _, child := range childs {
			res = append(res, MapItem{
				Id:   child,
				Name: MapDic[child],
				Next: MapChildDic[MapDic[child]],
			})
			childs2 = append(childs2, MapChildDic[MapDic[child]]...)
		}

		childs = childs2
	}

	return int32(step + 1), res
}
