package base

//定义ServiceLocator接口，ServiceLocator必须实现该接口
type IServiceLocator interface {
	//根据 id 获取服务实例
	GetService(id string) (service interface{}, err error)
	//设置服务
	SetService(id string, service interface{}) (err error)
	//删除服务
	DelService(id string) (service interface{}, err error)
	//根据 id 查看服务是否存在
	HasService(id string) (has bool, err error)
}

//服务定位模式（Service Locator Pattern）
//该模式使用一个称为"Service Locator"的中心注册表来处理请求并返回处理特定任务所需的必要信息。
//使用 Service Locator Pattern 来达成以下目标：
//	把类与依赖项解耦，从而使这些依赖项可被替换或者更新。
//	类在编译时并不知道依赖项的具体实现。
//	类的隔离性和可测试性非常好。
//	类无需负责依赖项的创建、定位和管理逻辑。
//	通过将应用程序分解为松耦合的模块，达成模块间的无依赖开发、测试、版本控制和部署。
type ServiceLocator struct {
	CComponent
	services map[string]interface{}
}