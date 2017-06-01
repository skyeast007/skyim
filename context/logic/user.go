package logic

import (
	"im/context"
	"strconv"
)

const cacheKeyPrefix = "user"

//TUser 用户信息映射结构体
type TUser struct {
	ID         int64
	Name       string
	Account    string
	Mobile     string
	Sign       string
	Password   string
	Gender     int
	Email      string
	Avatar     string
	Status     int
	CreateTime int64 `xorm:"created"`
	DeleteTime int64
	UpdateTime int64 `xorm:"updated"`
}

//GetUserInfoByAccount 根据帐号获取用户信息
func (u *TUser) GetUserInfoByAccount(account string) (bool, error) {
	ctx := context.NewCtx()
	var has bool
	var err error
	err = ctx.Redis.HGetAll(u.getCacheKeyByAccount(account), u)
	if err != nil {
		ctx.Log.Error("用户缓存读取失败：", err)
	}
	//从数据库读取
	if u.ID <= 0 {
		has, err = ctx.DB.Engine.Where("account=?", account).Get(u)
		if has {
			err = u.updateUserCache(u.getCacheKeyByAccount(account))
			if err != nil {
				ctx.Log.Warning("更新用户缓存失败:", err)
			}
		}
	} else {
		has = true
	}
	return has, err
}

//updateUserCache 更新用户缓存
func (u *TUser) updateUserCache(key string) error {
	ctx := context.NewCtx()
	return ctx.Redis.HMSet(key, u)
}

//getCacheKeyByAccount 获取帐号缓存key
func (u *TUser) getCacheKeyByAccount(account string) string {
	return cacheKeyPrefix + ":account:" + account
}

//getCacheKeyByUID 获取uid缓存key
func (u *TUser) getCacheKeyByUID(UID int64) string {
	return cacheKeyPrefix + ":uid:" + strconv.FormatInt(UID, 10)
}
