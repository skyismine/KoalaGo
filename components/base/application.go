package base

import (
	"github.com/skyismine/KoalaGo/components/logging"
	"github.com/skyismine/KoalaGo/components/router"
	"net/http"
)

type IApplication interface {
	Init()
	Run()
}

//Application 继承自 Module，是应用的入口模块
type Application struct {
	Module
	Router *router.Router
	Logger *logging.Logger
}

const (
	EVENT_BEFORE_REQUEST = "beforeRequest"
	EVENT_AFTER_REQUEST = "afterRequest"
)

//Application 状态
const (
	STATE_BEGIN = 0
	STATE_INIT = 1
	STATE_BEFORE_REQUEST = 2
	STATE_HANDLING_REQUEST = 3
	STATE_AFTER_REQUEST = 4
	STATE_SENDING_RESPONSE = 5
	STATE_END = 6
)

func (app *Application) Init() {
	r := app.GetService("router")
	if r == nil {
		panic("application Init error, router not configure")
	}
	app.Router = r.(*router.Router)
	l := app.GetService("logger")
	if l == nil {
		panic("application Init error, logger not configure")
	}
	app.Logger = l.(*logging.Logger)
}

func (app *Application) Run() {
	http.ListenAndServe(":10000", app.Router)
}