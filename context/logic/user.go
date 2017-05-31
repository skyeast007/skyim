package logic

import "im/context"

//User 用户信息映射结构体
type User struct {
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
	UpdateTime int64 `xorm:"update"`
}

//GetUserInfoByAccount 根据帐号获取用户信息
func (u *User) GetUserInfoByAccount(account string) (bool, error) {
	ctx := context.NewCtx()
	return ctx.DB.Engine.Where("account=?", account).Get(u)
}
