package models

import (
	tables "subway/gate/db"
	"subway/tool"
	"time"
)

var (
	RoleList    map[int]map[string]*Role
	AccountList map[string]*Account
)

func init() {
	RoleList = make(map[int]map[string]*Role)
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

func GetAccount(openId string) *Account {
	if u, ok := AccountList[openId]; ok {
		return u
	}

	t_a := tables.LoadAccount(openId)
	if t_a != nil {
		a := CreateAccountFromTableAccount(t_a)
		AccountList[a.OpenId] = a
		return AccountList[openId]
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

func Login(zoneId int, openId string) *Role {
	if account, ok := AccountList[openId]; ok {
		if zoneRoleMap, ok := RoleList[zoneId]; ok {
			if role, ok := zoneRoleMap[account.AccountId]; ok {
				return role
			}
		} else {
			RoleList[zoneId] = make(map[string]*Role)
		}

		t_role := tables.LoadRoleByAccount(zoneId, account.AccountId)
		if t_role != nil {
			role := CreateRoleFromTableRole(t_role)
			RoleList[zoneId][role.AccountId] = role
			return RoleList[zoneId][role.AccountId]
		}
	}
	return nil
}

func AddRole(zoneId int, openId string) *Role {
	if account, ok := AccountList[openId]; ok {
		if _, ok := RoleList[zoneId]; !ok {
			RoleList[zoneId] = make(map[string]*Role)
		}

		if role, ok := RoleList[zoneId][account.AccountId]; ok {
			return role
		} else {
			roleId := tool.UniqueId()
			RoleList[zoneId][account.AccountId] = &Role{
				ZoneId:    zoneId,
				RoleId:    roleId,
				AccountId: account.AccountId,
			}
			return RoleList[zoneId][account.AccountId]
		}
	}
	return nil
}

func PersistentUser() {
	accounts := make([]*tables.Account, 0)
	for _, a := range AccountList {
		accounts = append(accounts, CreateTableAccountFromAccount(a))
	}
	tables.PersistentAccount(accounts)

	roles := make([]*tables.Role, 0)
	for _, roleMap := range RoleList {
		for _, role := range roleMap {
			roles = append(roles, CreateTableRoleFromRole(role))
		}
	}
	tables.PersistentRole(roles)
}

func CreateTableAccountFromAccount(a *Account) *tables.Account {
	return &tables.Account{
		AccountId:  a.AccountId,
		Account:    a.Account,
		Password:   a.Password,
		OpenId:     a.OpenId,
		LoginTime:  a.LoginTime,
		LogoutTime: a.LogoutTime,
	}
}

func CreateAccountFromTableAccount(a *tables.Account) *Account {
	return &Account{
		AccountId:  a.AccountId,
		Account:    a.Account,
		Password:   a.Password,
		OpenId:     a.OpenId,
		LoginTime:  a.LoginTime,
		LogoutTime: a.LogoutTime,
	}
}

func CreateTableRoleFromRole(r *Role) *tables.Role {
	return &tables.Role{
		ZoneId:    r.ZoneId,
		RoleId:    r.RoleId,
		AccountId: r.AccountId,
	}
}

func CreateRoleFromTableRole(r *tables.Role) *Role {
	return &Role{
		ZoneId:    r.ZoneId,
		RoleId:    r.RoleId,
		AccountId: r.AccountId,
	}
}
