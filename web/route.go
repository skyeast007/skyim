package web

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"im/context"
	"im/web/handle"
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

func init() {

}

//Controller 控制器接口，所有适用此路由的控制器都必须实现本接口
type Controller interface {
	//init 在路由到控制器任何一个方法前都会先调用此方法
	Init(h *handle.HTTPRouteHandle)
	Get(h *handle.HTTPRouteHandle)
	Post(h *handle.HTTPRouteHandle)
	Put(h *handle.HTTPRouteHandle)
	Delete(h *handle.HTTPRouteHandle)
}

//atomRoute 原子路由
type atomRoute struct {
	controller Controller
	//0到3分别表示GET,POST,PUT,DELETE的正则参数索引
	regex  [4]*regexp.Regexp
	params [4]map[int]string
}

//Route 路由器
type Route struct {
	//静态文件服务器
	staticFileServer http.Handler
	controllers      map[string]*atomRoute
	ctx              *context.Context
}

//NewRoute 获取一个路由器实例
func NewRoute(ctx *context.Context) *Route {
	r := new(Route)
	r.staticFileServer = http.FileServer(http.Dir(ctx.Options.HTTPDocumentRoot))
	r.ctx = ctx
	r.controllers = make(map[string]*atomRoute)
	return r
}

//AddController 添加一个结构体
func (R *Route) AddController(uri string, c Controller, param map[string]string) bool {
	if uri == "" {
		R.ctx.Log.Error("路由添加错误，缺少确定的URI")
		return false
	}
	atom := new(atomRoute)
	atom.controller = c
	var j int
	params := make(map[int]string)
	//路由参数处理
	for k, v := range param {
		parts := strings.Split(v, "/")
		j = 0
		for i, part := range parts {
			if strings.HasPrefix(part, ":") {
				expr := "([^/]+)"
				//支持重新定义正则表达式，比如限定id为数字:/user/:id([0-9]+)
				if index := strings.Index(part, "("); index != -1 {
					expr = part[index:]
					part = part[:index]
				}
				params[j] = part
				parts[i] = expr
				j++
			}
		}
		pattern := strings.Join(parts, "/")
		regex, regexErr := regexp.Compile(pattern)
		if regexErr != nil {
			R.ctx.Log.Error("错误的路由正则", regexErr)
			return false
		}
		inx, err := R.getIndexByMothed(k)
		if err != nil {
			R.ctx.Log.Error("路由初始化错误:", err)
			return false
		}
		atom.regex[inx] = regex
		atom.params[inx] = params
	}
	R.controllers[uri] = atom
	return true
}

//获取索引
func (R *Route) getIndexByMothed(mothed string) (int, error) {
	var index = -1
	var err error
	err = nil
	switch strings.ToUpper(mothed) {
	case GET:
		index = 0
	case POST:
		index = 1
	case PUT:
		index = 2
	case DELETE:
		index = 3
	default:
		err = errors.New("不支持的方法:" + mothed)
	}
	return index, err
}

//FirstToUpper 首字母大写其余小写
func (R *Route) FirstToUpper(str string) string {
	if str == "" {
		return str
	}
	if str == "" {
		return str
	}
	s := []byte(str)
	s = append(bytes.ToUpper(s[:1]), bytes.ToLower(s[1:len(s)])...)
	return string(s)
}

//HTTPStatus HTTP状态返回
func (R *Route) HTTPStatus(w http.ResponseWriter, code int) {
	response := ""
	switch code {
	case 404:
		response = "Not Found"
	case 400:
		response = "Bad Request"
	}
	w.WriteHeader(code)
	w.Write([]byte(response))
}

//getCtx 获得总控模块
func (R *Route) getCtx() *context.Context {
	return R.ctx
}

//ServeHTTP HTTP服务核心处理函数
func (R *Route) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	R.ctx.Log.Debug("新请求:" + r.Method + ":" + r.URL.Path)
	if len(r.URL.Path) > 0 && !strings.HasSuffix(r.URL.Path, ".") {
		split := strings.Split(r.URL.Path, "/")
		//请求资源为空
		if split[1] == "" {
			R.HTTPStatus(w, 404)
			return
		}
		//检查请求的资源是否已定义
		if source, ok := R.controllers[split[1]]; ok {
			index, err := R.getIndexByMothed(r.Method)
			if err != nil {
				R.HTTPStatus(w, 400)
				return
			}
			//检查是否需要进行路由参数匹配
			if source.regex[index] != nil && source.regex[index].MatchString(r.URL.Path) && len(source.params[index]) > 0 {
				matches := source.regex[index].FindStringSubmatch(r.URL.Path)
				values := r.URL.Query()
				for i, match := range matches[1:] {
					values.Add(source.params[index][i], match)
				}
				r.URL.RawQuery = url.Values(values).Encode() + "&" + r.URL.RawQuery
			}
			method := R.FirstToUpper(r.Method)
			h := new(handle.HTTPRouteHandle)
			h.W = w
			h.R = r
			h.Ctx = R.getCtx()
			params := make([]reflect.Value, 1)
			params[0] = reflect.ValueOf(h)
			i := reflect.ValueOf(source.controller).MethodByName("Init")
			//调用初始化函数
			if i.IsValid() {
				i.Call(params)
			}
			//调用业务处理函数
			f := reflect.ValueOf(source.controller).MethodByName(method)
			if f.IsValid() {
				f.Call(params)
			} else {
				R.ctx.Log.Error("不存在的处理方法:" + method)
				R.HTTPStatus(w, 404)
			}
		} else {
			R.HTTPStatus(w, 404)
			return
		}
	} else {
		R.staticFileServer.ServeHTTP(w, r)
	}
}
