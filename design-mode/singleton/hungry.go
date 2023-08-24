package main

type demo struct{}

func newDemo() *demo {
	return &demo{}
}

type Instance interface {
	Work() string
}

func (d *demo) Work() string {
	return "demo is working..."
}

var d *demo

func init() {
	d = newDemo()
}

func GetDemo() Instance {
	return d
}

// 饥饿模式实现单例
