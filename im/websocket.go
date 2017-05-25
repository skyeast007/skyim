package im

import (
	"net/http"
	"runtime"
	"time"

	"im/context"

	"golang.org/x/net/websocket"
)

//ConnectPool 客户端连接池
var ConnectPool = make([]*Im, 0)

//Websocket 连接服务
type Websocket struct {
	ctx *context.Context
}

//StartWebsocketServer 启动websocket服务
func StartWebsocketServer(ctx *context.Context) {
	w := Websocket{ctx: ctx}
	http.Handle("/", websocket.Handler(w.Handle))
	ctx.Log.Info("im websocket 服务启动...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		ctx.Log.Fatal("websocket端口监听错误:", err)
	}
}

//Handle 连接处理
func (w *Websocket) Handle(ws *websocket.Conn) {
	im := new(Im)
	ConnectPool = append(ConnectPool, im)
	im.WebSocketConn = ws
	im.ctx = w.ctx
	im.ConnType = 1
	im.ConnTime = time.Now().Unix()
	go w.messageLoop(im)
}

//messageLoop 消息循环
func (w *Websocket) messageLoop(im *Im) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(im.WebSocketConn, &reply); err != nil {
			w.ctx.Log.Error("接收消息失败!")
			runtime.Gosched()
			continue
		} else {
			w.ctx.Log.Debug("收到消息...")
			// TODO: 执行消息处理
		}
	}
}
