package context

//Context ctx
type Context struct {
	Options *Options
	Tool    *Tool
	Log     *Log
}

//NewCtx 获取ctx
func NewCtx() *Context {
	c := new(Context)
	c.Options = NewOption()
	c.Tool = new(Tool)
	c.Log = new(Log)
	return c
}
