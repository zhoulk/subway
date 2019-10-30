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
	Info UserInfo

	Heros []*Hero
}

type UserInfo struct{
	Uid       string
	OpenId string
	Username string
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
	for _,u := range UserList {
		if u.Info.OpenId == openId {
			return u
		}
	}
	return nil
}

func AddUser(openId string, username string) *User{
	if _,ok := UserList[openId];!ok{
		uid := tool.UniqueId()
		UserList[uid] = &User{Info: UserInfo{uid, openId, username}}
		return UserList[openId]
	}
	return nil
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
