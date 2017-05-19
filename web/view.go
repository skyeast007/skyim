package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const DOCUMENT_ROOT = "/static"

// StaticHTTPServer 静态服务器
func StaticHTTPServer() {
	dir := GetStaticDirectory()
	h := http.FileServer(http.Dir(dir))

	if err := http.ListenAndServe(":3000", Service(h)); err != nil {
		log.Fatalln("HTTP服务启动失败:", err)
	} else {
		log.Println("HTTP服务启动成功!站点根目录:", dir)
	}
}

//Service 请求处理
func Service(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("新请求:", r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

//GetStaticDirectory 获取静态文件存放目录
func GetStaticDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatalln("根目录获取失败:",err)
	}
	return strings.Replace(dir, "\\", "/", -1) + DOCUMENT_ROOT
}

//Substr 字符切割
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
