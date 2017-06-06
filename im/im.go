package im

import (
	"bufio"
	"encoding/json"
	"errors"
	"net"
	"sync"

	"im/context"

	"golang.org/x/net/websocket"
)

//User 用户信息
type User struct {
	//UID 用户
	UID int64
	//Name 用户名
	Name       string
	Account    string
	Mobile     string
	Sign       string
	Gender     int
	Email      string
	Avatar     string
	CreateTime int64
	//Auth 授权信息
	auth *Auth
}

//TCP 连接信息
type TCP struct {
	writeLock sync.RWMutex
	metaLock  sync.RWMutex

	Reader *bufio.Reader
	Writer *bufio.Writer
}

//Im 客户端连接标识包含用户基本信息以及连接信息
type Im struct {
	//连接标识
	ID            int64
	WebSocketConn *websocket.Conn
	TCPConn       *net.TCPConn
	//ConnType 当前连接类型 1:websocket 2:TCP
	ConnType int8

	//ctx 全局操作对象
	ctx *context.Context
	//UserAgent 客户端自述,客户端类型
	UserAgent string

	//ConnTime 建立连接的时间
	ConnTime int64
	//LastReceiveTime 最近一次收到消息的时间
	LastReceiveTime int64
	//LastSendTime 最近一次发送消息的时间
	LastSendTime int64
	//ClientAddress 客户端地址
	ClientAddress net.Addr
	//通过TCP连接时建立此项
	TCP  *TCP
	User *User
}

//Response im客户端响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Request im客户端请求定义
type Request struct {
	Command   string      `json:"command"`
	Parameter interface{} `json:"param"`
}

//Encode 对数据进行json编码
func (i *Im) Encode(code int, msg string, data ...interface{}) []byte {
	response := new(Response)
	response.Code = code
	response.Msg = msg
	response.Data = data
	str, err := json.Marshal(*response)
	if err != nil {
		i.ctx.Log.Error("json编码错误:", err)
		str = []byte("")
	}
	return str
}

//Decode 客户端消息json解码
func (i *Im) Decode(data string) Request {
	var request Request
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		i.ctx.Log.Error("json解码错误:", err)
	}
	return request
}

//WebsocketSend 向客户端发送websocket消息
func (i *Im) WebsocketSend(msg []byte) error {
	if i.WebSocketConn == nil {
		return errors.New("连接未建立")
	}
	return websocket.Message.Send(i.WebSocketConn, string(msg))
}

//TCPSend 向客户端发送tcp消息
func (i *Im) TCPSend(msg string) {

}
