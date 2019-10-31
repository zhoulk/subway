package models

import (
	"errors"
	"subway/tool"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
}

type User struct {
	Info    UserInfo
	Profile UserProfile

	Heros []*Hero
	Copys []*Copy
}

type UserInfo struct {
	Uid      string
	OpenId   string
	Username string
}

type UserProfile struct {
	Gold          int64
	GuanKaId      int
	GoldGainSpeed int64 // 金币增长速度  个/秒
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func Login(openId string, username string) *User {
	for _, u := range UserList {
		if u.Info.OpenId == openId {
			return u
		}
	}
	return nil
}

func AddUser(openId string, username string) *User {
	if _, ok := UserList[openId]; !ok {
		uid := tool.UniqueId()
		UserList[uid] = &User{
			Info:    UserInfo{uid, openId, username},
			Profile: UserProfile{Gold: 10000, GuanKaId: 1, GoldGainSpeed: 1}}
		return UserList[uid]
	}
	return nil
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

func ChangeGuanKaId(uid string, gkId int) {
	u, _ := GetUser(uid)
	if u != nil {
		u.Profile.GoldGainSpeed = int64(gkId * 10)
	}
}
