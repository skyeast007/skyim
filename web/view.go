package web

import (
	"net/http"

	"im/context"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(ctx *context.Context) {
	h := http.FileServer(http.Dir(ctx.Options.HTTPDocumentRoot))

	if err := http.ListenAndServe(ctx.Options.HTTPAddress, Service(h, ctx)); err != nil {
		ctx.Log.Fatal("HTTP服务启动失败:", err)
	} else {
		ctx.Log.Info("HTTP服务启动成功!站点根目录:", ctx.Options.HTTPDocumentRoot)
	}
}

//Service 请求处理
func Service(h http.Handler, ctx *context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx.Log.Info("新请求:", r.URL.Path)
		h.ServeHTTP(w, r)
	})
}
