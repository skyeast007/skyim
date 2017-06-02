package controller

import "net/http"

//User 用户restfull操作接口实现
type User struct {
}

//Init 初始操作
func (u *User) Init(w http.ResponseWriter, r *http.Request) {}

//Get get请求获取一个用户信息
func (u *User) Get(w http.ResponseWriter, r *http.Request) {}

//Post post请求，新建一个用户信息
func (u *User) Post(w http.ResponseWriter, r *http.Request) {}

//Put 更新一个用户信息
func (u *User) Put(w http.ResponseWriter, r *http.Request) {}

//Delete 删除一个用户
func (u *User) Delete(w http.ResponseWriter, r *http.Request) {}
