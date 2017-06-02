package web

import (
	"fmt"
	"im/context"
	"net/http"
	"reflect"
	"strings"
)

const (
	//CONNECT CONNECT请求
	CONNECT = "CONNECT"
	//DELETE DELETE请求
	DELETE = "DELETE"
	//GET GET请求
	GET = "GET"
	//HEAD HEAD请求
	HEAD = "HEAD"
	//OPTIONS OPTIONS请求
	OPTIONS = "OPTIONS"
	//PATCH PATCH请求
	PATCH = "PATCH"
	//POST POST请求
	POST = "POST"
	//PUT PUT请求
	PUT = "PUT"
	//TRACE TRACE请求
	TRACE = "TRACE"
)

//commonly used mime-types
const (
	applicationJSON = "application/json"
)

func init() {

}

//Controller 控制器接口，所有适用此路由的控制器都必须实现本接口
type Controller interface {
	//init 在路由到控制器任何一个方法前都会先调用此方法
	Init(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

//Route 路由器
type Route struct {
	//静态文件服务器
	staticFileServer http.Handler
	controllers      map[string]*Controller
	ctx              *context.Context
}

//NewRoute 获取一个路由器实例
func NewRoute(ctx *context.Context) *Route {
	r := new(Route)
	r.staticFileServer = http.FileServer(http.Dir(ctx.Options.HTTPDocumentRoot))
	r.ctx = ctx
	return r
}

//AddController 添加一个结构体
func (s *Route) AddController(c Controller) {
	fmt.Println(reflect.TypeOf(c))
}

//ServeHTTP HTTP服务核心处理函数
func (s *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.ctx.Log.Debug("新请求:" + r.URL.Path)
	if len(r.URL.Path) > 0 && !strings.HasSuffix(r.URL.Path, ".") {

	} else {
		s.staticFileServer.ServeHTTP(w, r)
	}
}
