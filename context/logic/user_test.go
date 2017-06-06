package logic

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func Test_GetUserInfoByAccount(t *testing.T) {
	User := new(TUser)
	has := User.GetUserInfoByAccount("tttlkkkl")
	if has != true {
		t.Error("数据获取失败")
	}
	fmt.Println("has:", has)
	fmt.Println(User)
}
func Test_PasswordSalt(t *testing.T) {
	User := new(TUser)
	User.ID = 1
	var p string
	for i := 0; i < 100; i++ {
		rand.New(rand.NewSource(time.Now().UnixNano()))
		p = strconv.Itoa(rand.Int())
		fmt.Println("原有:", p)
		has1 := User.PasswordSalt(p)
		fmt.Println("has1:", has1)
		for j := 0; j < 10; j++ {
			sp := User.PasswordSalt(p)
			if sp != has1 {
				t.Error("结果不一致")
			}
		}
	}
}
