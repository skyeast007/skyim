package im

import (
	"fmt"
	"im/context"
	"testing"

	"golang.org/x/net/websocket"
)

//Test_WebsocketServer 测试websocket服务
func Test_WebsocketServer(t *testing.T) {

}

//Test_WebsocketClient 测试客户端
func Test_WebsocketClient(t *testing.T) {
	ctx := context.NewCtx()
	origin := "http://127.0.0.1"
	url := "ws://127.0.0.1" + ctx.Options.WebSoctetAddress
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		t.Error(err)
	}
	if _, err = ws.Write([]byte("hello, world!\n")); err != nil {
		t.Error(err)
	}
	var msg = make([]byte, 512)
	var n int
	if n, err = ws.Read(msg); err != nil {
		t.Error(err)
	}
	fmt.Printf("Received: %s.\n", msg[:n])
}
