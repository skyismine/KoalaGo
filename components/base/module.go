package base

type IModule interface {
	SetController(id string, controller *Controller)
	GetController(id string) (controller *Controller)
	DelController(id string)
	HasController(id string) (has bool)
}

//Module 继承自 ServiceLocator 用来提供服务定位
//通过 ServiceLocator 达到解耦和控制反转的目的
type Module struct {
	ServiceLocator
	Controllers map[string]*Controller
}

const (
	EVENT_BEFORE_ACTION = "beforeAction"
	EVENT_AFTER_ACTION = "afterAction"
)

func NewModule() IModule {
	return &Module{
		Controllers:make(map[string]*Controller),
	}
}

func (module *Module) SetController(id string, controller *Controller) {
	module.Controllers[id] = controller
}

func (module *Module) GetController(id string) (controller *Controller) {
	return module.Controllers[id]
}

func (module *Module) DelController(id string) {
	delete(module.Controllers, id)
}

func (module *Module) HasController(id string) (has bool) {
	_, ok := module.Controllers[id]
	return ok
}