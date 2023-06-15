package main

import (
	"encoding/json"
	"testing"
)

type Student struct {
	Name   string
	Age    int
	Gender bool
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s  = Student{"张三", 18, true}
	c1 = Class{
		Id:       "1(2)班",
		Students: []Student{s, s, s},
	}
)

// 单元测试
// go test -v -run=TestJson json_test.go
func TestJson(t *testing.T) {
	bytes, err := json.Marshal(c1)
	if err != nil {
		t.Fail()
	}
	var c2 Class
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		t.Fail()
	}
}

// 基准测试(性能测试)
// go test -benchmem -bench=BenchmarkJson json_test.go
func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c1)
		var c2 Class
		_ = json.Unmarshal(bytes, &c2)
	}
}
