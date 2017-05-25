package context

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
