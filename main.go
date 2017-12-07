package main

import (
	_ "begoceph/routers"
	"flag"
	"math/rand"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	level := beego.AppConfig.String("logLevel")
	if "" != level {
		flag.Set("v", level)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	// Add the base router for service entrypoint
	beego.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("TenxStorageAgent API service is running.\n"))
	})
	beego.Run()
}
