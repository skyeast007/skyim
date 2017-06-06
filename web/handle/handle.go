package handle

import (
	"encoding/json"
	"im/context"
	"io/ioutil"
	"net/http"
	"strconv"
)

//json
const (
	applicationJSON = "application/json"
)

//Response web json响应结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//HTTPRouteHandle 路由处理
type HTTPRouteHandle struct {
	W   http.ResponseWriter
	R   *http.Request
	Ctx *context.Context
}

//JSONResponse json响应
func (h *HTTPRouteHandle) JSONResponse(code int, msg string, data ...interface{}) {
	response := Response{code, msg, data}
	content, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		http.Error(h.W, err.Error(), http.StatusInternalServerError)
		return
	}
	h.W.Header().Set("Content-Length", strconv.Itoa(len(content)))
	h.W.Header().Set("Content-Type", applicationJSON)
	h.W.Write(content)
}

//JSONRead 将客户端提交的json数据解析到对应的结构中
func (h *HTTPRouteHandle) JSONRead(v interface{}) error {
	body, err := ioutil.ReadAll(h.R.Body)
	h.R.Body.Close()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
