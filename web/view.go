package web

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// StaticHTTPServer 静态服务器
func StaticHTTPServer() {
	http.Handle("/", http.FileServer(http.Dir(GetStaticDirectory())))
	http.ListenAndServe(":8080", nil)
}

//GetStaticDirectory 获取静态文件存放目录
func GetStaticDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dir = strings.Replace(dir, "\\", "/", -1)
	return Substr(dir, 0, strings.LastIndex(dir, "/"))
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
