package simple_factory

import "fmt"

// 定义API接口
type API interface {
	Say(name string) string
}

// 通过不同的类型返回不同的API实现
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI实现
type hiAPI struct {
}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI实现
type helloAPI struct {
}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
