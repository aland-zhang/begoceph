package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type BaseController struct {
	beego.Controller
	UserName string
	Token    string
}

const ApiPrefix = "/api/v1"

// init token and username
func (bs *BaseController) Init(ct *context.Context, controllerName, actionName string, app interface{}) {
	bs.Init(ct, controllerName, actionName, app)
}

// check request
func (bs *BaseController) Prepare() {
}
