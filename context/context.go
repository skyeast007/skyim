package context

//全局变量
var ctx *Context

//Context ctx
type Context struct {
	Options *Options
	Tool    *Tool
	Log     *Log
	GUID    *GUID
	DB      *Database
	Redis   *Redis
	//Version 版本信息
	Version string
}

//NewCtx 获取ctx
func NewCtx() *Context {
	if ctx == nil {
		c := new(Context)
		c.Options = NewOption()
		c.Tool = new(Tool)
		c.Log = new(Log)
		c.GUID = new(GUID)
		c.DB = NewDatabase(c.Options, c.Log)
		c.Redis = NewRedisPool(c.Options, c.Log)
		c.Version = "1.0"

		ctx = c
	}
	return ctx
}
