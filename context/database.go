package context

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//Database 数据库操作
type Database struct {
	Engine *xorm.Engine
}

//NewDatabase 初始化数据库引擎
func NewDatabase(o *Options, l *Log) *Database {
	DB := new(Database)
	var err error
	var connectString = o.DatabaseUser + ":" + o.DatabasePassword + "@/" + o.DatabaseName + "?charset=utf8"
	DB.Engine, err = xorm.NewEngine(o.DatabaseType, connectString)
	if err != nil {
		l.Fatal("数据库连接错误:", err)
	}
	return DB
}
