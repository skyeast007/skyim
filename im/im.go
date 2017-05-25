package im

import (
	"bufio"
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
	Name string
	//Auth 授权信息
	Auth *Auth
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
