package web

import (
	"im/context"
	"im/web/controller"
	"testing"
)

func Test_AddController(t *testing.T) {
	route := NewRoute(context.NewCtx())
	var u Controller
	u = new(controller.User)
	route.AddController(u)
}
