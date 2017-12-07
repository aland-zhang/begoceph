package controllers

import (
	"encoding/base64"
	"flag"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/golang/glog"
)

type BaseController struct {
	beego.Controller
	UserName string
	Token    string
}

const ApiPrefix = "/api/v1"

func init() {
	// allow specify config file path
	configPath := flag.String("appconfig", "conf/app.conf", "config path")

	flag.Parse()
	flag.Set("logtostderr", "true")

	err := beego.LoadAppConfig("ini", *configPath)
	if err != nil {
		panic("Load config file " + *configPath + " failed: " + err.Error())
	}
}

// init token and username
func (bs *BaseController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	bs.Init(ct, controllerName, actionName, app)
}

// check request User authentication
func (bs *BaseController) Prepare() {
	authorization := bs.Ctx.Request.Header.Get("authorization")
	if authorization == "" {
		glog.Error("need authorization in header")
		return
	}
	if authorization != bs.tokenFactory() {
		glog.Error("no authorized")
		return
	}
	bs.UserName = bs.ConfUserName()
	bs.Token = bs.ConfToken()
	return
}

func (bs *BaseController) tokenFactory() string {
	ut := "Basic " + beego.AppConfig.String("username") + ":" + beego.AppConfig.String("token")
	return base64.StdEncoding.EncodeToString([]byte(ut))
}

func (bs *BaseController) ConfUserName() string {
	return beego.AppConfig.String("username")
}

func (bs *BaseController) ConfToken() string {
	return beego.AppConfig.String("token")
}
