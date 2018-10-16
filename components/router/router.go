package router

import (
	"fmt"
	"github.com/skyismine/KoalaGo/components/base"
	"github.com/skyismine/KoalaGo/components/urlmanager"
	"github.com/skyismine/KoalaGo/components/utils"
	"net/http"
	"regexp"
)

type FilterFun func(w http.ResponseWriter, r *http.Request) bool

type IRouter interface {
	http.Handler
	SetModule(id string, module *base.Module)
	GetModule(id string) (module *base.Module)
	DelModule(id string)
	HasModule(id string) (has bool)
}

type Router struct {
	UrlManager urlmanager.IUrlManager
	CatchAll map[string][]FilterFun
	Modules map[string]*base.Module
}

func New(app *base.IApplication) IRouter {
	um := app
	return &Router{
		nil,
		make(map[string][]FilterFun),
		make(map[string]*base.Module),
	}
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	moduleid, controllerid, actionid := router.UrlManager.ParseRequest(path)
	if len(moduleid) <= 0 || len(controllerid) <= 0 || len(actionid) <= 0 {
		utils.HttpRspJson(w, 404, utils.HTTP_RSP_JSON_RESULT_ERROR, "0001", fmt.Sprintf("access path not exists: %s", path), nil)
		return
	}
	// 根据路径匹配所有的过滤器，过滤器使用正则表达式来匹配
	filters := []FilterFun{}
	for k, v := range router.CatchAll {
		reg, err := regexp.Compile(k)
		if err == nil && reg.MatchString(path) {
			filters = append(filters, v...)
		}
	}
	// 调用所有的过滤器进行预处理
	for _, f := range filters {
		skip := f(w, r)
		if skip {
			return
		}
	}
	// 查找模块、控制器、动作，执行找到的动作函数
	module := router.GetModule(moduleid)
	if module == nil {
		utils.HttpRspJson(w, 500, utils.HTTP_RSP_JSON_RESULT_ERROR, "0002", fmt.Sprintf("module not exists: %s", path), nil)
		return
	}
	controller := module.GetController(controllerid)
	if controller == nil {
		utils.HttpRspJson(w, 500, utils.HTTP_RSP_JSON_RESULT_ERROR, "0002", fmt.Sprintf("controller not exists: %s", path), nil)
		return
	}
	action := controller.GetAction(actionid)
	if action == nil {
		utils.HttpRspJson(w, 500, utils.HTTP_RSP_JSON_RESULT_ERROR, "0002", fmt.Sprintf("action not exists: %s", path), nil)
		return
	}
	action(w, r)
}

func (router *Router) SetModule(id string, module *base.Module) {
	router.Modules[id] = module
}

func (router *Router) GetModule(id string) (module *base.Module) {
	return router.Modules[id]
}

func (router *Router) DelModule(id string) {
	delete(router.Modules, id)
}

func (router *Router) HasModule(id string) (has bool) {
	_, ok := router.Modules[id]
	return ok
}