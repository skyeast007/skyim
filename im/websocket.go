package im

import (
	"net/http"
	"time"

	"im/context"

	"golang.org/x/net/websocket"
)

//ConnectPool 客户端连接池
var ConnectPool = make([]*Im, 10000)

//clientConnectID 客户端连接标识
var clientConnectID = uint64(0)

//Websocket 连接服务
type Websocket struct {
	ctx *context.Context
}

//StartWebsocketServer 启动websocket服务
func StartWebsocketServer(ctx *context.Context) {
	w := Websocket{ctx: ctx}
	http.Handle("/", websocket.Handler(w.Handle))
	ctx.Log.Info("im websocket 服务启动...")
	if err := http.ListenAndServe(ctx.Options.WebSoctetAddress, nil); err != nil {
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
	im.ClientAddress = ws.RemoteAddr()
	im.User = new(User)
	im.User.auth = new(Auth)
	w.messageLoop(im)
}

//messageLoop 消息循环
func (w *Websocket) messageLoop(im *Im) {
	var err error
	for {
		var reply string
		if err = websocket.Message.Receive(im.WebSocketConn, &reply); err != nil {
			continue
		} else {
			w.ctx.Log.Debug("收到消息..." + reply)
			// TODO: 执行消息处理
			request := im.ctx.Decode(reply)
			if request.Command == "" {
				im.WebsocketSend(im.ctx.Encode(4001, "未知的指令"))
				continue
			}
			if !im.User.auth.IsAuth && request.Command != "auth" {
				im.WebsocketSend(im.ctx.Encode(4002, "未授权的访问"))
				continue
			}
			w.handleMessage(request, im)
		}
	}
}

//handleMessage 消息处理
func (w *Websocket) handleMessage(r context.Request, im *Im) {
	switch r.Command {
	case "auth":
		err := im.User.auth.Auth(im, r.Parameter)
		if err != nil {
			im.WebsocketSend(im.ctx.Encode(4004, err.Error()))
		}
		ID := w.ctx.GUID.GetIncreaseID(&clientConnectID)
		im.ID = int64(ID)
		//加入连接池
		ConnectPool[ID] = im
		im.WebsocketSend(im.ctx.Encode(0, "success", im.User))
	default:
		im.WebsocketSend(im.ctx.Encode(4003, "不支持的指令"))
	}
}
