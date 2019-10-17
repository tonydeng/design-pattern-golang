package singleton

import "sync"

// 单例模式类
type Singleton struct {
}

var singleton *Singleton
var once sync.Once

// 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
