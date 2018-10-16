package base

import "net/http"

type ActionFun func(w http.ResponseWriter, r *http.Request)

type IController interface {
	SetAction(id string, fun ActionFun)
	GetAction(id string) (fun ActionFun)
	DelAction(id string)
	HasAction(id string) (has bool)
}

type Controller struct {
	CComponent
	Actions map[string]ActionFun
}

func NewController() IController {
	return &Controller{
		Actions:make(map[string]ActionFun),
	}
}

func (controller *Controller) SetAction(id string, fun ActionFun) {
	controller.Actions[id] = fun
}

func (controller *Controller) GetAction(id string) (fun ActionFun) {
	return controller.Actions[id]
}

func (controller *Controller) DelAction(id string) {
	delete(controller.Actions, id)
}

func (controller *Controller) HasAction(id string) (has bool) {
	_, ok := controller.Actions[id]
	return ok
}

