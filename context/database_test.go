package context

import (
	"fmt"
	"testing"
)

//Test_DatabaseConnection 测试数据库连接
func Test_DatabaseConnection(t *testing.T) {
	ctx := NewCtx()
	DB := NewDatabase(ctx.Options, ctx.Log)
	info, err := DB.Engine.DBMetas()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(info)
}
