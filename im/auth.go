package im

import (
	"errors"
	"im/context/logic"
	"time"
)

//Auth 授权信息
type Auth struct {
	//AuthTime 授权时间
	AuthTime int64
	//IsAuth 是否授权通过
	IsAuth bool
}

//Auth 授权
func (a *Auth) Auth(im *Im, param interface{}) error {
	params, ok := param.(map[string]interface{})
	if !ok {
		return errors.New("参数错误")
	}
	var account, password string
	if a, ok := params["account"]; ok {
		if a == "" {
			return errors.New("帐号信息不能为空")
		}
		account = a.(string)

	} else {
		return errors.New("缺少帐号信息")
	}
	if p, ok := params["password"]; ok {
		if p == "" {
			return errors.New("密码信息不能为空")
		}
		password = p.(string)

	} else {
		return errors.New("缺少密码信息")
	}
	User := new(logic.TUser)
	has := User.GetUserInfoByAccount(account)
	if !has {
		return errors.New("不存在的用户信息")
	}
	//记录每次验证时间
	im.User.auth.AuthTime = time.Now().Unix()
	if password != User.Password {
		return errors.New("用户名或密码错误")
	}
	im.User.auth.IsAuth = true
	im.User.UID = User.ID
	im.User.Name = User.Name
	im.User.Account = User.Account
	im.User.Mobile = User.Mobile
	im.User.Sign = User.Sign
	im.User.Email = User.Email
	im.User.Avatar = User.Avatar
	im.User.CreateTime = User.CreateTime
	return nil
}
