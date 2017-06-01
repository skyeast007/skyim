package im

import (
	"fmt"
	"im/context/logic"
)

//Auth 授权信息
type Auth struct {
	//AuthTime 授权时间
	AuthTime int64
	//IsAuth 是否授权通过
	IsAuth bool
}

//Auth 授权
func (a *Auth) Auth(im *Im, param interface{}) {
	User := new(logic.User)
	user, err := User.GetUserInfoByAccount("tttlkkkl")
	fmt.Println(user, err)
}
