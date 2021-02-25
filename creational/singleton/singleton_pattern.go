package singleton

import (
	"sync"
	"sync/atomic"
)

// 定义单例类型：工程中使用一般为线程池、网络资源等
type Singleton struct {}

// 饿汉模式
var Instance = &Singleton{}

// 懒汉模式（lazy-loading）
var instance *Singleton
func GetInstance() *Singleton {
	if instance == nil{
		instance = &Singleton{}
	}
	return instance
}

// 简单加锁
var mu sync.Mutex
func ConcurrentGetInstance() *Singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil{
		instance = &Singleton{}
	}
	return instance
}

// double-check-lock: 只做介绍，强烈不建议使用
var valotile uint32
func DCLGetInstance() *Singleton {
	if atomic.LoadUint32(&valotile) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
	if valotile == 0 {
		if instance == nil {
			instance = &Singleton{}
			atomic.StoreUint32(&valotile, 1)
		}
	}
	return instance
}

// Once: double-check-lock
var once sync.Once
func OnceGetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}