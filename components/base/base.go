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



