package decorator

import "testing"

func TestExampleDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	c = WrapAddDecorator(c, 10)
	c = WrapMulDecorator(c, 8)

	res := c.Callc()

	t.Logf("res %d\n", res)
}
