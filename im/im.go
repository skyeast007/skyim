package im

import (
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

// Client 客户端连接标识
type Client struct {
	Id          *websocket.Conn
	ConnectTime int64
}

//ConnectPool 客户端连接池
var ConnectPool = make([]*Client, 0)

func ImServer() {
	//go eventPool()
	http.Handle("/", websocket.Handler(Handle))
	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatalln("websocket端口监听错误:", err)
	} else {
		log.Println("im websocket 服务启动...")
	}
}

//Handle 处理
func Handle(ws *websocket.Conn) {
	client := Client{Id: ws, ConnectTime: time.Now().Unix()}
	ConnectPool = append(ConnectPool, &client)
}
func eventPool() {
	//var err error
	for {
		for client := range ConnectPool {
			//var reply string
			println("ws:", client)
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
