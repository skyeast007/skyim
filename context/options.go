package context

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

//Options 系统配置选项
type Options struct {
	//LogLevel 0:all,1:debug,2:info,3:warning,4:error,5:Fatal
	LogLevel int8 `toml:"log_level"`
	//日志目录，默认std输出
	LogPath string `toml:"log_path"`
	//WebSoctetAddress websocket服务监听地址 "127.0.0.1:3000"
	WebSoctetAddress string `toml:"web_socket_address"`
	//TCPAddress TCP服务监听地址
	TCPAddress string `toml:"tcp_address"`
	//HTTPAddress http服务监听地址
	HTTPAddress string `toml:"http_address"`
	//HTTPDocumentRoot http静态服务器绝对地址
	HTTPDocumentRoot string `toml:"http_document_root"`
	//RedisAddress redis连接地址
	RedisAddress string `toml:"redis_address"`
	//RedisAuth redis授权密码
	RedisAuth string `toml:"redis_auth"`
	//DatabaseType数据库类型，默认mysql
	DatabaseType string `toml:"database_type"`
	//DatabaseName 数据库名称
	DatabaseName string `toml:"database_name"`
	//DatabaseUser 数据库用户名
	DatabaseUser string `toml:"database_user"`
	//DatabasePassword 数据库密码
	DatabasePassword string `toml:"database_password"`
	//DatabaseAddress 数据库地址
	DatabaseAddress string `toml:"database_address"`
}
type config map[string]interface{}

//NewOption 初始化配置
func NewOption() *Options {
	var configFile string
	var err error
	var appPath string
	var o *Options
	t := Tool{}
	appPath, err = t.GetAppDirectory()
	flagFile := flag.String("c", "", "配置文件路径")
	flag.Parse()
	if len(*flagFile) == 0 {

		configFile = appPath + "/config.toml"
	} else {
		configFile = *flagFile
	}
	if err != nil {
		log.Println(err, "将使用默认配置...")
	}
	o, err = configFileparse(configFile)
	if err != nil {
		log.Println("解析配置文件失败:", err, "将使用默认配置...")
	}
	if o.LogLevel > 5 || o.LogLevel < 0 {
		o.LogLevel = 0
	}
	if len(o.LogPath) > 0 {
		_, err = os.Stat(o.LogPath)
		if os.IsExist(err) == false {
			log.Fatalln("日志目录不存在:", err)
		}
		if os.IsPermission(err) == false {
			log.Fatalln("日志目录权限不足:", err)
		}
	}
	if o.WebSoctetAddress == "" {
		o.WebSoctetAddress = ":3000"
	}
	if o.TCPAddress == "" {
		o.TCPAddress = ":3001"
	}
	if o.HTTPAddress == "" {
		o.HTTPAddress = ":3002"
	}
	if o.HTTPDocumentRoot == "" {
		o.HTTPDocumentRoot = appPath + "/static"
	}
	println(o.HTTPDocumentRoot)
	_, err = os.Stat(o.HTTPDocumentRoot)
	if os.IsExist(err) == false {
		log.Fatalln("HTTP站点根目录不存在:", err)
	}
	if os.IsPermission(err) == false {
		log.Fatalln("HTTP站点根目录权限不足:", err)
	}
	if o.RedisAddress == "" {
		o.RedisAddress = "127.0.0.1:6379"
	}
	if o.DatabaseType == "" {
		o.DatabaseType = "mysql"
	}
	if o.DatabaseName == "" {
		o.DatabaseName = "skyim"
	}
	if o.DatabaseUser == "" {
		o.DatabaseUser = "root"
	}
	if o.DatabasePassword == "" {
		o.DatabasePassword = "root"
	}
	if o.DatabaseAddress == "" {
		o.DatabaseAddress = "127.0.0.1:3306"
	}
	return o
}

//ConfigFileparse 配置文件解析
func configFileparse(filePath string) (*Options, error) {
	o := new(Options)
	_, err := toml.DecodeFile(filePath, &o)
	return o, err
}
