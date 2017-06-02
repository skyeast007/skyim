package web

import (
	"net/http"

	"im/context"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(ctx *context.Context) {
	ctx.Log.Info("HTTP服务启动...站点根目录:" + ctx.Options.HTTPDocumentRoot)
	ctx.Log.Info("端口:" + ctx.Options.HTTPAddress)
	route := NewRoute(ctx)
	if err := http.ListenAndServe(ctx.Options.HTTPAddress, route); err != nil {
		ctx.Log.Fatal("HTTP服务启动失败:", err)
	}
}
