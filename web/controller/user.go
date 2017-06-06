package controller

import (
	"fmt"
	"regexp"

	"github.com/gorilla/sessions"

	"im/context/logic"
	"im/web/handle"
)

//User 用户restfull操作接口实现
type User struct {
}

//Init 初始操作
func (u *User) Init(h *handle.HTTPRouteHandle) {
	//h.W.Write([]byte("Init"))
}

//Get get请求获取一个用户信息
func (u *User) Get(h *handle.HTTPRouteHandle) {
	h.Session.Values["foo"] = "session test"
	fmt.Println("---id:", h.Session.ID)
	if err := sessions.Save(h.R, h.W); err != nil {
		fmt.Println(err)
	}
	fmt.Println("请求查询参数:", h.R.URL.Query(), "session set:", h.Session.Name())
	fmt.Println("session value:", h.Session.Values)
	h.W.Write([]byte("get"))
}

//Post post请求，新建一个用户信息 即注册
func (u *User) Post(h *handle.HTTPRouteHandle) {
	fmt.Println("session:", h.Session.Values)
	fmt.Println("session set:", h.Session.Name())
	account := h.R.FormValue("account")
	password := h.R.FormValue("password")
	rePassword := h.R.FormValue("rePassword")
	if account == "" {
		h.JSONResponse(6001, "账号不能为空哦")
		return
	}
	if match, _ := regexp.MatchString("^[a-zA-Z0-9_]+$", account); match != true {
		h.JSONResponse(6002, "账号不能为空账号只能是英文、数字或下划线~_~")
		return
	}
	if password == "" {
		h.JSONResponse(6003, "登录密码不能为空哦")
		return
	}
	if len(password) < 6 {
		h.JSONResponse(6004, "登录密码长度不能小于6个字符")
		return
	}
	if rePassword == "" {
		h.JSONResponse(6003, "请再次确认登录密码")
		return
	}
	if password != rePassword {
		h.JSONResponse(6005, "两次登录密码不一致")
		return
	}
	Us := new(logic.TUser)
	if Us.GetUserInfoByAccount(account) {
		h.JSONResponse(6007, "已存在相同账号")
		return
	}
	user := new(logic.TUser)
	user.Name = account
	user.Account = account
	user.Password = user.PasswordSalt(password)
	if ok := user.CreateUser(); ok > 0 {
		h.JSONResponse(0, "success")
	} else {
		h.JSONResponse(6006, "注册失败，请重试")
	}
}

//Put 更新一个用户信息
func (u *User) Put(h *handle.HTTPRouteHandle) {
	h.W.Write([]byte("put"))
}

//Delete 删除一个用户
func (u *User) Delete(h *handle.HTTPRouteHandle) {
	h.W.Write([]byte("delete"))
}
