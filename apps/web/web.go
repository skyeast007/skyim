package main

import "im/web"
import "im/context"

func main() {
	ctx := context.NewCtx()
	web.StartHTTPServer(ctx)
}
