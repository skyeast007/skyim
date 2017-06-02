package im

import (
	"fmt"
	"im/context"
	"testing"
)

func Test_Auth(t *testing.T) {
	var authInfo = "{\"command\":\"auth\",\"param\":{\"account\":123456,\"password\":\"password\"}}"
	ctx := context.NewCtx()
	param := ctx.Decode(authInfo)
	fmt.Println("xxxx", param)
	// im := new(Im)
	// im.User.auth = new(Auth)
	// err := im.User.auth.Auth(im, param.Parameter)
	// if err != nil {
	// 	t.Error(err)
	// }
}
