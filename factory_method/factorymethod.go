package factory_method

// 被封装的实际接口
type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

// 工厂接口
type OperatorFactory interface {
	Create() Operator
}

// Operator接口实现的基类，封装公用方法
type OperatorBase struct {
	a, b int
}

// 设置A
func (o *OperatorBase) SetA(a int) {
	o.a = a
}

// 设置B
func (o *OperatorBase) SetB(b int) {
	o.b = b
}

// PlusOperator的工厂类
type PlusOperatorFactory struct {
}

// Operator的实际加法实现
type PlusOperator struct {
	*OperatorBase
}

func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		OperatorBase: &OperatorBase{},
	}
}

// 获取结果
func (o PlusOperator) Result() int {
	return o.a + o.b
}

// MinusOperator的工厂类
type MinusOperatorFactory struct {
}

// Operation的实际减法实现
type MinusOperator struct {
	*OperatorBase
}

func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		OperatorBase: &OperatorBase{},
	}
}

//获取结果
func (o MinusOperator) Result() int {
	return o.a - o.b
}
