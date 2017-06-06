package logic

import (
	"crypto/sha1"
	"im/context"
	"io"
	"strconv"
)

//cacheKeyPrefix 缓存key前缀
const cacheKeyPrefix = "user"

//passwordSalt 密码加盐参数
const passwordSalt = "skyim"

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
	UpdateTime int64
}

//GetUserInfoByAccount 根据帐号获取用户信息
func (u *TUser) GetUserInfoByAccount(account string) bool {
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
	if err != nil {
		ctx.Log.Error("数据错误:", err)
	}
	return has
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

//PasswordSalt 进行密码加盐
func (u *TUser) PasswordSalt(password string) string {
	if password == "" {
		return ""
	}
	passwordByte := []byte(password)
	id := []byte(passwordSalt)
	passwordLen := len(passwordByte)
	idLen := len(id)
	newPAssword := make([]byte, passwordLen*(idLen+1))
	for k, v := range passwordByte {
		newPAssword[k] = v << 4
		for i, val := range id {
			newPAssword[k*idLen+i] = val << 5
		}
	}
	h := sha1.New()
	io.WriteString(h, string(newPAssword))
	return string(h.Sum(nil))
}

//CheckPassword 检查密码是否正确
func (u *TUser) CheckPassword(password string) bool {
	if u.ID <= 0 || u.Password == "" {
		return false
	}
	if u.Password == u.PasswordSalt(password) {
		return true
	}
	return false
}

//CreateUser 新建用户
func (u *TUser) CreateUser() int64 {
	ctx := context.NewCtx()
	id, err := ctx.DB.Engine.InsertOne(u)
	if err != nil {
		ctx.Log.Error("新建用户失败", err)
		id = int64(0)
	}
	u.ID = id
	return id
}
