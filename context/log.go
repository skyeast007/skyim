package context

import "log"

//Log 日志打印
type Log struct{}

const (
	//Debug 调试日志
	Debug = 1
	//Info 信息日志
	Info = 2
	//Warning 警告日志
	Warning = 3
	//Error 错误日志
	Error = 4
	//Fatal 致命错误日志,此类日志输出后程序退出
	Fatal = 5
)

//Debug 打印调试日志
func (l *Log) Debug(info string, args ...interface{}) {
	l.Log(Debug, info, args)
}

//Info 打印普通日志
func (l *Log) Info(info string, args ...interface{}) {
	l.Log(Info, info, args)
}

//Warning 打印警告日志
func (l *Log) Warning(info string, args ...interface{}) {
	l.Log(Warning, info, args)
}

//Error 打印错误日志
func (l *Log) Error(info string, args ...interface{}) {
	l.Log(Error, info, args)
}

//Fatal 打印致命错误日志
func (l *Log) Fatal(info string, args ...interface{}) {
	l.Log(Fatal, info, args)
}

//Log 日志记录
func (l *Log) Log(level int, info string, args ...interface{}) {
	levelString := "Info"
	switch level {
	case Debug:
		levelString = "Debug"
	case Info:
		levelString = "Info"
	case Warning:
		levelString = "Warning"
	case Error:
		levelString = "Error"
	case Fatal:
		levelString = "Fatal"
	}

	if level == Fatal {
		if len(args) > 1 {
			log.Fatalln(levelString, ":", info, "->", args)
		} else {
			log.Fatalln(levelString, ":", info)
		}
	} else {
		if len(args) > 1 {
			log.Println(levelString, ":", info, "->", args)
		} else {
			log.Println(levelString, ":", info)
		}
	}
}
