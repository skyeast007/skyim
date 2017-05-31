package web

import (
	"net/http"
	"regexp"

	"im/context"
)

//Register 注册一名用户
func Register(w http.ResponseWriter, r *http.Request, ctx *context.Context) bool {
	if r.Method != "POST" {
		w.Write([]byte("只支持POST访问"))
		return false
	}
	if r.FormValue("account") == "" {
		w.Write([]byte("账号不能为空哦"))
		return false
	}
	if match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", r.FormValue("account")); match != true {
		w.Write([]byte("账号不能为空账号只能是英文、数字或下划线~_~"))
		return false
	}
	if r.FormValue("password") == "" {
		w.Write([]byte("登录密码不能为空哦"))
		return false
	}
	user := new(User)
	user.Account = r.FormValue("account")
	user.Name = "sky_" + r.FormValue("account")
	return true
}

//Identifying 获取验证码
func Identifying(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
}
