package web

import (
	"fmt"
	"im/context"
	"im/web/controller"
	"reflect"
	"testing"
)

func Test_AddController(t *testing.T) {
	route := NewRoute(context.NewCtx())
	var u Controller
	u = new(controller.User)
	route.AddController("user", u, map[string]string{"GET": "/:uid/"})
}
func Test_ControllerReflect(t *testing.T) {
	user := new(controller.User)
	fmt.Println("type:", reflect.TypeOf(user))
	f := reflect.ValueOf(user).MethodByName("Get")
	fmt.Println("fff:", f)
	if f.IsValid() {

	}
}
