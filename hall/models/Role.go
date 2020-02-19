package models

import (
	tables "subway/hall/db"
	"subway/tool"
)

var (
	RoleInfoList map[string]*RoleInfo
)

func init() {
	RoleInfoList = make(map[string]*RoleInfo)
}

type RoleInfo struct {
	RoleId  string
	Name    string
	HeadUrl string
	Gold    int64
	Diamond int64
	Power   int64
	Level   int32
	Exp     int64
	Point   int32
}

func GetRoleInfo(roleId string) *RoleInfo {
	if r, ok := RoleInfoList[roleId]; ok {
		return r
	}

	t_r := tables.LoadRoleBaseInfo(roleId)
	if t_r != nil {
		r := CreateRoleInfoFromTableRoleInfo(t_r)
		RoleInfoList[roleId] = r
		return RoleInfoList[roleId]
	}

	return nil
}

func AddRoleInfo(roleInfo *RoleInfo) *RoleInfo {
	if _, ok := RoleInfoList[roleInfo.RoleId]; ok {
		RoleInfoList[roleInfo.RoleId] = roleInfo
	} else {
		roleInfo.RoleId = tool.UniqueId()
		RoleInfoList[roleInfo.RoleId] = roleInfo
	}
	return RoleInfoList[roleInfo.RoleId]
}

func PersistentRoleInfo() {
	roleInfos := make([]*tables.RoleBaseInfo, 0)
	for _, a := range RoleInfoList {
		roleInfos = append(roleInfos, CreateTableRoleInfoFromRoleInfo(a))
	}
	tables.PersistentRoleBaseInfo(roleInfos)
}

func CreateRoleInfoFromTableRoleInfo(a *tables.RoleBaseInfo) *RoleInfo {
	return &RoleInfo{
		RoleId:  a.RoleId,
		Name:    a.Name,
		HeadUrl: a.HeadUrl,
		Gold:    a.Gold,
		Diamond: a.Diamond,
		Power:   a.Power,
		Level:   a.Level,
		Exp:     a.Exp,
		Point:   a.Point,
	}
}

func CreateTableRoleInfoFromRoleInfo(a *RoleInfo) *tables.RoleBaseInfo {
	return &tables.RoleBaseInfo{
		RoleId:  a.RoleId,
		Name:    a.Name,
		HeadUrl: a.HeadUrl,
		Gold:    a.Gold,
		Diamond: a.Diamond,
		Power:   a.Power,
		Level:   a.Level,
		Exp:     a.Exp,
		Point:   a.Point,
	}
}
