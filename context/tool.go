package context

import (
	"errors"
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
	return strings.Replace(dir, "\\", "/", -1), err
}

//FileIsExist 判断一个文件是否存在
func (t *Tool) FileIsExist(file string) (bool, error) {
	_, err := os.Stat(file)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

//MapFillToStruct 将map数据填充到结构体中
func (t *Tool) MapFillToStruct(data map[string]interface{}, container interface{}) error {
	return errors.New("type")
}
