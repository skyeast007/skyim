package im

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/websocket"
)

// Client 客户端连接标识
type Client struct {
	ID          *websocket.Conn
	ConnectTime int64
}

//ConnectPool 客户端连接池
var ConnectPool = make([]*Client, 0)

//WebsocketServer websocket服务器
func WebsocketServer() {
	go eventPool()
	log.Println("xxxx")
	http.Handle("/", websocket.Handler(Handle))
	log.Println("im websocket 服务启动...")
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatalln("websocket端口监听错误:", err)
	}
}

//Handle 处理
func Handle(ws *websocket.Conn) {
	client := Client{ID: ws, ConnectTime: time.Now().Unix()}
	ConnectPool = append(ConnectPool, &client)
}
func eventPool() {
	//var err error
	for {
		for k, client := range ConnectPool {
			//var reply string
			println("ws_v:", client)
			println("ws_k", k)
			// if err = websocket.Message.Receive(client.Id, &reply); err != nil {
			// 	log.Println("接收消息失败!")
			// 	break
			// }
			// log.Println("收到消息:", reply)
			// if err = websocket.Message.Send(client.Id, "你好"); err != nil {
			// 	log.Println("消息发送失败")
			// 	break
			// }
		}
	}
}
