package logic

import (
	"fmt"
	"im/context"
	"reflect"
)

const cacheKeyPrefix = "user:"

//User 用户信息映射结构体
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
	redisData, err := ctx.Redis.Pool.Do("hgetall", "xxx")
	x:=[]interface{}
	ctx.Redis.ScanStruct(redisData, x)
	fmt.Println(redisData, "redis错误信息：", err, reflect.TypeOf(redisData))
	return ctx.DB.Engine.Where("account=?", account).Get(u)
}
