package web

import (
	"im/context"
	"net/http"
)

//Register 注册一名用户
func Register(w http.ResponseWriter, r *http.Request, ctx *context.Context) {

}

//Identifying 获取验证码
func Identifying(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	session := globalSessions.SessionStart(w, r)
}
