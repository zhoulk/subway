package models

import (
	"errors"
	"subway/db/tables"
	"subway/tool"
	"time"
)

var (
	UserList    map[string]*User
	AccountList map[string]*Account
)

func init() {
	UserList = make(map[string]*User)
	AccountList = make(map[string]*Account)
}

type Account struct {
	AccountId  string
	OpenId     string
	Account    string
	Password   string
	LoginTime  time.Time
	LogoutTime time.Time
}

type Role struct {
	ZoneId    int
	RoleId    string
	AccountId string
}

type User struct {
	Info    UserInfo
	Profile UserProfile

	Heros       []*Hero
	HeroDic     map[string]*Hero
	Copys       []*Copy
	CopyItems   []*CopyItem
	CopyDic     map[int]*Copy
	CopyItemDic map[int]*CopyItem
}

type UserInfo struct {
	ZoneId   int
	Uid      string
	OpenId   string
	Username string
}

type UserProfile struct {
	Gold     int64
	GuanKaId int
	Tech     int
}

const (
	IncreaseGoldReasonGK   int8 = 1 // 过关奖励
	IncreaseGoldReasonCopy int8 = 1 // 副本奖励

)

func (u *User) SetGuanKaId(gkId int) {
	u.Profile.GuanKaId = gkId
}

func (u *User) IncreaseGold(gold int64, reason int8) {
	u.Profile.Gold = u.Profile.Gold + gold
}

// 开启一个新的副本
func (u *User) AddCopyItem(cpItem *CopyItem) {
	if _, ok := u.CopyItemDic[cpItem.CopyItemId]; !ok {
		u.CopyItems = append(u.CopyItems, cpItem)
		u.CopyItemDic[cpItem.CopyItemId] = cpItem
	}
}

// 开启一个章节
func (u *User) AddCopy(cp *Copy) {
	if _, ok := u.CopyDic[cp.Info.CopyId]; !ok {
		u.Copys = append(u.Copys, cp)
		u.CopyDic[cp.Info.CopyId] = cp
	}
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}

	t_user := tables.LoadUserByUid(uid)
	t_baseInfo := tables.LoadUserBaseInfo(uid)
	t_extendInfo := tables.LoadUserExtendInfo(uid)
	if t_user != nil {
		UserList[t_user.Uid] = &User{
			Info: UserInfo{
				Uid:    t_user.Uid,
				OpenId: t_user.OpenId,
				ZoneId: t_user.ZoneId,
			},
			Profile: UserProfile{
				Gold:     t_baseInfo.Gold,
				GuanKaId: t_extendInfo.GuanKaId,
				Tech:     t_extendInfo.Tech,
			},
		}
		GetSelfHeros(t_user.Uid)
		return UserList[t_user.Uid], nil
	}

	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func GetAccount(openId string) *Account {
	if u, ok := AccountList[openId]; ok {
		return u
	}
	return nil
}

func AddAccount(openId string) *Account {
	if u, ok := AccountList[openId]; ok {
		return u
	}
	accountId := tool.UniqueId()
	AccountList[openId] = &Account{
		AccountId: accountId,
		OpenId:    openId,
	}
	return AccountList[openId]
}

func Login(zoneId int, openId string) *User {
	for _, u := range UserList {
		if u.Info.OpenId == openId {
			return u
		}
	}

	t_user := tables.LoadUser(zoneId, openId)
	if t_user != nil {
		u, _ := GetUser(t_user.Uid)
		UserList[t_user.Uid] = u
		return UserList[t_user.Uid]
	}

	return nil
}

func AddUser(zoneId int, openId string) *User {
	if _, ok := UserList[openId]; !ok {
		uid := tool.UniqueId()
		UserList[uid] = &User{
			Info:    UserInfo{ZoneId: zoneId, Uid: uid, OpenId: openId},
			Profile: UserProfile{Gold: 10000, GuanKaId: 1}}
		return UserList[uid]
	}
	return nil
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

func PersistentUser() {
}

// func CreateTableAccountFromAccount(a *Account) *tables.Account {
// 	return &tables.Account{
// 		AccountId:  a.AccountId,
// 		Account:    a.Account,
// 		Password:   a.Password,
// 		OpenId:     a.OpenId,
// 		LoginTime:  a.LoginTime,
// 		LogoutTime: a.LogoutTime,
// 	}
// }

func CreateTableUserFromUser(u *User) *tables.User {
	return &tables.User{
		Uid:    u.Info.Uid,
		OpenId: u.Info.OpenId,
		ZoneId: u.Info.ZoneId,
	}
}

func CreateUserBaseInfoFromUser(u *User) *tables.UserBaseInfo {
	return &tables.UserBaseInfo{
		Uid:  u.Info.Uid,
		Gold: u.Profile.Gold,
	}
}

func CreateUserExtendInfoFromUser(u *User) *tables.UserExtendInfo {
	return &tables.UserExtendInfo{
		Uid:      u.Info.Uid,
		GuanKaId: u.Profile.GuanKaId,
		Tech:     u.Profile.Tech,
	}
}
