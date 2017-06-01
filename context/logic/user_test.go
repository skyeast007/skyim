package logic

import (
	"fmt"
	"testing"
)

func Test_GetUserInfoByAccount(t *testing.T) {
	User := new(TUser)
	has, err := User.GetUserInfoByAccount("tttlkkkl")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("has:", has)
	fmt.Println(User)
}
