package main

import (
	"fmt"
	"time"
)

func f1() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("[f1 error]", err)
		}
	}()
	a, b := 3, 0
	fmt.Println(a, b)
	_ = a / b // panic
	fmt.Println("f1 finish")
}

func main() {
	go f1()
	time.Sleep(time.Second * 1)
	fmt.Println("main finish")
}
