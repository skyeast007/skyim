package main

import (
	"im/im"
	"im/web"
)

func main() {
	//启动http服务
	go web.StaticHTTPServer()
	//启动websocket服务
	go im.WebsocketServer()
	//阻止主进程退出
	select {}
}
