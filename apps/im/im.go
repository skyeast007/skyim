package main

import (
	"im/context"
	"im/im"
	"im/web"
	"runtime"
)

func main() {
	ctx := context.NewCtx()
	//启用全部cup并行
	runtime.GOMAXPROCS(runtime.NumCPU())
	//启动http服务
	go web.StartHTTPServer(ctx)
	//启动websocket服务
	go im.StartWebsocketServer(ctx)
	//阻止主进程退出
	select {}
}
