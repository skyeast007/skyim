package web

import (
	"net/http"
	"strings"

	"im/context"
)

//StartHTTPServer 启动http服务
func StartHTTPServer(ctx *context.Context) {
	h := http.FileServer(http.Dir(ctx.Options.HTTPDocumentRoot))
	ctx.Log.Info("HTTP服务启动...站点根目录:" + ctx.Options.HTTPDocumentRoot)
	ctx.Log.Info("端口:" + ctx.Options.HTTPAddress)
	if err := http.ListenAndServe(ctx.Options.HTTPAddress, Service(h, ctx)); err != nil {
		ctx.Log.Fatal("HTTP服务启动失败:", err)
	}
}

//Service 请求处理
func Service(h http.Handler, ctx *context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx.Log.Info("新请求:" + r.URL.Path)
		if len(r.URL.Path) > 0 && !strings.HasSuffix(r.URL.Path, ".") {
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
		Register(w, r, ctx)
	default:
		w.Write([]byte("不支持的请求"))
	}
}
