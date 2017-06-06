package web

import (
	"net/http"

	"im/context"
	"im/web/controller"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(ctx *context.Context) {
	ctx.Log.Info("HTTP服务启动...站点根目录:" + ctx.Options.HTTPDocumentRoot)
	ctx.Log.Info("端口:" + ctx.Options.HTTPAddress)
	route := NewRoute(ctx)
	//添加一个控制器路由
	route.AddController("user", new(controller.User), map[string]string{"GET": "/:uid([0-9]+)/:cid", "PUT": "/:uid/"})
	if err := http.ListenAndServe(ctx.Options.HTTPAddress, route); err != nil {
		ctx.Log.Fatal("HTTP服务启动失败:", err)
	}
}
