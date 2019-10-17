package decorator

type Component interface {
	Callc() int
}

type ConcreteComponent struct {
}

func (*ConcreteComponent) Callc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Callc() int {
	return d.Component.Callc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Callc() int {
	return d.Component.Callc() + d.num
}
