package context

import "encoding/json"

//Response 客户端响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//Request 客户端请求定义
type Request struct {
	Command   string      `json:"command"`
	Parameter interface{} `json:"param"`
}

//Context ctx
type Context struct {
	Options *Options
	Tool    *Tool
	Log     *Log
	//Version 版本信息
	Version string
}

//NewCtx 获取ctx
func NewCtx() *Context {
	c := new(Context)
	c.Options = NewOption()
	c.Tool = new(Tool)
	c.Log = new(Log)
	c.Version = "1.0"
	return c
}

//Encode 对数据进行json编码
func (c *Context) Encode(code int, msg string, data ...interface{}) []byte {
	response := new(Response)
	response.Code = code
	response.Msg = msg
	response.Data = data
	str, err := json.Marshal(*response)
	if err != nil {
		c.Log.Error("json编码错误:", err)
		str = []byte("")
	}
	return str
}

//Decode 客户端消息json解码
func (c *Context) Decode(data string) Request {
	var request Request
	err := json.Unmarshal([]byte(data), &request)
	if err != nil {
		c.Log.Error("json解码错误:", err)
	}
	return request
}
