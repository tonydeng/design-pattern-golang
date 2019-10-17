package adapter

// Target是适配的目标接口
type Target interface {
	Request() string
}

// Adaptee是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

// 被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// 被适配的目标接口实现
type adapteeImpl struct {
}

//目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

//转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

// Adapter的工厂方法
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}
