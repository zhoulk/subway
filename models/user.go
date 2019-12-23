package models

import (
	"errors"
	"subway/db/tables"
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

func Login(zoneId int, openId string, username string) *User {
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

func AddUser(zoneId int, openId string, username string) *User {
	if _, ok := UserList[openId]; !ok {
		uid := tool.UniqueId()
		UserList[uid] = &User{
			Info:    UserInfo{zoneId, uid, openId, username},
			Profile: UserProfile{Gold: 10000, GuanKaId: 1}}
		return UserList[uid]
	}
	return nil
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

func PersistentUser() {
	users := make([]*tables.User, 0)
	userHeros := make([]*tables.UserHero, 0)
	heroEquips := make([]*tables.HeroEquip, 0)
	heroSkills := make([]*tables.HeroSkill, 0)
	userBaseInfos := make([]*tables.UserBaseInfo, 0)
	userExtendInfos := make([]*tables.UserExtendInfo, 0)
	userCopys := make([]*tables.UserCopy, 0)
	userCopyItems := make([]*tables.UserCopyItem, 0)
	userBags := make([]*tables.UserBag, 0)

	for _, u := range UserList {
		users = append(users, CreateTableUserFromUser(u))

		userBaseInfos = append(userBaseInfos, CreateUserBaseInfoFromUser(u))
		userExtendInfos = append(userExtendInfos, CreateUserExtendInfoFromUser(u))

		for _, u_h := range u.Heros {
			userHeros = append(userHeros, CreateUserHeroFromHero(u.Info.Uid, u_h))

			for _, u_h_e := range u_h.Equips {
				heroEquips = append(heroEquips, CreateHeroEquipFromEquip(u_h.Uid, u_h.Info.Floor, u_h_e))
			}

			for _, u_h_s := range u_h.Skills {
				heroSkills = append(heroSkills, CreateHeroSkillFromSkill(u_h.Uid, u_h_s))
			}
		}

		for _, u_c := range u.Copys {
			userCopys = append(userCopys, CreateUserCopyFromCopy(u.Info.Uid, u_c))
		}

		for _, u_c_i := range u.CopyItems {
			userCopyItems = append(userCopyItems, CreateUserCopyItemFromCopyItem(u.Info.Uid, u_c_i))
		}

		bag := GetBag(u.Info.Uid)
		for _, u_b_i := range bag.Items {
			userBags = append(userBags, CreateUserBagFromBagItem(u.Info.Uid, u_b_i))
		}
	}

	tables.PersistentUser(users)
	tables.PersistentUserHero(userHeros)
	tables.PersistentHeroEquip(heroEquips)
	tables.PersistentHeroSkill(heroSkills)
	tables.PersistentUserBaseInfo(userBaseInfos)
	tables.PersistentUserExtendInfo(userExtendInfos)
	tables.PersistentUserCopy(userCopys)
	tables.PersistentUserCopyItem(userCopyItems)
	tables.PersistentUserBags(userBags)
}

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
