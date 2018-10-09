package base

//系统根接口，所有 struct 都必须实现该接口
type Configurable interface {

}

//base component 系统根结构，所有 struct 都必须继承
type BComponent struct {
	Configurable
}

//core component 系统核心结构，framework 中所有 struct 都必须继承
type CComponent struct {
	BComponent
}

//Module 继承自 ServiceLocator 用来提供服务定位
//通过 ServiceLocator 达到解耦和控制反转的目的
type Module struct {
	ServiceLocator
}

//Application 继承自 Module，是应用的入口模块
type Application struct {
	Module
}
