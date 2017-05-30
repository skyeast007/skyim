package web

import (
	"net/http"
	"strings"

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
		ctx.Log.Info("新请求:" + r.URL.Path)
		if strings.HasSuffix(r.URL.Path, ".") {
			route(w, r, ctx)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}

//route 对请求进行简单路由
func route(w http.ResponseWriter, r *http.Request, ctx *context.Context) {
	switch r.URL.Path {
	case "/register":

	}
}
