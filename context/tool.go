package context

import (
	"os"
	"path/filepath"
	"strings"
)

//Tool 工具函数封装
type Tool struct {
}

//GetAppDirectory 获取程序运行路径
func (t *Tool) GetAppDirectory() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	println(dir)
	return strings.Replace(dir, "\\", "/", -1), err
}
