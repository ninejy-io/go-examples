package main

import "sync"

type lazy struct{}

func newLazy() *lazy {
	return &lazy{}
}

type Instance2 interface {
	Work2() string
}

func (l *lazy) Work2() string {
	return "lazy is working..."
}

var l *lazy
var once sync.Once

// var lock sync.Mutex

// func GetLazy() Instance2 {
// 	if l != nil {
// 		return l
// 	}
// 	lock.Lock()
// 	defer lock.Unlock()
// 	// double check
// 	if l != nil {
// 		return l
// 	}
// 	l = newLazy()
// 	return l
// }

func GetLazy() Instance2 {
	once.Do(func() {
		l = newLazy()
	})
	return l
}

// 懒惰模式实现单例
